package clientutil

import (
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// SetupOption 配置 ServiceAPI 组件注册。
type SetupOption func(*setupOptions)

type setupOptions struct {
	clientOptionsProvider []func(setuputil.LauncherManager) ([]Option, error)
}

// WithClientOptionsProvider 注入按 LauncherManager 生成的客户端选项。
func WithClientOptionsProvider(provider func(setuputil.LauncherManager) ([]Option, error)) SetupOption {
	return func(o *setupOptions) {
		o.clientOptionsProvider = append(o.clientOptionsProvider, provider)
	}
}

// WithSetup 注册集群服务 API 组件。
func WithSetup(opts ...SetupOption) setuputil.Option {
	setupOpts := &setupOptions{}
	for i := range opts {
		opts[i](setupOpts)
	}
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentServiceAPI, func() (ServiceAPIManager, error) {
			apiConfigs, _, err := ToConfig(ctx.GetConfig().GetClusterServiceApi())
			if err != nil {
				return nil, err
			}
			loggerForMiddleware, err := ctx.GetLoggerForMiddleware()
			if err != nil {
				return nil, err
			}
			clientOpts := []Option{
				WithLogger(loggerForMiddleware),
				WithSkipRegistryCheck(),
			}
			for i := range setupOpts.clientOptionsProvider {
				moreOpts, err := setupOpts.clientOptionsProvider[i](ctx)
				if err != nil {
					return nil, err
				}
				clientOpts = append(clientOpts, moreOpts...)
			}
			return NewServiceAPIManager(apiConfigs, clientOpts...)
		}, ctx.GetLifecycle())
	})
}

// GetManager 从 LauncherManager 获取集群服务 API 管理器。
func GetManager(launcherManager setuputil.LauncherManager) (ServiceAPIManager, error) {
	return setuputil.GetComponentValue[ServiceAPIManager](launcherManager, setuputil.ComponentServiceAPI)
}
