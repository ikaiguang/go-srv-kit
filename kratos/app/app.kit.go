package apputil

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

const (
	baseContentType = "application"
)

var (
	_ = http.DefaultRequestDecoder
)

// Response 响应
// 关联更新 v1.Response
type Response struct {
	Code     int32             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`

	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

// ResponseEncoder http.DefaultResponseEncoder
func ResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	data := &Response{
		Code:      0,
		RequestId: w.Header().Get(headerutil.RequestID),
		Data:      v,
	}
	return http.DefaultResponseEncoder(w, r, data)
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// ErrorEncoder http.DefaultErrorEncoder
func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")

	data := &v1.Response{
		Code:      se.Code,
		Reason:    se.Reason,
		Message:   se.Message,
		Metadata:  se.Metadata,
		RequestId: w.Header().Get(headerutil.RequestID),
	}
	body, err := codec.Marshal(data)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	//w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// ResponseDecoder 解码结果
func ResponseDecoder(contentBody []byte, data interface{}) (response *Response, err error) {
	response = &Response{
		Data: data,
	}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// ErrorDecoder 解码结果
func ErrorDecoder(contentBody []byte) (response *v1.Response, err error) {
	response = &v1.Response{}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}
