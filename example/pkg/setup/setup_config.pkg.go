package setuppkg

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
	configv1 "github.com/ikaiguang/go-srv-kit/example/api/config/v1"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
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
func newConfigHandler(setupOpts ...Option) (Config, error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// 启动选项
	setupOpt := &options{}
	for i := range setupOpts {
		setupOpts[i](setupOpt)
	}

	stdlog.Println("|==================== 加载配置文件 开始 ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== 加载配置文件 结束 ====================|")
	// 配置路径
	confPath := _configFilepath
	if setupOpt.configPath != "" {
		confPath = setupOpt.configPath
	} else if confPath == "" {
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

	// App配置
	if s.conf.App == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : app")
		return err
	}

	// 服务配置
	if s.conf.Server == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : server")
		return err
	}

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
	return apputil.ParseEnv(appEnv)
}

// IsEnvDebug 是否调试模式
func (s *configuration) IsEnvDebug(appEnv envv1.Env) bool {
	return apputil.IsDebugMode(appEnv)
}

// Watch 监听
func (s *configuration) Watch(key string, o config.Observer) error {
	return s.handler.Watch(key, o)
}

// Close 关闭
func (s *configuration) Close() error {
	return s.handler.Close()
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

// ServerConfig 服务配置
func (s *configuration) ServerConfig() *confv1.Server {
	return s.conf.Server
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

// ServerAuthConfig APP验证配置
func (s *configuration) BusinessAuthConfig() *confv1.Business_Auth {
	if s.conf.Business == nil {
		return nil
	}
	return s.conf.Business.Auth
}

// ServerSettingConfig APP设置配置
func (s *configuration) BaseSettingConfig() *confv1.Base_Setting {
	if s.conf.Base == nil {
		return nil
	}
	return s.conf.Base.Setting
}

// LoggerConfigForConsole 日志配置 控制台
func (s *configuration) LoggerConfigForConsole() *confv1.Log_Console {
	if s.conf.Log == nil {
		return nil
	}
	return s.conf.Log.Console
}

// LoggerConfigForFile 日志配置 文件
func (s *configuration) LoggerConfigForFile() *confv1.Log_File {
	if s.conf.Log == nil {
		return nil
	}
	return s.conf.Log.File
}

// DataConfig data配置
func (s *configuration) DataConfig() *confv1.Data {
	return s.conf.Data
}

// MySQLConfig mysql配置
func (s *configuration) MySQLConfig() *confv1.Data_MySQL {
	if s.conf.Data == nil {
		return nil
	}
	return s.conf.Data.Mysql
}

// PostgresConfig mysql配置
func (s *configuration) PostgresConfig() *confv1.Data_PSQL {
	if s.conf.Data == nil {
		return nil
	}
	return s.conf.Data.Psql
}

// RedisConfig redis配置
func (s *configuration) RedisConfig() *confv1.Data_Redis {
	if s.conf.Data == nil {
		return nil
	}
	return s.conf.Data.Redis
}

// ConsulConfig consul配置
func (s *configuration) ConsulConfig() *confv1.Base_Consul {
	if s.conf.Base == nil {
		return nil
	}
	return s.conf.Base.Consul
}

// JaegerTracerConfig jaeger tracer 配置
func (s *configuration) JaegerTracerConfig() *confv1.Base_JaegerTracer {
	if s.conf.Base == nil {
		return nil
	}
	return s.conf.Base.JaegerTracer
}

// SnowflakeWorkerConfig snowflake worker 配置
func (s *configuration) SnowflakeWorkerConfig() *confv1.Base_SnowflakeWorker {
	if s.conf.Base == nil {
		return nil
	}
	return s.conf.Base.SnowflakeWorker
}
