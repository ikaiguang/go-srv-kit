package apppkg

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
	"google.golang.org/grpc/codes"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// SetResponseContentType ...
func SetResponseContentType(w stdhttp.ResponseWriter, codec encoding.Codec) {
	switch codec.Name() {
	case json.Name:
		w.Header().Set(headerpkg.ContentType, headerpkg.ContentTypeJSONUtf8)
	default:
		w.Header().Set(headerpkg.ContentType, ContentType(codec.Name()))
	}
}

// HTTPResponseInterface .
type HTTPResponseInterface interface {
	GetCode() int32
	GetReason() string
	GetMessage() string
	GetMetadata() map[string]string
}

// HTTPResponse 响应
// 关联更新 apppkg.Response
type HTTPResponse struct {
	Code     int32             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`

	Data interface{} `json:"data"`
}

func (x *HTTPResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HTTPResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *HTTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HTTPResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// IsSuccessCode 成功的响应码
func IsSuccessCode(code int32) bool {
	return code == OK || IsSuccessHTTPCode(int(code))
}

// IsSuccessHTTPCode 成功的HTTP响应吗
func IsSuccessHTTPCode(code int) bool {
	return code >= stdhttp.StatusOK && code < stdhttp.StatusMultipleChoices
}

// IsSuccessGRPCCode 成功的GRPC响应吗
func IsSuccessGRPCCode(code uint32) bool {
	return codes.Code(code) == codes.OK
}

// ToResponseError 转换为错误
func ToResponseError(response HTTPResponseInterface) *errors.Error {
	return &errors.Error{
		Status: errors.Status{
			Code:     response.GetCode(),
			Reason:   response.GetReason(),
			Message:  response.GetMessage(),
			Metadata: response.GetMetadata(),
		},
	}
}
