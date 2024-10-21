package clientutil

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdgrpc "google.golang.org/grpc"
	"strings"
)

var (
	uninitializedConsulClientError = errorpkg.ErrorBadRequest("uninitialized: consulClient == nil")
	uninitializedEtcdClientError   = errorpkg.ErrorBadRequest("uninitialized: etcdClient == nil")
	uninitializedGRPCConnError     = errorpkg.ErrorBadRequest("uninitialized: grpcConn == nil")
	uninitializedHTTPClientError   = errorpkg.ErrorBadRequest("uninitialized: httpClient == nil")
)

type ServiceAPIManager interface {
	// RegisterServiceAPIConfigs 注册服务API，覆盖更新
	RegisterServiceAPIConfigs(apis []*Config, opts ...Option) error
	// GetServiceAPIConfig 获取服务配置
	GetServiceAPIConfig(serviceName ServiceName) (*Config, error)
	// NewAPIConnection 实例化客户端链接
	NewAPIConnection(serviceName ServiceName) (ServiceAPIConnection, error)
}

type ServiceAPIConnection interface {
	GetTransportType() configpb.TransportTypeEnum_TransportType
	IsTransportType(tt configpb.TransportTypeEnum_TransportType) bool
	GetGRPCConnection() (*stdgrpc.ClientConn, error)
	GetHTTPClient() (*http.Client, error)
}

// ServiceName ...
type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

// Config ...
type Config struct {
	ServiceName   string                                   // 服务名称
	TransportType configpb.TransportTypeEnum_TransportType // 传输协议：http、grpc、...；默认: HTTP
	RegistryType  configpb.RegistryTypeEnum_RegistryType   // 注册类型：注册类型：endpoint、consul、...；配置中心配置：${registry_type}；例： Bootstrap.Consul
	ServiceTarget string                                   // 服务目标：endpoint或registry，例：http://127.0.0.1:8899、discovery:///${registry_endpoint}
}

func (s *Config) SetByPbClusterServiceApi(cfg *configpb.ClusterServiceApi) {
	s.ServiceName = cfg.GetServiceName()
	tt := strings.ToLower(cfg.GetTransportType())
	switch tt {
	default:
		s.TransportType = configpb.TransportTypeEnum_HTTP
	case "http":
		s.TransportType = configpb.TransportTypeEnum_HTTP
	case "grpc":
		s.TransportType = configpb.TransportTypeEnum_GRPC
	}
	rt := strings.ToLower(cfg.GetRegistryType())
	switch rt {
	default:
		s.RegistryType = configpb.RegistryTypeEnum_ENDPOINT
	case "endpoint":
		s.RegistryType = configpb.RegistryTypeEnum_ENDPOINT
	case "consul":
		s.RegistryType = configpb.RegistryTypeEnum_CONSUL
	case "etcd":
		s.RegistryType = configpb.RegistryTypeEnum_ETCD
	}
	s.ServiceTarget = cfg.GetServiceTarget()
}

func (s *Config) IsConsulRegistry() bool {
	return s.RegistryType == configpb.RegistryTypeEnum_CONSUL
}

func (s *Config) IsEtcdRegistry() bool {
	return s.RegistryType == configpb.RegistryTypeEnum_ETCD
}

func ToConfig(apiConfigs []*configpb.ClusterServiceApi) ([]*Config, map[configpb.RegistryTypeEnum_RegistryType]bool, error) {
	var (
		results = make([]*Config, 0, len(apiConfigs))
		diffRT  = make(map[configpb.RegistryTypeEnum_RegistryType]bool, len(apiConfigs))
	)
	for i := range apiConfigs {
		if err := apiConfigs[i].Validate(); err != nil {
			e := errorpkg.ErrorBadRequest("")
			return results, diffRT, errorpkg.Wrap(e, err)
		}
		conf := &Config{}
		conf.SetByPbClusterServiceApi(apiConfigs[i])
		results = append(results, conf)
		diffRT[conf.RegistryType] = true
	}
	return results, diffRT, nil
}

type clientConnection struct {
	transportType configpb.TransportTypeEnum_TransportType
	grpcConn      *stdgrpc.ClientConn
	httpClient    *http.Client
}

func (c *clientConnection) SetTransportType(tt configpb.TransportTypeEnum_TransportType) {
	_, ok := configpb.TransportTypeEnum_TransportType_name[int32(tt)]
	if ok {
		c.transportType = tt
	}
	if c.transportType == configpb.TransportTypeEnum_UNSPECIFIED {
		c.transportType = configpb.TransportTypeEnum_HTTP
	}
}

func (c *clientConnection) GetTransportType() configpb.TransportTypeEnum_TransportType {
	return c.transportType
}

func (c *clientConnection) IsTransportType(tt configpb.TransportTypeEnum_TransportType) bool {
	return c.transportType == tt
}

func (c *clientConnection) IsHTTPTransport() bool {
	return c.transportType == configpb.TransportTypeEnum_HTTP
}

func (c *clientConnection) IsGRCPTransport() bool {
	return c.transportType == configpb.TransportTypeEnum_GRPC
}

func (c *clientConnection) GetGRPCConnection() (*stdgrpc.ClientConn, error) {
	if c.grpcConn == nil {
		return nil, errorpkg.WithStack(uninitializedGRPCConnError)
	}
	return c.grpcConn, nil
}

func (c *clientConnection) GetHTTPClient() (*http.Client, error) {
	if c.httpClient == nil {
		return nil, errorpkg.WithStack(uninitializedHTTPClientError)
	}
	return c.httpClient, nil
}
