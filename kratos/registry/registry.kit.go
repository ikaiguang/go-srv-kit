package registrypkg

import (
	"time"

	consulregistry "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

const (
	DefaultTimeout = time.Minute
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
