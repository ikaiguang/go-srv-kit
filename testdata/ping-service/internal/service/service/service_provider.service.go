package service

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service/v1/services"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
)

// RegisterServices 注册服务
// @return Services 用于wire
// @return func() = cleanup 关闭资源
// @return error 错误
func RegisterServices(
	hs *http.Server, gs *grpc.Server,
	homeService *HomeService,
	websocketService *WebsocketService,
	pingService pingservicev1.SrvPingServer,
	testdataService testdataservicev1.SrvTestdataServer,
) (cleanuputil.CleanupManager, error) {
	// 先进后出
	var cleanupManager = cleanuputil.NewCleanupManager()
	// grpc
	if gs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：GRPC: PingServer")
		pingservicev1.RegisterSrvPingServer(gs, pingService)
		stdlog.Println("|*** REGISTER_ROUTER：GRPC: TestdataServer")
		testdataservicev1.RegisterSrvTestdataServer(gs, testdataService)

		//cleanupManager.Append(cleanup)
	}

	// http
	if hs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：HTTP: PingServer")
		pingservicev1.RegisterSrvPingHTTPServer(hs, pingService)
		stdlog.Println("|*** REGISTER_ROUTER：HTTP: TestdataServer")
		testdataservicev1.RegisterSrvTestdataHTTPServer(hs, testdataService)

		// special
		RegisterSpecialRouters(hs, homeService, websocketService)

		//cleanupManager.Append(cleanup)
	}

	return cleanupManager, nil
}

// 路由路径常量
const (
	routeRoot      = "/"
	routeWebsocket = "/ws/v1/testdata/websocket"
)

func RegisterSpecialRouters(hs *http.Server, homeService *HomeService, websocketService *WebsocketService) {
	// router
	router := mux.NewRouter()

	stdlog.Println("|*** REGISTER_ROUTER：Root(/)")
	router.HandleFunc(routeRoot, homeService.Homepage)
	hs.Handle(routeRoot, router)

	stdlog.Println("|*** REGISTER_ROUTER：Websocket")
	router.HandleFunc(routeWebsocket, websocketService.TestWebsocket)

	// router
	hs.Handle("/ws/v1/websocket", router)
}
