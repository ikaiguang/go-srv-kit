package clientutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
)

type DiscoveryFactory func() (registry.Discovery, error)

type option struct {
	logger            log.Logger
	discoveryFactory  map[configpb.RegistryTypeEnum_RegistryType]DiscoveryFactory
	skipRegistryCheck bool
}

type Option func(*option)

func WithLogger(logger log.Logger) Option {
	return func(opt *option) {
		opt.logger = logger
	}
}

func WithDiscoveryFactory(registryType configpb.RegistryTypeEnum_RegistryType, factory DiscoveryFactory) Option {
	return func(opt *option) {
		if opt.discoveryFactory == nil {
			opt.discoveryFactory = make(map[configpb.RegistryTypeEnum_RegistryType]DiscoveryFactory)
		}
		opt.discoveryFactory[registryType] = factory
	}
}

// WithSkipRegistryCheck 跳过注册中心健康检查
// 适用于服务启动时注册中心中目标服务尚未注册的场景
func WithSkipRegistryCheck() Option {
	return func(opt *option) {
		opt.skipRegistryCheck = true
	}
}
