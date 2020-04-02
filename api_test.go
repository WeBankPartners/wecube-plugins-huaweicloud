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
		Timeout: time.Second * 30,
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
	VpcId    string
	SubnetId string

	VmIdPostPaid string
	VmIpPostPaid string

	VmIdPrePaid string
	VmIpPrePaid string

	LbId string

	NatGatewayId string

	PeeringsId string

	PublicicIp string

	RdsId       string
	RdsBackupId string

	SecurityGroupId string
	SecurityRuleId  string

	RouteId string
}

type ResourceFunc func(string, *CreatedResources) error

type ResourceFuncEntry struct {
	TestApiName  string
	ResourcePath string
	Func         ResourceFunc
}

var resourceFuncTable = []ResourceFuncEntry{
	//create funcs
	{"createVpc", "/huaweicloud/v1/vpc/create", createVpc},

	//delete funcs
	{"destroyVpc", "/huaweicloud/v1/vpc/delete", destroyVpc},
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
	return nil
}

func destroyVpc(path string, createdResources *CreatedResources) error {
	inputs := plugins.VpcDeleteInputs{
		Inputs: []plugins.VpcDeleteInput{
			{
				CloudProviderParam: getCloudProviderParam(),
				Guid:               "123",
				Id:                 createdResources.VpcId,
			},
		},
	}
	outputs := plugins.VpcDeleteOutputs{}

	if err := doHttpRequest(path, inputs, &outputs); err != nil {
		return err
	}

	return nil
}

func TestApis(t *testing.T) {
	var createdResources CreatedResources
	if err := loadEnvironmentVars(); err != nil {
		t.Errorf("loadEnvironmentVars meet err=%v", err)
		return
	}

	failedCase := 0
	for i, entry := range resourceFuncTable {
		err := entry.Func(entry.ResourcePath, &createdResources)
		if err == nil {
			t.Logf("Test case%3d:%v run ok", i, entry.TestApiName)
		} else {
			failedCase++
			t.Logf("Test case%3d:%v run failed, err=%v", i, entry.TestApiName, err)
		}
	}

	t.Logf("createdResources=%++v", createdResources)
	t.Logf("run %v test case, %v failed", len(resourceFuncTable), failedCase)
}
