package setuphandler

import (
	"flag"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
)

// Setup 启动与配置
func Setup() (err error) {
	// parses the command-line flags
	flag.Parse()

	// setup
	handler := &setup{}

	return handler.Setup()
}

// setup .
type setup struct {
	// isInit 是否 初始化了
	isInit bool
	// conf 配置引导文件
	conf *confv1.Bootstrap

	// enableDebug 是否启用 调试模式
	enableDebug bool
	// enableLogConsole 是否启用 日志输出到控制台
	enableLogConsole bool
	// enableLogFile 是否启用 日志输出到文件
	enableLogFile bool
}

// Setup 配置
func (s *setup) Setup() (err error) {
	// 配置手柄
	configHandler, err := s.getConfigHandler()
	if err != nil {
		return err
	}

	// 加载配置
	conf := &confv1.Bootstrap{}
	if err = configHandler.Scan(s.conf); err != nil {
		err = pkgerrors.WithStack(err)
		return
	}

	// 初始化
	s.Init(conf)

	// 调试工具
	if err = s.setupDebugUtil(); err != nil {
		return err
	}

	// 日志工具
	if err = s.setupLogUtil(); err != nil {
		return err
	}
	return err
}

// IsDebugMode 是否调试模式
func (s *setup) IsDebugMode() bool {
	return s.enableDebug
}

// Init 初始化
func (s *setup) Init(conf *confv1.Bootstrap) {
	// 初始化
	s.isInit = true
	// 配置
	s.conf = proto.Clone(conf).(*confv1.Bootstrap)
	// enableDebug 是否启用 调试模式
	if s.conf.App != nil {
		s.enableDebug = s.IsEnvDebug(s.conf.App.Env)
	}

	// 日志
	if s.conf.Log != nil {
		// // enableLogConsole 是否启用 日志输出到文件
		if s.conf.Log.Console != nil {
			s.enableLogConsole = s.conf.Log.Console.Enable
		}
		// enableLogFile 是否启用 日志输出到文件
		if s.conf.Log.File != nil {
			s.enableLogFile = s.conf.Log.File.Enable
		}
	}
}

// IsEnvDebug 是否调试模式
func (s *setup) IsEnvDebug(appEnv envv1.Env) bool {
	switch appEnv {
	case envv1.Env_DEVELOP, envv1.Env_TESTING:
		return true
	default:
		return false
	}
}
