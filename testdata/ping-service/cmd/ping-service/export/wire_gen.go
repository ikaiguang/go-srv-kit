// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package serviceexporter

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ikaiguang/go-srv-kit/service/cleanup"
	"github.com/ikaiguang/go-srv-kit/service/setup"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/services"
	servicev1_2 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/repo"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/repo"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

// Injectors from wire.go:

func exportPingData(launcherManager setuputil.LauncherManager) (datarepo.PingDataRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	pingDataRepo := data.NewPingData(logger)
	return pingDataRepo, nil
}

func exportWebsocketBiz(launcherManager setuputil.LauncherManager) (bizrepo.WebsocketBizRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	websocketBizRepo := biz.NewWebsocketBiz(logger)
	return websocketBizRepo, nil
}

func exportPingBiz(launcherManager setuputil.LauncherManager) (bizrepo.PingBizRepo, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	serviceAPIManager, err := setuputil.GetServiceAPIManager(launcherManager)
	if err != nil {
		return nil, err
	}
	pingDataRepo, err := exportPingData(launcherManager)
	if err != nil {
		return nil, err
	}
	pingBizRepo := biz.NewPingBiz(logger, serviceAPIManager, pingDataRepo)
	return pingBizRepo, nil
}

func exportPingService(launcherManager setuputil.LauncherManager) (servicev1.SrvPingServer, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	pingBizRepo, err := exportPingBiz(launcherManager)
	if err != nil {
		return nil, err
	}
	srvPingServer := service.NewPingService(logger, pingBizRepo)
	return srvPingServer, nil
}

func exportTestdataService(launcherManager setuputil.LauncherManager) (servicev1_2.SrvTestdataServer, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	websocketBizRepo, err := exportWebsocketBiz(launcherManager)
	if err != nil {
		return nil, err
	}
	srvTestdataServer := service.NewTestdataService(logger, websocketBizRepo)
	return srvTestdataServer, nil
}

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	homeService := service.NewHomeService(logger)
	websocketBizRepo, err := exportWebsocketBiz(launcherManager)
	if err != nil {
		return nil, err
	}
	websocketService := service.NewWebsocketService(logger, websocketBizRepo)
	srvPingServer, err := exportPingService(launcherManager)
	if err != nil {
		return nil, err
	}
	srvTestdataServer, err := exportTestdataService(launcherManager)
	if err != nil {
		return nil, err
	}
	cleanupManager, err := service.RegisterServices(hs, gs, homeService, websocketService, srvPingServer, srvTestdataServer)
	if err != nil {
		return nil, err
	}
	return cleanupManager, nil
}
