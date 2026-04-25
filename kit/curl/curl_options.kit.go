package curlpkg

import (
	"time"
)

// options 配置可选项
type options struct {
	timeout            time.Duration
	insecureSkipVerify bool // 是否跳过 TLS 证书验证，默认 false
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

// WithInsecureSkipVerify 显式启用跳过 TLS 证书验证。
// 警告：启用此选项会使连接容易受到中间人攻击，仅在开发/测试环境中使用。
func WithInsecureSkipVerify() Option {
	return func(o *options) {
		o.insecureSkipVerify = true
	}
}
