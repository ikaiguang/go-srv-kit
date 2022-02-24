package apputil

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	responsev1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
)

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
