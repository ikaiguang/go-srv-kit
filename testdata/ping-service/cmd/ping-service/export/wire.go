//go:build wireinject
// +build wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/services"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	bizrepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	datarepo "github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/repo"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

func exportPingData(launcherManager setuputil.LauncherManager) (datarepo.PingDataRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// data
		data.NewPingData,
	))
	return nil, nil
}

func exportWebsocketBiz(launcherManager setuputil.LauncherManager) (bizrepo.WebsocketBizRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// biz
		biz.NewWebsocketBiz,
	))
	return nil, nil
}

func exportPingBiz(launcherManager setuputil.LauncherManager) (bizrepo.PingBizRepo, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		setuputil.GetServiceAPIManager,
		exportPingData,
		// biz
		biz.NewPingBiz,
	))
	return nil, nil
}

func exportPingService(launcherManager setuputil.LauncherManager) (pingservicev1.SrvPingServer, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		exportPingBiz,
		// service
		service.NewPingService,
	))
	return nil, nil
}

func exportTestdataService(launcherManager setuputil.LauncherManager) (testdataservicev1.SrvTestdataServer, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		exportWebsocketBiz,
		// service
		service.NewTestdataService,
	))
	return nil, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	panic(wire.Build(
		setuputil.GetLogger,
		// service
		exportPingService, exportTestdataService,
		// home
		service.NewHomeService,
		exportWebsocketBiz, service.NewWebsocketService,
		// register services
		service.RegisterServices,
	))
	return nil, nil
}
