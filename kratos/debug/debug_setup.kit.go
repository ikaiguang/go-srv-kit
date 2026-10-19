package debugpkg

import (
	"log/slog"

	logpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/log"
)

// options 配置可选项
type options struct {
	useJSONFormat bool
}

// Option is config option.
type Option func(*options)

// WithUseJSONFormat json format
func WithUseJSONFormat() Option {
	return func(o *options) {
		o.useJSONFormat = true
	}
}

// debug
var (
	// handler log handler
	handler = defaultHandler()
)

// Setup 启动
func Setup(logger logpkg.Logger) {
	handler = logpkg.ToSlog(logger)
}

// CloseDebug ...
func CloseDebug() {
	handler = defaultHandler()
}

// defaultHandler .
func defaultHandler() *slog.Logger {
	return logpkg.NewNopLogger()
}
