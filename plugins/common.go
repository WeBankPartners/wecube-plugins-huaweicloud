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
	"strconv"
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
	identifyMap, err := GetMapFromString(param.IdentityParams)
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

	cloudMap, err := GetMapFromString(param.CloudParams)
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

	identifyMap, _ := GetMapFromString(param.IdentityParams)
	cloudMap, _ := GetMapFromString(param.CloudParams)
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

func GetMapFromString(providerParams string) (map[string]string, error) {
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
			return rtnMap, fmt.Errorf("GetMapFromString meet illegal format param=%s", param)
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

func isValidInteger(value string,min int64,max int64)(int64,error){
	valInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil  {
		return 0,err
	}

	if valInt >max || valInt < min {
		return 0,fmt.Errorf("value(%v) is not between[%v,%v]",value,min,max)
	}

	return valInt,nil 
}

func isValidStringValue(prefix string,value string,validValues []string)error{
	for _,validValue:=range validValues {
		if validValue == value {
			return nil 
		}
	}
	return fmt.Errorf("%v value(%v) is not valid",prefix,value)
}
