package plugins

import (
	"fmt"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
	"github.com/sirupsen/logrus"
)

const (
	VPC_STATUS_OK         = "OK"
	VPC_STATUS_CREATING   = "CREATING"
	VPC_SERVICE_CLIENT_V1 = "v1"
	VPC_SERVICE_CLIENT_V2 = "v2"
)

var vpcActions = make(map[string]Action)

func init() {
	vpcActions["create"] = new(VpcCreateAction)
	vpcActions["delete"] = new(VpcDeleteAction)
}

func CreateVpcServiceClientV1(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
		return nil, err
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
	action, found := vpcActions[actionName]
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
	CloudProviderParam
	Guid                string `json:"guid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Cidr                string `json:"cidr,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Description    string `json:"description,omitempty"`
}

type VpcCreateOutputs struct {
	Outputs []VpcCreateOutput `json:"outputs,omitempty"`
}

type VpcCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
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
	if err := isCloudProvicerParamValid(input.CloudProviderParam); err != nil {
		return err
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

	// create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether the vpc exited
	var vpcInfo *vpcs.VPC
	if input.Id != "" {
		vpcInfo, _, err = isVpcExist(sc, input.Id)
		if err != nil {
			logrus.Errorf("Get vpc by vpcId meet error=%v", err)
			return
		}
		if vpcInfo != nil {
			output.Id = input.Id
			logrus.Infof("Get vpc by vpcId[vpcId=%v], vpcInfo=%++v", input.Id, *vpcInfo)
			return
		}
	}

	// create vpc
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

	resp, err := vpcs.Create(sc, request).Extract()
	if err != nil {
		logrus.Errorf("Create vpc failed, request=%v, error=%v", request, err)
		return
	}
	output.Id = resp.ID

	if err = waitVpcCreated(sc, resp.ID); err != nil {
		logrus.Errorf("waitVpcCreated failed, error=%v", err)
	}

	return
}

func isVpcExist(sc *gophercloud.ServiceClient, vpcId string) (*vpcs.VPC, bool, error) {
	vpc, err := vpcs.Get(sc, vpcId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	return vpc, true, nil
}

func getVpcStatus(sc *gophercloud.ServiceClient, vpcId string) (string, error) {
	vpcInfo, err := vpcs.Get(sc, vpcId).Extract()
	if err != nil {
		logrus.Errorf("getVpcStatus meet error =%v", err)
		return "", err
	}

	return vpcInfo.Status, nil
}

func waitVpcCreated(sc *gophercloud.ServiceClient, vpcId string) error {
	count := 1

	for {
		status, err := getVpcStatus(sc, vpcId)
		if err != nil {
			return err
		}
		if status == VPC_STATUS_OK {
			return nil
		}

		if count > 20 {
			logrus.Errorf("after %vs, waitVpcCreated is timeout,last status=%v", count*5, status)
			return fmt.Errorf("after %vs, waitVpcCreated is timeout,last status=%v", count*5, status)
		}
		time.Sleep(time.Second * 5)
		count++
	}
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
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type VpcDeleteOutputs struct {
	Outputs []VpcDeleteOutput `json:"outputs,omitempty"`
}

type VpcDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
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
	if err := isCloudProvicerParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("vpcId is empty")
	}
	return nil
}

func (action *VpcDeleteAction) deleteVpc(input *VpcDeleteInput) (output VpcDeleteOutput, err error) {
	defer func() {
		output.Guid = input.Guid
		output.Id = input.Id
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
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	//check vpc exist
	_, exist, err := isVpcExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	// Delete vpc
	resp := vpcs.Delete(sc, input.Id)
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
