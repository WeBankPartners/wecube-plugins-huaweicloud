package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/WeBankPartners/wecube-plugins-huaweicloud/plugins"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

const (
	//是否创建包年包月资源
	CREATE_PREPREPAID_RESOUCE = false
	VM_IMAGE_ID               = "7077ec61-7553-4890-8b33-364005a590b9"
	AVAILABLE_ZONE            = "cn-south-1c"
	REQUEST_TIMEOUT           = 900 //900s
)

type EnvironmentVars struct {
	PluginServerAddr string
	AccessKey        string
	SecretKey        string
	Region           string
	ProjectId        string
	DomainId         string
}

var envVars EnvironmentVars

func getCloudProviderParam() plugins.CloudProviderParam {
	identityParams := fmt.Sprintf("AccessKey=%v;SecretKey=%v;DomainId=%v", envVars.AccessKey, envVars.SecretKey, envVars.DomainId)
	cloudParams := fmt.Sprintf("CloudApiDomainName=myhuaweicloud.com;ProjectId=%v;Region=%v", envVars.ProjectId, envVars.Region)

	param := plugins.CloudProviderParam{
		IdentityParams: identityParams,
		CloudParams:    cloudParams,
	}

	return param
}

func loadEnvironmentVars() error {
	envVars.PluginServerAddr = os.Getenv("HUAWEI_PLUGIN_ADDRESS")
	if envVars.PluginServerAddr == "" {
		envVars.PluginServerAddr = "127.0.0.1:8083"
	}
	envVars.AccessKey = os.Getenv("ACCESS_KEY")
	if envVars.AccessKey == "" {
		return fmt.Errorf("get access_key from env failed")
	}

	envVars.SecretKey = os.Getenv("SECRET_KEY")
	if envVars.SecretKey == "" {
		return fmt.Errorf("get secret_key from env failed")
	}

	envVars.Region = os.Getenv("REGION")
	if envVars.Region == "" {
		return fmt.Errorf("get region from env failed")
	}

	envVars.ProjectId = os.Getenv("PROJECT_ID")
	if envVars.ProjectId == "" {
		return fmt.Errorf("get project_id from env failed")
	}

	envVars.DomainId = os.Getenv("DOMAIN_ID")
	if envVars.DomainId == "" {
		return fmt.Errorf("get domian_id from env failed")
	}

	return nil
}

func isValidPointer(response interface{}) error {
	if nil == response {
		return errors.New("input param should not be nil")
	}

	if kind := reflect.ValueOf(response).Type().Kind(); kind != reflect.Ptr {
		return errors.New("input param should be pointer type")
	}

	return nil
}

func doHttpRequest(urlPath string, request interface{}, response interface{}) error {
	if err := isValidPointer(response); err != nil {
		return err
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	url := "http://" + envVars.PluginServerAddr + urlPath
	httpRequest, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(requestBytes))
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * REQUEST_TIMEOUT,
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}
	if httpResponse != nil {
		defer httpResponse.Body.Close()
	}

	if httpResponse.StatusCode != 200 {
		return fmt.Errorf("Cmdb DoPostHttpRequest httpResponse.StatusCode != 200,statusCode=%v", httpResponse.StatusCode)
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	//logrus.Debugf("Http response, url =%s,response=%s", url, string(body))
	commonResp := plugins.PluginResponse{}
	err = json.Unmarshal(body, &commonResp)
	if err != nil {
		return err
	}
	if commonResp.ResultCode != "0" {
		return fmt.Errorf(commonResp.ResultMsg)
	}

	outputBytes, _ := json.Marshal(commonResp.Results)
	err = json.Unmarshal(outputBytes, response)
	if err != nil {
		return err
	}
	return nil
}

type CreatedResources struct {
	VpcId         string
	VpcIdForPeers string
	SubnetId      string

	VmIdPostPaid string
	VmIpPostPaid string

	VmIdPrePaid string
	VmIpPrePaid string

	InternalLbId string
	ExternalLbId string

	NatGatewayId string
	SnatRuleId   string

	PeeringsId string

	PublicIpId string

	RdsId       string
	RdsBackupId string

	SecurityGroupId     string
	SecurityGroupRuleId string

	RouteId string
}

type ResourceFunc func(string, *CreatedResources) error

type ResourceFuncEntry struct {
	TestApiName       string
	ResourcePath      string
	Func              ResourceFunc
	IsPrePaidResource bool
}

var resourceFuncTable = []ResourceFuncEntry{
	//create funcs
	{"createVpc", "/huaweicloud/v1/vpc/create", createVpc, false},
	{"createSubnet", "/huaweicloud/v1/subnet/create", createSubnet, false},
	{"createSecurityGroup", "/huaweicloud/v1/security-group/create", createSecurityGroup, false},
	{"addSecurityRule", "/huaweicloud/v1/security-group-rule/create", addSecurityGroupRule, false},
	{"createPostPaidVm", "/huaweicloud/v1/vm/create", createPostPaidVm, false},
	{"createPrePaidVm", "/huaweicloud/v1/vm/create", createPrePaidVm, true},
	{"stopVm", "/huaweicloud/v1/vm/stop", stopVm, false},
	{"startVm", "/huaweicloud/v1/vm/start", startVm, false},
	{"createPeerings", "/huaweicloud/v1/peerings/create", createPeerings, false},
	{"createPublicIp", "/huaweicloud/v1/public-ip/create", createPublicIp, false},
	{"createNatGateway", "/huaweicloud/v1/nat-gateway/create", createNatGateway, false},
	{"addSnatRule", "/huaweicloud/v1/nat-snat-rule/add", addSnatRule, false},
	{"addRoute", "/huaweicloud/v1/route/create", addRoute, false},
	{"createInternalLb", "/huaweicloud/v1/lb/create", createInternalLb, false},
	{"createExternalLb", "/huaweicloud/v1/lb/create", createExternalLb, false},
	{"addTargetToLb", "/huaweicloud/v1/lb-target/add-backtarget", addTargetToLb, false},

	//delete funcs
	{"deletePrePaidVm", "/huaweicloud/v1/vm/terminate", deletePrePaidVm, true},
	{"deleteTargetFromLb", "/huaweicloud/v1/lb-target/del-backtarget", deleteTargetFromLb, false},
	{"deleteInternalLb", "/huaweicloud/v1/lb/delete", deleteInternalLb, false},
	{"deleteExternalLb", "/huaweicloud/v1/lb/delete", deleteExternalLb, false},
	{"deleteSnatRule", "/huaweicloud/v1/nat-snat-rule/delete", deleteSnatRule, false},
	{"deleteNatGateway", "/huaweicloud/v1/nat-gateway/delete", deleteNatGateway, false},
	{"deleteRoute", "/huaweicloud/v1/route/delete", deleteRoute, false},
	{"deletePeerings", "/huaweicloud/v1/peerings/delete", deletePeerings, false},
	{"deletePublicIp", "/huaweicloud/v1/public-ip/delete", deletePublicIp, false},
	{"deletePostPaidVm", "/huaweicloud/v1/vm/terminate", deletePostPaidVm, false},
	{"deleteSecurityRule", "/huaweicloud/v1/security-group-rule/delete", deleteSecurityGroupRule, false},
	{"deleteSecurityGroup", "/huaweicloud/v1/security-group/delete", deleteSecurityGroup, false},
	{"deleteSubnet", "/huaweicloud/v1/subnet/delete", deleteSubnet, false},
	{"deleteVpc", "/huaweicloud/v1/vpc/delete", deleteVpc, false},
}

func createVpc(path string, createdResources *CreatedResources) error {
	inputs := plugins.VpcCreateInputs{
		Inputs: []plugins.VpcCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Name:               "apiTestCreated",
				Cidr:               "192.168.0.0/16",
			},
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "234",
				Name:               "apiTestCreatedForPeerings",
				Cidr:               "10.0.0.0/16",
			},
		},
	}
	outputs := plugins.VpcCreateOutputs{}

	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("vpcId is invalid")
	}

	createdResources.VpcId = outputs.Outputs[0].Id
	createdResources.VpcIdForPeers = outputs.Outputs[1].Id
	return nil
}

func deleteVpc(path string, createdResources *CreatedResources) error {
	inputs := plugins.VpcDeleteInputs{
		Inputs: []plugins.VpcDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.VpcId,
			},
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "234",
				Id:                 createdResources.VpcIdForPeers,
			},
		},
	}

	outputs := plugins.VpcDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createSubnet(path string, createdResources *CreatedResources) error {
	inputs := plugins.SubnetCreateInputs{
		Inputs: []plugins.SubnetCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				VpcId:              createdResources.VpcId,
				Name:               "testApiCreated",
				Cidr:               "192.168.1.0/24",
			},
		},
	}

	outputs := plugins.SubnetCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("subnetId is invalid")
	}

	createdResources.SubnetId = outputs.Outputs[0].Id
	return nil
}

func deleteSubnet(path string, createdResources *CreatedResources) error {
	inputs := plugins.SubnetDeleteInputs{
		Inputs: []plugins.SubnetDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				VpcId:              createdResources.VpcId,
				Id:                 createdResources.SubnetId,
			},
		},
	}

	outputs := plugins.VpcDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createSecurityGroup(path string, createdResources *CreatedResources) error {
	inputs := plugins.SecurityGroupCreateInputs{
		Inputs: []plugins.SecurityGroupCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				VpcId:              createdResources.VpcId,
				Name:               "testApiCreated",
			},
		},
	}

	outputs := plugins.SecurityGroupCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("securityGroupId is invalid")
	}

	createdResources.SecurityGroupId = outputs.Outputs[0].Id
	return nil
}

func addSecurityGroupRule(path string, createdResources *CreatedResources) error {
	inputs := plugins.SecurityGroupRuleCreateInputs{
		Inputs: []plugins.SecurityGroupRuleCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				SecurityGroupId:    createdResources.SecurityGroupId,
				Direction:          "egress",
				Protocol:           "tcp",
				PortRangeMin:       "8080",
				PortRangeMax:       "8080",
				RemoteIpPrefix:     "10.4.0.0/20",
			},
		},
	}

	outputs := plugins.SubnetCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("securityGroupRuleId is invalid")
	}

	createdResources.SecurityGroupRuleId = outputs.Outputs[0].Id
	return nil
}

func deleteSecurityGroup(path string, createdResources *CreatedResources) error {
	inputs := plugins.SecurityGroupDeleteInputs{
		Inputs: []plugins.SecurityGroupDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.SecurityGroupId,
			},
		},
	}

	outputs := plugins.SecurityGroupDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func deleteSecurityGroupRule(path string, createdResources *CreatedResources) error {
	inputs := plugins.SecurityGroupRuleDeleteInputs{
		Inputs: []plugins.SecurityGroupRuleDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.SecurityGroupRuleId,
			},
		},
	}

	outputs := plugins.SecurityGroupRuleDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createPostPaidVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmCreateInputs{
		Inputs: []plugins.VmCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Seed:               "seed",
				ImageId:            VM_IMAGE_ID,
				HostType:           "1c1g",
				SystemDiskSize:     "50",
				VpcId:              createdResources.VpcId,
				SubnetId:           createdResources.SubnetId,
				Name:               "testApiCreatedPostPaid",
				AvailabilityZone:   AVAILABLE_ZONE,
				SecurityGroups:     createdResources.SecurityGroupId,
				ChargeType:         "postPaid",
			},
		},
	}

	outputs := plugins.VmCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("vmId is invalid")
	}

	createdResources.VmIdPostPaid = outputs.Outputs[0].Id
	createdResources.VmIpPostPaid = outputs.Outputs[0].PrivateIp

	return nil
}

func createPrePaidVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmCreateInputs{
		Inputs: []plugins.VmCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Seed:               "seed",
				ImageId:            VM_IMAGE_ID,
				HostType:           "1c1g",
				SystemDiskSize:     "50",
				VpcId:              createdResources.VpcId,
				SubnetId:           createdResources.SubnetId,
				Name:               "testApiCreatedPrePaid",
				AvailabilityZone:   AVAILABLE_ZONE,
				SecurityGroups:     createdResources.SecurityGroupId,
				ChargeType:         "prePaid",
				PeriodType:         "month", //年或月
				PeriodNum:          "1",     //年有效值[1-9],月有效值[1-3]
			},
		},
	}

	outputs := plugins.VmCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}
	if outputs.Outputs[0].Id == "" {
		return fmt.Errorf("vmId is invalid")
	}

	createdResources.VmIdPrePaid = outputs.Outputs[0].Id
	createdResources.VmIpPrePaid = outputs.Outputs[0].PrivateIp

	return nil
}

func startVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmStartInputs{
		Inputs: []plugins.VmStartInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.VmIdPostPaid,
			},
		},
	}

	outputs := plugins.VmStartOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func stopVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmStopInputs{
		Inputs: []plugins.VmStopInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.VmIdPostPaid,
			},
		},
	}

	outputs := plugins.VmStopOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func deletePrePaidVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmDeleteInputs{
		Inputs: []plugins.VmDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "456",
				Id:                 createdResources.VmIdPrePaid,
			},
		},
	}

	outputs := plugins.VmStopOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func deletePostPaidVm(path string, createdResources *CreatedResources) error {
	inputs := plugins.VmDeleteInputs{
		Inputs: []plugins.VmDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.VmIdPostPaid,
			},
		},
	}

	outputs := plugins.VmStopOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createPeerings(path string, createdResources *CreatedResources) error {
	inputs := plugins.PeeringsCreateInputs{
		Inputs: []plugins.PeeringsCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Name:               "testApiCreated",
				LocalVpcId:         createdResources.VpcId,
				PeerVpcId:          createdResources.VpcIdForPeers,
			},
		},
	}

	outputs := plugins.PeeringsCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.PeeringsId = outputs.Outputs[0].Id
	return nil
}

func deletePeerings(path string, createdResources *CreatedResources) error {
	inputs := plugins.PeeringsDeleteInputs{
		Inputs: []plugins.PeeringsDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.PeeringsId,
			},
		},
	}

	outputs := plugins.PeeringsDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createPublicIp(path string, createdResources *CreatedResources) error {
	inputs := plugins.PublicIpCreateInputs{
		Inputs: []plugins.PublicIpCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				BandWidth:          "1",
			},
		},
	}

	outputs := plugins.PublicIpCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.PublicIpId = outputs.Outputs[0].Id
	return nil
}

func deletePublicIp(path string, createdResources *CreatedResources) error {
	inputs := plugins.PublicIpDeleteInputs{
		Inputs: []plugins.PublicIpDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.PublicIpId,
			},
		},
	}

	outputs := plugins.PublicIpDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createNatGateway(path string, createdResources *CreatedResources) error {
	inputs := plugins.NatCreateInputs{
		Inputs: []plugins.NatCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Name:               "testApiCreated",
				VpcId:              createdResources.VpcId,
				SubnetId:           createdResources.SubnetId,
			},
		},
	}

	outputs := plugins.NatCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.NatGatewayId = outputs.Outputs[0].Id
	return nil
}

func deleteNatGateway(path string, createdResources *CreatedResources) error {
	inputs := plugins.NatDeleteInputs{
		Inputs: []plugins.NatDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.NatGatewayId,
			},
		},
	}

	outputs := plugins.NatDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func addSnatRule(path string, createdResources *CreatedResources) error {
	inputs := plugins.AddSnatRuleInputs{
		Inputs: []plugins.AddSnatRuleInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				GatewayId:          createdResources.NatGatewayId,
				SubnetId:           createdResources.SubnetId,
				PublicIpId:         createdResources.PublicIpId,
			},
		},
	}

	outputs := plugins.AddSnatRuleOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.SnatRuleId = outputs.Outputs[0].Id
	return nil
}

func deleteSnatRule(path string, createdResources *CreatedResources) error {
	inputs := plugins.DeleteSnatRuleInputs{
		Inputs: []plugins.DeleteSnatRuleInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.SnatRuleId,
			},
		},
	}

	outputs := plugins.DeleteSnatRuleOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func addRoute(path string, createdResources *CreatedResources) error {
	inputs := plugins.RouteCreateInputs{
		Inputs: []plugins.RouteCreateInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Destination:        "10.0.0.0/16",
				Nexthop:            createdResources.PeeringsId,
				Type:               "peering",
				VpcId:              createdResources.VpcId,
			},
		},
	}

	outputs := plugins.RouteCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.RouteId = outputs.Outputs[0].Id
	return nil
}

func deleteRoute(path string, createdResources *CreatedResources) error {
	inputs := plugins.RouteDeleteInputs{
		Inputs: []plugins.RouteDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.RouteId,
			},
		},
	}

	outputs := plugins.RouteDeleteOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createInternalLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.CreateLbInputs{
		Inputs: []plugins.CreateLbInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Name:               "apiCreatedInternalLb",
				Type:               "Internal",
				SubnetId:           createdResources.SubnetId,
			},
		},
	}

	outputs := plugins.RouteCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.InternalLbId = outputs.Outputs[0].Id
	return nil
}

func deleteInternalLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.DeleteLbInputs{
		Inputs: []plugins.DeleteLbInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.InternalLbId,
				Type:               "Internal",
			},
		},
	}

	outputs := plugins.DeleteLbOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func createExternalLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.CreateLbInputs{
		Inputs: []plugins.CreateLbInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Name:               "apiCreatedExternalLb",
				Type:               "External",
				SubnetId:           createdResources.SubnetId,
				BandwidthSize:      "1",
			},
		},
	}

	outputs := plugins.RouteCreateOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	createdResources.ExternalLbId = outputs.Outputs[0].Id
	return nil
}

func deleteExternalLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.DeleteLbInputs{
		Inputs: []plugins.DeleteLbInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.ExternalLbId,
				Type:               "External",
			},
		},
	}

	outputs := plugins.DeleteLbOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func addTargetToLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.LbHostInputs{
		Inputs: []plugins.LbHostInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				LbId:               createdResources.InternalLbId,
				Port:               "9090",
				Protocol:           "tcp",
				HostIds:            createdResources.VmIdPostPaid,
				HostPorts:          "8090",
			},
		},
	}

	outputs := plugins.LbHostOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func deleteTargetFromLb(path string, createdResources *CreatedResources) error {
	inputs := plugins.LbHostInputs{
		Inputs: []plugins.LbHostInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				LbId:               createdResources.InternalLbId,
				Port:               "9090",
				Protocol:           "tcp",
				HostIds:            createdResources.VmIdPostPaid,
				HostPorts:          "8090",
			},
		},
	}

	outputs := plugins.LbHostOutputs{}
	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func TestApis(t *testing.T) {
	createdResources := CreatedResources{}

	if err := loadEnvironmentVars(); err != nil {
		t.Errorf("loadEnvironmentVars meet err=%v", err)
		return
	}

	totalCase, failedCase := 0, 0
	for _, entry := range resourceFuncTable {
		if entry.IsPrePaidResource && !CREATE_PREPREPAID_RESOUCE {
			continue
		}
		totalCase++
		err := entry.Func(entry.ResourcePath, &createdResources)
		if err == nil {
			t.Logf("Test case%3d:%v run ok", totalCase, entry.TestApiName)
		} else {
			failedCase++
			t.Logf("Test case%3d:%v run failed, err=%v", totalCase, entry.TestApiName, err)
		}
	}

	t.Logf("createdResources=%++v", createdResources)
	t.Logf("run %v test case, %v failed", totalCase, failedCase)
}
