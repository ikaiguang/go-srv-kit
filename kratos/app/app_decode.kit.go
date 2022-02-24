package apputil

import (
	v1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
)

// DecodeSuccessResponse 解码结果
func DecodeSuccessResponse(contentBody []byte, data interface{}) (response *Response, err error) {
	response = &Response{
		Data: data,
	}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// DecodeErrorResponse 解码结果
func DecodeErrorResponse(contentBody []byte) (response *v1.Response, err error) {
	response = &v1.Response{}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}
