package setup

import (
	"flag"
	stdlog "log"
	"sync"
)

var (
	// initEngineMutex 初始化
	initEngineMutex sync.Once
	engineInstance  Engine
)

// Init 启动与配置与设置存储Packages
func Init(opts ...Option) (err error) {
	initEngineMutex.Do(func() {
		engineInstance, err = Setup(opts...)
	})
	if err != nil {
		initEngineMutex = sync.Once{}
		return err
	}
	return err
}

// GetEngine 获取初始化后的引擎模块
func GetEngine() (Engine, error) {
	if err := Init(); err != nil {
		return nil, err
	}
	return engineInstance, nil
}

// Close .
func Close() error {
	if engineInstance == nil {
		return nil
	}
	return engineInstance.Close()
}

// Setup 启动与配置
func Setup(opts ...Option) (engineHandler Engine, err error) {
	// parses the command-line flags
	flag.Parse()

	// 初始化配置手柄
	configHandler, err := newConfigHandler(opts...)
	if err != nil {
		return engineHandler, err
	}

	// 开始配置
	stdlog.Println("|==================== 配置程序 开始 ====================|")
	defer stdlog.Println("|==================== 配置程序 结束 ====================|")

	// 启动手柄
	setupHandler := NewEngine(configHandler)

	// 设置调试工具
	if err = setupHandler.loadingDebugUtil(); err != nil {
		return engineHandler, err
	}

	// 设置日志工具
	if _, err = setupHandler.loadingLogHelper(); err != nil {
		return engineHandler, err
	}

	// mysql gorm 数据库
	//if _, err = setupHandler.GetMySQLGormDB(); err != nil {
	//	return engineHandler, err
	//}

	// postgres gorm 数据库
	//if _, err = setupHandler.GetPostgresGormDB(); err != nil {
	//	return engineHandler, err
	//}

	// redis 客户端
	//if _, err = setupHandler.GetRedisClient(); err != nil {
	//	return engineHandler, err
	//}
	return setupHandler, err
}
