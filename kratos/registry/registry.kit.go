package registrypkg

import (
	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	etcdregistry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

const (
	DefaultTimeout = time.Minute
)

// RegistryType ...
type RegistryType string

const (
	RegistryTypeLocal  RegistryType = "local"
	RegistryTypeConsul RegistryType = "consul"
	RegistryTypeEtcd   RegistryType = "etcd"
)

// NewConsulRegistry consul
func NewConsulRegistry(consulClient *api.Client, opts ...consulregistry.Option) (*consulregistry.Registry, error) {
	var registryOpts = []consulregistry.Option{
		consulregistry.WithHealthCheck(true),
		consulregistry.WithHeartbeat(true),
		consulregistry.WithTimeout(DefaultTimeout),
	}
	registryOpts = append(registryOpts, opts...)

	return consulregistry.New(consulClient, registryOpts...), nil
}

// NewEtcdRegistry creates etcd registry
func NewEtcdRegistry(etcdClient *clientv3.Client, opts ...etcdregistry.Option) (*etcdregistry.Registry, error) {
	var registryOpts = []etcdregistry.Option{
		etcdregistry.MaxRetry(3),
	}
	registryOpts = append(registryOpts, opts...)

	return etcdregistry.New(etcdClient, registryOpts...), nil
}
