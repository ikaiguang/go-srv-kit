package apppkg

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

var (
	_ encoding.Codec = (*multipartForm)(nil)
)

const (
	CodecNameMultipartForm = "form-data"    // 表单提交
	CodecNameOctetStream   = "octet-stream" // 字节流
	DefaultUploadMaxSize   = 20 << 20       // 20M
)

// RegisterCodec ...
func RegisterCodec() {
	encoding.RegisterCodec(&multipartForm{})
	encoding.RegisterCodec(&octetStream{})
}

// RequestDecoder ...
func RequestDecoder(r *http.Request, v interface{}) error {
	// 不解析 multipart/form-data
	//contentType := r.Header.Get(headerpkg.ContentType)
	//if strings.HasPrefix(contentType, headerpkg.ContentTypeMultipartForm) {
	//	return nil
	//}

	// 解码
	codec, ok := http.CodecForRequest(r, headerpkg.ContentType)
	if !ok {
		msg := fmt.Sprintf("[CODEC] unregister Content-Type: %s", r.Header.Get(headerpkg.ContentType))
		e := errorpkg.ErrorBadRequest(msg)
		return errorpkg.WithStack(e)
	}
	// 不解析 multipart/form-data : encoding.RegisterCodec(&multipartForm{})
	if codec.Name() == CodecNameMultipartForm {
		return nil
	}

	// read data
	data, err := io.ReadAll(r.Body)
	if len(data) > DefaultUploadMaxSize {
		e := errorpkg.ErrorRequestEntityTooLarge("REQUEST_ENTITY_TOO_LARGE")
		return errorpkg.WithStack(e)
	}

	// reset body.
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	if err != nil {
		e := errorpkg.ErrorInvalidParameter("[CODEC] invalid request body")
		return errorpkg.Wrap(e, err)
	}
	if len(data) == 0 {
		return nil
	}
	if err = codec.Unmarshal(data, v); err != nil {
		e := errorpkg.ErrorInvalidParameter("[CODEC] unmarshal request body failed")
		return errorpkg.Wrap(e, err)
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

func (f *multipartForm) Name() string {
	return CodecNameMultipartForm
}

func (f *multipartForm) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}
func (f *multipartForm) Unmarshal(data []byte, v interface{}) error {
	return nil
}

type octetStream struct{}

func (s *octetStream) Name() string {
	return CodecNameOctetStream
}
func (s *octetStream) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}
func (s *octetStream) Unmarshal(data []byte, v interface{}) error {
	return nil
}
