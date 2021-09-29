package setuphandler

import (
	"flag"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	"github.com/ikaiguang/go-srv-kit/log/helper"
	setuputil "github.com/ikaiguang/go-srv-kit/setup"
)

const (
	_defaultConfigFilepath = "./configs"
)

var (
	// 配置文件 所在的目录
	_configFilepath string
)

func init() {
	flag.StringVar(&_configFilepath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// Setup 启动与配置
func Setup() (err error) {
	// parses the command-line flags
	flag.Parse()

	// 启动手柄
	handler := &up{}
	if err = handler.initialization(); err != nil {
		return err
	}

	// 设置调试工具
	if err = handler.setupDebugUtil(); err != nil {
		return err
	}

	// 设置日志工具
	if err = handler.setupLogUtil(); err != nil {
		return err
	}
	return err
}

// up 启动手柄
type up struct {
	config setuputil.Config
}

// initialization 初始化
func (s *up) initialization() (err error) {
	s.config, err = s.getConfigHandler()
	if err != nil {
		return err
	}
	return err
}

// getConfigHandler 配置手柄
func (s *up) getConfigHandler() (setuputil.Config, error) {
	// 配置路径
	confPath := _configFilepath
	if confPath == "" {
		confPath = _defaultConfigFilepath
	}
	stdlog.Println("*** | 配置文件路径：", confPath)

	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	return setuputil.NewConfiguration(opts...)
}

// setupDebugUtil 设置调试工具
func (s *up) setupDebugUtil() error {
	stdlog.Println("*** | 加载调试工具：", s.config.IsDebugMode())
	if !s.config.IsDebugMode() {
		return nil
	}
	return debugutil.Setup()
}

// setupLogUtil 设置日志工具
func (s *up) setupLogUtil() (err error) {
	// loggers
	var loggers []log.Logger
	defer func() {
		if len(loggers) == 0 {
			stdlog.Println("*** | 未加载日志工具：", s.config.IsDebugMode())
		}
	}()

	// 配置
	loggerConfig := s.config.LoggerConfig()
	if loggerConfig == nil {
		return err
	}

	// 日志 输出到控制台
	if s.config.EnableLoggingConsole() && loggerConfig.Console != nil {
		stdlog.Println("*** | 加载日志工具：日志输出到控制台")
		stdLoggerConfig := &logutil.ConfigStd{
			Level:      logutil.ParseLevel(loggerConfig.Console.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,
		}
		stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
		if err != nil {
			return err
		}
		loggers = append(loggers, stdLogger)
	}

	// 日志 输出到文件
	if s.config.EnableLoggingFile() && loggerConfig.File != nil {
		stdlog.Println("*** | 加载日志工具：日志输出到文件")
		// file logger
		fileLoggerConfig := &logutil.ConfigFile{
			Level:      logutil.ParseLevel(loggerConfig.File.Level),
			CallerSkip: logutil.DefaultCallerSkip + 2,

			Dir:      loggerConfig.File.Dir,
			Filename: loggerConfig.File.Filename + "_util",

			RotateTime: loggerConfig.File.RotateTime.AsDuration(),
			RotateSize: loggerConfig.File.RotateSize,

			StorageCounter: uint(loggerConfig.File.StorageCounter),
			StorageAge:     loggerConfig.File.StorageAge.AsDuration(),
		}
		fileLogger, err := logutil.NewFileLogger(fileLoggerConfig)
		if err != nil {
			panic(err)
		}
		loggers = append(loggers, fileLogger)
	}

	// 日志工具
	if len(loggers) == 0 {
		return err
	}
	multiLogger := log.MultiLogger(loggers...)
	loghelper.Setup(multiLogger)
	return err
}
