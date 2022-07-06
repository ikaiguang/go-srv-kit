package setup

import (
	"flag"
	stdlog "log"
)

// Setup 启动与配置
func Setup(opts ...Option) (modulesHandler Modules, err error) {
	// parses the command-line flags
	flag.Parse()

	// 初始化配置手柄
	configHandler, err := newConfigHandler(opts...)
	if err != nil {
		return modulesHandler, err
	}

	// 开始配置
	stdlog.Println("|==================== 配置程序 开始 ====================|")
	defer stdlog.Println("|==================== 配置程序 结束 ====================|")

	// 启动手柄
	setupHandler := NewModules(configHandler)

	// 设置调试工具
	if err = setupHandler.loadingDebugUtil(); err != nil {
		return modulesHandler, err
	}

	// 设置日志工具
	if _, err = setupHandler.loadingLogHelper(); err != nil {
		return modulesHandler, err
	}

	// mysql gorm 数据库
	//if _, err = setupHandler.GetMySQLGormDB(); err != nil {
	//	return modulesHandler, err
	//}

	// postgres gorm 数据库
	//if _, err = setupHandler.GetPostgresGormDB(); err != nil {
	//	return modulesHandler, err
	//}

	// redis 客户端
	//if _, err = setupHandler.GetRedisClient(); err != nil {
	//	return modulesHandler, err
	//}
	return setupHandler, err
}
