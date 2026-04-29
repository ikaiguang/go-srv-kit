package postgresutil

import (
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"gorm.io/gorm"
)

// WithSetup 注册 PostgreSQL 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentPostgres, func() (PostgresManager, error) {
			loggerManager, err := ctx.GetLoggerComp().Get()
			if err != nil {
				return nil, err
			}
			return NewPostgresManager(ctx.GetConfig().GetPsql(), loggerManager)
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentPostgres, func(name string) func() (PostgresManager, error) {
			return func() (PostgresManager, error) {
				psqlConfig, ok := ctx.GetConfig().GetPsqlInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentPostgres, name)
				}
				loggerManager, err := ctx.GetLoggerComp().Get()
				if err != nil {
					return nil, err
				}
				return NewPostgresManager(psqlConfig, loggerManager)
			}
		}, ctx.GetLifecycle())
	})
}

// GetDB 从 LauncherManager 获取默认 PostgreSQL 连接。
func GetDB(launcherManager setuputil.LauncherManager) (*gorm.DB, error) {
	mgr, err := setuputil.GetComponentValue[PostgresManager](launcherManager, setuputil.ComponentPostgres)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}

// GetNamedDB 从 LauncherManager 获取命名 PostgreSQL 连接。
func GetNamedDB(launcherManager setuputil.LauncherManager, name string) (*gorm.DB, error) {
	mgr, err := setuputil.GetNamedComponentValue[PostgresManager](launcherManager, setuputil.ComponentPostgres, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetDB()
}
