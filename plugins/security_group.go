package plugins

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var SecurityGroupActions = make(map[string]Action)

func init() {
	SecurityGroupActions["create"] = new(VpcCreateAction)
	SecurityGroupActions["delete"] = new(VpcDeleteAction)
}

type SecurityGroupPlugin struct {
}

func (plugin *SecurityGroupPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := SecurityGroupActions[actionName]
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
	Guid         string `json:"guid,omitempty"`
	IdentiyParam string `json:"identiy_param,omitempty"`
	Cloudpram    string `json:"cloudpram,omitempty"`
	ProjectId    string `json:"project_id,omitempty"`
	DomainId     string `json:"domainId,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	VpcId        string `json:"vpc_id,omitempty"`

	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type SecurityGroupCreateOutputs struct {
	Outputs []VpcCreateOutput `json:"outputs,omitempty"`
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
	if input.Guid == "" {
		return fmt.Errorf("Guid is empty")
	}
	if input.Name == "" {
		return fmt.Errorf("Name is empty")
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
