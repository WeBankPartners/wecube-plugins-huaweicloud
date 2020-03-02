package plugins

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygrouprules"

	"github.com/sirupsen/logrus"
)

const (
	RULE_DIRECTION_EGRESS  = "egress"
	RULE_DIRECTION_INGRESS = "ingress"
)

var securityGroupRuleActions = make(map[string]Action)

func init() {
	securityGroupRuleActions["create"] = new(SecurityGroupRuleCreateAction)
	securityGroupRuleActions["delete"] = new(SecurityGroupRuleDeleteAction)
}

type SecurityGroupRulePlugin struct {
}

func (plugin *SecurityGroupRulePlugin) GetActionByName(actionName string) (Action, error) {
	action, found := securityGroupRuleActions[actionName]
	if !found {
		logrus.Errorf("SecurityGroupRule plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("SecurityGroupRule plugin,action = %s not found", actionName)
	}

	return action, nil
}

type SecurityGroupRuleCreateInputs struct {
	Inputs []SecurityGroupRuleCreateInput `json:"inputs,omitempty"`
}

type SecurityGroupRuleCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid                string `json:"guid,omitempty"`
	Id                  string `json:"id,omitempty"`
	SecurityGroupId     string `json:"security_group_id,omitempty"`
	Direction           string `json:"direction,omitempty"`
	Protocol            string `json:"protocol,omitempty"`
	PortRangeMin        string `json:"port_range_min,omitempty"`
	RemoteIpPrefix      string `json:"remote_ip_prefix,omitempty"`
	PortRangeMax        string `json:"port_range_max,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type SecurityGroupRuleCreateOutputs struct {
	Outputs []SecurityGroupRuleCreateOutput `json:"outputs,omitempty"`
}

type SecurityGroupRuleCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupRuleCreateAction struct {
}

func (action *SecurityGroupRuleCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs SecurityGroupRuleCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *SecurityGroupRuleCreateAction) checkCreateRuleParams(input SecurityGroupRuleCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.SecurityGroupId == "" {
		return fmt.Errorf("SecurityGroupId is empty")
	}

	if input.Direction == "" {
		return fmt.Errorf("Direction is empty")
	} else {
		input.Direction = strings.ToLower(input.Direction)
		if input.Direction != RULE_DIRECTION_EGRESS && input.Direction != RULE_DIRECTION_INGRESS {
			return fmt.Errorf("Direction is wrong")
		}
	}

	if err := checkPortRangeParams(input.PortRangeMin, input.PortRangeMax); err != nil {
		return err
	}

	return nil
}

func (action *SecurityGroupRuleCreateAction) createRule(input *SecurityGroupRuleCreateInput) (output SecurityGroupRuleCreateOutput, err error) {
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

	if err = action.checkCreateRuleParams(*input); err != nil {
		logrus.Errorf("checkCreateRuleParams meet error=%v", err)
		return
	}

	// create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether rule is exist.
	if input.Id != "" {
		var ruleInfo *securitygrouprules.SecurityGroupRule
		if ruleInfo, _, err = isRuleExist(sc, input.Id); err != nil {
			logrus.Errorf("check rule meet error=%v", err)
			return
		}
		if ruleInfo != nil {
			output.Id = ruleInfo.ID
			logrus.Infof("the rule[id=%v] is exist", input.Id)
			return
		}
	}

	// create security group rule
	request := securitygrouprules.CreateOpts{
		SecurityGroupId: input.SecurityGroupId,
		Direction:       strings.ToLower(input.Direction),
	}
	if input.Protocol != "" {
		request.Protocol = strings.ToLower(input.Protocol)
	}

	if input.PortRangeMin != "" {
		portMin, _ := strconv.Atoi(input.PortRangeMin)
		request.PortRangeMin = &portMin
	}
	if input.PortRangeMax != "" {
		portMax, _ := strconv.Atoi(input.PortRangeMax)
		request.PortRangeMax = &portMax
	}

	if input.RemoteIpPrefix != "" {
		request.RemoteIpPrefix = input.RemoteIpPrefix
	}

	response, err := securitygrouprules.Create(sc, request).Extract()
	if err != nil {
		logrus.Errorf("create security group rule meet error=%v", err)
		return
	}
	output.Id = response.ID

	return
}

func checkPortRangeParams(portRangeMin, portRangeMax string) error {
	var err error
	var portMin, portMax int

	if portRangeMin != "" {
		if portMin, err = strconv.Atoi(portRangeMin); err != nil {
			return err
		}
		if portMin <= -1 || portMin > 65535 {
			return fmt.Errorf("port range should be -1 ~ 65535")
		}
	}

	if portRangeMax != "" {
		if portMax, err = strconv.Atoi(portRangeMax); err != nil {
			return err
		}
		if portMax <= -1 || portMax > 65535 {
			return fmt.Errorf("port range should be -1 ~ 65535")
		}
	}
	if portMin > portMax {
		return fmt.Errorf("portRangeMin should be <= portRangeMax")
	}
	return nil
}

func isRuleExist(sc *gophercloud.ServiceClient, ruleId string) (*securitygrouprules.SecurityGroupRule, bool, error) {
	ruleInfo, err := securitygrouprules.Get(sc, ruleId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not found") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	return ruleInfo, true, nil
}

func (action *SecurityGroupRuleCreateAction) Do(inputs interface{}) (interface{}, error) {
	rules, _ := inputs.(SecurityGroupRuleCreateInputs)
	outputs := SecurityGroupRuleCreateOutputs{}
	var finalErr error
	for _, rule := range rules.Inputs {
		output, err := action.createRule(&rule)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroup rules = %v are created", rules)
	return &outputs, finalErr
}

type SecurityGroupRuleDeleteInputs struct {
	Inputs []SecurityGroupRuleDeleteInput `json:"inputs,omitempty"`
}

type SecurityGroupRuleDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupRuleDeleteOutputs struct {
	Outputs []SecurityGroupRuleDeleteOutput `json:"outputs,omitempty"`
}

type SecurityGroupRuleDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupRuleDeleteAction struct {
}

func (action *SecurityGroupRuleDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs SecurityGroupRuleDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *SecurityGroupRuleDeleteAction) checkDeleteRuleParams(input SecurityGroupRuleDeleteInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("SecurityGroupRule id is empty")
	}

	return nil
}

func (action *SecurityGroupRuleDeleteAction) deleteRule(input *SecurityGroupRuleDeleteInput) (output SecurityGroupRuleDeleteOutput, err error) {
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

	if err = action.checkDeleteRuleParams(*input); err != nil {
		logrus.Errorf("SecurityGroupRuleDeleteAction checkDeleteRuleParams meet error=%v", err)
		return
	}

	// Create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether securitygroup rule is exist
	_, exist, err := isRuleExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	// Delete securitygroup
	response := securitygrouprules.Delete(sc, input.Id)
	if response.Err != nil {
		err = response.Err
		logrus.Errorf("Delete securitygroup rule[securitygrouprule=%v] failed, error=%v", input.Id, err)
		return
	}

	return
}

func (action *SecurityGroupRuleDeleteAction) Do(inputs interface{}) (interface{}, error) {
	rules, _ := inputs.(SecurityGroupRuleDeleteInputs)

	outputs := SecurityGroupRuleDeleteOutputs{}
	var finalErr error
	for _, rule := range rules.Inputs {
		output, err := action.deleteRule(&rule)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroup rules = %v are deleted", rules)
	return &outputs, finalErr
}
