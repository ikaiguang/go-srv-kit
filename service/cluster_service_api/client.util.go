package clientutil

import (
	"context"
	"net/url"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	connectionpkg "github.com/ikaiguang/go-srv-kit/kit/connection"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"google.golang.org/grpc/resolver"
)

type serviceAPIManager struct {
	opt         *option
	log         *log.Helper
	configMap   map[ServiceName]*Config
	configMutex sync.RWMutex
}

func NewServiceAPIManager(apiConfigs []*Config, opts ...Option) (ServiceAPIManager, error) {
	o := &option{}
	o.logger, _ = logpkg.NewDummyLogger()
	for i := range opts {
		opts[i](o)
	}
	logHelper := log.NewHelper(log.With(o.logger, "module", "client/util/ServiceAPIManager"))
	manager := &serviceAPIManager{
		opt:         o,
		log:         logHelper,
		configMap:   make(map[ServiceName]*Config),
		configMutex: sync.RWMutex{},
	}
	err := manager.RegisterServiceAPIConfigs(apiConfigs)
	if err != nil {
		return nil, err
	}
	return manager, nil
}

// RegisterServiceAPIConfigs 注册服务API，覆盖已有服务
func (s *serviceAPIManager) RegisterServiceAPIConfigs(apiConfigs []*Config, opts ...Option) error {
	for i := range opts {
		opts[i](s.opt)
	}

	s.configMutex.Lock()
	defer s.configMutex.Unlock()

	var (
		hasConsulRegistry, hasEtcdRegistry bool
	)
	for i := range apiConfigs {
		s.configMap[ServiceName(apiConfigs[i].ServiceName)] = apiConfigs[i]
		if apiConfigs[i].IsConsulRegistry() {
			hasConsulRegistry = true
		} else if apiConfigs[i].IsEtcdRegistry() {
			hasEtcdRegistry = true
		}
	}
	if hasConsulRegistry && s.opt.discoveryFactory[configpb.RegistryTypeEnum_CONSUL] == nil {
		return errorpkg.WithStack(uninitializedConsulDiscoveryError)
	}
	if hasEtcdRegistry && s.opt.discoveryFactory[configpb.RegistryTypeEnum_ETCD] == nil {
		return errorpkg.WithStack(uninitializedEtcdDiscoveryError)
	}
	return nil
}

func (s *serviceAPIManager) NewAPIConnection(serviceName ServiceName) (ServiceAPIConnection, error) {
	apiConfig, err := s.GetServiceAPIConfig(serviceName)
	if err != nil {
		return nil, err
	}
	conn := &clientConnection{}
	conn.SetTransportType(apiConfig.TransportType)
	switch apiConfig.TransportType {
	default:
		conn.httpClient, err = s.NewHTTPClient(apiConfig)
		if err != nil {
			return nil, err
		}
	case configpb.TransportTypeEnum_HTTP:
		conn.httpClient, err = s.NewHTTPClient(apiConfig)
		if err != nil {
			return nil, err
		}
	case configpb.TransportTypeEnum_GRPC:
		conn.grpcConn, err = s.NewGRPCConnection(apiConfig)
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}

func (s *serviceAPIManager) GetServiceAPIConfig(serviceName ServiceName) (*Config, error) {
	if serviceName.String() == "" {
		e := errorpkg.ErrorBadRequest("service name cannot be empty")
		return nil, errorpkg.WithStack(e)
	}
	s.configMutex.RLock()
	defer s.configMutex.RUnlock()
	conf, ok := s.configMap[serviceName]
	if !ok {
		e := errorpkg.ErrorRecordNotFound("service configuration not found; ServiceName: %s", serviceName.String())
		return nil, errorpkg.WithStack(e)
	}
	if conf == nil {
		e := errorpkg.ErrorInternalError("service configuration error: config == nil")
		return nil, errorpkg.WithStack(e)
	}
	return conf, nil
}

func (s *serviceAPIManager) checkGeneralEndpointValidity(serviceTarget string) error {
	ok, err := connectionpkg.CheckEndpointValidity(serviceTarget)
	if err != nil {
		e := errorpkg.ErrorServiceUnavailable("endpoint validity check failed; serviceTarget: %s", serviceTarget)
		if s.opt.skipRegistryCheck {
			logpkg.Warnw("skip endpoint check", "serviceTarget", serviceTarget, "err", e)
			return nil
		}
		return errorpkg.Wrap(e, err)
	}
	if !ok {
		e := errorpkg.ErrorServiceUnavailable("checkGeneralEndpointValidity is not ok; serviceTarget: %s", serviceTarget)
		if s.opt.skipRegistryCheck {
			logpkg.Warnw("skip endpoint check", "serviceTarget", serviceTarget, "err", e)
			return nil
		}
		return errorpkg.WithStack(e)
	}
	return nil
}

func (s *serviceAPIManager) getAndCheckRegistryDiscovery(apiConfig *Config, serviceTarget string) (registry.Discovery, error) {
	r, err := s.getRegistryDiscovery(apiConfig)
	if err != nil {
		return nil, err
	}

	// 跳过健康检查
	//if s.opt.skipRegistryCheck {
	//	return r, nil
	//}

	// target
	u, err := url.Parse(serviceTarget)
	if err != nil {
		e := errorpkg.ErrorInternalServer("failed to parse service target; serviceTarget: %s", serviceTarget)
		if s.opt.skipRegistryCheck {
			logpkg.Warnw("skip registry check", "serviceTarget", serviceTarget, "err", e)
			return r, nil
		}
		return nil, errorpkg.Wrap(e, err)
	}
	target := resolver.Target{URL: *u}
	// 健康检查
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = r.GetService(ctx, target.Endpoint())
	if err != nil {
		e := errorpkg.ErrorServiceUnavailable("registry discovery failed; serviceTarget: %s", serviceTarget)
		if s.opt.skipRegistryCheck {
			logpkg.Warnw("skip registry check", "serviceTarget", serviceTarget, "err", e)
			return r, nil
		}
		return nil, errorpkg.Wrap(e, err)
	}
	return r, nil
}

func (s *serviceAPIManager) getRegistryDiscovery(apiConfig *Config) (registry.Discovery, error) {
	switch apiConfig.RegistryType {
	default:
		e := errorpkg.ErrorUnimplemented("unsupported registry type")
		return nil, errorpkg.WithStack(e)
	case configpb.RegistryTypeEnum_CONSUL:
		return s.getRegistryDiscoveryByFactory(configpb.RegistryTypeEnum_CONSUL)
	case configpb.RegistryTypeEnum_ETCD:
		return s.getRegistryDiscoveryByFactory(configpb.RegistryTypeEnum_ETCD)
	}
}

func (s *serviceAPIManager) getRegistryDiscoveryByFactory(registryType configpb.RegistryTypeEnum_RegistryType) (registry.Discovery, error) {
	factory := s.opt.discoveryFactory[registryType]
	if factory == nil {
		return nil, errorpkg.WithStack(uninitializedDiscoveryFactoryError(registryType.String()))
	}
	r, err := factory()
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return nil, errorpkg.Wrap(e, err)
	}
	return r, nil
}
