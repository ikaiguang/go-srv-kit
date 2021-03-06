package apputil

import (
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	baseerror "github.com/ikaiguang/go-srv-kit/api/base/error"
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

// IsSuccessCode 成功的响应码
func IsSuccessCode(code int32) bool {
	if code == OK {
		return true
	}
	return IsSuccessHTTPCode(int(code))
}

// IsSuccessHTTPCode 成功的HTTP响应吗
func IsSuccessHTTPCode(code int) bool {
	if code >= stdhttp.StatusOK && code < stdhttp.StatusMultipleChoices {
		return true
	}
	return false
}

// IsSuccessGRPCCode 成功的GRPC响应吗
func IsSuccessGRPCCode(code uint32) bool {
	return codes.Code(code) == codes.OK
}

// ToError 转换为错误
func ToError(response ResponseInterface) *errors.Error {
	return &errors.Error{
		Status: errors.Status{
			Code:     response.GetCode(),
			Reason:   response.GetReason(),
			Message:  response.GetMessage(),
			Metadata: response.GetMetadata(),
		},
	}
}

// HTTPError 转换为错误
func HTTPError(code int, message string) *errors.Error {
	return &errors.Error{
		Status: errors.Status{
			Code:    int32(code),
			Reason:  baseerror.ERROR_STATUS_REQUEST_FAILURE.String(),
			Message: message,
			Metadata: map[string]string{
				"status": stdhttp.StatusText(code),
			},
		},
	}
}
