package registryutil

import (
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

// RegistryType ...
type RegistryType string

const (
	RegistryTypeLocal  RegistryType = "local"
	RegistryTypeConsul RegistryType = "consul"
)

// NewConsulRegistry consul
func NewConsulRegistry(consulClient *api.Client) (*consul.Registry, error) {
	var opts = []consul.Option{
		consul.WithHealthCheck(true),
		consul.WithHeartbeat(true),
	}

	return consul.New(consulClient, opts...), nil
}
