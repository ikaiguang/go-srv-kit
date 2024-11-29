package ospkg

import "runtime"

// IsWindows 判断是否是windows系统
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
