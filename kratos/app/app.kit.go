package apppkg

import (
	"strings"
	"sync/atomic"

	"github.com/go-kratos/kratos/v3/transport/http"
)

var (
	_ = http.DefaultRequestDecoder
	_ = http.DefaultErrorEncoder
	_ = http.DefaultResponseEncoder
	_ = http.DefaultResponseDecoder
)

var appRuntimeEnv atomic.Int32

func init() {
	appRuntimeEnv.Store(int32(RuntimeEnvEnum_PRODUCTION))
}

// SetRuntimeEnv ...
func SetRuntimeEnv(env RuntimeEnvEnum_RuntimeEnv) {
	appRuntimeEnv.Store(int32(env))
}

// GetRuntimeEnv ...
func GetRuntimeEnv() RuntimeEnvEnum_RuntimeEnv {
	return RuntimeEnvEnum_RuntimeEnv(appRuntimeEnv.Load())
}

// IsDebugMode ...
func IsDebugMode() bool {
	env := GetRuntimeEnv()
	return env == RuntimeEnvEnum_LOCAL ||
		env == RuntimeEnvEnum_DEVELOP ||
		env == RuntimeEnvEnum_TESTING
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
