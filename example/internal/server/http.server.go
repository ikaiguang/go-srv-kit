package servers

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	middlewareutil "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(packages setup.Packages) (srv *http.Server, err error) {
	httpConfig := packages.HTTPConfig()
	stdlog.Printf("|*** 加载HTTP服务：%s\n", httpConfig.Addr)

	// 日志
	logger, _, err := packages.Logger()
	if err != nil {
		return srv, err
	}

	// options
	var opts = []http.ServerOption{
		http.Logger(logger),
	}
	if httpConfig.Network != "" {
		opts = append(opts, http.Network(httpConfig.Network))
	}
	if httpConfig.Addr != "" {
		opts = append(opts, http.Address(httpConfig.Addr))
	}
	if httpConfig.Timeout != nil {
		opts = append(opts, http.Timeout(httpConfig.Timeout.AsDuration()))
	}

	// 响应
	opts = append(opts, http.ResponseEncoder(apputil.ResponseEncoder))
	opts = append(opts, http.ErrorEncoder(apputil.ErrorEncoder))

	// ===== 中间件 =====
	var middlewareSlice = []middleware.Middleware{
		recovery.Recovery(),
	}
	// 中间件日志
	loggerMiddle, _, err := packages.LoggerMiddleware()
	if err != nil {
		return srv, err
	}
	// 日志输出
	//middlewareSlice = append(middlewareSlice, logging.Server(loggerMiddle))
	// 请求头
	middlewareSlice = append(middlewareSlice, middlewareutil.RequestHeader())
	// 错误追踪
	if packages.IsDebugMode() {
		middlewareSlice = append(middlewareSlice, middlewareutil.ErrorStack(loggerMiddle))
	}

	// 中间件选项
	opts = append(opts, http.Middleware(middlewareSlice...))

	// 服务
	srv = http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv, err
}

// RegisterHTTPRoute 注册路由
func RegisterHTTPRoute(packages setup.Packages, srv *http.Server) (err error) {
	stdlog.Println("|*** 注册HTTP路由：...")
	return err
}
