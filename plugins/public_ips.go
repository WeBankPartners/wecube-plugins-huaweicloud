package plugins

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/publicips"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	PUBLIC_IP_TYPE_5_BGP       = "5_bgp"
	BANDWIDTH_SHARE_TYPE_PER   = "PER"   //独占带宽
	BANDWIDTH_SHARE_TYPE_WHOLE = "WHOLE" //共享带宽

	BANDWIDTH_SIZE_MAX = 2000
	BANDWIDTH_SIZE_MIN = 1
)

var publicIpActions = make(map[string]Action)

func init() {
	publicIpActions["create"] = new(PublicIpCreateAction)
	publicIpActions["delete"] = new(PublicIpDeleteAction)
}

type PublicIpPlugin struct {
}

func (plugin *PublicIpPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := publicIpActions[actionName]
	if !found {
		logrus.Errorf("publicIp plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("publicIp plugin,action = %s not found", actionName)
	}
	return action, nil
}

type PublicIpCreateInputs struct {
	Inputs []PublicIpCreateInput `json:"inputs,omitempty"`
}

type PublicIpCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`

	BandWidth string `json:"band_width,omitempty"`
}

type PublicIpCreateOutputs struct {
	Outputs []PublicIpCreateOutput `json:"outputs,omitempty"`
}

type PublicIpCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
	Ip   string `json:"ip,omitempty"`
}

type PublicIpCreateAction struct {
}

func (action *PublicIpCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs PublicIpCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func createPluginPublicIp(input PublicIpCreateInput) (output PublicIpCreateOutput, err error) {
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

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id != "" {
		exist := false
		exist, err = isPublicIpExist(input.CloudProviderParam, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	resp, err := createPublicIp(input.CloudProviderParam, input.BandWidth, "")
	if err != nil {
		return
	}
	output.Id = resp.ID
	output.Ip = resp.PublicIpAddress
	return
}

func (action *PublicIpCreateAction) Do(inputs interface{}) (interface{}, error) {
	publicIps, _ := inputs.(PublicIpCreateInputs)
	outputs := PublicIpCreateOutputs{}
	var finalErr error

	for _, input := range publicIps.Inputs {
		output, err := createPluginPublicIp(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all publicIp = %v are created", publicIps)
	return &outputs, finalErr
}

//-----------delete publicIp action---------------//
type PublicIpDeleteInputs struct {
	Inputs []PublicIpDeleteInput `json:"inputs,omitempty"`
}

type PublicIpDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type PublicIpDeleteOutputs struct {
	Outputs []PublicIpDeleteOutput `json:"outputs,omitempty"`
}

type PublicIpDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
}

type PublicIpDeleteAction struct {
}

func (action *PublicIpDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs PublicIpDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deletePluginPublicIp(input PublicIpDeleteInput) (output PublicIpDeleteOutput, err error) {
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

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id == "" {
		err = fmt.Errorf("empty id")
		return
	}

	err = deletePublicIp(input.CloudProviderParam, input.Id)
	return
}

func (action *PublicIpDeleteAction) Do(inputs interface{}) (interface{}, error) {
	publicIps, _ := inputs.(PublicIpDeleteInputs)
	outputs := PublicIpDeleteOutputs{}
	var finalErr error

	for _, input := range publicIps.Inputs {
		output, err := deletePluginPublicIp(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all publicIp = %v are deleted", publicIps)
	return &outputs, finalErr
}

func isPublicIpExist(params CloudProviderParam, id string) (bool, error) {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return false, err
	}

	_, err = publicips.Get(sc, id).Extract()
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

func getPublicIpInfo(params CloudProviderParam, id string) (*publicips.PublicIP, error) {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return nil, err
	}

	publicIp, err := publicips.Get(sc, id).Extract()
	if err != nil {
		logrus.Errorf("getPublicIp meet err=%v", err)
	}
	return publicIp, err
}

func createPublicIp(params CloudProviderParam, bandwidthSize string, enterpriseProjectId string) (*publicips.PublicIPCreateResp, error) {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return nil, err
	}

	size, _ := strconv.Atoi(bandwidthSize)
	resp, err := publicips.Create(sc, publicips.CreateOpts{
		Publicip: publicips.PublicIPRequest{
			Type:      PUBLIC_IP_TYPE_5_BGP,
			IPVersion: 4,
		},
		Bandwidth: publicips.BandWidth{
			Name:      "wecubeCreated",
			ShareType: BANDWIDTH_SHARE_TYPE_PER,
			Size:      size,
		},
		EnterpriseProjectId: enterpriseProjectId,
	}).Extract()
	if err != nil {
		logrus.Errorf("createPublicIp meet err=%v", err)
	}

	return resp, err
}

func updatePublicIpPortId(params CloudProviderParam, lbId string, portId string) error {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return err
	}

	_, err = publicips.Update(sc, lbId, publicips.UpdateOpts{
		PortId: portId,
	}).Extract()
	if err != nil {
		logrus.Errorf("updatePublicIpPortId meet err=%v", err)
	}

	return err
}

func deletePublicIp(params CloudProviderParam, id string) error {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return err
	}

	resp := publicips.Delete(sc, id)
	if resp.Err != nil {
		logrus.Errorf("deletePublicIp meet err=%v", err)
		return resp.Err
	}

	return nil
}

func getPublicIpByPortId(params CloudProviderParam, portId string) (*publicips.PublicIP, error) {
	sc, err := CreateVpcServiceClientV1(params)
	if err != nil {
		return nil, err
	}

	allPages, err := publicips.List(sc, publicips.ListOpts{
		Limit: 100,
	}).AllPages()
	if err != nil {
		logrus.Errorf("getPublicIpByPortId list meet err=%v", err)
		return nil, err
	}

	publicipList, err := publicips.ExtractPublicIPs(allPages)
	if err != nil {
		logrus.Errorf("getPublicIpByPortId ExtractPublicIPs meet err=%v", err)
		return nil, err
	}

	for _, resp := range publicipList {
		if resp.PortId == portId {
			return &resp, nil
		}
	}
	return nil, fmt.Errorf("can't found publicIp by portId(%v)", portId)
}
