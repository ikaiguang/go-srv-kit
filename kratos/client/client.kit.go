package clientpkg

import (
	"context"
	stdjson "encoding/json"
	"io"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	apppkg "github.com/ikaiguang/go-srv-kit/kratos/app"
	stdgrpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
		http.WithResponseDecoder(ResponseDecoder),
		http.WithEndpoint(endpoint),
	}
	for i := range opts {
		httpOpts = append(httpOpts, opts[i])
	}
	return http.NewClient(ctx, httpOpts...)
}

// ResponseDecoder http.DefaultResponseDecoder
func ResponseDecoder(ctx context.Context, res *stdhttp.Response, v interface{}) error {
	defer func() { _ = res.Body.Close() }()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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
		return data.Data.UnmarshalTo(m)
	default:
		unknownData := &apppkg.ResponseData{}
		if err = data.Data.UnmarshalTo(unknownData); err != nil {
			return err
		}
		return stdjson.Unmarshal([]byte(unknownData.Data), v)
	}

	//return http.CodecForResponse(res).Unmarshal(data, v)
}
