package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	clientpkg "github.com/ikaiguang/go-srv-kit/kratos/client"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	"strings"
)

func (s *serviceAPIManager) NewHTTPClient(apiConfig *Config, otherOpts ...http.ClientOption) (*http.Client, error) {
	var opts = []http.ClientOption{
		http.WithTimeout(DefaultTimeout),
	}
	opts = append(opts, apputil.ClientDecoderEncoder()...)

	// 中间件
	logHelper := log.NewHelper(s.opt.logger)
	opts = append(opts, http.WithMiddleware(middlewarepkg.DefaultClientMiddlewares(logHelper)...))

	// 服务端点
	endpointOpts, err := s.getHTTPEndpointOptions(apiConfig)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)
	logHelper.Infow(
		"msg", "NewHTTPClient",
		"client.serviceName", apiConfig.ServiceName,
		"client.transportType", apiConfig.TransportType.String(),
		"client.registryType", apiConfig.RegistryType.String(),
		"client.serviceTarget", apiConfig.ServiceTarget,
	)

	// 其他
	opts = append(opts, otherOpts...)

	// http 链接
	conn, err := clientpkg.NewHTTPClient(context.Background(), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return conn, nil
}

// getHTTPEndpointOptions 获取服务端点
func (s *serviceAPIManager) getHTTPEndpointOptions(apiConfig *Config) ([]http.ClientOption, error) {
	var opts []http.ClientOption

	// endpoint
	endpoint := apiConfig.ServiceTarget

	// registry
	switch apiConfig.RegistryType {
	case configpb.RegistryTypeEnum_CONSUL, configpb.RegistryTypeEnum_ETCD:
		r, err := s.getAndCheckRegistryDiscovery(apiConfig, apiConfig.ServiceTarget)
		if err != nil {
			return nil, err
		}
		opts = append(opts, http.WithDiscovery(r))
	default:
		if !strings.Contains(endpoint, Sep) {
			e := errorpkg.ErrorInvalidParameter("invalid ServiceTarget")
			return nil, errorpkg.WithStack(e)
		}
		err := s.checkGeneralEndpointValidity(apiConfig.ServiceTarget)
		if err != nil {
			return nil, err
		}
	}
	opts = append(opts, http.WithEndpoint(endpoint))
	return opts, nil
}
