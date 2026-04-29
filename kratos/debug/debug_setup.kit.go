package debugpkg

import (
	"github.com/go-kratos/kratos/v2/log"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
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
func Setup(logger log.Logger) {
	handler = log.NewHelper(logger)
}

// CloseDebug ...
func CloseDebug() {
	handler = defaultHandler()
}

// defaultHandler .
func defaultHandler() *log.Helper {
	logger, _ := logpkg.NewDummyLogger()

	return log.NewHelper(logger)
}
