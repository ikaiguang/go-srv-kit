package debugutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// debug
var (
	// handler log handler
	handler = defaultHandler()
)

// Setup 启动
func Setup() (err error) {
	// std logger
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip + 1,
	}
	stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	handler = log.NewHelper(stdLogger)

	return err
}

// defaultHandler .
func defaultHandler() *log.Helper {
	logger, _ := logutil.NewDummyLogger()

	return log.NewHelper(logger)
}
