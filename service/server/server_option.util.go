package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	authpkg "github.com/ikaiguang/go-srv-kit/kratos/auth"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
)

var (
	_httpOptions []http.ServerOption
	_grpcOptions []grpc.ServerOption
)

// AppOptionProvider 基于 LauncherManager 生成 Kratos App 选项。
type AppOptionProvider func(launcherManager setuputil.LauncherManager) ([]kratos.Option, error)

// AuthManagerProvider 按需获取认证管理器。
type AuthManagerProvider func(launcherManager setuputil.LauncherManager) (authpkg.AuthRepo, error)

// JaegerExporterProvider 按需获取 Jaeger exporter。
type JaegerExporterProvider func(launcherManager setuputil.LauncherManager) (*otlptrace.Exporter, error)

// Option 配置 all-in-one 启动流程。
type Option func(*options)

type options struct {
	setupOptions           []setuputil.Option
	appOptionProviders     []AppOptionProvider
	authManagerProvider    AuthManagerProvider
	jaegerExporterProvider JaegerExporterProvider
}

// WithSetupOptions 注入 LauncherManager 组件注册选项。
func WithSetupOptions(opts ...setuputil.Option) Option {
	return func(o *options) {
		o.setupOptions = append(o.setupOptions, opts...)
	}
}

// WithAppOptionProvider 注入 Kratos App 选项提供者。
func WithAppOptionProvider(providers ...AppOptionProvider) Option {
	return func(o *options) {
		o.appOptionProviders = append(o.appOptionProviders, providers...)
	}
}

// WithAuthManagerProvider 注入认证管理器提供者。
func WithAuthManagerProvider(provider AuthManagerProvider) Option {
	return func(o *options) {
		o.authManagerProvider = provider
	}
}

// WithJaegerExporterProvider 注入 Jaeger exporter 提供者。
func WithJaegerExporterProvider(provider JaegerExporterProvider) Option {
	return func(o *options) {
		o.jaegerExporterProvider = provider
	}
}

func RegisterHTTPServerOption(opts ...http.ServerOption) {
	_httpOptions = append(_httpOptions, opts...)
}

func RegisterGRPCServerOption(opts ...grpc.ServerOption) {
	_grpcOptions = append(_grpcOptions, opts...)
}

// InjectHTTPServerOptions 注入 HTTP 服务器选项
func InjectHTTPServerOptions(opts *[]http.ServerOption) {
	*opts = append(*opts, _httpOptions...)
}

// InjectGRPCServerOptions 注入 GRPC 服务器选项
func InjectGRPCServerOptions(opts *[]grpc.ServerOption) {
	*opts = append(*opts, _grpcOptions...)
}
