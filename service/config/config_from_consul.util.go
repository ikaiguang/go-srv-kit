package configutil

import (
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	consulapi "github.com/hashicorp/consul/api"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/proto"
	stdlog "log"
)

// LoadingConfigFromConsul 从consul中加载配置
// 首先：读取服务base配置
// 然后：读取本服务配置
// 最后：使用本服务配置 覆盖 base 配置
func LoadingConfigFromConsul(consulClient *consulapi.Client, appConfig *configpb.App, loadingOpts ...Option) (*configpb.Bootstrap, error) {
	var bootstrap = &configpb.Bootstrap{}
	loadOpts := &options{}
	for i := range loadingOpts {
		loadingOpts[i](loadOpts)
	}

	// 通用配置
	generalPath := appConfig.GetConfigPathForGeneral()
	if generalPath != "" {
		conf, err := loadingConfigFromConsul(consulClient, generalPath, loadOpts.configs...)
		if err != nil {
			return nil, err
		}
		bootstrap = conf
	} else {
		stdlog.Println("|*** INFO: no general configuration path configured")
	}

	// 服务配置 合并与覆盖
	serverPath := appConfig.GetConfigPathForServer()
	if serverPath != "" {
		// new configs
		newOtherConfigs := make([]proto.Message, 0, len(loadOpts.configs))
		for i := range loadOpts.configs {
			newOtherConfigs = append(newOtherConfigs, proto.Clone(loadOpts.configs[i]))
		}
		// scan
		conf, err := loadingConfigFromConsul(consulClient, serverPath, newOtherConfigs...)
		if err != nil {
			return nil, err
		}
		// merge
		MergeConfig(bootstrap, conf)
		for i := range newOtherConfigs {
			MergeConfig(loadOpts.configs[i], newOtherConfigs[i])
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

	// 配置source
	cs, err := consul.New(consulClient, consul.WithPath(consulConfigPath))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	kvs, err := cs.Load()
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	if len(kvs) == 0 {
		e := errorpkg.ErrorRecordNotFound("consul configuration not found; path: %s", consulConfigPath)
		return nil, errorpkg.WithStack(e)
	}

	// options
	var opts []config.Option
	opts = append(opts, config.WithSource(cs))

	handler := config.New(opts...)
	defer func() {
		stdlog.Println("|*** LOADING: COMPLETE : consul configuration path: ", consulConfigPath)
		_ = handler.Close()
	}()

	// 加载配置
	if err = handler.Load(); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}

	// 读取配置文件
	conf := &configpb.Bootstrap{}
	if err = handler.Scan(conf); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}
	for i := range otherConfigs {
		if otherConfigs[i] == nil {
			continue
		}
		if err = handler.Scan(otherConfigs[i]); err != nil {
			err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
			return nil, err
		}
	}
	return conf, nil
}

func newConsulClient(cfg *configpb.Consul) (*consulapi.Client, error) {
	cc, err := consulpkg.NewConsulClient(ToConsulConfig(cfg))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return cc, nil
}

// ToConsulConfig ...
func ToConsulConfig(cfg *configpb.Consul) *consulpkg.Config {
	return &consulpkg.Config{
		Scheme:             cfg.Scheme,
		Address:            cfg.Address,
		PathPrefix:         cfg.PathPrefix,
		Datacenter:         cfg.Datacenter,
		WaitTime:           cfg.WaitTime,
		Token:              cfg.Token,
		Namespace:          cfg.Namespace,
		Partition:          cfg.Partition,
		WithHttpBasicAuth:  cfg.WithHttpBasicAuth,
		AuthUsername:       cfg.AuthUsername,
		AuthPassword:       cfg.AuthPassword,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		TlsAddress:         cfg.TlsAddress,
		TlsCaPem:           cfg.TlsCaPem,
		TlsCertPem:         cfg.TlsCertPem,
		TlsKeyPem:          cfg.TlsKeyPem,
	}
}
