package apputil

import (
	stdjson "encoding/json"
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding/json"
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

	// nil
	if v == nil {
		//respData := &responsev1.Response{
		//	Code:      OK,
		//	RequestId: headerutil.GetRequestID(r.Header),
		//	//Data:      v,
		//}
		//respData.Code = stdhttp.StatusInternalServerError
		//respData.Reason = baseerror.ERROR_STATUS_NO_CONTENT.String()
		//respData.Metadata = map[string]string{"data": "null"}
		return nil
	}

	// 响应
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		stdhttp.Redirect(w, r, url, code)
		return nil
	}

	// 响应结果
	respData := &responsev1.Response{
		Code:      OK,
		RequestId: headerutil.GetRequestID(r.Header),
		//Data:      v,
	}
	var resultMessage proto.Message
	if vMessage, ok := v.(proto.Message); ok {
		// message
		resultMessage = vMessage
	} else {
		// unknown
		vBytes, _ := stdjson.Marshal(v)
		resultMessage = &responsev1.Data{
			Data: string(vBytes),
		}
	}
	anyData, err := anypb.New(resultMessage)
	if err != nil {
		respData.Code = stdhttp.StatusInternalServerError
		respData.Reason = baseerror.ERROR_STATUS_NO_CONTENT.String()
		respData.Metadata = map[string]string{"error": err.Error()}
	} else {
		respData.Data = anyData
	}

	// return
	codec, _ := http.CodecForRequest(r, "Accept")
	dataBytes, err := codec.Marshal(respData)
	if err != nil {
		return err
	}
	switch codec.Name() {
	case json.Name:
		w.Header().Set("Content-Type", headerutil.ContentTypeJSONUtf8)
	default:
		w.Header().Set("Content-Type", ContentType(codec.Name()))
	}
	w.WriteHeader(stdhttp.StatusOK)
	_, err = w.Write(dataBytes)
	if err != nil {
		return err
	}
	return nil

	// 参考
	//return http.DefaultResponseEncoder(w, r, respData)
}

// ErrorEncoder http.DefaultErrorEncoder
func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	if headerutil.GetIsWebsocket(r.Header) {
		return
	}

	// 响应错误
	se := errors.FromError(err)
	data := &responsev1.Response{
		Code:      se.Code,
		Reason:    se.Reason,
		Message:   se.Message,
		Metadata:  se.Metadata,
		RequestId: headerutil.GetRequestID(r.Header),
	}

	codec, _ := http.CodecForRequest(r, "Accept")
	//body, err := codec.Marshal(se)
	body, err := codec.Marshal(data)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		return
	}
	switch codec.Name() {
	case json.Name:
		w.Header().Set("Content-Type", headerutil.ContentTypeJSONUtf8)
	default:
		w.Header().Set("Content-Type", ContentType(codec.Name()))
	}
	w.WriteHeader(stdhttp.StatusOK)
	//w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)

	// 参考
	//_ = http.DefaultResponseEncoder(w, r, data)
	//http.DefaultErrorEncoder(w, r, err)
	return
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}
