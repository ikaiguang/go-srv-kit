package servers

import (
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"

	routes "github.com/ikaiguang/go-srv-kit/example/internal/route"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
)

// NewApp .
func NewApp(engineHandler setup.Engine) (app *kratos.App, err error) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		return app, err
	}
	log.SetLogger(logger)

	// 服务
	hs, err := NewHTTPServer(engineHandler)
	if err != nil {
		return app, err
	}
	gs, err := NewGRPCServer(engineHandler)
	if err != nil {
		return app, err
	}

	// 路由
	err = routes.RegisterRoutes(engineHandler, hs, gs)
	if err != nil {
		return app, err
	}

	// 服务
	var servers []transport.Server
	if cfg := engineHandler.HTTPConfig(); cfg != nil && cfg.Enable {
		servers = append(servers, hs)
	}
	if cfg := engineHandler.GRPCConfig(); cfg != nil && cfg.Enable {
		servers = append(servers, gs)
	}
	if len(servers) == 0 {
		err = pkgerrors.New("服务列表为空")
		return app, err
	}

	// app
	var (
		appConfig  = engineHandler.AppConfig()
		appID      = apputil.ID(appConfig)
		appOptions = []kratos.Option{
			kratos.ID(appID),
			kratos.Name(appID),
			kratos.Version(appConfig.Version),
			kratos.Metadata(appConfig.Metadata),
			kratos.Logger(logger),
			kratos.Server(servers...),
		}
	)

	// 启用服务注册中心
	if appConfig.Setting != nil && appConfig.Setting.EnableServiceRegistry {
		stdlog.Println("|*** 加载：服务注册与发现")
		consulClient, err := engineHandler.GetConsulClient()
		if err != nil {
			return app, err
		}
		r := consul.New(consulClient)
		appOptions = append(appOptions, kratos.Registrar(r))
	}

	app = kratos.New(appOptions...)
	return app, err
}
