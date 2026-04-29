package redisutil

import (
	"github.com/redis/go-redis/v9"

	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// WithSetup 注册 Redis 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentRedis, func() (RedisManager, error) {
			return NewRedisManager(ctx.GetConfig().GetRedis())
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentRedis, func(name string) func() (RedisManager, error) {
			return func() (RedisManager, error) {
				redisConfig, ok := ctx.GetConfig().GetRedisInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentRedis, name)
				}
				return NewRedisManager(redisConfig)
			}
		}, ctx.GetLifecycle())
	})
}

// GetClient 从 LauncherManager 获取默认 Redis 客户端。
func GetClient(launcherManager setuputil.LauncherManager) (redis.UniversalClient, error) {
	mgr, err := setuputil.GetComponentValue[RedisManager](launcherManager, setuputil.ComponentRedis)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedClient 从 LauncherManager 获取命名 Redis 客户端。
func GetNamedClient(launcherManager setuputil.LauncherManager, name string) (redis.UniversalClient, error) {
	mgr, err := setuputil.GetNamedComponentValue[RedisManager](launcherManager, setuputil.ComponentRedis, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}
