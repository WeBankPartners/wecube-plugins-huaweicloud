package plugins

import (
	"fmt"
	"strings"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/natgateways"
	"github.com/sirupsen/logrus"
)

func createNatServiceClient(params CloudProviderParam) (*golangsdk.ServiceClient, error) {
	client, err := createGolangSdkProviderClient(params)
	if err != nil {
		logrus.Errorf("get golangsdk provider client failed, error=%v", err)
		return nil, err
	}

	cloudMap, _ := GetMapFromString(params.CloudParams)
	sc, err := openstack.NewNatV2(client, golangsdk.EndpointOpts{
		Region: cloudMap[CLOUD_PARAM_REGION],
	})
	if err != nil {
		logrus.Errorf("createNatServiceClient meet err=%v", err)
		return nil, err
	}
	return sc, err
}

var natActions = make(map[string]Action)

func init() {
	natActions["create"] = new(NatCreateAction)
	natActions["delete"] = new(NatDeleteAction)
}

type NatPlugin struct {
}

func (plugin *NatPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := natActions[actionName]
	if !found {
		logrus.Errorf("natGateway plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("natGateway plugin,action = %s not found", actionName)
	}
	return action, nil
}

type NatCreateInputs struct {
	Inputs []NatCreateInput `json:"inputs,omitempty"`
}

type NatCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	VpcId    string `json:"vpc_id,omitempty"`
	SubnetId string `json:"subnet_id,omitempty"`
}

type NatCreateOutputs struct {
	Outputs []NatCreateOutput `json:"outputs,omitempty"`
}

type NatCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type NatCreateAction struct {
}

func (action *NatCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs NatCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkNatGatewayCreateParam(input NatCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.Name == "" {
		return fmt.Errorf("name is empty")
	}

	if input.VpcId == "" {
		return fmt.Errorf("vpcId is empty")
	}

	if input.SubnetId == "" {
		return fmt.Errorf("subnetId is empty")
	}
	return nil
}

func isNatGatewayExist(sc *golangsdk.ServiceClient, id string) (bool, error) {
	_, err := natgateways.Get(sc, id).Extract()
	if err != nil {
		if strings.Contains(err.Error(), "No Nat Gateway exist") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createNatGateway(input NatCreateInput) (output NatCreateOutput, err error) {
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

	if err = checkNatGatewayCreateParam(input); err != nil {
		return
	}

	sc, err := createNatServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	//check if exist
	if input.Id != "" {
		exist := false
		exist, err = isNatGatewayExist(sc, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	//create natgateway
	cloudMap, _ := GetMapFromString(input.CloudProviderParam.CloudParams)
	opts := natgateways.CreateOpts{
		TenantID:          cloudMap[CLOUD_PARAM_PROJECT_ID],
		Name:              input.Name,
		Spec:              "1",
		RouterID:          input.VpcId,
		InternalNetworkID: input.SubnetId,
	}

	fmt.Printf("opts=%++v\n", opts)

	result, err := natgateways.Create(sc, opts).Extract()
	if err != nil {
		return
	}
	output.Id = result.ID
	return
}

func (action *NatCreateAction) Do(inputs interface{}) (interface{}, error) {
	gateways, _ := inputs.(NatCreateInputs)
	outputs := NatCreateOutputs{}
	var finalErr error

	for _, input := range gateways.Inputs {
		output, err := createNatGateway(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all natGateway = %v are created", gateways)
	return &outputs, finalErr
}

//---------delete nat gateway------------//
type NatDeleteInputs struct {
	Inputs []NatDeleteInput `json:"inputs,omitempty"`
}

type NatDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type NatDeleteOutputs struct {
	Outputs []NatDeleteOutput `json:"outputs,omitempty"`
}

type NatDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type NatDeleteAction struct {
}

func (action *NatDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs NatDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deleteNatGateway(input NatDeleteInput) (output NatDeleteOutput, err error) {
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

	sc, err := createNatServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	exist, err := isNatGatewayExist(sc, input.Id)
	if err != nil || !exist {
		return
	}

	if err = natgateways.Delete(sc, input.Id).ExtractErr(); err != nil {
		logrus.Errorf("natgateway(%v) delete failed,err=%v", input.Id, err)
	}

	return
}

func (action *NatDeleteAction) Do(inputs interface{}) (interface{}, error) {
	gateways, _ := inputs.(NatDeleteInputs)
	outputs := NatDeleteOutputs{}
	var finalErr error

	for _, input := range gateways.Inputs {
		output, err := deleteNatGateway(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all natGateway = %v are created", gateways)
	return &outputs, finalErr
}
