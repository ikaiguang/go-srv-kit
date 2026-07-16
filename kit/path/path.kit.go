package pathpkg

import (
	"path/filepath"
	"runtime"
)

// PackageDir returns the directory containing this package's source file.
func PackageDir() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		return "."
	}
	return filepath.Dir(f)
}

// Deprecated: use PackageDir instead.
func Path() string { return PackageDir() }
