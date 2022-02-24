package pingroute

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1"
	pingsrv "github.com/ikaiguang/go-srv-kit/example/internal/application/service/ping"
)

// RegisterRoutes 注册路由
func RegisterRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {
	ping := pingsrv.NewPingService(logger)

	pingv1.RegisterSrvPingHTTPServer(hs, ping)
	pingv1.RegisterSrvPingServer(gs, ping)
}
