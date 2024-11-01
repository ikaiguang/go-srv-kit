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
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/biz/biz"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/data/data"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/service/service"
)

// Injectors from wire.go:

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
	logger, err := setuputil.GetLogger(launcherManager)
	if err != nil {
		return nil, err
	}
	homeService := service.NewHomeService(logger)
	websocketBizRepo := biz.NewWebsocketBiz(logger)
	websocketService := service.NewWebsocketService(logger, websocketBizRepo)
	serviceAPIManager, err := setuputil.GetServiceAPIManager(launcherManager)
	if err != nil {
		return nil, err
	}
	pingDataRepo := data.NewPingData(logger)
	pingBizRepo := biz.NewPingBiz(logger, serviceAPIManager, pingDataRepo)
	srvPingServer := service.NewPingService(logger, pingBizRepo)
	srvTestdataServer := service.NewTestdataService(logger, websocketBizRepo)
	cleanupManager, err := service.RegisterServices(hs, gs, homeService, websocketService, srvPingServer, srvTestdataServer)
	if err != nil {
		return nil, err
	}
	return cleanupManager, nil
}
