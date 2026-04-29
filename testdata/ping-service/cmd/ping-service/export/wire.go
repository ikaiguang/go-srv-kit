//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		clientutil.GetManager,
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
