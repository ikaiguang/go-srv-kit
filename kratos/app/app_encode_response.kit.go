package apppkg

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

const (
	OK = 0

	baseContentType = "application"
)

// SuccessResponseEncoder http.DefaultResponseEncoder
func SuccessResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v any) error {
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
		e := errorpkg.ErrorInternalServer("marshal response failed")
		return errorpkg.Wrap(e, err)
	}
	_, err = w.Write(data)
	if err != nil {
		e := errorpkg.ErrorInternalServer("write response failed")
		return errorpkg.Wrap(e, err)
	}
	return nil
}

// ErrorResponseEncoder http.DefaultErrorEncoder
func ErrorResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	SetResponseContentType(w, codec)

	se := errorpkg.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if se.Code < _minInfoLevelCode {
		w.WriteHeader(int(se.Code))
	} else {
		w.WriteHeader(stdhttp.StatusBadRequest)
	}
	_, _ = w.Write(body)
}

// SuccessResponseDecoder http.DefaultResponseDecoder
func SuccessResponseDecoder(ctx context.Context, res *stdhttp.Response, v any) error {
	return http.DefaultResponseDecoder(ctx, res, v)
}

// ErrorResponseDecode ...
func ErrorResponseDecode(ctx context.Context, res *stdhttp.Response) error {
	return http.DefaultErrorDecoder(ctx, res)
}

// SuccessResponseDecoderBody http.DefaultResponseDecoder
func SuccessResponseDecoderBody(contentType string, data []byte, v any) error {
	codec := encoding.GetCodec(ContentSubtype(contentType))
	if codec == nil {
		codec = encoding.GetCodec(json.Name)
	}
	return codec.Unmarshal(data, v)
}

// SuccessResponseDecoderBody2 http.DefaultResponseDecoder
func SuccessResponseDecoderBody2(header stdhttp.Header, data []byte, v any) error {
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
