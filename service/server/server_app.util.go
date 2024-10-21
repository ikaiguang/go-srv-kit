package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	registrypkg "github.com/ikaiguang/go-srv-kit/kratos/registry"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	tracerutil "github.com/ikaiguang/go-srv-kit/service/tracer"
	stdlog "log"
	"net/url"
)

// NewApp .
func NewApp(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (*kratos.App, error) {
	var (
		conf               = launcherManager.GetConfig()
		appConfig          = conf.GetApp()
		enableJaegerTracer = conf.GetSetting().GetEnableJaegerTracer()
	)
	if enableJaegerTracer {
		exp, err := launcherManager.GetJaegerExporter()
		if err != nil {
			return nil, err
		}
		err = tracerutil.InitTracerWithJaegerExporter(appConfig, exp)
		if err != nil {
			return nil, err
		}
	} else {
		err := tracerutil.InitTracer(appConfig)
		if err != nil {
			return nil, err
		}
	}

	// logger
	logger, err := launcherManager.GetLogger()
	if err != nil {
		return nil, err
	}

	// 服务
	var servers []transport.Server

	// http
	httpConfig := conf.GetServer().GetHttp()
	if httpConfig.GetEnable() {
		servers = append(servers, hs)
	}

	// grpc
	grpcConfig := conf.GetServer().GetGrpc()
	if grpcConfig.GetEnable() {
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		e := errorpkg.ErrorInvalidParameter("server list cannot be empty")
		return nil, errorpkg.WithStack(e)
	}

	// appid
	appID := apputil.ID(apputil.ToAppConfig(appConfig))
	appConfig.Id = appID
	if appConfig.GetMetadata() == nil {
		appConfig.Metadata = make(map[string]string)
	}
	appConfig.GetMetadata()["id"] = appID

	// app
	var (
		appOptions = []kratos.Option{
			kratos.ID(appID),
			kratos.Name(appID),
			kratos.Version(appConfig.GetServerVersion()),
			kratos.Metadata(appConfig.GetMetadata()),
			kratos.Logger(logger),
			kratos.Server(servers...),
		}
	)

	// 服务注册，如果为空，自动获取服务的Endpoint
	// 如： http://192.168.100.200:10001、grpc://192.168.100.200:10002
	// 如： http://xxx-service.namespace.svc.cluster.local:10001、grpc://xxx-service.namespace:10002
	var endpoints = make([]*url.URL, 0, len(appConfig.GetRegistryEndpoints()))
	for _, item := range appConfig.GetRegistryEndpoints() {
		u, err := url.Parse(item)
		if err != nil {
			e := errorpkg.ErrorInvalidParameter(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		endpoints = append(endpoints, u)
	}
	if len(endpoints) > 0 {
		appOptions = append(appOptions, kratos.Endpoint(endpoints...))
	}

	// 启用服务注册中心
	if conf.GetSetting().GetEnableConsulRegistry() {
		stdlog.Println("|*** LOADING: ServiceRegistry: ...")
		consulClient, err := launcherManager.GetConsulClient()
		if err != nil {
			return nil, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return nil, err
		}
		appOptions = append(appOptions, kratos.Registrar(r))
	}

	return kratos.New(appOptions...), err
}
