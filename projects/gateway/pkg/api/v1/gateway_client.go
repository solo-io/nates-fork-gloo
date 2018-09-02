package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type GatewayClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Gateway, error)
	Write(resource *Gateway, opts clients.WriteOpts) (*Gateway, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (GatewayList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan GatewayList, <-chan error, error)
}

type gatewayClient struct {
	rc clients.ResourceClient
}

func NewGatewayClient(rcFactory factory.ResourceClientFactory) (GatewayClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Gateway{},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Gateway resource client")
	}
	return &gatewayClient{
		rc: rc,
	}, nil
}

func (client *gatewayClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *gatewayClient) Register() error {
	return client.rc.Register()
}

func (client *gatewayClient) Read(namespace, name string, opts clients.ReadOpts) (*Gateway, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Gateway), nil
}

func (client *gatewayClient) Write(gateway *Gateway, opts clients.WriteOpts) (*Gateway, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(gateway, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Gateway), nil
}

func (client *gatewayClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *gatewayClient) List(namespace string, opts clients.ListOpts) (GatewayList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToGateway(resourceList), nil
}

func (client *gatewayClient) Watch(namespace string, opts clients.WatchOpts) (<-chan GatewayList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	gatewaysChan := make(chan GatewayList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				gatewaysChan <- convertToGateway(resourceList)
			case <-opts.Ctx.Done():
				close(gatewaysChan)
				return
			}
		}
	}()
	return gatewaysChan, errs, nil
}

func convertToGateway(resources resources.ResourceList) GatewayList {
	var gatewayList GatewayList
	for _, resource := range resources {
		gatewayList = append(gatewayList, resource.(*Gateway))
	}
	return gatewayList
}
