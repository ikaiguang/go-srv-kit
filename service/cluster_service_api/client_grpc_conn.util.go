package clientutil

import (
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	stdgrpc "google.golang.org/grpc"
)

var (
	DefaultTimeout     = time.Minute      // 请求超时（建议用 ctx 控制）
	DefaultDialTimeout = 10 * time.Second // 拨号超时

)

const (
	Sep = "://"
)

func (s *serviceAPIManager) NewGRPCConnection(apiConfig *Config, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	var opts = []grpc.ClientOption{
		grpc.WithHealthCheck(true),
		grpc.WithPrintDiscoveryDebugLog(true),
	}

	// 中间件
	logHelper := log.NewHelper(s.opt.logger)
	opts = append(opts, grpc.WithMiddleware(middlewarepkg.DefaultClientMiddlewares(logHelper)...))

	// 服务端点
	endpointOpts, err := s.getGRPCEndpointOptions(apiConfig)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)
	logHelper.Infow(
		"msg", "NewGRPCConnection",
		"client.serviceName", apiConfig.ServiceName,
		"client.transportType", apiConfig.TransportType.String(),
		"client.registryType", apiConfig.RegistryType.String(),
		"client.serviceTarget", apiConfig.ServiceTarget,
	)

	// 其他
	opts = append(opts, otherOpts...)

	// grpc 链接（带拨号超时）
	dialCtx, dialCancel := context.WithTimeout(context.Background(), DefaultDialTimeout)
	defer dialCancel()
	conn, err := grpc.DialInsecure(dialCtx, opts...)
	if err != nil {
		e := errorpkg.ErrorInternalServer("failed to create grpc connection")
		return nil, errorpkg.Wrap(e, err)
	}
	return conn, nil
}

func (s *serviceAPIManager) getGRPCEndpointOptions(apiConfig *Config) ([]grpc.ClientOption, error) {
	var opts []grpc.ClientOption

	// endpoint
	endpoint := apiConfig.ServiceTarget

	// registry
	switch apiConfig.RegistryType {
	case configpb.RegistryTypeEnum_CONSUL, configpb.RegistryTypeEnum_ETCD:
		r, err := s.getAndCheckRegistryDiscovery(apiConfig, apiConfig.ServiceTarget)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithDiscovery(r))
	default:
		if i := strings.Index(endpoint, Sep); i >= 0 {
			endpoint = endpoint[i+len(Sep):]
		}
		err := s.checkGeneralEndpointValidity(apiConfig.ServiceTarget)
		if err != nil {
			return nil, err
		}
	}
	opts = append(opts, grpc.WithEndpoint(endpoint))
	return opts, nil
}
