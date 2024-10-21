package configtestdata

import (
	"path/filepath"
	"runtime"
)

func CurrentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

func ConfigPath() string {
	return filepath.Join(CurrentPath(), "configs")
}
