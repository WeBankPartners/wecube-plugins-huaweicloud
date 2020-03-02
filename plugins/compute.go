package plugins

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/sirupsen/logrus"
)

func createComputeV2Client(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("Failed to get the NewComputeV2 client: ", err)
		return nil, err
	}
	return client, nil
}
