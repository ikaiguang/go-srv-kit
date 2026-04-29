package rabbitmqutil

import (
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// WithSetup 注册 RabbitMQ 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentRabbitmq, func() (RabbitmqManager, error) {
			loggerManager, err := ctx.GetLoggerComp().Get()
			if err != nil {
				return nil, err
			}
			return NewRabbitmqManager(ctx.GetConfig().GetRabbitmq(), loggerManager)
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentRabbitmq, func(name string) func() (RabbitmqManager, error) {
			return func() (RabbitmqManager, error) {
				rabbitmqConfig, ok := ctx.GetConfig().GetRabbitmqInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentRabbitmq, name)
				}
				loggerManager, err := ctx.GetLoggerComp().Get()
				if err != nil {
					return nil, err
				}
				return NewRabbitmqManager(rabbitmqConfig, loggerManager)
			}
		}, ctx.GetLifecycle())
	})
}

// GetConn 从 LauncherManager 获取默认 RabbitMQ 连接。
func GetConn(launcherManager setuputil.LauncherManager) (*rabbitmqpkg.ConnectionWrapper, error) {
	mgr, err := setuputil.GetComponentValue[RabbitmqManager](launcherManager, setuputil.ComponentRabbitmq)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedConn 从 LauncherManager 获取命名 RabbitMQ 连接。
func GetNamedConn(launcherManager setuputil.LauncherManager, name string) (*rabbitmqpkg.ConnectionWrapper, error) {
	mgr, err := setuputil.GetNamedComponentValue[RabbitmqManager](launcherManager, setuputil.ComponentRabbitmq, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}
