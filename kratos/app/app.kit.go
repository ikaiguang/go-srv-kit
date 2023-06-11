package apppkg

var (
	_appRuntimeEnv = RuntimeEnvEnum_PRODUCTION
)

// SetRuntimeEnv ...
func SetRuntimeEnv(env RuntimeEnvEnum_RuntimeEnv) {
	_appRuntimeEnv = env
}

// GetRuntimeEnv ...
func GetRuntimeEnv() RuntimeEnvEnum_RuntimeEnv {
	return _appRuntimeEnv
}

// IsDebugMode ...
func IsDebugMode() bool {
	return _appRuntimeEnv == RuntimeEnvEnum_LOCAL ||
		_appRuntimeEnv == RuntimeEnvEnum_DEVELOP ||
		_appRuntimeEnv == RuntimeEnvEnum_TESTING
}
