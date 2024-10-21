package loggerutil

import (
	stderrors "errors"
	"github.com/go-kratos/kratos/v2/log"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"io"
	stdlog "log"
	"sync"
)

type loggerManager struct {
	appConfig *configpb.App
	conf      *configpb.Log

	// 不要直接使用 s.writer, 请使用 GetWriter()
	writer     io.Writer
	writerOnce sync.Once

	// 不要直接使用 s.loggerXxx, 请使用 GetLoggers()
	loggerOnce          sync.Once
	logger              log.Logger
	loggerForMiddleware log.Logger
	loggerForHelper     log.Logger
	loggerCloser        io.Closer
}

type Loggers struct {
	Logger              log.Logger
	LoggerForMiddleware log.Logger
	LoggerForHelper     log.Logger
}

type LoggerManager interface {
	EnableConsole() bool
	EnableFile() bool
	GetWriter() (io.Writer, error)
	GetLogger() (log.Logger, error)
	GetLoggerForMiddleware() (log.Logger, error)
	GetLoggerForHelper() (log.Logger, error)
	Close() error
}

func NewLoggerManager(conf *configpb.Log, appConfig *configpb.App) (LoggerManager, error) {
	if appConfig == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = app")
		return nil, errorpkg.WithStack(e)
	} else if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = log")
		return nil, errorpkg.WithStack(e)
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
