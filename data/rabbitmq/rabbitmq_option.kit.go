package rabbitmqpkg

import (
	"crypto/tls"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-kratos/kratos/v2/log"
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

func WithKratosLogger(logger log.Logger) Option {
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
