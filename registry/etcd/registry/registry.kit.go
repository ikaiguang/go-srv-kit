package registrypkg

import (
	"errors"

	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v3"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdRegistry creates etcd registry
func NewEtcdRegistry(etcdClient *clientv3.Client, opts ...etcdregistry.Option) (*etcdregistry.Registry, error) {
	if etcdClient == nil {
		return nil, errors.New("etcd client is nil")
	}

	var registryOpts = []etcdregistry.Option{
		etcdregistry.MaxRetry(3),
	}
	registryOpts = append(registryOpts, opts...)

	return etcdregistry.New(etcdClient, registryOpts...), nil
}
