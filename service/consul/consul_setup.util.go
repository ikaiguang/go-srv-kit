package consulutil

import (
	consulapi "github.com/hashicorp/consul/api"

	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// WithSetup 注册 Consul 组件。
func WithSetup() setuputil.Option {
	return setuputil.WithComponentRegistrar(func(ctx setuputil.ComponentRegistrarContext) {
		setuputil.RegisterComponent(ctx.GetRegistry(), setuputil.ComponentConsul, func() (ConsulManager, error) {
			return NewConsulManager(ctx.GetConfig().GetConsul())
		}, ctx.GetLifecycle())
		setuputil.RegisterComponentGroup(ctx.GetRegistry(), setuputil.ComponentConsul, func(name string) func() (ConsulManager, error) {
			return func() (ConsulManager, error) {
				consulConfig, ok := ctx.GetConfig().GetConsulInstances()[name]
				if !ok {
					return nil, setuputil.ComponentNotFoundError(setuputil.ComponentConsul, name)
				}
				return NewConsulManager(consulConfig)
			}
		}, ctx.GetLifecycle())
	})
}

// GetClient 从 LauncherManager 获取默认 Consul 客户端。
func GetClient(launcherManager setuputil.LauncherManager) (*consulapi.Client, error) {
	mgr, err := setuputil.GetComponentValue[ConsulManager](launcherManager, setuputil.ComponentConsul)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}

// GetNamedClient 从 LauncherManager 获取命名 Consul 客户端。
func GetNamedClient(launcherManager setuputil.LauncherManager, name string) (*consulapi.Client, error) {
	mgr, err := setuputil.GetNamedComponentValue[ConsulManager](launcherManager, setuputil.ComponentConsul, name)
	if err != nil {
		return nil, err
	}
	return mgr.GetClient()
}
