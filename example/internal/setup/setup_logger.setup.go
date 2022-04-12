package setup

import (
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"

	logutil "github.com/ikaiguang/go-srv-kit/log"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
)

// Logger 日志处理示例
func (s *modules) Logger() (log.Logger, []func() error, error) {
	var (
		err error
	)
	s.loggerMutex.Do(func() {
		s.logger, s.loggerCloseFnSlice, err = s.loadingLogger()
	})
	if err != nil {
		s.loggerMutex = sync.Once{}
		return nil, nil, err
	}
	return s.logger, s.loggerCloseFnSlice, err
}

// LoggerHelper 日志处理示例
func (s *modules) LoggerHelper() (log.Logger, []func() error, error) {
	var err error
	s.loggerHelperMutex.Do(func() {
		s.loggerHelper, s.loggerHelperCloseFnSlice, err = s.loadingLoggerHelper()
	})
	if err != nil {
		s.loggerHelperMutex = sync.Once{}
		return nil, nil, err
	}
	return s.loggerHelper, s.loggerHelperCloseFnSlice, err
}

// LoggerMiddleware 中间件的日志处理示例
func (s *modules) LoggerMiddleware() (log.Logger, []func() error, error) {
	var err error
	s.loggerMiddlewareMutex.Do(func() {
		s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err = s.loadingLoggerMiddleware()
	})
	if err != nil {
		s.loggerMiddlewareMutex = sync.Once{}
		return nil, nil, err
	}
	return s.loggerMiddleware, s.loggerMiddlewareCloseFnSlice, err
}

// loadingLogHelper 加载日志工具
func (s *modules) loadingLogHelper() (closeFnSlice []func() error, err error) {
	loggerInstance, closeFnSlice, err := s.LoggerHelper()
	if err != nil {
		return closeFnSlice, err
	}
	if loggerInstance == nil {
		stdlog.Println("|*** 未加载日志工具")
		return closeFnSlice, err
	}

	// 日志
	if s.Config.EnableLoggingConsole() && s.LoggerConfigForConsole() != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到控制台")
	}
	if s.Config.EnableLoggingFile() && s.LoggerConfigForFile() != nil {
		stdlog.Println("|*** 加载日志工具：日志输出到文件")
	}

	loghelper.Setup(loggerInstance)
	return closeFnSlice, err
}

// loadingLogger 初始化日志输出实例
func (s *modules) loadingLogger() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip + 1
	return s.loadingLoggerWithCallerSkip(skip)
}

// loadingLoggerHelper 初始化日志工具输出实例
func (s *modules) loadingLoggerHelper() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip + 2
	return s.loadingLoggerWithCallerSkip(skip)
}

// loadingLoggerMiddleware 初始化中间价的日志输出实例
func (s *modules) loadingLoggerMiddleware() (logger log.Logger, closeFnSlice []func() error, err error) {
	skip := logutil.DefaultCallerSkip - 1
	return s.loadingLoggerWithCallerSkip(skip)
}

// loadingLoggerWithCallerSkip 初始化日志输出实例
func (s *modules) loadingLoggerWithCallerSkip(skip int) (logger log.Logger, closeFnSlice []func() error, err error) {
	// loggers
	var loggers []log.Logger

	// DummyLogger
	stdLogger, err := logutil.NewDummyLogger()
	if err != nil {
		return logger, closeFnSlice, err
	}

	// 配置
	if !s.EnableLoggingConsole() && !s.EnableLoggingFile() {
		fakeLogger := log.MultiLogger(stdLogger)
		return fakeLogger, closeFnSlice, err
	}

	// 日志 输出到控制台
	loggerConfigForConsole := s.LoggerConfigForConsole()
	if s.Config.EnableLoggingConsole() && loggerConfigForConsole != nil {
		stdLoggerConfig := &logutil.ConfigStd{
			Level:      logutil.ParseLevel(loggerConfigForConsole.Level),
			CallerSkip: skip,
		}
		stdLoggerImpl, err := logutil.NewStdLogger(stdLoggerConfig)
		if err != nil {
			return logger, closeFnSlice, err
		}
		closeFnSlice = append(closeFnSlice, stdLoggerImpl.Sync)
		stdLogger = stdLoggerImpl
	}
	loggers = append(loggers, stdLogger)

	// 日志 输出到文件
	loggerConfigForFile := s.LoggerConfigForFile()
	if s.Config.EnableLoggingFile() && loggerConfigForFile != nil {
		// file logger
		fileLoggerConfig := &logutil.ConfigFile{
			Level:      logutil.ParseLevel(loggerConfigForFile.Level),
			CallerSkip: skip,

			Dir:      loggerConfigForFile.Dir,
			Filename: loggerConfigForFile.Filename,

			RotateTime: loggerConfigForFile.RotateTime.AsDuration(),
			RotateSize: loggerConfigForFile.RotateSize,

			StorageCounter: uint(loggerConfigForFile.StorageCounter),
			StorageAge:     loggerConfigForFile.StorageAge.AsDuration(),
		}
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return logger, closeFnSlice, err
		}
		fileLogger, err := logutil.NewFileLogger(
			fileLoggerConfig,
			logutil.WithWriter(writer),
		)
		closeFnSlice = append(closeFnSlice, fileLogger.Sync)
		if err != nil {
			return logger, closeFnSlice, err
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	if len(loggers) == 0 {
		return logger, closeFnSlice, err
	}
	return log.MultiLogger(loggers...), closeFnSlice, err
}
