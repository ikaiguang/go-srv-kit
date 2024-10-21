package loggerutil

import (
	"io"
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
)

func (s *loggerManager) GetLoggers() (*Loggers, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return &Loggers{
		Logger:              s.logger,
		LoggerForMiddleware: s.loggerForMiddleware,
		LoggerForHelper:     s.loggerForHelper,
	}, nil
}

func (s *loggerManager) GetLogger() (log.Logger, error) {
	loggers, err := s.GetLoggers()
	if err != nil {
		return nil, err
	}
	return loggers.Logger, nil
}

func (s *loggerManager) GetLoggerForMiddleware() (log.Logger, error) {
	loggers, err := s.GetLoggers()
	if err != nil {
		return nil, err
	}
	return loggers.LoggerForMiddleware, nil
}

func (s *loggerManager) GetLoggerForHelper() (log.Logger, error) {
	loggers, err := s.GetLoggers()
	if err != nil {
		return nil, err
	}
	return loggers.LoggerForHelper, nil
}

func (s *loggerManager) setupLoggerOnce() error {
	var err error
	s.loggerOnce.Do(func() {
		err = s.setupLogger()
	})
	if err != nil {
		s.loggerOnce = sync.Once{}
	}
	return err
}

func (s *loggerManager) setupLogger() error {
	cleanup := &closer{
		cs: make([]io.Closer, 0, 6),
	}
	// 日志
	if s.conf.GetConsole().GetEnable() {
		stdlog.Println("|*** LOADING: ConsoleLogger: ...")
	}
	if s.conf.GetFile().GetEnable() {
		stdlog.Println("|*** LOADING: FileLogger: ...")
	}

	// logger
	loggerSkip := logpkg.CallerSkipForLogger
	logger, loggerClosers, err := s.loadingLoggerWithCallerSkip(loggerSkip)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// for middleware
	// 20240806 等同与 logger
	//middlewareSkip := logpkg.CallerSkipForMiddleware + 1
	middlewareSkip := logpkg.CallerSkipForLogger
	loggerForMiddleware, loggerClosers, err := s.loadingLoggerWithCallerSkip(middlewareSkip)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// for helper
	helperSkip := logpkg.CallerSkipForHelper
	loggerForHelper, loggerClosers, err := s.loadingLoggerWithCallerSkip(helperSkip)
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// prefix
	prefixKvs := s.withLoggerPrefix()
	logger = log.With(logger, prefixKvs...)
	loggerForMiddleware = log.With(loggerForMiddleware, prefixKvs...)
	loggerForHelper = log.With(loggerForHelper, prefixKvs...)

	s.logger = logger
	s.loggerForMiddleware = loggerForMiddleware
	s.loggerForHelper = loggerForHelper
	s.loggerCloser = cleanup
	return nil
}

func (s *loggerManager) withLoggerPrefix() []interface{} {
	var kvs = NewServiceInfo(s.appConfig).Kvs()
	for _, kv := range NewTracerInfo().Kvs() {
		kvs = append(kvs, kv)
	}
	return kvs
}

func (s *loggerManager) loadingLoggerWithCallerSkip(skip int) (logger log.Logger, closeFnSlice []io.Closer, err error) {
	// loggers
	var loggers []log.Logger

	// DummyLogger
	stdLogger, err := logpkg.NewDummyLogger()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return logger, closeFnSlice, errorpkg.WithStack(e)
	}

	// 日志 输出到控制台
	consoleLoggerConfig := s.conf.GetConsole()
	if consoleLoggerConfig != nil && consoleLoggerConfig.GetEnable() {
		stdLoggerConfig := &logpkg.ConfigStd{
			Level:      logpkg.ParseLevel(consoleLoggerConfig.GetLevel()),
			CallerSkip: skip,
		}
		stdLoggerImpl, err := logpkg.NewStdLogger(stdLoggerConfig)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return logger, closeFnSlice, errorpkg.WithStack(e)
		}
		//closeFnSlice = append(closeFnSlice, stdLoggerImpl)
		stdLogger = stdLoggerImpl
	}
	// 覆盖 stdLogger
	loggers = append(loggers, stdLogger)

	// 日志 输出到文件
	fileLoggerConfig := s.conf.GetFile()
	if fileLoggerConfig != nil && fileLoggerConfig.GetEnable() {
		// file logger
		fileLoggerConfig := &logpkg.ConfigFile{
			Level:      logpkg.ParseLevel(fileLoggerConfig.GetLevel()),
			CallerSkip: skip,

			Dir:      fileLoggerConfig.GetDir(),
			Filename: fileLoggerConfig.GetFilename(),

			RotateTime: fileLoggerConfig.GetRotateTime().AsDuration(),
			RotateSize: fileLoggerConfig.GetRotateSize(),

			StorageCounter: uint(fileLoggerConfig.GetStorageCounter()),
			StorageAge:     fileLoggerConfig.GetStorageAge().AsDuration(),
		}
		writer, err := s.GetWriter()
		if err != nil {
			return logger, closeFnSlice, err
		}
		fileLogger, err := logpkg.NewFileLogger(fileLoggerConfig, logpkg.WithWriter(writer))
		closeFnSlice = append(closeFnSlice, fileLogger)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return logger, closeFnSlice, errorpkg.WithStack(e)
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	return logpkg.NewMultiLogger(loggers...), closeFnSlice, err
}
