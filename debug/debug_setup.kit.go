package debugpkg

import (
	"io"

	"github.com/go-kratos/kratos/v2/log"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
)

// debug
var (
	// handler log handler
	handler = defaultHandler()
)

// Setup 启动
func Setup() (closer io.Closer, err error) {
	// std logger
	stdLoggerConfig := &logpkg.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logpkg.DefaultCallerSkip,
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
