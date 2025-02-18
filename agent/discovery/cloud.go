package discovery

import (
	"fmt"

	"github.com/trento-project/trento/internal/cloud"
	"github.com/trento-project/trento/internal/consul"
	"github.com/trento-project/trento/internal/hosts"
)

const CloudDiscoveryId string = "cloud_discovery"

type CloudDiscovery struct {
	id        string
	discovery BaseDiscovery
}

func NewCloudDiscovery(client consul.Client) CloudDiscovery {
	r := CloudDiscovery{}
	r.id = CloudDiscoveryId
	r.discovery = NewDiscovery(client)
	return r
}

func (d CloudDiscovery) GetId() string {
	return d.id
}

func (d CloudDiscovery) Discover() (string, error) {
	cloudData, err := cloud.NewCloudInstance()
	if err != nil {
		return "", err
	}

	err = cloudData.Store(d.discovery.client)
	if err != nil {
		return "", err
	}

	err = storeCloudMetadata(d.discovery.client, cloudData.Provider)
	if err != nil {
		return "", err
	}

	if cloudData.Provider == "" {
		return "No cloud provider discovered on this host", nil
	}

	return fmt.Sprintf("Cloud provider %s discovered", cloudData.Provider), nil
}

func storeCloudMetadata(client consul.Client, cloudProvider string) error {
	metadata := hosts.Metadata{
		CloudProvider: cloudProvider,
	}
	err := metadata.Store(client)
	if err != nil {
		return err
	}

	return nil
}
