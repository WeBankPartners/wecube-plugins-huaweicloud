package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
)

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
