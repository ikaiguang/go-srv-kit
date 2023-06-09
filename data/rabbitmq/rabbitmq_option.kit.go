package rabbitmqpkg

import (
	"crypto/tls"

	"github.com/ThreeDotsLabs/watermill"
)

// options ...
type options struct {
	// isNonDurable 非持久性
	isNonDurable bool
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

// WithTLSConfig tls
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(o *options) {
		o.tlsConfig = tlsConfig
	}
}
