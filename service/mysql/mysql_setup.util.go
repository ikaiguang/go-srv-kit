package mysqlutil

import (
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"gorm.io/gorm"
)

// WithSetup 注册 MySQL 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentMysql, func() (MysqlManager, error) {
			loggerManager, err := ctx.GetLoggerComp().Get()
			if err != nil {
				return nil, err
			}
			return NewMysqlManager(ctx.GetConfig().GetMysql(), loggerManager)
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentMysql, func(name string) func() (MysqlManager, error) {
			return func() (MysqlManager, error) {
				mysqlConfig, ok := ctx.GetConfig().GetMysqlInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentMysql, name)
				}
				loggerManager, err := ctx.GetLoggerComp().Get()
				if err != nil {
					return nil, err
				}
				return NewMysqlManager(mysqlConfig, loggerManager)
			}
		}, ctx.GetLifecycle())
	})
}

// GetDB 从 LauncherManager 获取默认 MySQL 连接。
func GetDB(launcherManager setuputil.LauncherManager) (*gorm.DB, error) {
	mgr, err := setuputil.GetComponentValue[MysqlManager](launcherManager, setuputil.ComponentMysql)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// GetNamedDB 从 LauncherManager 获取命名 MySQL 连接。
func GetNamedDB(launcherManager setuputil.LauncherManager, name string) (*gorm.DB, error) {
	mgr, err := setuputil.GetNamedComponentValue[MysqlManager](launcherManager, setuputil.ComponentMysql, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}
