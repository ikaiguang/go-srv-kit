package debugpkg

import (
	"io"

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
func Setup(opts ...Option) (closer io.Closer, err error) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}
	// std logger
	stdLoggerConfig := &logpkg.ConfigStd{
		Level:          log.LevelDebug,
		CallerSkip:     logpkg.DefaultCallerSkip,
		UseJSONEncoder: opt.useJSONFormat,
	}
	stdLogger, err := logpkg.NewStdLogger(stdLoggerConfig)
	if err != nil {
		return closer, err
	}
	closer = stdLogger
	handler = log.NewHelper(stdLogger)

	return closer, err
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
