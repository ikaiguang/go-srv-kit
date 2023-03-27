package setuppkg

import (
	"flag"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
)

// New 启动与配置
func New(opts ...Option) (engineHandler Engine, err error) {
	// parses the command-line flags
	flag.Parse()

	// 初始化配置手柄
	configHandler, err := newConfigHandler(opts...)
	if err != nil {
		return engineHandler, pkgerrors.WithStack(err)
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
	if cfg := setupHandler.Config.MySQLConfig(); cfg != nil && cfg.Enable {
		if _, err = setupHandler.GetMySQLGormDB(); err != nil {
			return engineHandler, err
		}
	}

	// postgres gorm 数据库
	if cfg := setupHandler.Config.PostgresConfig(); cfg != nil && cfg.Enable {
		if _, err = setupHandler.GetPostgresGormDB(); err != nil {
			return engineHandler, err
		}
	}

	// redis 客户端
	if cfg := setupHandler.Config.RedisConfig(); cfg != nil && cfg.Enable {
		redisCC, err := setupHandler.GetRedisClient()
		if err != nil {
			return engineHandler, err
		}
		// 验证Token工具
		_ = setupHandler.GetAuthTokenRepo(redisCC)
	}

	// consul 客户端
	if cfg := setupHandler.Config.ConsulConfig(); cfg != nil && cfg.Enable {
		_, err = setupHandler.GetConsulClient()
		if err != nil {
			return engineHandler, err
		}
	}

	// jaeger trace
	if cfg := setupHandler.Config.JaegerTracerConfig(); cfg != nil && cfg.Enable {
		_, err = setupHandler.GetJaegerTraceExporter()
		if err != nil {
			return engineHandler, err
		}
	}

	// 雪花算法
	if cfg := setupHandler.Config.BaseSettingConfig(); cfg != nil && cfg.EnableSnowflakeWorker {
		err = setupHandler.loadingSnowflakeWorker()
		if err != nil {
			return engineHandler, err
		}
	}

	// 监听配置 app
	//if err = setupHandler.watchConfigApp(); err != nil {
	//	return engineHandler, err
	//}

	// 监听配置 data
	//if err = setupHandler.watchConfigData(); err != nil {
	//	return engineHandler, err
	//}

	return setupHandler, err
}
