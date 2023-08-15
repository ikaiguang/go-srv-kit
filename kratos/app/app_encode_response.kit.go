package apppkg

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ikaiguang/go-srv-kit/kratos/error"
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
	// w.Header().Set("Content-Type", ContentType(codec.Name()))
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
	//w.Header().Set("Content-Type", ContentType(codec.Name()))
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
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// DefaultResponseDecoder http.DefaultResponseDecoder
func DefaultResponseDecoder(ctx context.Context, res *stdhttp.Response, v interface{}) error {
	return http.DefaultResponseDecoder(ctx, res, v)
}
