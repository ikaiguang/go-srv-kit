package apppkg

// RuntimeEnv ...
type RuntimeEnv string

const (
	RuntimeEnvLocal      RuntimeEnv = "LOCAL"      // 本地开发
	RuntimeEnvDevelop    RuntimeEnv = "DEVELOP"    // 开发环境
	RuntimeEnvTesting    RuntimeEnv = "TESTING"    // 测试环境
	RuntimeEnvPreview    RuntimeEnv = "PREVIEW"    // 预发布环境
	RuntimeEnvProduction RuntimeEnv = "PRODUCTION" // 生产环境
)

var (
	_appRuntimeEnv = RuntimeEnvProduction
)

// SetRuntimeEnv ...
func SetRuntimeEnv(env RuntimeEnv) {
	_appRuntimeEnv = env
}

// GetRuntimeEnv ...
func GetRuntimeEnv() RuntimeEnv {
	return _appRuntimeEnv
}

// IsDebugMode ...
func IsDebugMode() bool {
	return _appRuntimeEnv == RuntimeEnvLocal ||
		_appRuntimeEnv == RuntimeEnvDevelop ||
		_appRuntimeEnv == RuntimeEnvTesting
}
