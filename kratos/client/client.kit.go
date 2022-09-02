package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdgrpc "google.golang.org/grpc"

	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
)

// NewHTTPClient ...
func NewHTTPClient(ctx context.Context, opts ...http.ClientOption) (*http.Client, error) {
	return http.NewClient(ctx, opts...)
}

// NewGRPCClient ...
func NewGRPCClient(ctx context.Context, insecure bool, opts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	if insecure {
		return grpc.DialInsecure(ctx, opts...)
	}
	return grpc.Dial(ctx, opts...)
}

// NewSampleHTTPClient ...
func NewSampleHTTPClient(ctx context.Context, endpoint string, opts ...http.ClientOption) (*http.Client, error) {
	var httpOpts = []http.ClientOption{
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithResponseDecoder(apputil.ResponseDecoder),
		http.WithEndpoint(endpoint),
	}
	for i := range opts {
		httpOpts = append(httpOpts, opts[i])
	}
	return http.NewClient(ctx, httpOpts...)
}
