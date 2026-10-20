package clientpkg

import (
	"context"
	"crypto/tls"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"io"
	stdhttp "net/http"
	"reflect"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/v3/app"
	stdgrpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const DefaultResponseMaxSize = 20 << 20

var errNilHTTPResponse = errors.New("HTTP response or body is nil")
var errNilResponseTarget = errors.New("response target is nil")

// NewHTTPClient ...
func NewHTTPClient(ctx context.Context, opts ...http.ClientOption) (*http.Client, error) {
	return http.NewClient(ctx, opts...)
}

// NewGRPCClient ...
func NewGRPCClient(ctx context.Context, insecure bool, opts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	if insecure {
		opts = append(opts, grpc.WithTLSConfig(nil))
	} else {
		tlsOpt := grpc.WithTLSConfig(&tls.Config{MinVersion: tls.VersionTLS12})
		opts = append([]grpc.ClientOption{tlsOpt}, opts...)
	}
	return grpc.NewClient(ctx, opts...)
}

// NewSampleHTTPClient ...
func NewSampleHTTPClient(ctx context.Context, endpoint string, opts ...http.ClientOption) (*http.Client, error) {
	var httpOpts = []http.ClientOption{
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithResponseDecoder(ResponseDecoder),
		http.WithEndpoint(endpoint),
	}
	for i := range opts {
		httpOpts = append(httpOpts, opts[i])
	}
	return http.NewClient(ctx, httpOpts...)
}

// ResponseDecoder http.DefaultResponseDecoder
func ResponseDecoder(_ context.Context, res *stdhttp.Response, v any) error {
	if res == nil || res.Body == nil {
		return errNilHTTPResponse
	}
	defer func() { _ = res.Body.Close() }()
	bodyBytes, err := io.ReadAll(io.LimitReader(res.Body, DefaultResponseMaxSize+1))
	if err != nil {
		return err
	}
	if len(bodyBytes) > DefaultResponseMaxSize {
		return fmt.Errorf("response body exceeds %d bytes", DefaultResponseMaxSize)
	}

	// 解析数据
	data := &apppkg.Response{}
	if err = http.CodecForResponse(res).Unmarshal(bodyBytes, data); err != nil {
		return err
	}

	// 解密
	if data.Data == nil {
		return nil
	}
	switch m := v.(type) {
	case proto.Message:
		value := reflect.ValueOf(m)
		if value.Kind() == reflect.Pointer && value.IsNil() {
			return errNilResponseTarget
		}
		return data.Data.UnmarshalTo(m)
	default:
		unknownData := &apppkg.ResponseData{}
		if err = data.Data.UnmarshalTo(unknownData); err != nil {
			return err
		}
		return stdjson.Unmarshal([]byte(unknownData.Data), v)
	}
}
