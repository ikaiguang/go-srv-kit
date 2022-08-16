package apputil

import (
	"context"
	stdjson "encoding/json"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	stdhttp "net/http"

	responsev1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
)

// DecodeProto 解码结果
func DecodeProto(contentBody []byte, pbMessage proto.Message) (response *responsev1.Response, err error) {
	response = &responsev1.Response{}
	err = protojson.Unmarshal(contentBody, response)
	if err != nil {
		return response, err
	}

	// 解密
	if response.Data == nil {
		return response, err
	}
	err = response.Data.UnmarshalTo(pbMessage)
	if err != nil {
		return response, err
	}
	return response, err
}

// DecodeResponse 解码结果
func DecodeResponse(contentBody []byte, data interface{}) (response *Response, err error) {
	response = &Response{
		Data: data,
	}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// DecodeError 解码结果
func DecodeError(contentBody []byte) (response *responsev1.Response, err error) {
	response = &responsev1.Response{}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// ResponseDecoder http.DefaultResponseDecoder
func ResponseDecoder(ctx context.Context, res *stdhttp.Response, v interface{}) error {
	defer func() { _ = res.Body.Close() }()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 解析数据
	data := &responsev1.Response{}
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
		unknownData := &responsev1.Data{}
		if err = data.Data.UnmarshalTo(unknownData); err != nil {
			return err
		}
		return stdjson.Unmarshal([]byte(unknownData.Data), v)
	}

	//return http.CodecForResponse(res).Unmarshal(data, v)
}
