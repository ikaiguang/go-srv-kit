package registrypkg

import (
	"errors"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v3"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdRegistry creates an etcd registry.
// Deprecated: use NewRegistry.
func NewEtcdRegistry(etcdClient *clientv3.Client, opts ...etcdregistry.Option) (*etcdregistry.Registry, error) {
	return NewRegistry(etcdClient, opts...)
}

// NewRegistry creates an etcd registry with the package defaults.
func NewRegistry(etcdClient *clientv3.Client, opts ...etcdregistry.Option) (*etcdregistry.Registry, error) {
	if etcdClient == nil {
		return nil, errors.New("etcd client is nil")
	}

	var registryOpts = []etcdregistry.Option{
		etcdregistry.MaxRetry(3),
	}
	registryOpts = append(registryOpts, opts...)

	return etcdregistry.New(etcdClient, registryOpts...), nil
}
