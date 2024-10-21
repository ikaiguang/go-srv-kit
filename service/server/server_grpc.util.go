package serverutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	stdlog "log"
)

var _ metadata.Option

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	launcherManager setuputil.LauncherManager,
	authWhiteList map[string]middlewareutil.TransportServiceKind,
) (*grpc.Server, error) {
	grpcConfig := configutil.GRPCConfig(launcherManager.GetConfig())
	if !grpcConfig.GetEnable() {
		return nil, nil
	}
	stdlog.Printf("|*** LOADING：GRPC Server：%s\n", grpcConfig.GetAddr())

	// loggerForMiddleware
	loggerForMiddleware, err := launcherManager.GetLoggerForMiddleware()
	if err != nil {
		return nil, err
	}

	// options
	var opts []grpc.ServerOption
	if grpcConfig.Network != "" {
		opts = append(opts, grpc.Network(grpcConfig.GetNetwork()))
	}
	if grpcConfig.Addr != "" {
		opts = append(opts, grpc.Address(grpcConfig.GetAddr()))
	}
	if grpcConfig.Timeout != nil {
		opts = append(opts, grpc.Timeout(grpcConfig.GetTimeout().AsDuration()))
	}

	// ===== 中间件 =====
	var (
		logHelper       = log.NewHelper(loggerForMiddleware)
		middlewareSlice = middlewareutil.DefaultServerMiddlewares(logHelper)
	)

	// setting
	settingConfig := configutil.SettingConfig(launcherManager.GetConfig())
	if settingConfig.GetEnableAuthMiddleware() {
		stdlog.Println("|*** LOADING：AuthMiddleware：GRPC")
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
	opts = append(opts, grpc.Middleware(middlewareSlice...))

	//v1.RegisterGreeterServer(srv, greeter)
	return grpc.NewServer(opts...), err
}
