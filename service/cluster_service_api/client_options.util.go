package clientutil

import (
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type option struct {
	logger            log.Logger
	consulClient      *consulapi.Client
	etcdClient        *clientv3.Client
	skipRegistryCheck bool
}

type Option func(*option)

func WithLogger(logger log.Logger) Option {
	return func(opt *option) {
		opt.logger = logger
	}
}

func WithConsulClient(consulClient *consulapi.Client) Option {
	return func(opt *option) {
		opt.consulClient = consulClient
	}
}

func WithEtcdClient(etcdClient *clientv3.Client) Option {
	return func(opt *option) {
		opt.etcdClient = etcdClient
	}
}

// WithSkipRegistryCheck 跳过注册中心健康检查
// 适用于服务启动时注册中心中目标服务尚未注册的场景
func WithSkipRegistryCheck() Option {
	return func(opt *option) {
		opt.skipRegistryCheck = true
	}
}
