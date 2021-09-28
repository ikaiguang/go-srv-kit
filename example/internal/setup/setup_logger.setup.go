package setuphandler

import (
	"github.com/go-kratos/kratos/v2/log"
	pkgerrors "github.com/pkg/errors"

	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// setupLogUtil 设置日志工具
func (s *setup) setupLogUtil() (err error) {
	if s.isInit {
		return pkgerrors.New("处理手柄未初始化")
	}
	// loggers
	var loggers []log.Logger

	// 日志 输出到控制台
	if s.enableLogConsole {
		stdLoggerConfig := &logutil.ConfigStd{
			Level:      logutil.ParseLevel(s.conf.Log.Console.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,
		}
		stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
		if err != nil {
			return err
		}
		loggers = append(loggers, stdLogger)
	}

	// 日志 输出到文件
	if s.enableLogFile {
		// file logger
		fileLoggerConfig := &logutil.ConfigFile{
			Level:      logutil.ParseLevel(s.conf.Log.File.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,

			Dir:      s.conf.Log.File.Dir,
			Filename: s.conf.Log.File.Filename + "_util",

			RotateTime: s.conf.Log.File.RotateTime.AsDuration(),
			RotateSize: s.conf.Log.File.RotateSize,

			StorageCounter: uint(s.conf.Log.File.StorageCounter),
			StorageAge:     s.conf.Log.File.StorageAge.AsDuration(),
		}
		fileLogger, err := logutil.NewFileLogger(fileLoggerConfig)
		if err != nil {
			panic(err)
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志 工具
	if len(loggers) > 0 {
		multiLogger := log.MultiLogger(loggers...)
		logutil.Setup(multiLogger)
	}
	return err
}
