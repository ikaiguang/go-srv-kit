package debugutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"io"

	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// debug
var (
	// handler log handler
	handler = defaultHandler()
)

// Setup 启动
func Setup() (closer io.Closer, err error) {
	// std logger
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip,
	}
	stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
	if err != nil {
		return closer, err
	}
	closer = stdLogger
	handler = log.NewHelper(stdLogger)

	return closer, err
}

// defaultHandler .
func defaultHandler() *log.Helper {
	logger, _ := logutil.NewDummyLogger()

	return log.NewHelper(logger)
}
