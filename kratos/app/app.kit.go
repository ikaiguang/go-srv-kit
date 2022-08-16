package apputil

import (
	stdhttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	envv1 "github.com/ikaiguang/go-srv-kit/api/env/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"

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

// ID ...
func ID(appConfig *confv1.App) string {
	identifier := appConfig.Name
	identifier += ":" + ParseEnv(appConfig.Env).String()
	if appConfig.EnvBranch != "" {
		branchString := strings.Replace(appConfig.EnvBranch, " ", ":", -1)
		identifier += ":" + branchString
	}
	if appConfig.Version != "" {
		identifier += ":" + appConfig.Version
	}
	return identifier
}

// ParseEnv ...
func ParseEnv(appEnv string) (envEnum envv1.Env) {
	envInt32, ok := envv1.Env_value[strings.ToUpper(appEnv)]
	if ok {
		envEnum = envv1.Env(envInt32)
	}
	if envEnum == envv1.Env_UNKNOWN {
		envEnum = envv1.Env_PRODUCTION
		return envEnum
	}
	return envEnum
}

// IsDebugMode ...
func IsDebugMode(appEnv envv1.Env) bool {
	switch appEnv {
	case envv1.Env_DEVELOP, envv1.Env_TESTING:
		return true
	default:
		return false
	}
}

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
			Reason:  errorv1.ERROR_REQUEST_FAILED.String(),
			Message: message,
			Metadata: map[string]string{
				"status": stdhttp.StatusText(code),
			},
		},
	}
}
