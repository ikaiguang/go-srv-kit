package testdata

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdgrpc "google.golang.org/grpc"
)

const (
	_defaultHTTPAddr     = "127.0.0.1:8081"
	_defaultHTTPEndpoint = "http://" + _defaultHTTPAddr

	_defaultGRPCAddr = "127.0.0.1:9091"
)

// HTTPEndpoint .
func HTTPEndpoint() string {
	return _defaultHTTPEndpoint
}

// GRPCAddr .
func GRPCAddr() string {
	return _defaultGRPCAddr
}

// GenURL .
func GenURL(urlPath string) string {
	return HTTPEndpoint() + "/" + strings.TrimPrefix(urlPath, "/")
}

// NewHTTPClient 客户端 http
func NewHTTPClient(ctx context.Context, opts ...http.ClientOption) (*http.Client, error) {
	return http.NewClient(ctx, opts...)
}

// NewGRPCConn 链接 grpc
func NewGRPCConn(ctx context.Context, opts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	return grpc.Dial(ctx, opts...)
}

// NewGRPCInsecureConn 链接 grpc
func NewGRPCInsecureConn(ctx context.Context, opts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	return grpc.DialInsecure(ctx, opts...)
}
