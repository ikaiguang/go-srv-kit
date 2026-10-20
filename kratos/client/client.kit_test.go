package clientpkg

import (
	"context"
	"io"
	stdhttp "net/http"
	"strings"
	"testing"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/v3/app"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"
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

type countingReader struct {
	remaining int64
	read      int64
}

func (r *countingReader) Read(p []byte) (int, error) {
	if r.remaining == 0 {
		return 0, io.EOF
	}
	n := int64(len(p))
	if n > r.remaining {
		n = r.remaining
	}
	r.remaining -= n
	r.read += n
	return int(n), nil
}

func TestResponseDecoderLimitsBodyRead(t *testing.T) {
	body := &countingReader{remaining: DefaultResponseMaxSize * 2}
	response := &stdhttp.Response{Body: io.NopCloser(body)}

	err := ResponseDecoder(context.Background(), response, &struct{}{})
	require.Error(t, err)
	require.LessOrEqual(t, body.read, int64(DefaultResponseMaxSize+1))
}

func TestResponseDecoderRejectsNilResponse(t *testing.T) {
	require.ErrorIs(t, ResponseDecoder(context.Background(), nil, nil), errNilHTTPResponse)
	require.ErrorIs(t, ResponseDecoder(context.Background(), &stdhttp.Response{}, nil), errNilHTTPResponse)
}

func TestResponseDecoderRejectsTypedNilProtoTarget(t *testing.T) {
	data, err := anypb.New(&apppkg.Response{Message: "data"})
	require.NoError(t, err)
	body, err := apppkg.MarshalJSON(&apppkg.Response{Data: data})
	require.NoError(t, err)
	response := &stdhttp.Response{
		Body:   io.NopCloser(strings.NewReader(string(body))),
		Header: stdhttp.Header{"Content-Type": []string{"application/json"}},
	}
	var target *apppkg.Response

	require.NotPanics(t, func() {
		require.ErrorIs(t, ResponseDecoder(context.Background(), response, target), errNilResponseTarget)
	})
}
