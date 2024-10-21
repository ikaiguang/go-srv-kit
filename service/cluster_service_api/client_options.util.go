package clientutil

import (
	"github.com/go-kratos/kratos/v2/log"
	consulapi "github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type option struct {
	logger       log.Logger
	consulClient *consulapi.Client
	etcdClient   *clientv3.Client
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
