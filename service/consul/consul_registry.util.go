package consulutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

// RegistryAppOptions 为 Kratos App 注入 Consul 服务注册。
func RegistryAppOptions(launcherManager setuputil.LauncherManager) ([]kratos.Option, error) {
	stdlog.Println("|*** LOADING: ServiceRegistry: consul")
	consulClient, err := GetClient(launcherManager)
	if err != nil {
		return nil, err
	}
	r, err := registrypkg.NewConsulRegistry(consulClient)
	if err != nil {
		return nil, err
	}
	return []kratos.Option{kratos.Registrar(r)}, nil
}

// WithRegistryAppOptions 将 Consul 服务注册接入 all-in-one 启动。
func WithRegistryAppOptions() serverutil.Option {
	return serverutil.WithAppOptionProvider(RegistryAppOptions)
}

// ServiceAPIDiscoveryOptions 为 cluster_service_api 注入 Consul discovery。
func ServiceAPIDiscoveryOptions(launcherManager setuputil.LauncherManager) ([]clientutil.Option, error) {
	consulClient, err := GetClient(launcherManager)
	if err != nil {
		return nil, err
	}
	return []clientutil.Option{
		clientutil.WithDiscoveryFactory(configpb.RegistryTypeEnum_CONSUL, func() (registry.Discovery, error) {
			return registrypkg.NewConsulRegistry(consulClient)
		}),
	}, nil
}
