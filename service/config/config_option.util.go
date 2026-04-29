package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"google.golang.org/protobuf/proto"
)

// Option is config option.
type Option func(*options)

// ConsulConfigLoader 从 Consul 加载配置的函数类型
// 由 service/consul 子模块注入，避免 service/config 硬依赖 Consul
type ConsulConfigLoader func(consulConfig *configpb.Consul, appConfig *configpb.App, loadingOpts ...Option) (*configpb.Bootstrap, error)

// options 配置可选项
type options struct {
	configs            []proto.Message
	consulConfigLoader ConsulConfigLoader
}

// WithOtherConfig 附加其他配置
func WithOtherConfig(configs ...proto.Message) Option {
	return func(o *options) {
		o.configs = append(o.configs, configs...)
	}
}

// WithConsulConfigLoader 注入 Consul 配置加载器
// 当 config_method = "consul" 时使用此加载器从 Consul 读取配置
func WithConsulConfigLoader(loader ConsulConfigLoader) Option {
	return func(o *options) {
		o.consulConfigLoader = loader
	}
}
