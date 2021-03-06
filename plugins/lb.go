package plugins

import (
	"fmt"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
	"github.com/sirupsen/logrus"
)

const (
	LB_TYPE_INTERNAL = "Internal"
	LB_TYPE_EXTERNAL = "External"
)

var lbActions = make(map[string]Action)

func init() {
	lbActions["create"] = new(CreateLbAction)
	lbActions["delete"] = new(DeleteLbAction)
}

func createLbServiceClient(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		return nil, err
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("createLbServiceClient meet err=%v", err)
		return nil, err
	}
	return sc, err
}

type LoadbalancerPlugin struct {
}

func (plugin *LoadbalancerPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := lbActions[actionName]
	if !found {
		logrus.Errorf("loadbalancer plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("loadbalancer plugin,action = %s not found", actionName)
	}
	return action, nil
}

type CreateLbAction struct {
}

type CreateLbInputs struct {
	Inputs []CreateLbInput `json:"inputs,omitempty"`
}

type CreateLbInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`

	Name string `json:"name"`
	//VpcId          string `json:"vpc_id"`
	Type     string `json:"type"`
	SubnetId string `json:"subnet_id"`

	//external lb param
	BandwidthSize       string `json:"bandwidth_size"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type CreateLbOutputs struct {
	Outputs []CreateLbOutput `json:"outputs,omitempty"`
}

type CreateLbOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
	Vip  string `json:"vip,omitempty"`
}

func (action *CreateLbAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs CreateLbInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkLbCreateParams(input CreateLbInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Name == "" {
		return fmt.Errorf("empty name")
	}

	if err := isValidStringValue("lb type", input.Type, []string{LB_TYPE_INTERNAL, LB_TYPE_EXTERNAL}); err != nil {
		return err
	}

	if input.SubnetId == "" {
		return fmt.Errorf("empty subnetId")
	}
	if input.Type == LB_TYPE_EXTERNAL {
		if _, err := isValidInteger(input.BandwidthSize, BANDWIDTH_SIZE_MIN, BANDWIDTH_SIZE_MAX); err != nil {
			return fmt.Errorf("bandwidth size(%v) is not in [%v,%v]", input.BandwidthSize, BANDWIDTH_SIZE_MIN, BANDWIDTH_SIZE_MAX)
		}
	}

	return nil
}

func getLbInfoById(cloudProviderParam CloudProviderParam, id string) (*loadbalancers.LoadBalancer, error) {
	sc, err := createLbServiceClient(cloudProviderParam)
	if err != nil {
		return nil, err
	}

	lbInfo, err := loadbalancers.Get(sc, id).Extract()
	if err != nil {
		logrus.Errorf("getLbInfoById meet err=%v", err)
	}
	return lbInfo, err
}

func isLbExist(cloudProviderParam CloudProviderParam, id string) (bool, error) {
	_, err := getLbInfoById(cloudProviderParam, id)
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func waitLbCreateOk(cloudProviderParam CloudProviderParam, id string) error {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		lbInfo, err := getLbInfoById(cloudProviderParam, id)
		if err != nil {
			return err
		}
		if lbInfo.ProvisioningStatus == "ERROR" {
			return fmt.Errorf("waitLb createOk,meet status ==ERROR")
		}

		if lbInfo.ProvisioningStatus == "ACTIVE" {
			break
		}
	}
	return nil
}

func getLbIpAddress(input CreateLbInput, id string) (string, error) {
	lbInfo, err := getLbInfoById(input.CloudProviderParam, id)
	if err != nil {
		return "", err
	}
	if input.Type == LB_TYPE_INTERNAL {
		return lbInfo.VipAddress, nil
	}

	//https://support.huaweicloud.com/api-elb/zh-cn_topic_0096561535.html
	publicIpInfo, err := createPublicIp(input.CloudProviderParam, input.BandwidthSize, input.EnterpriseProjectId, "createdForLb")
	if err != nil {
		return "", err
	}

	//update publicIp,bind public ip to port id
	if err = updatePublicIpPortId(input.CloudProviderParam, publicIpInfo.ID, lbInfo.VipPortID); err != nil {
		return "", err
	}

	return publicIpInfo.PublicIpAddress, nil
}

func createLb(input CreateLbInput) (output CreateLbOutput, err error) {
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

	if err = checkLbCreateParams(input); err != nil {
		return
	}
	if input.Id != "" {
		exist := false
		exist, err = isLbExist(input.CloudProviderParam, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	neutronSubnetID, err := getSubnetIdByNetworkId(input.CloudProviderParam, input.SubnetId)
	if err != nil {
		return
	}

	trueVlaue := true
	opts := loadbalancers.CreateOpts{
		Name:         input.Name,
		AdminStateUp: &trueVlaue,
		Provider:     "vlb",
		VipSubnetID:  neutronSubnetID,
	}

	sc, err := createLbServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	resp, err := loadbalancers.Create(sc, opts).Extract()
	if err != nil {
		logrus.Errorf("createLb meet err=%v", err)
		return
	}
	output.Id = resp.ID
	if err = waitLbCreateOk(input.CloudProviderParam, resp.ID); err != nil {
		return
	}
	output.Vip, err = getLbIpAddress(input, resp.ID)
	return
}

func (action *CreateLbAction) Do(inputs interface{}) (interface{}, error) {
	lbs, _ := inputs.(CreateLbInputs)
	outputs := CreateLbOutputs{}
	var finalErr error

	for _, input := range lbs.Inputs {
		output, err := createLb(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

//--------------delete lb------------------//
type DeleteLbAction struct {
}

type DeleteLbInputs struct {
	Inputs []DeleteLbInput `json:"inputs,omitempty"`
}

type DeleteLbInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type DeleteLbOutputs struct {
	Outputs []DeleteLbOutput `json:"outputs,omitempty"`
}

type DeleteLbOutput struct {
	CallBackParameter
	Result
	Id   string `json:"id,omitempty"`
	Guid string `json:"guid,omitempty"`
}

func (action *DeleteLbAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DeleteLbInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deleteLbPublicIp(cloudProviderParam CloudProviderParam, id string) error {
	lbInfo, err := getLbInfoById(cloudProviderParam, id)
	if err != nil {
		return err
	}

	publicIp, err := getPublicIpByPortId(cloudProviderParam, lbInfo.VipPortID)
	if err != nil {
		return err
	}

	err = deletePublicIp(cloudProviderParam, publicIp.ID)
	return err
}

func deleteLbListenerAndPools(input DeleteLbInput) error {
	lbInfo, err := getLbInfoById(input.CloudProviderParam, input.Id)
	if err != nil {
		return err
	}

	for _, pool := range lbInfo.Pools {
		if err = deleteLbPools(input.CloudProviderParam, pool.ID); err != nil {
			return err
		}
	}
	for _, listener := range lbInfo.Listeners {
		if err = deleteLbListener(input.CloudProviderParam, listener.ID); err != nil {
			return err
		}
	}

	return nil
}

func deleteLb(input DeleteLbInput) (output DeleteLbOutput, err error) {
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

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id == "" {
		err = fmt.Errorf("empty id")
		return
	}
	if err = isValidStringValue("lb type", input.Type, []string{LB_TYPE_INTERNAL, LB_TYPE_EXTERNAL}); err != nil {
		return
	}

	exist, err := isLbExist(input.CloudProviderParam, input.Id)
	if err != nil || !exist {
		return
	}

	if input.Type == LB_TYPE_EXTERNAL {
		if err = deleteLbPublicIp(input.CloudProviderParam, input.Id); err != nil {
			return
		}
	}

	if err = deleteLbListenerAndPools(input); err != nil {
		return
	}

	//delete lb
	sc, err := createLbServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	err = loadbalancers.Delete(sc, input.Id).ExtractErr()
	if err != nil {
		logrus.Errorf("delete lb failed ,err=%v", err)
	}

	waitLbDeleteOk(input.CloudProviderParam, input.Id)

	return
}

func waitLbDeleteOk(cloudProviderParam CloudProviderParam, id string) {
	count := 0
	for {
		time.Sleep(time.Second * 5)
		exist, err := isLbExist(cloudProviderParam, id)
		if err != nil || !exist {
			break
		}

		count++
		if count > 10 {
			break
		}
	}
}

func (action *DeleteLbAction) Do(inputs interface{}) (interface{}, error) {
	lbs, _ := inputs.(DeleteLbInputs)
	outputs := DeleteLbOutputs{}
	var finalErr error

	for _, input := range lbs.Inputs {
		output, err := deleteLb(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}
