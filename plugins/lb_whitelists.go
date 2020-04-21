package plugins

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/lbaas_v2/whitelists"
	"github.com/sirupsen/logrus"
)

func createLbGolangSdkServiceClient(params CloudProviderParam) (*golangsdk.ServiceClient, error) {
	client, err := createGolangSdkProviderClient(params)
	if err != nil {
		logrus.Errorf("get golangsdk provider client failed, error=%v", err)
		return nil, err
	}

	cloudMap, _ := GetMapFromString(params.CloudParams)
	sc, err := openstack.NewNetworkV2(client, golangsdk.EndpointOpts{
		Region: cloudMap[CLOUD_PARAM_REGION],
	})
	if err != nil {
		logrus.Errorf("createNatServiceClient meet err=%v", err)
		return nil, err
	}
	return sc, err
}

var whitelistActions = make(map[string]Action)

func init() {
	whitelistActions["create"] = new(WhitelistCreateAction)
	whitelistActions["delete"] = new(WhitelistDeleteAction)
	whitelistActions["add"] = new(WhitelistAddAction)
	whitelistActions["remove"] = new(WhitelistRemoveAction)
}

type LbWhitelistPlugin struct {
}

func (plugin *LbWhitelistPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := whitelistActions[actionName]
	if !found {
		logrus.Errorf("lb-whitelist plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("lb-whitelist plugin,action = %s not found", actionName)
	}
	return action, nil
}

type WhitelistCreateInputs struct {
	Inputs []WhitelistCreateInput `json:"inputs,omitempty"`
}

type WhitelistCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid       string `json:"guid,omitempty"`
	Id         string `json:"id,omitempty"`
	ListenerId string `json:"listener_id,omitempty"`
	Whitelist  string `json:"whitelist_ips,omitempty"`
}

type WhitelistCreateOutputs struct {
	Outputs []WhitelistCreateOutput `json:"outputs,omitempty"`
}

type WhitelistCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type WhitelistCreateAction struct {
}

func (action *WhitelistCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs WhitelistCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkWhitelistCreateParams(input WhitelistCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.ListenerId == "" {
		return fmt.Errorf("listener_id is empty")
	}
	if input.Whitelist == "" {
		return fmt.Errorf("whitelist is empty")
	}
	return nil
}

func isWhitelistExist(sc *golangsdk.ServiceClient, id string) (*whitelists.Whitelist, bool, error) {
	whitelistInfo, err := whitelists.Get(sc, id).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}

	return whitelistInfo, true, nil
}

func createWhitelist(input *WhitelistCreateInput) (output WhitelistCreateOutput, err error) {
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

	if err = checkWhitelistCreateParams(*input); err != nil {
		return
	}

	list, err := GetArrayFromString(input.Whitelist, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	sc, err := createLbGolangSdkServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	var exist bool
	if input.Id != "" {
		_, exist, err = isWhitelistExist(sc, input.Id)
		if err != nil {
			return
		}
		if exist {
			output.Id = input.Id
			return
		}
	}

	opts := whitelists.CreateOpts{
		ListenerId: input.ListenerId,
		Whitelist:  strings.Join(list, ","),
	}
	listInfo, err := whitelists.Create(sc, opts).Extract()
	if err != nil {
		return
	}
	output.Id = listInfo.ID

	return
}

func (action *WhitelistCreateAction) Do(inputs interface{}) (interface{}, error) {
	whitelists, _ := inputs.(WhitelistCreateInputs)
	outputs := WhitelistCreateOutputs{}
	var finalErr error

	for _, input := range whitelists.Inputs {
		output, err := createWhitelist(&input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

type WhitelistAddInputs struct {
	Inputs []WhitelistAddInput `json:"inputs,omitempty"`
}

type WhitelistAddInput struct {
	CallBackParameter
	CloudProviderParam
	Guid      string `json:"guid,omitempty"`
	Id        string `json:"id,omitempty"`
	Whitelist string `json:"whitelist_ips,omitempty`
}

type WhitelistAddOutputs struct {
	Outputs []WhitelistAddOutput `json:"outputs,omitempty"`
}

type WhitelistAddOutput struct {
	CallBackParameter
	Result
	Guid string `json:"omitempty"`
}

type WhitelistAddAction struct {
}

func (action *WhitelistAddAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs WhitelistAddInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkWhitelistAddParams(input WhitelistAddInput) error {
	err := isCloudProviderParamValid(input.CloudProviderParam)
	if err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("whiltelist id is empty")
	}

	return nil
}

func addWhitelist(input *WhitelistAddInput) (output WhitelistAddOutput, err error) {
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

	if err = checkWhitelistAddParams(*input); err != nil {
		return
	}

	var inputList []string
	if input.Whitelist != "" {
		inputList, err = GetArrayFromString(input.Whitelist, ARRAY_SIZE_REAL, 0)
		if err != nil {
			return
		}
	}

	sc, err := createLbGolangSdkServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	// get origin whitelist
	listInfo, exist, err := isWhitelistExist(sc, input.Id)
	if err != nil {
		return
	}
	if !exist {
		err = fmt.Errorf("whitelist[%v] is not exist", input.Id)
		return
	}

	origin, err := GetArrayFromString(listInfo.Whitelist, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	// merge the whitelist origin and input
	list := MergeTwoArraysString(origin, inputList)

	opts := whitelists.UpdateOpts{}
	if input.Whitelist != "" {
		opts.Whitelist = strings.Join(list, ",")
	}
	logrus.Infof("lb-witelist update opts=%v", opts)
	_, err = whitelists.Update(sc, input.Id, opts).Extract()

	return
}

func (action *WhitelistAddAction) Do(inputs interface{}) (interface{}, error) {
	whitelists, _ := inputs.(WhitelistAddInputs)
	outputs := WhitelistAddOutputs{}
	var finalErr error

	for _, input := range whitelists.Inputs {
		output, err := addWhitelist(&input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

type WhitelistRemoveInputs struct {
	Inputs []WhitelistRemoveInput `json:"inputs,omitempty"`
}

type WhitelistRemoveInput struct {
	CallBackParameter
	CloudProviderParam
	Guid      string `json:"guid,omitempty"`
	Id        string `json:"id,omitempty"`
	Whitelist string `json:"whitelist_ips,omitempty`
}

type WhitelistRemoveOutputs struct {
	Outputs []WhitelistRemoveOutput `json:"outputs,omitempty"`
}

type WhitelistRemoveOutput struct {
	CallBackParameter
	Result
	Guid string `json:"omitempty"`
}

type WhitelistRemoveAction struct {
}

func (action *WhitelistRemoveAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs WhitelistRemoveInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkWhitelistRemoveParams(input WhitelistRemoveInput) error {
	err := isCloudProviderParamValid(input.CloudProviderParam)
	if err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("whiltelist id is empty")
	}

	return nil
}

func removeWhitelist(input *WhitelistRemoveInput) (output WhitelistRemoveOutput, err error) {
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

	if err = checkWhitelistRemoveParams(*input); err != nil {
		return
	}

	var inputList []string
	if input.Whitelist != "" {
		inputList, err = GetArrayFromString(input.Whitelist, ARRAY_SIZE_REAL, 0)
		if err != nil {
			return
		}
	}

	sc, err := createLbGolangSdkServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	// get origin whitelist
	listInfo, exist, err := isWhitelistExist(sc, input.Id)
	if err != nil {
		return
	}
	if !exist {
		err = fmt.Errorf("whitelist[%v] is not exist", input.Id)
		return
	}

	origin, err := GetArrayFromString(listInfo.Whitelist, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	// cull the whitelist input from origin
	list := CullTwoArraysString(origin, inputList)

	opts := whitelists.UpdateOpts{}
	if input.Whitelist != "" {
		opts.Whitelist = strings.Join(list, ",")
	}
	logrus.Infof("lb-witelist update opts=%v", opts)
	_, err = whitelists.Update(sc, input.Id, opts).Extract()

	return
}

func (action *WhitelistRemoveAction) Do(inputs interface{}) (interface{}, error) {
	whitelists, _ := inputs.(WhitelistRemoveInputs)
	outputs := WhitelistRemoveOutputs{}
	var finalErr error

	for _, input := range whitelists.Inputs {
		output, err := removeWhitelist(&input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

type WhitelistDeleteInputs struct {
	Inputs []WhitelistDeleteInput `json:"inputs,omitempty"`
}

type WhitelistDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type WhitelistDeleteOutputs struct {
	Outputs []WhitelistDeleteOutput `json:"outputs,omitempty"`
}

type WhitelistDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
}

type WhitelistDeleteAction struct {
}

func (action *WhitelistDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs WhitelistDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkWhitelistDeleteParams(input WhitelistDeleteInput) error {
	err := isCloudProviderParamValid(input.CloudProviderParam)
	if err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("whitelist id is empty")
	}
	return nil
}

func deleteWhitelist(input *WhitelistDeleteInput) (output WhitelistDeleteOutput, err error) {
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

	if err = checkWhitelistDeleteParams(*input); err != nil {
		return
	}

	sc, err := createLbGolangSdkServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check wether the whitelist is exist.
	_, exist, err := isWhitelistExist(sc, input.Id)
	if err != nil {
		return
	}
	if !exist {
		err = fmt.Errorf("the whitelist[%v] could not be found", input.Id)
	}

	err = whitelists.Delete(sc, input.Id).ExtractErr()

	return
}

func (action *WhitelistDeleteAction) Do(inputs interface{}) (interface{}, error) {
	whitelists, _ := inputs.(WhitelistDeleteInputs)
	outputs := WhitelistDeleteOutputs{}
	var finalErr error

	for _, input := range whitelists.Inputs {
		output, err := deleteWhitelist(&input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}
