package apputil

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/ikaiguang/go-srv-kit/api/response/v1"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// ResponseEncoder http.DefaultResponseEncoder
func ResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	w.WriteHeader(stdhttp.StatusOK)

	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	if headerutil.GetIsWebsocket(r.Header) {
		return nil
	}

	// 响应结果
	data := &Response{
		Code:      OK,
		RequestId: headerutil.GetRequestID(r.Header),
		Data:      v,
	}
	return http.DefaultResponseEncoder(w, r, data)
}

// ErrorEncoder http.DefaultErrorEncoder
func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	w.WriteHeader(stdhttp.StatusOK)

	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	if headerutil.GetIsWebsocket(r.Header) {
		return
	}

	// 响应错误
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")

	data := &v1.Response{
		Code:      se.Code,
		Reason:    se.Reason,
		Message:   se.Message,
		Metadata:  se.Metadata,
		RequestId: headerutil.GetRequestID(r.Header),
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

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}
