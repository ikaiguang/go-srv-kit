package setup

import (
	"flag"
	stdlog "log"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	pkgerrors "github.com/pkg/errors"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
	configv1 "github.com/ikaiguang/go-srv-kit/example/api/config/v1"
)

const (
	_defaultConfigFilepath = "./configs"
)

var (
	// _configFilepath 配置文件 所在的目录
	_configFilepath string
)

func init() {
	flag.StringVar(&_configFilepath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// newConfigHandler 初始化配置手柄
func newConfigHandler() (Config, error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	stdlog.Println("|==================== 加载配置文件 开始 ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== 加载配置文件 结束 ====================|")
	// 配置路径
	confPath := _configFilepath
	if confPath == "" {
		confPath = _defaultConfigFilepath
	}
	log.Infof("配置文件路径: %s\n", confPath)

	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	return NewConfiguration(opts...)
}

// configuration 实现ConfigInterface
type configuration struct {
	// handler 配置处理手柄
	handler config.Config
	// conf 配置引导文件
	conf *configv1.Bootstrap

	// env app环境
	env envv1.Env

	// enableDebug 是否启用 调试模式
	enableDebug bool
	// enableLoggingConsole 是否启用 日志输出到控制台
	enableLoggingConsole bool
	// enableLoggingFile 是否启用 日志输出到文件
	enableLoggingFile bool
}

// NewConfiguration 配置处理手柄
func NewConfiguration(opts ...config.Option) (Config, error) {
	handler := &configuration{}
	if err := handler.New(opts...); err != nil {
		return nil, err
	}
	return handler, nil
}

// New 配置处理手柄
func (s *configuration) New(opts ...config.Option) (err error) {
	// 处理手柄
	s.handler = config.New(opts...)

	// 加载配置
	if err = s.handler.Load(); err != nil {
		err = pkgerrors.WithStack(err)
		return
	}

	// 读取配置文件
	s.conf = &configv1.Bootstrap{}
	if err = s.handler.Scan(s.conf); err != nil {
		err = pkgerrors.WithStack(err)
		return
	}

	// 初始化
	s.initialization()

	return
}

// initialization 初始化
func (s *configuration) initialization() {
	// app环境
	s.env = envv1.Env_PRODUCTION
	if s.conf.App != nil {
		// app环境
		s.env = s.ParseEnv(s.conf.App.Env)
		// enableDebug 是否启用 调试模式
		s.enableDebug = s.IsEnvDebug(s.env)
	}

	// 日志
	if s.conf.Log != nil {
		// // enableLogConsole 是否启用 日志输出到文件
		if s.conf.Log.Console != nil {
			s.enableLoggingConsole = s.conf.Log.Console.Enable
		}
		// enableLogFile 是否启用 日志输出到文件
		if s.conf.Log.File != nil {
			s.enableLoggingFile = s.conf.Log.File.Enable
		}
	}
}

// ParseEnv 解析环境
func (s *configuration) ParseEnv(appEnv string) envv1.Env {
	switch strings.ToUpper(appEnv) {
	case envv1.Env_DEVELOP.String():
		return envv1.Env_DEVELOP
	case envv1.Env_TESTING.String():
		return envv1.Env_TESTING
	case envv1.Env_PREVIEW.String():
		return envv1.Env_PREVIEW
	case envv1.Env_PRODUCTION.String():
		return envv1.Env_PRODUCTION
	default:
		return envv1.Env_PRODUCTION
	}
}

// IsEnvDebug 是否调试模式
func (s *configuration) IsEnvDebug(appEnv envv1.Env) bool {
	switch appEnv {
	case envv1.Env_DEVELOP, envv1.Env_TESTING:
		return true
	default:
		return false
	}
}

// Env app环境
func (s *configuration) Env() envv1.Env {
	return s.env
}

// IsDebugMode 是否启用 调试模式
func (s *configuration) IsDebugMode() bool {
	return s.enableDebug
}

// EnableLoggingConsole 是否启用 日志输出到控制台
func (s *configuration) EnableLoggingConsole() bool {
	return s.enableLoggingConsole
}

// EnableLoggingFile 是否启用 日志输出到文件
func (s *configuration) EnableLoggingFile() bool {
	return s.enableLoggingFile
}

// AppConfig APP配置
func (s *configuration) AppConfig() *confv1.App {
	return s.conf.App
}

// LoggerConfigForConsole 日志配置 控制台
func (s *configuration) LoggerConfigForConsole() *confv1.Log_Console {
	if s.conf.Log.Console == nil {
		return nil
	}
	return s.conf.Log.Console
}

// LoggerConfigForFile 日志配置 文件
func (s *configuration) LoggerConfigForFile() *confv1.Log_File {
	if s.conf.Log.File == nil {
		return nil
	}
	return s.conf.Log.File
}

// MySQLConfig mysql配置
func (s *configuration) MySQLConfig() *confv1.Data_MySQL {
	if s.conf.Data == nil {
		return nil
	}
	return s.conf.Data.Mysql
}

// RedisConfig redis配置
func (s *configuration) RedisConfig() *confv1.Data_Redis {
	if s.conf.Data == nil {
		return nil
	}
	return s.conf.Data.Redis
}

// HTTPConfig http配置
func (s *configuration) HTTPConfig() *confv1.Server_HTTP {
	if s.conf.Server == nil {
		return nil
	}
	return s.conf.Server.Http
}

// GRPCConfig grpc配置
func (s *configuration) GRPCConfig() *confv1.Server_GRPC {
	if s.conf.Server == nil {
		return nil
	}
	return s.conf.Server.Grpc
}
