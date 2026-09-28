package registrypkg

import (
	"errors"
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v3"
	"github.com/hashicorp/consul/api"
)

const (
	DefaultTimeout = time.Minute
)

// NewConsulRegistry creates a Consul registry.
// Deprecated: use NewRegistry.
func NewConsulRegistry(consulClient *api.Client, opts ...consulregistry.Option) (*consulregistry.Registry, error) {
	return NewRegistry(consulClient, opts...)
}

// NewRegistry creates a Consul registry with the package defaults.
func NewRegistry(consulClient *api.Client, opts ...consulregistry.Option) (*consulregistry.Registry, error) {
	if consulClient == nil {
		return nil, errors.New("consul client is nil")
	}

	var registryOpts = []consulregistry.Option{
		consulregistry.WithHealthCheck(true),
		consulregistry.WithHeartbeat(true),
		consulregistry.WithTimeout(DefaultTimeout),
	}
	registryOpts = append(registryOpts, opts...)

	return consulregistry.New(consulClient, registryOpts...), nil
}
