package consulutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	consulapi "github.com/hashicorp/consul/api"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	"google.golang.org/protobuf/proto"
)

// NewConfigLoader 创建 Consul 配置加载器。
func NewConfigLoader() configutil.ConsulConfigLoader {
	return func(consulConfig *configpb.Consul, appConfig *configpb.App, loadingOpts ...configutil.Option) (*configpb.Bootstrap, error) {
		consulClient, err := newConfigConsulClient(consulConfig)
		if err != nil {
			return nil, err
		}
		return LoadingConfigFromConsul(consulClient, appConfig, loadingOpts...)
	}
}

// LoadingConfigFromConsul 从 Consul 中加载配置。
func LoadingConfigFromConsul(consulClient *consulapi.Client, appConfig *configpb.App, loadingOpts ...configutil.Option) (*configpb.Bootstrap, error) {
	bootstrap := &configpb.Bootstrap{}
	loadOpts := configutil.NewLoadOptions(loadingOpts...)

	generalPath := appConfig.GetConfigPathForGeneral()
	if generalPath != "" {
		conf, err := loadingConfigFromConsul(consulClient, generalPath, loadOpts.Configs...)
		if err != nil {
			return nil, err
		}
		bootstrap = conf
	} else {
		stdlog.Println("|*** INFO: no general configuration path configured")
	}

	serverPath := appConfig.GetConfigPathForServer()
	if serverPath != "" {
		newOtherConfigs := make([]proto.Message, 0, len(loadOpts.Configs))
		for i := range loadOpts.Configs {
			newOtherConfigs = append(newOtherConfigs, proto.Clone(loadOpts.Configs[i]))
		}
		conf, err := loadingConfigFromConsul(consulClient, serverPath, newOtherConfigs...)
		if err != nil {
			return nil, err
		}
		configutil.MergeConfig(bootstrap, conf)
		for i := range newOtherConfigs {
			configutil.MergeConfig(loadOpts.Configs[i], newOtherConfigs[i])
		}
	} else {
		stdlog.Println("|*** INFO: this service configuration path is not configured")
	}

	return bootstrap, nil
}

func loadingConfigFromConsul(consulClient *consulapi.Client, consulConfigPath string, otherConfigs ...proto.Message) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING CONSUL CONFIGURATION : START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING CONSUL CONFIGURATION : END ====================|")
	stdlog.Println("|*** LOADING: consul configuration path: ", consulConfigPath)

	cs, err := consul.New(consulClient, consul.WithPath(consulConfigPath))
	if err != nil {
		e := errorpkg.ErrorInternalError("%s", err.Error())
		return nil, errorpkg.WithStack(e)
	}
	kvs, err := cs.Load()
	if err != nil {
		e := errorpkg.ErrorInternalError("%s", err.Error())
		return nil, errorpkg.WithStack(e)
	}
	if len(kvs) == 0 {
		e := errorpkg.ErrorRecordNotFound("consul configuration not found; path: %s", consulConfigPath)
		return nil, errorpkg.WithStack(e)
	}

	handler := config.New(config.WithSource(cs))
	defer func() {
		stdlog.Println("|*** LOADING: COMPLETE : consul configuration path: ", consulConfigPath)
		_ = handler.Close()
	}()

	if err = handler.Load(); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError("%s", err.Error()))
		return nil, err
	}

	conf := &configpb.Bootstrap{}
	if err = handler.Scan(conf); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError("%s", err.Error()))
		return nil, err
	}
	for i := range otherConfigs {
		if otherConfigs[i] == nil {
			continue
		}
		if err = handler.Scan(otherConfigs[i]); err != nil {
			err = errorpkg.WithStack(errorpkg.ErrorInternalError("%s", err.Error()))
			return nil, err
		}
	}
	return conf, nil
}

func newConfigConsulClient(cfg *configpb.Consul) (*consulapi.Client, error) {
	cc, err := consulpkg.NewConsulClient(ToConsulConfig(cfg))
	if err != nil {
		e := errorpkg.ErrorInternalError("%s", err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return cc, nil
}
