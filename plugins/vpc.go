package plugins

import (
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/routes"
	"github.com/sirupsen/logrus"
)

const (
	VPC_STATUS_OK         = "OK"
	VPC_STATUS_CREATING   = "CREATING"
	VPC_SERVICE_CLIENT_V1 = "v1"
	VPC_SERVICE_CLIENT_V2 = "v2"
)

var VpcActions = make(map[string]Action)

func init() {
	VpcActions["create"] = new(VpcCreateAction)
	VpcActions["delete"] = new(VpcDeleteAction)
}

func CreateVpcServiceClient(projectId, domainId, cloudpram, identiyParam, version string) (*gophercloud.ServiceClient, error) {
	provider, err := GetGopherCloudProviderClient(projectId, domainId, cloudpram, identiyParam)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
	}

	//Initialization service client
	if version == VPC_SERVICE_CLIENT_V2 {
		sc, err := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})
		if err != nil {
			logrus.Errorf("Get vpc v2 client failed, error=%v", err)
			return nil, err
		}
		return sc, nil
	}

	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("Get vpc v1 client failed, error=%v", err)
		return nil, err
	}

	return sc, nil
}

type VpcPlugin struct {
}

func (plugin *VpcPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := VpcActions[actionName]
	if !found {
		logrus.Errorf("VPC plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("VPC plugin,action = %s not found", actionName)
	}

	return action, nil
}

type VpcCreateInputs struct {
	Inputs []VpcCreateInput `json:"inputs,omitempty"`
}

type VpcCreateInput struct {
	CallBackParameter
	Guid         string `json:"guid,omitempty"`
	IdentiyParam string `json:"identiy_param,omitempty"`
	Cloudpram    string `json:"cloudpram,omitempty"`
	ProjectId    string `json:"project_id,omitempty"`
	DomainId     string `json:"domainId,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Cidr         string `json:"cidr,omitempty"`

	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Description    string `json:"description,omitempty"`
}

type VpcCreateOutputs struct {
	Outputs []VpcCreateOutput `json:"outputs,omitempty"`
}

type VpcCreateOutput struct {
	CallBackParameter
	Result
	Guid         string `json:"guid,omitempty"`
	RouteTableId string `json:"route_table_id,omitempty"`
	Id           string `json:"id,omitempty"`
}

type VpcCreateAction struct {
}

func (action *VpcCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VpcCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *VpcCreateAction) checkCreateVpcParam(input VpcCreateInput) error {
	if input.Guid == "" {
		return fmt.Errorf("Guid is empty")
	}
	if input.DomainId == "" {
		return fmt.Errorf("DomainId is empty")
	}
	if input.IdentiyParam == "" {
		return fmt.Errorf("IdentiyParam is empty")
	}
	if input.Cloudpram == "" {
		return fmt.Errorf("Cloudpram is empty")
	}
	if input.ProjectId == "" {
		return fmt.Errorf("ProjectId is empty")
	}

	return nil
}

func (action *VpcCreateAction) createVpc(input *VpcCreateInput) (output VpcCreateOutput, err error) {
	defer func() {
		output.Guid = input.Guid
		output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	err = action.checkCreateVpcParam(*input)
	if err != nil {
		logrus.Errorf("VpcCreateAction checkCreateVpcParam meet error=%v", err)
		return
	}

	// Create vpc service client
	sc1, err := CreateVpcServiceClient(input.ProjectId, input.DomainId, input.Cloudpram, input.IdentiyParam, VPC_SERVICE_CLIENT_V1)
	if err != nil {
		logrus.Errorf("CreateVpcServiceClient[%v] meet error=%v", VPC_SERVICE_CLIENT_V1, err)
		return
	}

	sc2, err := CreateVpcServiceClient(input.ProjectId, input.DomainId, input.Cloudpram, input.IdentiyParam, VPC_SERVICE_CLIENT_V2)
	if err != nil {
		logrus.Errorf("CreateVpcServiceClient[%v] meet error=%v", VPC_SERVICE_CLIENT_V2, err)
		return
	}

	// Check whether the vpc exited
	var vpcInfo *vpcs.VPC
	var routes []routes.Route
	if input.Id != "" {
		vpcInfo, err = vpcs.Get(sc1, input.Id).Extract()
		if err != nil {
			logrus.Errorf("Get vpc by vpcId meet error=%v", err)
			return
		}
		logrus.Infof("Get vpc by vpcId[vpcId=%v], vpcInfo=%++v", input.Id, *vpcInfo)

		routes, err = GetRoutesByVpcId(sc2, input.Id)
		if err != nil {
			logrus.Errorf("GetRoutesByVpcId[vpcId=%v] meet error=%v", input.Id, err)
			return
		}
		if len(routes) > 0 {
			output.RouteTableId = routes[0].ID
		}
		return
	}

	// Create vpc
	request := vpcs.CreateOpts{}
	if input.Name != "" {
		request.Name = input.Name
	}
	if input.Cidr != "" {
		request.Cidr = input.Cidr
	}
	if input.EnterpriseProjectId != "" {
		request.EnterpriseProjectId = input.EnterpriseProjectId
	}

	resp, err := vpcs.Create(sc1, request).Extract()
	if err != nil {
		logrus.Errorf("Create vpc failed, request=%v, error=%v", request, err)
		return
	}

	// Check whether vpc status is OK and get the first routeId of the vpc.
	output.RouteTableId, err = action.waitVpcCreated(sc1, sc2, resp.ID, 30)
	if err != nil {
		logrus.Errorf("Create vpc failed, error=%v", err)
		return
	}

	return
}

func (action *VpcCreateAction) waitVpcCreated(sc1, sc2 *gophercloud.ServiceClient, vpcId string, timeout int) (routeId string, err error) {
	var vpcInfo *vpcs.VPC
	var routes []routes.Route
	count := 1
	for {
		vpcInfo, err = vpcs.Get(sc1, vpcId).Extract()
		if err == nil && vpcInfo.Status == VPC_STATUS_OK {
			logrus.Infof("Get vpc by vpcId[vpcId=%v], vpcInfo=%++v", vpcId, *vpcInfo)
			routes, err = GetRoutesByVpcId(sc2, vpcId)
			if err != nil {
				logrus.Errorf("GetRoutesByVpcId[vpcId=%v] meet error=%v", vpcId, err)
				return
			}

			if len(routes) > 0 {
				routeId = routes[0].ID
			}
			return
		}
		if count > timeout {
			logrus.Errorf("waitVpcCreated is timeout[%v s]...", timeout*5)
			err = fmt.Errorf("waitVpcCreated is timeout[%v s]...", timeout*5)
			return
		}
		time.Sleep(time.Second * 5)
		count++
	}

	return
}

func GetRoutesByVpcId(sc2 *gophercloud.ServiceClient, vpcId string) ([]routes.Route, error) {
	allPages, err := routes.List(sc2, routes.ListOpts{
		VpcID: vpcId,
	}).AllPages()
	if err != nil {
		logrus.Errorf("GetRoutesByVpcId meet error=%v", err)
		return []routes.Route{}, err
	}

	result, err := routes.ExtractRoutes(allPages)
	if err != nil {
		logrus.Errorf("GetRoutesByVpcId meet error=%v", err)
		return []routes.Route{}, err
	}

	return result, nil
}

func (action *VpcCreateAction) Do(inputs interface{}) (interface{}, error) {
	vpcs, _ := inputs.(VpcCreateInputs)
	outputs := VpcCreateOutputs{}
	var finalErr error
	for _, vpc := range vpcs.Inputs {
		vpcOutput, err := action.createVpc(&vpc)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, vpcOutput)
	}

	logrus.Infof("all vpcs = %v are created", vpcs)
	return &outputs, finalErr
}

type VpcDeleteInputs struct {
	Inputs []VpcDeleteInput `json:"inputs,omitempty"`
}

type VpcDeleteInput struct {
	CallBackParameter
	Guid         string `json:"guid,omitempty"`
	Cloudpram    string `json:"cloudpram,omitempty"`
	IdentiyParam string `json:"identiy_param,omitempty"`
	ProjectId    string `json:"project_id,omitempty"`
	DomainId     string `json:"domainId,omitempty"`
	Id           string `json:"id,omitempty"`
}

type VpcDeleteOutputs struct {
	Outputs []VpcDeleteOutput `json:"outputs,omitempty"`
}

type VpcDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
}

type VpcDeleteAction struct {
}

func (action *VpcDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VpcDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *VpcDeleteAction) checkDeleteVpcParam(input VpcDeleteInput) error {
	if input.Guid == "" {
		return fmt.Errorf("Guid is empty")
	}
	if input.ProjectId == "" {
		return fmt.Errorf("ProjectId is empty")
	}
	if input.IdentiyParam == "" {
		return fmt.Errorf("IdentiyParam is empty")
	}
	if input.Cloudpram == "" {
		return fmt.Errorf("Cloudpram is empty")
	}
	if input.DomainId == "" {
		return fmt.Errorf("DomainId is empty")
	}
	if input.Id == "" {
		return fmt.Errorf("Id is empty")
	}

	return nil
}

func (action *VpcDeleteAction) deleteVpc(input *VpcDeleteInput) (output VpcDeleteOutput, err error) {
	defer func() {
		output.Guid = input.Guid
		output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	err = action.checkDeleteVpcParam(*input)
	if err != nil {
		logrus.Errorf("VpcDeleteAction checkDeleteVpcParam meet error=%v", err)
		return
	}

	// Create vpc service client
	sc1, err := CreateVpcServiceClient(input.ProjectId, input.DomainId, input.Cloudpram, input.IdentiyParam, VPC_SERVICE_CLIENT_V1)
	if err != nil {
		logrus.Errorf("CreateVpcServiceClient[%v] meet error=%v", VPC_SERVICE_CLIENT_V1, err)
		return
	}

	// Delete vpc
	resp := vpcs.Delete(sc1, input.Id)
	if resp.Err != nil {
		err = resp.Err
		logrus.Errorf("Delete vpc[vpcId=%v] failed, error=%v", input.Id, err)
		return
	}

	return
}

func (action *VpcDeleteAction) Do(inputs interface{}) (interface{}, error) {
	vpcs, _ := inputs.(VpcDeleteInputs)
	outputs := VpcDeleteOutputs{}
	var finalErr error
	for _, vpc := range vpcs.Inputs {
		vpcOutput, err := action.deleteVpc(&vpc)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, vpcOutput)
	}

	logrus.Infof("all vpcs = %v are deleted", vpcs)
	return &outputs, finalErr
}
