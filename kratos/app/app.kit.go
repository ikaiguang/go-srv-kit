package apputil

import (
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/codes"
)

const (
	OK = 0

	baseContentType = "application"
)

var (
	_ = http.DefaultRequestDecoder
	_ = http.DefaultResponseEncoder
	_ = http.DefaultErrorEncoder
)

// Response 响应
// 关联更新 responsev1.Response
type Response struct {
	Code     int32             `json:"code"`
	Reason   string            `json:"reason"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`

	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
}

// IsSuccessCode 成功的响应码
func IsSuccessCode(code int32) bool {
	if code == OK {
		return true
	}
	return IsSuccessHTTPCode(code)
}

// IsSuccessHTTPCode 成功的HTTP响应吗
func IsSuccessHTTPCode(code int32) bool {
	if code >= stdhttp.StatusOK && code < stdhttp.StatusMultipleChoices {
		return true
	}
	return false
}

// IsSuccessGRPCCode 成功的GRPC响应吗
func IsSuccessGRPCCode(code uint32) bool {
	return codes.Code(code) == codes.OK
}
