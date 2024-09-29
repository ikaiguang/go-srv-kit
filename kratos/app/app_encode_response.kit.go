package apppkg

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ikaiguang/go-srv-kit/kratos/error"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

const (
	OK = 0

	baseContentType = "application"
)

// SuccessResponseEncoder http.DefaultResponseEncoder
func SuccessResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	//return http.DefaultResponseEncoder(w, r, v)
	//return successResponseEncoder(w, r, v)
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		stdhttp.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	// w.Header().Set(headerpkg.ContentType, ContentType(codec.Name()))
	SetResponseContentType(w, codec)
	w.WriteHeader(stdhttp.StatusOK)

	data, err := codec.Marshal(v)
	if err != nil {
		e := errorpkg.ErrorInternalServer(errorpkg.ERROR_INTERNAL_SERVER.String())
		return errorpkg.Wrap(e, err)
	}
	_, err = w.Write(data)
	if err != nil {
		e := errorpkg.ErrorInternalServer(errorpkg.ERROR_INTERNAL_SERVER.String())
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// ErrorResponseEncoder http.DefaultErrorEncoder
func ErrorResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	// http.DefaultErrorEncoder(w, r, err)
	// errorResponseEncoder(w, r, err)
	codec, _ := http.CodecForRequest(r, "Accept")
	//w.Header().Set(headerpkg.ContentType, ContentType(codec.Name()))
	SetResponseContentType(w, codec)

	se := errorpkg.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	//if !IsDebugMode() {
	//	se.Metadata = nil
	//}
	if se.Code < _minInfoLevelCode {
		w.WriteHeader(int(se.Code))
	} else {
		w.WriteHeader(stdhttp.StatusBadRequest)
	}
	_, _ = w.Write(body)
}

// SuccessResponseDecoder http.DefaultResponseDecoder
func SuccessResponseDecoder(ctx context.Context, res *stdhttp.Response, v interface{}) error {
	return http.DefaultResponseDecoder(ctx, res, v)
}

// ErrorResponseDecode ...
func ErrorResponseDecode(ctx context.Context, res *stdhttp.Response) error {
	return http.DefaultErrorDecoder(ctx, res)
}

// SuccessResponseDecoderBody http.DefaultResponseDecoder
func SuccessResponseDecoderBody(contentType string, data []byte, v interface{}) error {
	codec := encoding.GetCodec(ContentSubtype(contentType))
	if codec == nil {
		codec = encoding.GetCodec(json.Name)
	}
	return codec.Unmarshal(data, v)
}

// SuccessResponseDecoderBody2 http.DefaultResponseDecoder
func SuccessResponseDecoderBody2(header stdhttp.Header, data []byte, v interface{}) error {
	return SuccessResponseDecoderBody(header.Get(headerpkg.ContentType), data, v)
}

// ErrorResponseDecodeBody ...
func ErrorResponseDecodeBody(contentType string, statusCode int, data []byte) (*errors.Error, error) {
	e := new(errors.Error)
	codec := encoding.GetCodec(ContentSubtype(contentType))
	if codec == nil {
		codec = encoding.GetCodec(json.Name)
	}
	if err := codec.Unmarshal(data, e); err != nil {
		e = errorpkg.ErrorInternalServer("[CODEC] unmarshal failed")
		return nil, errorpkg.Wrap(e, err)
	}
	e.Code = int32(statusCode)
	return e, nil
}

// ErrorResponseDecodeBody2 ...
func ErrorResponseDecodeBody2(header stdhttp.Header, statusCode int, data []byte) (*errors.Error, error) {
	return ErrorResponseDecodeBody(header.Get(headerpkg.ContentType), statusCode, data)
}
