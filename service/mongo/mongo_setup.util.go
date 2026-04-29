package mongoutil

import (
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// WithSetup 注册 MongoDB 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentMongo, func() (MongoManager, error) {
			loggerManager, err := ctx.GetLoggerComp().Get()
			if err != nil {
				return nil, err
			}
			return NewMongoManager(ctx.GetConfig().GetMongo(), loggerManager)
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentMongo, func(name string) func() (MongoManager, error) {
			return func() (MongoManager, error) {
				mongoConfig, ok := ctx.GetConfig().GetMongoInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentMongo, name)
				}
				loggerManager, err := ctx.GetLoggerComp().Get()
				if err != nil {
					return nil, err
				}
				return NewMongoManager(mongoConfig, loggerManager)
			}
		}, ctx.GetLifecycle())
	})
}

// GetClient 从 LauncherManager 获取默认 MongoDB 客户端。
func GetClient(launcherManager setuputil.LauncherManager) (*mongo.Client, error) {
	mgr, err := setuputil.GetComponentValue[MongoManager](launcherManager, setuputil.ComponentMongo)
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}

// GetNamedClient 从 LauncherManager 获取命名 MongoDB 客户端。
func GetNamedClient(launcherManager setuputil.LauncherManager, name string) (*mongo.Client, error) {
	mgr, err := setuputil.GetNamedComponentValue[MongoManager](launcherManager, setuputil.ComponentMongo, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetMongoClient()
}
