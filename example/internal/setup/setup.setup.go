package setup

import (
	"flag"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
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
func Setup() (packages Packages, err error) {
	// parses the command-line flags
	flag.Parse()

	// 开始配置
	stdlog.Println("|==================== 配置程序 开始 ====================|")
	defer stdlog.Println("|==================== 配置程序 结束 ====================|")

	// 初始化配置手柄
	configHandler, err := newConfigHandler()
	if err != nil {
		return packages, err
	}

	// 启动手柄
	upHandler := newUpHandler(configHandler)

	// 设置调试工具
	if err = upHandler.setupDebugUtil(); err != nil {
		return packages, err
	}

	// 设置日志工具
	if err = upHandler.setupLogUtil(); err != nil {
		return packages, err
	}

	// mysql gorm 数据库
	if _, err = upHandler.MysqlGormDB(); err != nil {
		return packages, err
	}

	// redis 客户端
	if _, err = upHandler.RedisClient(); err != nil {
		return packages, err
	}

	return upHandler, err
}

// newConfigHandler 初始化配置手柄
func newConfigHandler() (Config, error) {
	// 配置路径
	confPath := _configFilepath
	if confPath == "" {
		confPath = _defaultConfigFilepath
	}
	stdlog.Printf("|*** 配置文件路径：%s\n", confPath)

	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	return NewConfiguration(opts...)
}
