package serverutil

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

var _ metadata.Option

func errorpkgMissingProvider(name string) error {
	e := errorpkg.ErrorBadRequest("%s provider is required; import the corresponding service module and pass its all-in-one option", name)
	return errorpkg.WithStack(e)
}

// NewHTTPServer new HTTP server.
func NewHTTPServer(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
	serverOpts ...http.ServerOption,
) (*http.Server, error) {
	return newHTTPServerWithOptions(launcherManager, authWhiteList, nil, serverOpts...)
}

func newHTTPServerWithOptions(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
	runOpts *options,
	serverOpts ...http.ServerOption,
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
		if runOpts == nil || runOpts.authManagerProvider == nil {
			return nil, errorpkgMissingProvider("auth manager")
		}
		authManager, err := runOpts.authManagerProvider(launcherManager)
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

	// other
	InjectHTTPServerOptions(&opts)
	opts = append(opts, serverOpts...)

	return http.NewServer(opts...), err
}
