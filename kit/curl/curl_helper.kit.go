package curlutil

import (
	"fmt"
	"net/http"
)

// IsSuccessCode .
func IsSuccessCode(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}

// ErrRequestFailure ...
func ErrRequestFailure(code int) error {
	return fmt.Errorf("request failure; code=%d", code)
}
