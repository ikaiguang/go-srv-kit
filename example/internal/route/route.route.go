package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingroute "github.com/ikaiguang/go-srv-kit/example/internal/route/ping"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(packages setup.Packages, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// 日志
	logger, err := packages.Logger()
	if err != nil {
		return err
	}

	stdlog.Println("|*** 注册路由：ping")
	pingroute.RegisterRoutes(hs, gs, logger)

	return err
}
