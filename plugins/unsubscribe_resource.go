package plugins

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	bbs "github.com/gophercloud/gophercloud/openstack/bss/v1/periodresource"
	bbsIntl "github.com/gophercloud/gophercloud/openstack/bssintl/v1/periodresource"
	"github.com/sirupsen/logrus"
)

const (
	UNSUBSCRIBE_RESOURCE_UNSUBTYPE   = 1
	UNSUBSCRIBE_RESOURCE_REASON_TYPE = 5
	UNSUBSCRIBE_RESOURCE_REASON      = "api test"
)

func createBbsServiceClientV1(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("NewBSSV1 failed, error=%v", err)
		return nil, err
	}

	return sc, nil
}

func createBbsIntlServiceClientV1(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("NewBSSIntlV1 failed, error=%v", err)
		return nil, err
	}

	return sc, nil
}

func unsubscribeByResourceId(client *gophercloud.ServiceClient, resourceIds []string) (string, error) {
	unSubType := UNSUBSCRIBE_RESOURCE_UNSUBTYPE
	reasonType := UNSUBSCRIBE_RESOURCE_REASON_TYPE

	opts := bbs.UnsubscribeByResourceIdOpts{
		ResourceIds:           resourceIds,
		UnSubType:             &unSubType,
		UnsubscribeReasonType: &reasonType,
		UnsubscribeReason:     UNSUBSCRIBE_RESOURCE_REASON,
	}
	detailRsp, err := bbs.UnsubscribeByResourceId(client, opts).Extract()
	if err != nil {
		return "", err
	}
	if len(detailRsp.OrderIds) == 0 {
		return "", fmt.Errorf("return no orderID")
	}
	return detailRsp.OrderIds[0], nil
}

func unsubscribeByResourceIdIntl(client *gophercloud.ServiceClient, resourceIds []string) (string, error) {
	unSubType := UNSUBSCRIBE_RESOURCE_UNSUBTYPE
	reasonType := UNSUBSCRIBE_RESOURCE_REASON_TYPE

	opts := bbsIntl.UnsubscribeByResourceIdOpts{
		ResourceIds:           resourceIds,
		UnSubType:             &unSubType,
		UnsubscribeReasonType: &reasonType,
		UnsubscribeReason:     UNSUBSCRIBE_RESOURCE_REASON,
	}
	detailRsp, err := bbsIntl.UnsubscribeByResourceId(client, opts).Extract()
	if err != nil {
		return "", err
	}
	if len(detailRsp.OrderIds) == 0 {
		return "", fmt.Errorf("return no orderID")
	}
	return detailRsp.OrderIds[0], nil
}
