package apputil

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	baseerror "github.com/ikaiguang/go-srv-kit/api/base/error"
	responsev1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// ResponseEncoder http.DefaultResponseEncoder
func ResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	if headerutil.GetIsWebsocket(r.Header) {
		return nil
	}
	w.WriteHeader(stdhttp.StatusOK)

	// 响应结果
	data := &responsev1.Response{
		Code:      OK,
		RequestId: headerutil.GetRequestID(r.Header),
		//Data:      v,
	}
	if v != nil {
		if vMessage, ok := v.(proto.Message); ok {
			anyData, err := anypb.New(vMessage)
			if err != nil {
				data.Code = stdhttp.StatusInternalServerError
				data.Reason = baseerror.ERROR_STATUS_NO_CONTENT.String()
				data.Metadata = map[string]string{"error": err.Error()}
			} else {
				data.Data = anyData
			}
		}
	} else {
		data.Code = stdhttp.StatusInternalServerError
		data.Reason = baseerror.ERROR_STATUS_NO_CONTENT.String()
		data.Metadata = map[string]string{"data": "null"}
	}
	return http.DefaultResponseEncoder(w, r, data)
}

// ErrorEncoder http.DefaultErrorEncoder
func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	if headerutil.GetIsWebsocket(r.Header) {
		return
	}
	w.WriteHeader(stdhttp.StatusOK)

	// 响应错误
	se := errors.FromError(err)
	data := &responsev1.Response{
		Code:      se.Code,
		Reason:    se.Reason,
		Message:   se.Message,
		Metadata:  se.Metadata,
		RequestId: headerutil.GetRequestID(r.Header),
	}
	_ = http.DefaultResponseEncoder(w, r, data)
	return
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}
