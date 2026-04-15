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
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
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
	if hasConsulRegistry && s.opt.consulClient == nil {
		return errorpkg.WithStack(uninitializedConsulClientError)
	}
	if hasEtcdRegistry && s.opt.etcdClient == nil {
		return errorpkg.WithStack(uninitializedEtcdClientError)
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
		e := errorpkg.ErrorServiceUnavailable("endpoint validity check failed")
		return errorpkg.Wrap(e, err)
	}
	if !ok {
		e := errorpkg.ErrorServiceUnavailable("checkGeneralEndpointValidity is not ok")
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
	if s.opt.skipRegistryCheck {
		return r, nil
	}

	// target
	u, err := url.Parse(serviceTarget)
	if err != nil {
		e := errorpkg.ErrorInternalServer("failed to parse service target")
		return nil, errorpkg.Wrap(e, err)
	}
	target := resolver.Target{URL: *u}
	// 健康检查
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = r.GetService(ctx, target.Endpoint())
	if err != nil {
		e := errorpkg.ErrorServiceUnavailable("registry discovery failed")
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
		if s.opt.consulClient == nil {
			return nil, errorpkg.WithStack(uninitializedConsulClientError)
		}
		r, err := registrypkg.NewConsulRegistry(s.opt.consulClient)
		if err != nil {
			e := errorpkg.ErrorInternalServer("")
			return nil, errorpkg.Wrap(e, err)
		}
		return r, nil
	case configpb.RegistryTypeEnum_ETCD:
		if s.opt.etcdClient == nil {
			return nil, errorpkg.WithStack(uninitializedEtcdClientError)
		}
		r, err := registrypkg.NewEtcdRegistry(s.opt.etcdClient)
		if err != nil {
			e := errorpkg.ErrorInternalServer("")
			return nil, errorpkg.Wrap(e, err)
		}
		return r, nil
	}
}
