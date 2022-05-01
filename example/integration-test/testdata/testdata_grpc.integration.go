package testdata

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	stdgrpc "google.golang.org/grpc"
)

// DefaultGRPCConn grpc链接
func DefaultGRPCConn() (*stdgrpc.ClientConn, error) {
	ctx := context.Background()
	var opts = []grpc.ClientOption{
		grpc.WithEndpoint(_defaultGRPCAddr),
	}
	return grpc.DialInsecure(ctx, opts...)
}
