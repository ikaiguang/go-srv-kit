//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (serverutil.ServiceInterface, error) {
	panic(wire.Build(
		// service
		setuputil.GetLogger,
		setuputil.GetServiceAPIManager,
		data.NewPingData,
		biz.NewWebsocketBiz, biz.NewPingBiz,
		service.NewHomeService, service.NewWebsocketService,
		service.NewPingService, service.NewTestdataService,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
