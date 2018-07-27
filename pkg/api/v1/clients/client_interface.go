package clients

import (
	"context"
	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

const DefaultNamespace = "default"

var DefaultRefreshRate = time.Second * 30

func DefaultNamespaceIfEmpty(namespace string) string {
	if namespace == "" {
		return DefaultNamespace
	}
	return namespace
}

type ResourceClient interface {
	Kind() string
	NewResource() resources.Resource
	Register() error
	Read(namespace, name string, opts ReadOpts) (resources.Resource, error)
	Write(resource resources.Resource, opts WriteOpts) (resources.Resource, error)
	Delete(namespace, name string, opts DeleteOpts) error
	List(namespace string, opts ListOpts) ([]resources.Resource, error)
	Watch(namespace string, opts WatchOpts) (<-chan []resources.Resource, <-chan error, error)
}

type ReadOpts struct {
	Ctx context.Context
}

func (o ReadOpts) WithDefaults() ReadOpts {
	if o.Ctx == nil {
		o.Ctx = context.TODO()
	}
	return o
}

type WriteOpts struct {
	Ctx               context.Context
	OverwriteExisting bool
}

func (o WriteOpts) WithDefaults() WriteOpts {
	if o.Ctx == nil {
		o.Ctx = context.TODO()
	}
	return o
}

type DeleteOpts struct {
	Ctx            context.Context
	IgnoreNotExist bool
}

func (o DeleteOpts) WithDefaults() DeleteOpts {
	if o.Ctx == nil {
		o.Ctx = context.TODO()
	}
	return o
}

type ListOpts struct {
	Ctx      context.Context
	Selector map[string]string
}

func (o ListOpts) WithDefaults() ListOpts {
	if o.Ctx == nil {
		o.Ctx = context.TODO()
	}
	return o
}

type WatchOpts struct {
	Ctx         context.Context
	Selector    map[string]string
	RefreshRate time.Duration
}

func (o WatchOpts) WithDefaults() WatchOpts {
	if o.Ctx == nil {
		o.Ctx = context.TODO()
	}
	if o.RefreshRate == 0 {
		o.RefreshRate = DefaultRefreshRate
	}
	return o
}
