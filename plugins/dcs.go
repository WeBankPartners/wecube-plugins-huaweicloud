package plugins

import (
	"fmt"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/sirupsen/logrus"
)

func createDcsServiceClient(params CloudProviderParam) (*golangsdk.ServiceClient, error) {
	if err := isCloudProviderParamValid(params); err != nil {
		return nil, err
	}

	identifyMap, _ := GetMapFromString(params.IdentityParams)
	cloudMap, _ := GetMapFromString(params.CloudParams)
	identityURL := "https://iam." + cloudMap[CLOUD_PARAM_REGION] + "." + cloudMap[CLOUD_PARAM_CLOUD_DOAMIN_NAME] + "/v3"

	opts := golangsdk.AKSKAuthOptions{
		IdentityEndpoint: identityURL,
		AccessKey:        identifyMap[IDENTITY_ACCESS_KEY],
		SecretKey:        identifyMap[IDENTITY_SECRET_KEY],
		//DomainID:         identifyMap[IDENTITY_DOMAIN_ID],
		ProjectId: cloudMap[CLOUD_PARAM_PROJECT_ID],
		Domain:    cloudMap[CLOUD_PARAM_CLOUD_DOAMIN_NAME],
		Region:    cloudMap[CLOUD_PARAM_REGION],
	}
	client, err := openstack.NewClient(identityURL)
	if err != nil {
		logrus.Errorf("new client failed err=%v", err)
		return nil, err
	}
	err = openstack.Authenticate(client, opts)
	if err != nil {
		logrus.Errorf("createDcsServiceClient auth failed err=%v", err)
		return nil, err
	}
	sc, err := openstack.NewDCSServiceV1(client, golangsdk.EndpointOpts{
		Region: cloudMap[CLOUD_PARAM_REGION],
	})
	if err != nil {
		logrus.Errorf("createDcsServiceClient meet err=%v", err)
		return nil, err
	}
	return sc, err
}

var dcsActions = make(map[string]Action)

func init() {
	dcsActions["create"] = new(DcsCreateAction)
	dcsActions["delete"] = new(DcsDeleteAction)
}

type DcsPlugin struct {
}

func (plugin *DcsPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := dcsActions[actionName]
	if !found {
		logrus.Errorf("dcs plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("dcs plugin,action = %s not found", actionName)
	}
	return action, nil
}

type DcsCreateInputs struct {
	Inputs []DcsCreateInput `json:"inputs,omitempty"`
}

type DcsCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid          string `json:"guid,omitempty"`
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	EngineVersion string `json:"engine_version,omitempty"`
	Capacity      string `json:"capacity,omitempty"`
	Password      string `json:"password,omitempty"`
	// EngineType       string `json:"engine_type,omitempty"`
	// NoPasswordAccess string `json:"no_password_access,omitempty"`
	// PublicipId       string `json:"publicip_id,omitempty"`

	VpcId               string `json:"vpc_id,omitempty"`
	SubnetId            string `json:"subnet_id,omitempty"`
	SecurityGroupId     string `json:"security_group_id,omitempty"`
	AvailableZones      string `json:"az,omitempty"`
	ProductId           string `json:"product_id,omitempty"`
	PrivateIp           string `json:"private_ip,omitempty"`
	Labels              string `json:"labels,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	ChargeType          string `json:"charge_type,omitempty"`

	//包年包月
	PeriodType  string `json:"period_type,omitempty"`   //年或月
	PeriodNum   string `json:"period_num,omitempty"`    //年有效值[1-9],月有效值[1-3]
	IsAutoRenew string `json:"is_auto_renew,omitempty"` //是否自动续费
}

type DcsCreateOutputs struct {
	Outputs []DcsCreateOutput `json:"outputs,omitempty"`
}

type DcsCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DcsCreateAction struct {
}

func (action *DcsCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DcsCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *DcsCreateAction) checkCreateDcsParams(input DcsCreateInput) error {
	return nil
}

func (action *DcsCreateAction) createDcs(input *DcsCreateInput) (output DcsCreateOutput, err error) {
	return
}

func (action *DcsCreateAction) Do(inputs interface{}) (interface{}, error) {
	return nil, nil
}

type DcsDeleteInputs struct {
	Inputs []DcsDeleteInput `json:"inputs,omitempty"`
}

type DcsDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DcsDeleteOutputs struct {
	Outputs []DcsDeleteOutput `json:"outputs,omitempty"`
}

type DcsDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DcsDeleteAction struct {
}

func (action *DcsDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DcsDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *DcsDeleteAction) checkDeleteDcsParams(input DcsDeleteInput) error {
	return nil
}

func (action *DcsDeleteAction) deleteDcs(input *DcsDeleteInput) (output DcsDeleteOutput, err error) {
	return
}

func (action *DcsDeleteAction) Do(inputs interface{}) (interface{}, error) {
	return nil, nil
}
