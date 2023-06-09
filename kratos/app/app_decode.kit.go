package apppkg

import (
	"context"
	stdjson "encoding/json"
	"io"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// DecodeProtobufResponse 解码结果
func DecodeProtobufResponse(contentBody []byte, pbMessage proto.Message) (response *Response, err error) {
	response = &Response{}
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

// DecodeHTTPResponse 解码结果
func DecodeHTTPResponse(contentBody []byte, data interface{}) (response *HTTPResponse, err error) {
	response = &HTTPResponse{
		Data: data,
	}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// DecodeError 解码结果
func DecodeError(contentBody []byte) (response *Response, err error) {
	response = &Response{}
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
	data := &Response{}
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
		unknownData := &ResponseData{}
		if err = data.Data.UnmarshalTo(unknownData); err != nil {
			return err
		}
		return stdjson.Unmarshal([]byte(unknownData.Data), v)
	}

	//return http.CodecForResponse(res).Unmarshal(data, v)
}
