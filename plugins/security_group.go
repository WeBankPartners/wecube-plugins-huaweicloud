package plugins

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygroups"
	"github.com/sirupsen/logrus"
)

var securityGroupActions = make(map[string]Action)

func init() {
	securityGroupActions["create"] = new(SecurityGroupCreateAction)
	securityGroupActions["delete"] = new(SecurityGroupDeleteAction)
}

type SecurityGroupPlugin struct {
}

func (plugin *SecurityGroupPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := securityGroupActions[actionName]
	if !found {
		logrus.Errorf("SecurityGroup plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("SecurityGroup plugin,action = %s not found", actionName)
	}

	return action, nil
}

type SecurityGroupCreateInputs struct {
	Inputs []SecurityGroupCreateInput `json:"inputs,omitempty"`
}

type SecurityGroupCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid                string `json:"guid,omitempty"`
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	VpcId               string `json:"vpc_id,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type SecurityGroupCreateOutputs struct {
	Outputs []SecurityGroupCreateOutput `json:"outputs,omitempty"`
}

type SecurityGroupCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupCreateAction struct {
}

func (action *SecurityGroupCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs SecurityGroupCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *SecurityGroupCreateAction) checkCreateSecurityGroupParam(input SecurityGroupCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Name == "" {
		return fmt.Errorf("SecurityGroup name is empty")
	}

	return nil
}

func (action *SecurityGroupCreateAction) createSecurityGroup(input *SecurityGroupCreateInput) (output SecurityGroupCreateOutput, err error) {
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

	if err = action.checkCreateSecurityGroupParam(*input); err != nil {
		logrus.Errorf("SecurityGroupCreateAction checkCreateSecurityGroupParam meet error=%v", err)
		return
	}

	// create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// Check whether the securitygroup exited
	var securitygroupInfo *securitygroups.SecurityGroup
	if input.Id != "" {
		securitygroupInfo, _, err = isSecurityGroupExist(sc, input.Id)
		if err != nil {
			logrus.Errorf("Get securitygroup by securitygroupId[%v] meet error=%v", input.Id, err)
			return
		}
		if securitygroupInfo != nil {
			output.Id = securitygroupInfo.ID
			logrus.Infof("Get securitygroup by securitygroupId[%v], securitygroupInfo=%++v", input.Id, *securitygroupInfo)
			return
		}
	}

	// Create securitygroup
	request := securitygroups.CreateOpts{
		Name: input.Name,
	}
	if input.VpcId != "" {
		request.VpcId = input.VpcId
	}
	if input.EnterpriseProjectId != "" {
		request.EnterpriseProjectId = input.EnterpriseProjectId
	}
	response, err := securitygroups.Create(sc, request).Extract()
	if err != nil {
		logrus.Errorf("Create securitygroup failed, request=%v, error=%v", request, err)
		return
	}

	output.Id = response.ID
	return
}

func (action *SecurityGroupCreateAction) Do(inputs interface{}) (interface{}, error) {
	securitygroups, _ := inputs.(SecurityGroupCreateInputs)
	outputs := SecurityGroupCreateOutputs{}
	var finalErr error
	for _, securitygroup := range securitygroups.Inputs {
		output, err := action.createSecurityGroup(&securitygroup)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroups = %v are created", securitygroups)
	return &outputs, finalErr
}

type SecurityGroupDeleteInputs struct {
	Inputs []SecurityGroupDeleteInput `json:"inputs,omitempty"`
}

type SecurityGroupDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupDeleteOutputs struct {
	Outputs []SecurityGroupDeleteOutput `json:"outputs,omitempty"`
}

type SecurityGroupDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type SecurityGroupDeleteAction struct {
}

func (action *SecurityGroupDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs SecurityGroupDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *SecurityGroupDeleteAction) checkDeleteSecurityGroupParam(input SecurityGroupDeleteInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("SecurityGroup id is empty")
	}

	return nil
}

func (action *SecurityGroupDeleteAction) deleteSecurityGroup(input *SecurityGroupDeleteInput) (output SecurityGroupDeleteOutput, err error) {
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

	if err = action.checkDeleteSecurityGroupParam(*input); err != nil {
		logrus.Errorf("SecurityGroupDeleteAction checkDeleteSecurityGroupParam meet error=%v", err)
		return
	}

	// Create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	//check securitygroup exist
	_, exist, err := isSecurityGroupExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	// Delete securitygroup
	response := securitygroups.Delete(sc, input.Id)
	if response.Err != nil {
		err = response.Err
		logrus.Errorf("Delete securitygroup[securitygroup=%v] failed, error=%v", input.Id, err)
		return
	}

	return
}

func isSecurityGroupExist(sc *gophercloud.ServiceClient, securitygroupId string) (*securitygroups.SecurityGroup, bool, error) {
	securitygroupInfo, err := securitygroups.Get(sc, securitygroupId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "does not exist") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	return securitygroupInfo, true, nil
}

func (action *SecurityGroupDeleteAction) Do(inputs interface{}) (interface{}, error) {
	securitygroups, _ := inputs.(SecurityGroupDeleteInputs)

	outputs := SecurityGroupDeleteOutputs{}
	var finalErr error
	logrus.Infof("securitygroups.Inputs=%v", securitygroups.Inputs)
	for _, securitygroup := range securitygroups.Inputs {
		output, err := action.deleteSecurityGroup(&securitygroup)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securitygroups = %v are deleted", securitygroups)
	return &outputs, finalErr
}
