package apppkg

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

var (
	_ encoding.Codec = (*multipartForm)(nil)
)

const (
	codecNameMultipartForm = "form-data"

	DefaultUploadMaxSize = 10 << 20 // 20M
)

// RegisterCodec ...
func RegisterCodec() {
	encoding.RegisterCodec(&multipartForm{})
}

// RequestDecoder ...
func RequestDecoder(r *http.Request, v interface{}) error {
	// 不解析 multipart/form-data
	//contentType := r.Header.Get(headerpkg.ContentType)
	//if strings.HasPrefix(contentType, headerpkg.ContentTypeMultipartForm) {
	//	return nil
	//}

	// 解码
	codec, ok := http.CodecForRequest(r, "Content-Type")
	if !ok {
		return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
	}
	// 不解析 multipart/form-data : encoding.RegisterCodec(&multipartForm{})
	if codec.Name() == codecNameMultipartForm {
		return nil
	}

	// read data
	data, err := io.ReadAll(r.Body)

	// reset body.
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err = codec.Unmarshal(data, v); err != nil {
		code := errorpkg.ERROR_INVALID_PARAMETER
		reason := code.String()
		message := fmt.Sprintf("body unmarshal %s", err.Error())
		return errorpkg.New(int(code), reason, message)
	}
	return nil
}

// ContentSubtype returns the content-subtype for the given content-type.  The
// given content-type must be a valid content-type that starts with
// but no content-subtype will be returned.
// according rfc7231.
// contentType is assumed to be lowercase already.
func ContentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}

// multipartForm multipart/form-data headerpkg.ContentTypeMultipartForm
type multipartForm struct{}

func (f *multipartForm) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (f *multipartForm) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (f *multipartForm) Name() string {
	return codecNameMultipartForm
}
