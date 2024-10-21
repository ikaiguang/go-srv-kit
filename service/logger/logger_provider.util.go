package loggerutil

import (
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"io"
	"sync"
)

var (
	singletonMutex         sync.Once
	singletonLoggerManager LoggerManager
)

func NewSingletonLoggerManager(conf *configpb.Log, appConfig *configpb.App) (LoggerManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLoggerManager, err = NewLoggerManager(conf, appConfig)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonLoggerManager, err
}

func GetWriter(loggerManager LoggerManager) (io.Writer, error) {
	return loggerManager.GetWriter()
}

func GetLogger(loggerManager LoggerManager) (log.Logger, error) {
	return loggerManager.GetLogger()
}

func GetLoggerForMiddleware(loggerManager LoggerManager) (log.Logger, error) {
	return loggerManager.GetLoggerForMiddleware()
}

func GetLoggerForHelper(loggerManager LoggerManager) (log.Logger, error) {
	return loggerManager.GetLoggerForHelper()
}
