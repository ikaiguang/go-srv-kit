package servers

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
	middlewareutil "github.com/ikaiguang/go-srv-kit/kratos/middleware"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(packages setup.Packages) (srv *grpc.Server, err error) {
	c := packages.ServerConfig()
	stdlog.Printf("|*** 加载GRPC服务：%s\n", c.Grpc.Addr)

	// 日志
	logger, _, err := packages.Logger()
	if err != nil {
		return srv, err
	}

	// options
	var opts = []grpc.ServerOption{
		grpc.Logger(logger),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

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
	// 错误追踪
	if packages.IsDebugMode() {
		middlewareSlice = append(middlewareSlice, middlewareutil.ErrorStack(loggerMiddle))
	}

	// 中间件选项
	opts = append(opts, grpc.Middleware(middlewareSlice...))

	// 服务
	srv = grpc.NewServer(opts...)
	//v1.RegisterGreeterServer(srv, greeter)

	return srv, err
}

// RegisterGRPCRoute 注册路由
func RegisterGRPCRoute(packages setup.Packages, srv *grpc.Server) (err error) {
	stdlog.Println("|*** 注册GRPC路由：...")
	return err
}
