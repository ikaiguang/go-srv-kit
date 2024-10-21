package configutil

import "google.golang.org/protobuf/proto"

// Option is config option.
type Option func(*options)

// options 配置可选项
type options struct {
	configs []proto.Message
}

func WithOtherConfig(configs ...proto.Message) Option {
	return func(o *options) {
		o.configs = append(o.configs, configs...)
	}
}
