package serverutil

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	_httpOptions []http.ServerOption
	_grpcOptions []grpc.ServerOption
)

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
