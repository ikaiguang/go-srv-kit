package clientpkg

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/stretchr/testify/require"
)

func TestNewGRPCClientSecurityModes(t *testing.T) {
	for _, insecure := range []bool{true, false} {
		conn, err := NewGRPCClient(
			context.Background(),
			insecure,
			grpc.WithEndpoint("127.0.0.1:1"),
		)
		require.NoError(t, err)
		require.NoError(t, conn.Close())
	}
}
