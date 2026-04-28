package etcdregistry

import (
	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// NewEtcdRegistry creates etcd registry
func NewEtcdRegistry(etcdClient *clientv3.Client, opts ...etcdregistry.Option) (*etcdregistry.Registry, error) {
	var registryOpts = []etcdregistry.Option{
		etcdregistry.MaxRetry(3),
	}
	registryOpts = append(registryOpts, opts...)

	return etcdregistry.New(etcdClient, registryOpts...), nil
}
