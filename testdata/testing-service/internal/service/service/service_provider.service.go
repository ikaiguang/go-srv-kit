package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	servicev1 "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/testdata-service/v1/services"
	stdlog "log"
)

// RegisterServices 注册服务
// @return Services 用于wire
// @return func() = cleanup 关闭资源
// @return error 错误
func RegisterServices(
	hs *http.Server, gs *grpc.Server,
	testingV1Service servicev1.SrvTestdataServer,
) (serverutil.ServiceInterface, error) {
	// 先进后出
	var cleanup = func() {}
	// grpc
	if gs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：GRPC: testingV1Service")
		servicev1.RegisterSrvTestdataServer(gs, testingV1Service)

		// cleanup example
		//cleanup = func() {
		//	cleanup()
		//}
	}

	// http
	if hs != nil {
		stdlog.Println("|*** REGISTER_ROUTER：HTTP: testingV1Service")
		servicev1.RegisterSrvTestdataHTTPServer(hs, testingV1Service)

		// special
		//RegisterSpecialRouters(hs, homeService, websocketService)

		// cleanup example
		//cleanup = func() {
		//	cleanup()
		//}
	}

	return serverutil.NewServiceInterface(cleanup), nil
}

//func RegisterSpecialRouters(hs *http.Server, homeService *HomeService, websocketService *WebsocketService) {
//	// router
//	router := mux.NewRouter()
//
//	stdlog.Println("|*** REGISTER_ROUTER：Root(/)")
//	router.HandleFunc("/", homeService.Homepage)
//	hs.Handle("/", router)
//
//	stdlog.Println("|*** REGISTER_ROUTER：Websocket")
//	router.HandleFunc("/ws/v1/testdata/websocket", websocketService.TestWebsocket)
//
//	// router
//	hs.Handle("/ws/v1/websocket", router)
//}
