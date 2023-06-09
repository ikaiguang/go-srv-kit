package curlpkg

import (
	"time"
)

// options 配置可选项
type options struct {
	timeout time.Duration
}

// Option is config option.
type Option func(*options)

// WithTimeoutDefault 默认时间选项
func WithTimeoutDefault() Option {
	return WithTimeout(DefaultTimeout)
}

// WithTimeout http.Client.Timeout
func WithTimeout(duration time.Duration) Option {
	return func(o *options) {
		o.timeout = duration
	}
}
