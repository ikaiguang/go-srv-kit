package apppkg

import (
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	_ = http.DefaultRequestDecoder
	_ = http.DefaultErrorEncoder
	_ = http.DefaultResponseEncoder
	_ = http.DefaultResponseDecoder
)

var (
	appRuntimeEnv = RuntimeEnvEnum_PRODUCTION
)

// SetRuntimeEnv ...
func SetRuntimeEnv(env RuntimeEnvEnum_RuntimeEnv) {
	appRuntimeEnv = env
}

// GetRuntimeEnv ...
func GetRuntimeEnv() RuntimeEnvEnum_RuntimeEnv {
	return appRuntimeEnv
}

// IsDebugMode ...
func IsDebugMode() bool {
	return appRuntimeEnv == RuntimeEnvEnum_LOCAL ||
		appRuntimeEnv == RuntimeEnvEnum_DEVELOP ||
		appRuntimeEnv == RuntimeEnvEnum_TESTING
}

// IsLocalMode ...
func IsLocalMode() bool {
	return GetRuntimeEnv() == RuntimeEnvEnum_LOCAL
}

// ParseEnv ...
func ParseEnv(appEnv string) RuntimeEnvEnum_RuntimeEnv {
	envInt32, ok := RuntimeEnvEnum_RuntimeEnv_value[strings.ToUpper(appEnv)]
	if !ok {
		return RuntimeEnvEnum_PRODUCTION
	}
	env := RuntimeEnvEnum_RuntimeEnv(envInt32)
	if env == RuntimeEnvEnum_UNKNOWN {
		env = RuntimeEnvEnum_PRODUCTION
	}
	return env
}
