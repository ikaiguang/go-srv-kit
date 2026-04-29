package rabbitmqpkg

import (
	"crypto/tls"

	"github.com/ThreeDotsLabs/watermill"
)

// options ...
type options struct {
	isNonDurable bool // 非持久性
	logger       watermill.LoggerAdapter

	tlsConfig *tls.Config
}

// Option is config option.
type Option func(*options)

// WithLogger 日志
func WithLogger(logger watermill.LoggerAdapter) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// WithKratosLogger 接受与 kratos log.Logger 签名相同的 Logger 接口。
// kratos 的 log.Logger 自动满足此接口（Go 隐式接口实现）。
func WithKratosLogger(logger Logger) Option {
	return func(o *options) {
		o.logger = NewLogger(logger)
	}
}

func WithNonDurable() Option {
	return func(o *options) {
		o.isNonDurable = true
	}
}

// WithTLSConfig tls
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(o *options) {
		o.tlsConfig = tlsConfig
	}
}
