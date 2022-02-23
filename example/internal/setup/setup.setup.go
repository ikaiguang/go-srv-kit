package setup

import (
	"flag"
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	_defaultConfigFilepath = "./configs"
)

var (
	// _configFilepath 配置文件 所在的目录
	_configFilepath string

	// initModulesMutex 初始化
	initModulesMutex sync.Once
	modulesInstance  Modules
)

func init() {
	flag.StringVar(&_configFilepath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// GetModules 获取初始化后的模块
func GetModules() (Modules, error) {
	if err := Init(); err != nil {
		return nil, err
	}
	return modulesInstance, nil
}

// Init 启动与配置与设置存储Packages
func Init() (err error) {
	initModulesMutex.Do(func() {
		modulesInstance, err = Setup()
	})
	if err != nil {
		initModulesMutex = sync.Once{}
		return err
	}
	if modulesInstance != nil {
		return err
	}

	modulesInstance, err = Setup()
	if err != nil {
		return err
	}
	return err
}

// Close .
func Close() error {
	if modulesInstance == nil {
		return nil
	}
	return modulesInstance.Close()
}

// Setup 启动与配置
func Setup() (modulesHandler Modules, err error) {
	// parses the command-line flags
	flag.Parse()

	// 初始化配置手柄
	configHandler, err := newConfigHandler()
	if err != nil {
		return modulesHandler, err
	}

	// 开始配置
	stdlog.Println("|==================== 配置程序 开始 ====================|")
	defer stdlog.Println("|==================== 配置程序 结束 ====================|")

	// 启动手柄
	setupHandler := newModuleHandler(configHandler)

	// 设置调试工具
	if err = setupHandler.loadingDebugUtil(); err != nil {
		return modulesHandler, err
	}

	// 设置日志工具
	if _, err = setupHandler.loadingLogHelper(); err != nil {
		return modulesHandler, err
	}

	// mysql gorm 数据库
	if _, err = setupHandler.MysqlGormDB(); err != nil {
		return modulesHandler, err
	}

	// redis 客户端
	if _, err = setupHandler.RedisClient(); err != nil {
		return modulesHandler, err
	}
	return setupHandler, err
}

// newConfigHandler 初始化配置手柄
func newConfigHandler() (Config, error) {
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
