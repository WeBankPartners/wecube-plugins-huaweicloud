package plugins

import (
	"fmt"
	"strings"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/snatrules"
	"github.com/sirupsen/logrus"
)

var snatRuleActions = make(map[string]Action)

func init() {
	snatRuleActions["add"] = new(AddSnatRuleAction)
	snatRuleActions["delete"] = new(DeleteSnatRuleAction)
}

type SnatRulePlugin struct {
}

func (plugin *SnatRulePlugin) GetActionByName(actionName string) (Action, error) {
	action, found := snatRuleActions[actionName]
	if !found {
		logrus.Errorf("snatRule plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("snatRule plugin,action = %s not found", actionName)
	}
	return action, nil
}

//------------add snat rule----------------//
type AddSnatRuleInputs struct {
	Inputs []AddSnatRuleInput `json:"inputs,omitempty"`
}

type AddSnatRuleInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`

	GatewayId  string `json:"gateway_id,omitempty"`
	SubnetId   string `json:"subnet_id,omitempty"`
	PublicIpId string `json:"public_ip_id,omitempty"`
}

type AddSnatRuleOutputs struct {
	Outputs []AddSnatRuleOutput `json:"outputs,omitempty"`
}

type AddSnatRuleOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type AddSnatRuleAction struct {
}

func (action *AddSnatRuleAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs AddSnatRuleInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func isSnatRuleExist(sc *golangsdk.ServiceClient, id string) (bool, error) {
	_, err := snatrules.Get(sc, id).Extract()
	if err != nil {
		if strings.Contains(err.Error(), "No Snat Rule exist") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func checkAddSnatParam(input AddSnatRuleInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.SubnetId == "" {
		return fmt.Errorf("subnetId is empty")
	}
	if input.GatewayId == "" {
		return fmt.Errorf("gatewayId is empty")
	}
	if input.PublicIpId == "" {
		return fmt.Errorf("publicIpId is empty")
	}

	return nil
}

func addSnatRule(input AddSnatRuleInput) (output AddSnatRuleOutput, err error) {
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

	if err = checkAddSnatParam(input); err != nil {
		return
	}

	sc, err := createNatServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	if input.Id != "" {
		exist := false
		exist, err = isSnatRuleExist(sc, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	opts := snatrules.CreateOpts{
		NatGatewayID: input.GatewayId,
		NetworkID:    input.SubnetId,
		FloatingIPID: input.PublicIpId,
		SourceType:   0,
	}
	resp, err := snatrules.Create(sc, opts).Extract()
	if err != nil {
		logrus.Errorf("create snat rule failed,err=%v", err)
		return
	}

	output.Id = resp.ID
	return
}

func (action *AddSnatRuleAction) Do(inputs interface{}) (interface{}, error) {
	rules, _ := inputs.(AddSnatRuleInputs)
	outputs := AddSnatRuleOutputs{}
	var finalErr error

	for _, input := range rules.Inputs {
		output, err := addSnatRule(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all snat rule = %v are created", rules)
	return &outputs, finalErr
}

//------------delete snat rule----------------//

type DeleteSnatRuleInputs struct {
	Inputs []DeleteSnatRuleInput `json:"inputs,omitempty"`
}

type DeleteSnatRuleInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DeleteSnatRuleOutputs struct {
	Outputs []DeleteSnatRuleOutput `json:"outputs,omitempty"`
}

type DeleteSnatRuleOutput struct {
	CallBackParameter
	Result
	Id   string `json:"id,omitempty"`
	Guid string `json:"guid,omitempty"`
}

type DeleteSnatRuleAction struct {
}

func (action *DeleteSnatRuleAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DeleteSnatRuleInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deleteSnatRule(input DeleteSnatRuleInput) (output DeleteSnatRuleOutput, err error) {
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

	sc, err := createNatServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	exist, err := isSnatRuleExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	if err = snatrules.Delete(sc, input.Id).ExtractErr(); err != nil {
		logrus.Errorf("delete snat rule failed err=%v", err)
	}

	return
}

func (action *DeleteSnatRuleAction) Do(inputs interface{}) (interface{}, error) {
	rules, _ := inputs.(DeleteSnatRuleInputs)
	outputs := DeleteSnatRuleOutputs{}
	var finalErr error

	for _, input := range rules.Inputs {
		output, err := deleteSnatRule(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all snat rule = %v are delete", rules)
	return &outputs, finalErr
}
