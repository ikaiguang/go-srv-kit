package connectionpkg

import (
	"errors"
	"net"
	"strings"
)

// IsConnectionClosedError reports whether err represents a closed network connection.
func IsConnectionClosedError(err error) bool {
	if errors.Is(err, net.ErrClosed) {
		return true
	}
	var operationError *net.OpError
	if errors.As(err, &operationError) {
		return strings.Contains(operationError.Error(), "use of closed network connection")
	}
	return false
}

// Deprecated: use IsConnectionClosedError instead.
func IsConnCloseErr(err error) bool {
	return IsConnectionClosedError(err)
}
