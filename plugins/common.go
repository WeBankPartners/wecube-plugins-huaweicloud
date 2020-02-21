package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/sirupsen/logrus"
)

const (
	RESULT_CODE_SUCCESS = "0"
	RESULT_CODE_ERROR   = "1"
	CLOUD_PROVIDER      = "myhuaweicloud.com"

	//identity param info
	IDENTITY_SECRET_KEY = "SecretKey"
	IDENTITY_ACCESS_KEY = "AccessKey"
	IDENTITY_DOMAIN_ID  = "DomainId"

	//cloud param info
	CLOUD_PARAM_CLOUD_DOAMIN_NAME = "CloudApiDomainName"
	CLOUD_PARAM_PROJECT_ID        = "ProjectId"
	CLOUD_PARAM_REGION            = "Region"
)

type CloudProviderParam struct {
	IdentityParams string `json:"identity_params"`
	CloudParams    string `json:"cloud_params"`
}

type CallBackParameter struct {
	Parameter string `json:"callbackParameter,omitempty"`
}

type Result struct {
	Code    string `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func UnmarshalJson(source interface{}, target interface{}) error {
	reader, ok := source.(io.Reader)
	if !ok {
		return fmt.Errorf("the source to be unmarshaled is not a io.reader type")
	}

	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("parse http request (%v) meet error (%v)", reader, err)
	}

	if err = json.Unmarshal(bodyBytes, target); err != nil {
		return fmt.Errorf("unmarshal http request (%v) meet error (%v)", reader, err)
	}
	return nil
}

func isMapHasKeys(inputMap map[string]string, keys []string, mapName string) error {
	for _, key := range keys {
		val, ok := inputMap[key]
		if !ok {
			return fmt.Errorf("%s do not have value of key[%v]", mapName, key)
		}
		if val == "" {
			return fmt.Errorf("%s key[%v] have empty value", mapName, key)
		}
	}
	return nil
}

func isCloudProvicerParamValid(param CloudProviderParam) error {
	identifyMap, err := GetMapFromProviderParams(param.IdentityParams)
	if err != nil {
		return err
	}
	identityKeys := []string{
		IDENTITY_ACCESS_KEY,
		IDENTITY_SECRET_KEY, IDENTITY_DOMAIN_ID,
	}
	if err = isMapHasKeys(identifyMap, identityKeys, "IdentityParams"); err != nil {
		return err
	}

	cloudMap, err := GetMapFromProviderParams(param.CloudParams)
	if err != nil {
		return err
	}
	cloudKeys := []string{
		CLOUD_PARAM_PROJECT_ID, CLOUD_PARAM_CLOUD_DOAMIN_NAME,
		CLOUD_PARAM_REGION,
	}
	if err = isMapHasKeys(cloudMap, cloudKeys, "CloudParams"); err != nil {
		return err
	}
	return nil
}

func createGopherCloudProviderClient(param CloudProviderParam) (*gophercloud.ProviderClient, error) {
	if err := isCloudProvicerParamValid(param); err != nil {
		return nil, err
	}

	identifyMap, _ := GetMapFromProviderParams(param.IdentityParams)
	cloudMap, _ := GetMapFromProviderParams(param.CloudParams)
	identityURL := "https://iam." + cloudMap[CLOUD_PARAM_REGION] + "." + cloudMap[CLOUD_PARAM_CLOUD_DOAMIN_NAME] + "." + "/v3"

	opts := aksk.AKSKOptions{
		IdentityEndpoint: identityURL,
		AccessKey:        identifyMap[IDENTITY_ACCESS_KEY],
		SecretKey:        identifyMap[IDENTITY_SECRET_KEY],
		DomainID:         identifyMap[IDENTITY_DOMAIN_ID],
		ProjectID:        cloudMap[CLOUD_PARAM_PROJECT_ID],
		Cloud:            cloudMap[CLOUD_PARAM_CLOUD_DOAMIN_NAME],
		Region:           cloudMap[CLOUD_PARAM_REGION],
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logrus.Errorf("Openstack authenticated client failed, error=%v", err)
		return nil, err
	}
	return provider, nil
}

func GetGopherCloudProviderClient(projectId, domainId, cloudparam, identiyParam string) (*gophercloud.ProviderClient, error) {
	cloudparamMap, err := GetMapFromProviderParams(cloudparam)
	if err != nil {
		logrus.Errorf("GetMapFromProviderParams[cloudparam=%v] meet error=%v", cloudparam, err)
		return nil, err
	}

	identiyParamMap, err := GetMapFromProviderParams(identiyParam)
	if err != nil {
		logrus.Errorf("GetMapFromProviderParams[identiyParam=%v] meet error=%v", identiyParam, err)
		return nil, err
	}

	identityEndpointArray := []string{
		"https://iam",
		cloudparamMap["Region"],
		CLOUD_PROVIDER,
		"/v3",
	}
	identityEndpoint := strings.Join(identityEndpointArray, ".")

	// AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: identityEndpoint,
		ProjectID:        projectId,
		AccessKey:        identiyParamMap["SecretId"],
		SecretKey:        identiyParamMap["SecretKey"],
		Cloud:            CLOUD_PROVIDER,
		Region:           cloudparamMap["Region"],
		DomainID:         domainId,
	}

	// Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logrus.Errorf("Openstack authenticated client failed, error=%v", err)
		return nil, err
	}
	return provider, nil
}

func GetMapFromProviderParams(providerParams string) (map[string]string, error) {
	rtnMap := make(map[string]string)
	params := strings.Split(providerParams, ";")

	if len(params) == 0 {
		return rtnMap, nil
	}

	for _, param := range params {
		afterTrimParam := strings.Trim(param, " ")
		kv := strings.Split(afterTrimParam, "=")
		if len(kv) == 2 {
			rtnMap[kv[0]] = kv[1]
		} else {
			return rtnMap, fmt.Errorf("GetMapFromProviderParams meet illegal format param=%s", param)
		}
	}
	return rtnMap, nil
}

func isValidCidr(cidr string) error {
	if _, _, err := net.ParseCIDR(cidr); err != nil {
		return fmt.Errorf("cidr(%v) is invalid", cidr)
	}
	return nil
}

func getCidrGatewayIp(cidr string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	ip = ip.Mask(ipnet.Mask)
	ip[len(ip)-1] = ip[len(ip)-1] + 1
	return ip.String(), nil
}
