package pathutil

import (
	"path/filepath"
	"runtime"
)

// Path 当前目录
func Path() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		return "."
	}
	return filepath.Dir(f)
}
