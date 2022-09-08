package servers

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdlog "log"

	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/example/pkg/middleware"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	headermiddle "github.com/ikaiguang/go-srv-kit/kratos/middleware/header"
	logmiddle "github.com/ikaiguang/go-srv-kit/kratos/middleware/log"
)

var _ metadata.Option

// NewHTTPServer new HTTP server.
func NewHTTPServer(engineHandler setup.Engine) (srv *http.Server, err error) {
	httpConfig := engineHandler.HTTPConfig()
	stdlog.Printf("|*** 加载：HTTP服务：%s\n", httpConfig.Addr)

	// 日志
	//logger, _, err := engineHandler.Logger()
	//if err != nil {
	//	return srv, err
	//}

	// options
	var opts []http.ServerOption
	//var opts = []http.ServerOption{
	//	http.Filter(middlewarepkg.NewCORS()),
	//}
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
		//metadata.Server(),
	}
	// tracer
	settingConfig := engineHandler.ServerSettingConfig()
	if settingConfig != nil && settingConfig.EnableServiceTracer {
		stdlog.Println("|*** 加载：服务追踪：HTTP")
		if err = middlewarepkg.SetTracerProvider(engineHandler); err != nil {
			return srv, err
		}
		middlewareSlice = append(middlewareSlice, tracing.Server())
	}
	// 请求头
	middlewareSlice = append(middlewareSlice, headermiddle.RequestHeader())
	// 中间件日志
	middleLogger, _, err := engineHandler.LoggerMiddleware()
	if err != nil {
		return srv, err
	}
	// 日志输出
	//errorutil.DefaultStackTracerDepth += 2
	middlewareSlice = append(middlewareSlice, logmiddle.ServerLog(
		middleLogger,
		//logmiddle.WithDefaultSkip(),
	))
	// jwt
	//stdlog.Println("|*** 加载：JWT中间件：HTTP")
	//jwtMiddleware, err := middlewarepkg.NewJWTMiddleware(engineHandler)
	//if err != nil {
	//	return srv, err
	//}
	//middlewareSlice = append(middlewareSlice, jwtMiddleware)

	// 中间件选项
	opts = append(opts, http.Middleware(middlewareSlice...))

	// 服务
	srv = http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv, err
}

// RegisterHTTPRoute 注册路由
func RegisterHTTPRoute(engineHandler setup.Engine, srv *http.Server) (err error) {
	stdlog.Println("|*** 注册HTTP路由：...")
	return err
}
