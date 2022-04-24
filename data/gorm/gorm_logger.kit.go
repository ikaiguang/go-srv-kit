package gormutil

import (
	"strings"

	"gorm.io/gorm/logger"
)

// NewLogger 数据库日志
func NewLogger(conf *logger.Config, writers ...logger.Writer) logger.Interface {
	if len(writers) == 0 {
		return logger.New(NewStdWriter(), *conf)
	}
	w := &multiWriter{
		writers: writers,
	}
	return logger.New(w, *conf)
}

// NewLoggerForConn 数据库链接日志
func NewLoggerForConn(opt *ConnOption) logger.Interface {
	loggerConfig := &logger.Config{
		LogLevel:                  opt.LoggerLevel,
		SlowThreshold:             opt.SlowThreshold,
		Colorful:                  opt.LoggerColorful,
		IgnoreRecordNotFoundError: opt.IgnoreRecordNotFoundError,
	}
	if !opt.LoggerEnable {
		return logger.New(NewDummyWriter(), *loggerConfig)
	}
	if len(opt.LoggerWriters) == 0 {
		return logger.New(NewStdWriter(), *loggerConfig)
	}
	return NewLogger(loggerConfig, opt.LoggerWriters...)
}

// ParseLoggerLevel 日志级别
func ParseLoggerLevel(lv string) logger.LogLevel {
	switch strings.ToUpper(lv) {
	case "DEBUG":
		return logger.Info
	case "INFO":
		return logger.Info
	case "WARN":
		return logger.Warn
	case "ERROR":
		return logger.Error
	case "FATAL":
		return logger.Info
	default:
		return logger.Info
	}
}
