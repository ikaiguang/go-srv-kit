package apppkg

import (
	"context"
	stdjson "encoding/json"
	"io"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// CustomSuccessResponseEncoder http.DefaultResponseEncoder
func CustomSuccessResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	//if headerpkg.GetIsWebsocket(r.Header) {
	//	return nil
	//}

	// nil
	if v == nil {
		//respData := &Response{
		//	Code:      OK,
		//	RequestId: headerpkg.GetRequestID(r.Header),
		//	//Data:      v,
		//}
		//respData.Code = stdhttp.StatusInternalServerError
		//respData.Reason = errorv1.ERROR_NO_CONTENT.String()
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
	respData := &Response{
		Code: OK,
		//RequestId: headerpkg.GetRequestID(r.Header),
		//Data:      v,
	}
	var resultMessage proto.Message
	if vMessage, ok := v.(proto.Message); ok {
		// message
		resultMessage = vMessage
	} else {
		// unknown
		vBytes, _ := stdjson.Marshal(v)
		resultMessage = &ResponseData{
			Data: string(vBytes),
		}
	}
	anyData, err := anypb.New(resultMessage)
	if err != nil {
		respData.Code = stdhttp.StatusInternalServerError
		respData.Reason = errorpkg.ERROR_INTERNAL_SERVER.String()
		respData.Metadata = map[string]string{"error": err.Error()}
	} else {
		respData.Data = anyData
	}

	// return
	codec, _ := http.CodecForRequest(r, "Accept")
	SetResponseContentType(w, codec)
	w.WriteHeader(stdhttp.StatusOK)

	// return
	dataBytes, err := codec.Marshal(respData)
	if err != nil {
		e := errorpkg.ErrorInternalServer(errorpkg.ERROR_INTERNAL_SERVER.String())
		return errorpkg.Wrap(e, err)
	}
	_, err = w.Write(dataBytes)
	if err != nil {
		e := errorpkg.ErrorInternalServer(errorpkg.ERROR_INTERNAL_SERVER.String())
		return errorpkg.Wrap(e, err)
	}
	return nil

	// 参考
	//return http.DefaultResponseEncoder(w, r, respData)
}

// CustomErrorResponseEncoder http.DefaultErrorEncoder
func CustomErrorResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	// 在websocket时日志干扰：http: superfluous response.WriteHeader call from xxx(file:line)
	// 在websocket时日志干扰：http: response.Write on hijacked connection from
	// is websocket
	//if headerpkg.GetIsWebsocket(r.Header) {
	//	return
	//}

	// 响应错误
	se := errorpkg.FromError(err)
	data := &Response{
		Code:     se.Code,
		Reason:   se.Reason,
		Message:  se.Message,
		Metadata: se.Metadata,
		//RequestId: headerpkg.GetRequestID(r.Header),
	}
	if !IsDebugMode() {
		data.Metadata = nil
	}

	codec, _ := http.CodecForRequest(r, "Accept")
	SetResponseContentType(w, codec)

	// // return
	//body, err := codec.Marshal(se)
	body, err := codec.Marshal(data)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	//w.WriteHeader(stdhttp.StatusOK)
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)

	// 参考
	//_ = http.DefaultResponseEncoder(w, r, data)
	//http.DefaultErrorEncoder(w, r, err)
	return
}

// CustomResponseDecoder http.DefaultResponseDecoder
func CustomResponseDecoder(ctx context.Context, res *stdhttp.Response, v interface{}) error {
	//return http.CodecForResponse(res).Unmarshal(data, v)
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
}

// CustomDecodeProtobufResponse 解码结果
func CustomDecodeProtobufResponse(contentBody []byte, pbMessage proto.Message) (response *Response, err error) {
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

// CustomDecodeHTTPResponse 解码结果
func CustomDecodeHTTPResponse(contentBody []byte, data interface{}) (response *HTTPResponse, err error) {
	response = &HTTPResponse{
		Data: data,
	}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}

// CustomDecodeResponseError 解码结果
func CustomDecodeResponseError(contentBody []byte) (response *Response, err error) {
	response = &Response{}
	err = UnmarshalJSON(contentBody, response)
	if err != nil {
		return response, err
	}
	return response, err
}
