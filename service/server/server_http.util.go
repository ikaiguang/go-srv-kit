package serverutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	stdlog "log"
)

var _ metadata.Option

// NewHTTPServer new HTTP server.
func NewHTTPServer(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (*http.Server, error) {
	httpConfig := configutil.HTTPConfig(launcherManager.GetConfig())
	if !httpConfig.GetEnable() {
		return nil, nil
	}
	stdlog.Printf("|*** LOADING：HTTP Server：%s\n", httpConfig.GetAddr())

	// loggerForMiddleware
	loggerForMiddleware, err := launcherManager.GetLoggerForMiddleware()
	if err != nil {
		return nil, err
	}

	// options
	var opts []http.ServerOption
	//var opts = []http.ServerOption{
	//	http.Filter(middlewareutil.NewCORS()),
	//}
	if httpConfig.Network != "" {
		opts = append(opts, http.Network(httpConfig.GetNetwork()))
	}
	if httpConfig.Addr != "" {
		opts = append(opts, http.Address(httpConfig.GetAddr()))
	}
	if httpConfig.Timeout != nil {
		opts = append(opts, http.Timeout(httpConfig.GetTimeout().AsDuration()))
	}

	// 编码 与 解码
	opts = append(opts, apputil.ServerDecoderEncoder()...)
	opts = append(opts, apppkg.NotFound404())

	// ===== 中间件 =====
	var (
		logHelper       = log.NewHelper(loggerForMiddleware)
		middlewareSlice = middlewareutil.DefaultServerMiddlewares(logHelper)
	)

	// setting
	settingConfig := configutil.SettingConfig(launcherManager.GetConfig())
	if settingConfig.GetEnableAuthMiddleware() {
		stdlog.Println("|*** LOADING：AuthMiddleware：HTTP")
		// authManager
		authManager, err := launcherManager.GetAuthManager()
		if err != nil {
			return nil, err
		}
		jwtMiddleware, err := middlewareutil.NewAuthMiddleware(authManager, authWhiteList)
		if err != nil {
			return nil, err
		}
		middlewareSlice = append(middlewareSlice, jwtMiddleware)
	}

	// 中间件选项
	opts = append(opts, http.Middleware(middlewareSlice...))

	//v1.RegisterGreeterHTTPServer(srv, greeter)
	return http.NewServer(opts...), err
}
