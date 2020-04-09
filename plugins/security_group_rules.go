package plugins

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygrouprules"
	"strconv"
	"strings"
	"time"

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
	Port                string `json:"port,omitempty"`
	RemoteIpPrefix      string `json:"remote_ip_prefix,omitempty"`
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

	if input.RemoteIpPrefix == "" {
		return fmt.Errorf("RemoteIpPrefix is empty")
	}

	if input.Port == "" {
		return fmt.Errorf("Port is empty")
	}

	if input.Protocol == "" {
		return fmt.Errorf("Protocol is empty")
	}

	return nil
}

func getPort(port string) (int, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil || portInt >= 65535 {
		return 0, fmt.Errorf("port(%s) is invalid", port)
	}

	return portInt, nil
}

func getPortMinAndMax(port string) (int, int, error) {
	port = strings.TrimSpace(port)
	if strings.EqualFold(port, "ALL") {
		return 0, 0, nil
	}

	//single port
	portInt, err := strconv.Atoi(port)
	if err == nil && portInt <= 65535 {
		return portInt, portInt, nil

	}

	//range port
	portRange := strings.Split(port, "-")
	if len(portRange) == 2 {
		firstPort, firstErr := getPort(portRange[0])
		lastPort, lastErr := getPort(portRange[1])
		if firstErr == nil && lastErr == nil && firstPort < lastPort {
			return firstPort, lastPort, nil
		}
	}

	return 0, 0, fmt.Errorf("port(%v) is unsupported port format", port)
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

	// check whether rule is sc,input.Idexist.
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

	portMin, portMax, err := getPortMinAndMax(input.Port)
	if err != nil {
		return
	}

	request.PortRangeMin = &portMin
	request.PortRangeMax = &portMax
	request.Protocol = strings.ToLower(input.Protocol)
	request.RemoteIpPrefix = input.RemoteIpPrefix
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

func waitSecurityRuleDeleteOk(sc *gophercloud.ServiceClient, ruleId string) {
	count := 0
	for {
		time.Sleep(time.Second * 5)
		_, exist, err := isRuleExist(sc, ruleId)
		if err != nil || !exist {
			break
		}

		count++
		if count > 10 {
			break
		}
	}

}

func isRuleExist(sc *gophercloud.ServiceClient, ruleId string) (*securitygrouprules.SecurityGroupRule, bool, error) {
	ruleInfo, err := securitygrouprules.Get(sc, ruleId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
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
	}

	waitSecurityRuleDeleteOk(sc, input.Id)

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
