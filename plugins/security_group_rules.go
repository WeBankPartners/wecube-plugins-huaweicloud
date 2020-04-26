package plugins

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	Inputs []SecurityGroupRuleInput `json:"inputs,omitempty"`
}

type SecurityGroupRuleInput struct {
	CallBackParameter
	CloudProviderParam
	Guid            string `json:"guid,omitempty"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
	Direction       string `json:"direction,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Port            string `json:"port,omitempty"`
	RemoteIpPrefix  string `json:"remote_ip_prefix,omitempty"`
}

type SecurityGroupRuleCreateOutputs struct {
	Outputs []SecurityGroupRuleCreateOutput `json:"outputs,omitempty"`
}

type SecurityGroupRuleCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
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

func checkRuleInputParams(input SecurityGroupRuleInput) error {
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
	if err != nil || portInt > 65535 {
		return 0, fmt.Errorf("port(%s) is invalid", port)
	}

	return portInt, nil
}

func getPortMinAndMax(port string) (int, int, error) {
	port = strings.TrimSpace(port)
	if strings.EqualFold(strings.ToUpper(port), "ALL") {
		return 1, 65535, nil
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
		if firstErr == nil && lastErr == nil && firstPort <= lastPort {
			return firstPort, lastPort, nil
		}
	}

	return 0, 0, fmt.Errorf("port(%v) is unsupported port format", port)
}

func createRule(input *SecurityGroupRuleInput) (output SecurityGroupRuleCreateOutput, err error) {
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

	if err = checkRuleInputParams(*input); err != nil {
		logrus.Errorf("checkCreateRuleParams meet error=%v", err)
		return
	}

	newInputs, err := extractSecurityGroupRules(*input)
	if err != nil {
		return
	}

	// create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	newCreateRules := []SecurityGroupRuleInput{}
	for _, newInput := range newInputs {
		var portMin, portMax int

		// create security group rule
		request := securitygrouprules.CreateOpts{
			SecurityGroupId: newInput.SecurityGroupId,
			Direction:       strings.ToLower(newInput.Direction),
		}

		portMin, portMax, err = getPortMinAndMax(newInput.Port)
		if err != nil {
			break
		}

		request.PortRangeMin = &portMin
		request.PortRangeMax = &portMax
		request.Protocol = strings.ToLower(newInput.Protocol)
		request.RemoteIpPrefix = newInput.RemoteIpPrefix
		_, err = securitygrouprules.Create(sc, request).Extract()

		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				if strings.Contains(ue.ErrorCode(), "Com.409") {
					err = nil
				} else {
					logrus.Errorf("create security group rule meet error=%v", err)
					break
				}
			}
		} else {
			newCreateRules = append(newCreateRules, newInput)
		}
	}

	if err != nil {
		logrus.Infof("clean up newCreateRules=%++v", newCreateRules)
		for _, newCreateRule := range newCreateRules {
			_, deleteErr := deleteRule(&newCreateRule)
			if deleteErr != nil {
				err = fmt.Errorf("create rule meet error=%v && clean up created rules meet error=%v", err, deleteErr)
				return
			}
		}
	}

	return
}

func extractSecurityGroupRules(input SecurityGroupRuleInput) (newInputs []SecurityGroupRuleInput, err error) {
	ruleIps, err := GetArrayFromString(input.RemoteIpPrefix, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	ports, err := GetArrayFromString(input.Port, ARRAY_SIZE_AS_EXPECTED, len(ruleIps))
	if err != nil {
		return
	}

	protocols, err := GetArrayFromString(input.Protocol, ARRAY_SIZE_AS_EXPECTED, len(ruleIps))
	if err != nil {
		return
	}

	for i, ip := range ruleIps {
		rule := SecurityGroupRuleInput{
			Port:               ports[i],
			Protocol:           protocols[i],
			Direction:          input.Direction,
			SecurityGroupId:    input.SecurityGroupId,
			CloudProviderParam: input.CloudProviderParam,
		}
		if !strings.Contains(ip, "/") {
			rule.RemoteIpPrefix = ip + "/32"
		} else {
			rule.RemoteIpPrefix = ip
		}
		newInputs = append(newInputs, rule)
	}

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

func waitSecurityRuleDeleteOk(sc *gophercloud.ServiceClient, ruleId string) error {
	count := 0
	for {
		time.Sleep(time.Second * 1)
		_, exist, err := isRuleExist(sc, ruleId)
		if err != nil {
			return err
		}
		if !exist {
			return nil
		}

		count++
		if count > 10 {
			break
		}
	}

	return fmt.Errorf("after %vs, delete rule[%v] time out", count, ruleId)
}

func isRuleExist(sc *gophercloud.ServiceClient, ruleId string) (*securitygrouprules.SecurityGroupRule, bool, error) {
	ruleInfo, err := securitygrouprules.Get(sc, ruleId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "does not exist") {
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
		output, err := createRule(&rule)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroup rules = %v are created", rules)
	return &outputs, finalErr
}

type SecurityGroupRuleDeleteInputs struct {
	Inputs []SecurityGroupRuleInput `json:"inputs,omitempty"`
}

type SecurityGroupRuleDeleteOutputs struct {
	Outputs []SecurityGroupRuleDeleteOutput `json:"outputs,omitempty"`
}

type SecurityGroupRuleDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
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

func getSecurityGroupRule(input SecurityGroupRuleInput) (string, error) {
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return "", err
	}

	opts := securitygrouprules.ListOpts{
		SecurityGroupId: input.SecurityGroupId,
	}
	allPages, err := securitygrouprules.List(sc, opts).AllPages()
	if err != nil {
		return "", err
	}
	resp, err := securitygrouprules.ExtractSecurityGroupRules(allPages)
	if err != nil {
		return "", err
	}

	minPort, maxPort, err := getPortMinAndMax(input.Port)
	if err != nil {
		return "", err
	}
	for _, rule := range resp {
		if rule.Protocol != input.Protocol {
			continue
		}
		if *rule.PortRangeMax != maxPort || *rule.PortRangeMin != minPort {
			continue
		}
		if rule.Direction != input.Direction {
			continue
		}
		if rule.RemoteIpPrefix != input.RemoteIpPrefix {
			continue
		}

		logrus.Infof("the security group rule id = %v", rule.ID)
		return rule.ID, nil
	}
	return "", fmt.Errorf("could not find the security group rule[%++v]", input)
}

func deleteRule(input *SecurityGroupRuleInput) (output SecurityGroupRuleDeleteOutput, err error) {
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

	if err = checkRuleInputParams(*input); err != nil {
		logrus.Errorf("SecurityGroupRuleDeleteAction checkDeleteRuleParams meet error=%v", err)
		return
	}

	newInputs, err := extractSecurityGroupRules(*input)
	if err != nil {
		return
	}

	// Create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	for _, newInput := range newInputs {
		// get the rule Id
		var ruleId string
		ruleId, err = getSecurityGroupRule(newInput)
		if err != nil {
			if strings.Contains(err.Error(), "could not find the security group rule") {
				err = nil
				continue
			}
			return
		}

		// Delete securitygroup
		response := securitygrouprules.Delete(sc, ruleId)
		if response.Err != nil {
			err = response.Err
			logrus.Errorf("Delete securitygroup rule[securitygrouprule=%v] failed, error=%v", ruleId, err)
			return
		}

		// err = waitSecurityRuleDeleteOk(sc, ruleId)
		// if err != nil {
		// 	return
		// }
	}

	return
}

func (action *SecurityGroupRuleDeleteAction) Do(inputs interface{}) (interface{}, error) {
	rules, _ := inputs.(SecurityGroupRuleDeleteInputs)

	outputs := SecurityGroupRuleDeleteOutputs{}
	var finalErr error
	for _, rule := range rules.Inputs {
		output, err := deleteRule(&rule)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroup rules = %v are deleted", rules)
	return &outputs, finalErr
}
