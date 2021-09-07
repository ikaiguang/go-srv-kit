package logutil

import "github.com/go-kratos/kratos/v2/log"

// logger 实现日志接口：log.Logger
type logger struct{}

// Log 输出日志
func (*logger) Log(level log.Level, keyvals ...interface{}) (err error) {
	return err
}

// NewLogger .
func NewLogger() (loggerImpl log.Logger, err error) {
	return loggerImpl, err
}
