package apppkg

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	stdhttp "net/http"
)

var (
	_ = http.NotFoundHandler
)

func NotFound404() http.ServerOption {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/", func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		msg := "404 page not found"
		e := errorpkg.ErrorNotFound(msg)
		codec, _ := http.CodecForRequest(r, "Accept")
		// w.Header().Set(headerpkg.ContentType, ContentType(codec.Name()))
		SetResponseContentType(w, codec)
		w.WriteHeader(stdhttp.StatusNotFound)
		data, err := codec.Marshal(e)
		if err != nil {
			_, _ = w.Write([]byte(msg))
		}
		_, _ = w.Write(data)
	})
	return http.NotFoundHandler(mux)
}
