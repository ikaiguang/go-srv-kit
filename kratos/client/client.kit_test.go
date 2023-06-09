package clientpkg

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/require"

	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// go test -v -count=1 ./kratos/client -run=TestNewSampleHTTPClient
func TestNewSampleHTTPClient(t *testing.T) {
	type args struct {
		ctx      context.Context
		endpoint string
		opts     []http.ClientOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#ping-hello",
			args: args{
				ctx:      context.Background(),
				endpoint: "127.0.0.1:50627",
				opts:     nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSampleHTTPClient(tt.args.ctx, tt.args.endpoint, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSampleHTTPClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			pingClient := pingservicev1.NewSrvPingHTTPClient(got)
			reply, err := pingClient.Ping(context.Background(), &pingv1.PingReq{Message: "http"})
			require.NoError(t, err)
			t.Logf("[http] Ping %s\n", reply.Message)
		})
	}
}

// go test -v -count=1 ./kratos/client -run=TestNewHTTPClient
func TestNewHTTPClient(t *testing.T) {
	ctx := context.Background()
	var httpOptions = []http.ClientOption{
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithResponseDecoder(ResponseDecoder),
	}

	httpOptions = append(httpOptions, http.WithEndpoint("127.0.0.1:50627"))

	got, err := NewHTTPClient(ctx, httpOptions...)
	require.NoError(t, err)

	pingClient := pingservicev1.NewSrvPingHTTPClient(got)
	reply, err := pingClient.Ping(context.Background(), &pingv1.PingReq{Message: "endpoint"})
	require.NoError(t, err)
	t.Logf("[http] Ping %s\n", reply.Message)
}
