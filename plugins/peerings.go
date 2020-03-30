package plugins

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/peerings"
	"github.com/sirupsen/logrus"
)

var peeringsActions = make(map[string]Action)

func init() {
	peeringsActions["create"] = new(PeeringsCreateAction)
	peeringsActions["delete"] = new(PeeringsDeleteAction)
}

type PeeringsPlugin struct {
}

func (plugin *PeeringsPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := peeringsActions[actionName]
	if !found {
		logrus.Errorf("peerings plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("peerings plugin,action = %s not found", actionName)
	}
	return action, nil
}

type PeeringsCreateInputs struct {
	Inputs []PeeringsCreateInput `json:"inputs,omitempty"`
}

type PeeringsCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid       string `json:"guid,omitempty"`
	Id         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	LocalVpcId string `json:"local_vpc_id,omitempty"`
	PeerVpcId  string `json:"peer_vpc_id,omitempty"`
}

type PeeringsCreateOutputs struct {
	Outputs []PeeringsCreateOutput `json:"outputs,omitempty"`
}

type PeeringsCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type PeeringsCreateAction struct {
}

func (action *PeeringsCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs PeeringsCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkPeeringsCreateParam(input PeeringsCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if input.LocalVpcId == "" {
		return fmt.Errorf("localVpcId is empty")
	}
	if input.PeerVpcId == "" {
		return fmt.Errorf("peerVpcId is empty")
	}
	return nil
}

func isPeeringsExist(sc *gophercloud.ServiceClient, id string) (bool, error) {
	_, err := peerings.Get(sc, id).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "No VPC peering exist") {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func createPeerings(input PeeringsCreateInput) (output PeeringsCreateOutput, err error) {
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

	if err = checkPeeringsCreateParam(input); err != nil {
		return
	}

	sc, err := createVpcServiceClientV2(input.CloudProviderParam)
	if err != nil {
		return
	}

	if input.Id != "" {
		exist := false
		exist, err = isPeeringsExist(sc, input.Id)
		if err != nil {
			return
		}
		if exist {
			output.Id = input.Id
			return
		}
	}

	opts := peerings.CreateOpts{
		Name: input.Name,
		RequestVpcInfo: peerings.VPCInfo{
			VpcID: input.LocalVpcId,
		},
		AcceptVpcInfo: peerings.VPCInfo{
			VpcID: input.PeerVpcId,
		},
	}

	result, createErr := peerings.Create(sc, opts).Extract()
	if createErr != nil {
		err = createErr
		return
	}
	output.Id = result.ID
	return
}

func (action *PeeringsCreateAction) Do(inputs interface{}) (interface{}, error) {
	peerings, _ := inputs.(PeeringsCreateInputs)
	outputs := PeeringsCreateOutputs{}
	var finalErr error

	for _, input := range peerings.Inputs {
		output, err := createPeerings(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all peerings = %v are created", peerings)
	return &outputs, finalErr
}

type PeeringsDeleteInputs struct {
	Inputs []PeeringsDeleteInput `json:"inputs,omitempty"`
}

type PeeringsDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type PeeringsDeleteOutputs struct {
	Outputs []PeeringsDeleteOutput `json:"outputs,omitempty"`
}

type PeeringsDeleteOutput struct {
	CallBackParameter
	Result
	Id   string `json:"id,omitempty"`
	Guid string `json:"guid,omitempty"`
}

type PeeringsDeleteAction struct {
}

func (action *PeeringsDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs PeeringsDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deletePeerings(input PeeringsDeleteInput) (output PeeringsDeleteOutput, err error) {
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
		err = fmt.Errorf("empty peer id")
		return
	}

	sc, err := createVpcServiceClientV2(input.CloudProviderParam)
	if err != nil {
		return
	}

	exist, err := isPeeringsExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	err = peerings.Delete(sc, input.Id).ExtractErr()
	if err != nil {
		logrus.Errorf("delete peerings[id=%v] failed, error=%v", input.Id, err)
	}
	return
}

func (action *PeeringsDeleteAction) Do(inputs interface{}) (interface{}, error) {
	peerings, _ := inputs.(PeeringsDeleteInputs)
	outputs := PeeringsDeleteOutputs{}
	var finalErr error

	for _, input := range peerings.Inputs {
		output, err := deletePeerings(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all peerings = %v are deleted", peerings)
	return &outputs, finalErr
}
