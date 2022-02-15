package servers

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(packages setup.Packages) (srv *http.Server, err error) {
	c := packages.ServerConfig()
	stdlog.Printf("|*** 加载HTTP服务：%s\n", c.Http.Addr)

	// 日志
	logger, _, err := packages.Logger()
	if err != nil {
		return srv, err
	}

	// http
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Logger(logger),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv = http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv, err
}

// RegisterHTTPRoute 注册路由
func RegisterHTTPRoute(packages setup.Packages, srv *http.Server) (err error) {
	stdlog.Println("|*** 注册HTTP路由：...")
	return err
}
