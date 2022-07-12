package errorutil

import (
	"runtime"
)

// RecoverStack ...
func RecoverStack() string {
	buf := make([]byte, 64<<10) //nolint:gomnd
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
