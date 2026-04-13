//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setupv2 "github.com/ikaiguang/go-srv-kit/service/setup_v2"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

func exportServices(launcherManager setupv2.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setupv2.GetLogger,
		setupv2.GetServiceAPIManager,
		// data
		data.NewPingData,
		// biz
		biz.NewPingBiz,
		biz.NewWebsocketBiz,
		// service
		service.NewPingService,
		service.NewTestdataService,
		service.NewHomeService,
		service.NewWebsocketService,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
