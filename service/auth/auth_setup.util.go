package authutil

import (
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// WithSetup 注册认证组件；需要同时注册 redisutil.WithSetup()。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentAuth, func() (AuthInstance, error) {
			loggerManager, err := ctx.GetLoggerComp().Get()
			if err != nil {
				return nil, err
			}
			redisClient, err := redisutil.GetClient(ctx)
			if err != nil {
				return nil, err
			}
			return NewAuthInstance(ctx.GetConfig().GetEncrypt().GetTokenEncrypt(), redisClient, loggerManager)
		}, ctx.GetLifecycle())
	})
}

// GetTokenManager 从 LauncherManager 获取 Token 管理器。
func GetTokenManager(launcherManager setuputil.LauncherManager) (authpkg.TokenManager, error) {
	authInstance, err := setuputil.GetComponentValue[AuthInstance](launcherManager, setuputil.ComponentAuth)
	if err != nil {
		return nil, err
	}
	return authInstance.GetTokenManger()
}

// GetAuthManager 从 LauncherManager 获取认证管理器。
func GetAuthManager(launcherManager setuputil.LauncherManager) (authpkg.AuthRepo, error) {
	authInstance, err := setuputil.GetComponentValue[AuthInstance](launcherManager, setuputil.ComponentAuth)
	if err != nil {
		return nil, err
	}
	return authInstance.GetAuthManger()
}
