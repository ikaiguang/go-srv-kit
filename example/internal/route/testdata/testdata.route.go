package testdataroute

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1"
	testdatasrv "github.com/ikaiguang/go-srv-kit/example/internal/application/service/testdata"
)

// RegisterRoutes 注册路由
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {
	stdlog.Println("|*** 注册路由：testdata")

	testdata := testdatasrv.NewTestService(logger)

	v1.RegisterSrvTestdataHTTPServer(hs, testdata)
	v1.RegisterSrvTestdataServer(gs, testdata)
}
