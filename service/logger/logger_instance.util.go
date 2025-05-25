package loggerutil

import (
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"io"
	stdlog "log"
	"sync"
)

func (s *loggerManager) GetLogger() (log.Logger, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return s.logger, nil
}

func (s *loggerManager) GetLoggerForMiddleware() (log.Logger, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return s.loggerForMiddleware, nil
}

func (s *loggerManager) GetLoggerForHelper() (log.Logger, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return s.loggerForHelper, nil
}

func (s *loggerManager) GetLoggerForGORM() (log.Logger, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return s.loggerForGORM, nil
}

func (s *loggerManager) GetLoggerForRabbitmq() (log.Logger, error) {
	err := s.setupLoggerOnce()
	if err != nil {
		return nil, err
	}
	return s.loggerForRabbitmq, nil
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

	writer, err := s.GetWriter()
	if err != nil {
		return err
	}
	fileLoggerConf := s.conf.GetFile()

	// logger
	loggerSkip := logpkg.CallerSkipForLogger
	logger, loggerClosers, err := s.loadingLoggerWithCallerSkip(loggerSkip, fileLoggerConf, writer)
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
	loggerForMiddleware, loggerClosers, err := s.loadingLoggerWithCallerSkip(middlewareSkip, fileLoggerConf, writer)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// for helper
	helperSkip := logpkg.CallerSkipForHelper
	loggerForHelper, loggerClosers, err := s.loadingLoggerWithCallerSkip(helperSkip, fileLoggerConf, writer)
	if err != nil {
		return err
	}
	for i := range loggerClosers {
		cleanup.cs = append(cleanup.cs, loggerClosers[i])
	}

	// prefix
	prefixKvs := s.withLoggerPrefix()
	logger = log.With(logger, prefixKvs...)
	loggerForMiddleware = log.With(loggerForMiddleware, prefixKvs...)
	loggerForHelper = log.With(loggerForHelper, prefixKvs...)

	// logger
	s.logger = logger
	s.loggerForMiddleware = loggerForMiddleware
	s.loggerForHelper = loggerForHelper

	// for gorm
	gormLoggerConf := s.conf.GetGorm()
	if gormLoggerConf.GetEnable() {
		gormLoggerWriter, err := s.GetWriterForGORM()
		if err != nil {
			return err
		}
		loggerForGORM, loggerClosers, err := s.loadingLoggerWithCallerSkip(helperSkip, gormLoggerConf, gormLoggerWriter)
		if err != nil {
			return err
		}
		for i := range loggerClosers {
			cleanup.cs = append(cleanup.cs, loggerClosers[i])
		}
		loggerForGORM = log.With(loggerForGORM, prefixKvs...)
		s.loggerForGORM = loggerForGORM
	} else {
		s.loggerForGORM = logger
	}

	// for rabbitmq
	rabbitmqLoggerConf := s.conf.GetRabbitmq()
	if rabbitmqLoggerConf.GetEnable() {
		rabbitmqLoggerWriter, err := s.GetWriterForRabbitmq()
		if err != nil {
			return err
		}
		loggerForRabbitmq, loggerClosers, err := s.loadingLoggerWithCallerSkip(helperSkip, rabbitmqLoggerConf, rabbitmqLoggerWriter)
		if err != nil {
			return err
		}
		for i := range loggerClosers {
			cleanup.cs = append(cleanup.cs, loggerClosers[i])
		}
		loggerForRabbitmq = log.With(loggerForRabbitmq, prefixKvs...)
		s.loggerForRabbitmq = loggerForRabbitmq
	} else {
		s.loggerForRabbitmq = logger
	}

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

func (s *loggerManager) loadingLoggerWithCallerSkip(skip int, fileLoggerConf *configpb.Log_File, writer io.Writer) (logger log.Logger, closeFnSlice []io.Closer, err error) {
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
	if fileLoggerConf != nil && fileLoggerConf.GetEnable() {
		// file logger
		loggerConfig := &logpkg.ConfigFile{
			Level:      logpkg.ParseLevel(fileLoggerConf.GetLevel()),
			CallerSkip: skip,

			Dir:      fileLoggerConf.GetDir(),
			Filename: fileLoggerConf.GetFilename(),

			RotateTime: fileLoggerConf.GetRotateTime().AsDuration(),
			RotateSize: fileLoggerConf.GetRotateSize(),

			StorageCounter: uint(fileLoggerConf.GetStorageCounter()),
			StorageAge:     fileLoggerConf.GetStorageAge().AsDuration(),
		}
		fileLogger, err := logpkg.NewFileLogger(loggerConfig, logpkg.WithWriter(writer))
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
