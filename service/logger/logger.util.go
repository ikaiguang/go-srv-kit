package loggerutil

import (
	stderrors "errors"
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"io"
	stdlog "log"
	"sync"
)

type loggerManager struct {
	appConfig *configpb.App
	conf      *configpb.Log

	// 不要直接使用 s.writer, 请使用 GetWriter()
	writer                io.Writer
	writerOnce            sync.Once
	writerForGORM         io.Writer
	writerForGORMOnce     sync.Once
	writerForRabbitmq     io.Writer
	writerForRabbitmqOnce sync.Once

	// 不要直接使用 s.loggerXxx, 请使用 GetLoggers()
	loggerOnce          sync.Once
	logger              log.Logger
	loggerForMiddleware log.Logger
	loggerForHelper     log.Logger
	loggerForGORM       log.Logger
	loggerForRabbitmq   log.Logger
	loggerCloser        io.Closer
}

type Loggers struct {
	Logger              log.Logger
	LoggerForMiddleware log.Logger
	LoggerForHelper     log.Logger
	LoggerForGORM       log.Logger
	LoggerForRabbitmq   log.Logger
}

type LoggerManager interface {
	EnableConsole() bool
	EnableFile() bool
	EnableGORM() bool
	EnableRabbitmq() bool

	GetWriter() (io.Writer, error)
	GetWriterForGORM() (io.Writer, error)
	GetWriterForRabbitmq() (io.Writer, error)

	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)
	GetLoggerForGORM() (log.Logger, error)
	GetLoggerForRabbitmq() (log.Logger, error)

	Close() error
}

func NewLoggerManager(conf *configpb.Log, appConfig *configpb.App) (LoggerManager, error) {
	if appConfig == nil {
		stdlog.Println("[CONFIGURATION] Configuration not found, key = app; Use default configuration")
		//e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = app")
		//e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = app")
		//return nil, errorpkg.WithStack(e)
		appConfig = &_defaultAppConfig
	}
	if conf == nil {
		stdlog.Println("[CONFIGURATION] Configuration not found, key = log; Use default configuration")
		//e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = log")
		//return nil, errorpkg.WithStack(e)
		conf = _defaultLogConfig
	}
	manager := &loggerManager{
		appConfig: appConfig,
		conf:      conf,
	}
	_, err := manager.GetLogger()
	if err != nil {
		return nil, err
	}
	return manager, nil
}

func (s *loggerManager) EnableConsole() bool {
	return s.conf.GetConsole().GetEnable()
}

func (s *loggerManager) EnableFile() bool {
	return s.conf.GetFile().GetEnable()
}

func (s *loggerManager) EnableGORM() bool {
	return s.conf.GetGorm().GetEnable()
}

func (s *loggerManager) EnableRabbitmq() bool {
	return s.conf.GetRabbitmq().GetEnable()
}

func (s *loggerManager) Close() error {
	var errs []error

	// loggers
	if s.loggerCloser != nil {
		stdlog.Println("|*** STOP: close: Logger")
		if err := s.loggerCloser.Close(); err != nil {
			stdlog.Println("|*** STOP: close: Logger failed: ", err.Error())
			errs = append(errs, err)
		}
	}

	// writer
	if s.writer != nil {
		if writerCloser, ok := s.writer.(io.Closer); ok {
			stdlog.Println("|*** STOP: close: Writer")
			if err := writerCloser.Close(); err != nil {
				stdlog.Println("|*** STOP: close: Writer failed: ", err.Error())
				errs = append(errs, err)
			}
		}
	}
	if s.writerForGORM != nil {
		if writerCloser, ok := s.writerForGORM.(io.Closer); ok {
			stdlog.Println("|*** STOP: close: gorm Writer")
			if err := writerCloser.Close(); err != nil {
				stdlog.Println("|*** STOP: close: gorm Writer failed: ", err.Error())
				errs = append(errs, err)
			}
		}
	}
	if s.writerForRabbitmq != nil {
		if writerCloser, ok := s.writer.(io.Closer); ok {
			stdlog.Println("|*** STOP: close: rabbitmq Writer")
			if err := writerCloser.Close(); err != nil {
				stdlog.Println("|*** STOP: close: rabbitmq Writer failed: ", err.Error())
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}

type closer struct {
	cs []io.Closer
}

func (c *closer) Close() error {
	var errs []error
	for _, v := range c.cs {
		if err := v.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return stderrors.Join(errs...)
	}
	return nil
}
