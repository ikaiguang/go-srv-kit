package connectionpkg

import (
	"net"
	"strings"
)

// IsConnCloseErr .
func IsConnCloseErr(err error) bool {
	if readErr, ok := err.(*net.OpError); ok {
		return strings.Contains(readErr.Error(), "use of closed network connection")
	}
	return false
}
