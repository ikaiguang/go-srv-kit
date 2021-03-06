package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingroute "github.com/ikaiguang/go-srv-kit/example/internal/route/ping"
	testdataroute "github.com/ikaiguang/go-srv-kit/example/internal/route/testdata"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		return err
	}

	pingroute.RegisterRoutes(hs, gs, logger)
	testdataroute.RegisterRoutes(hs, gs, logger)

	return err
}
