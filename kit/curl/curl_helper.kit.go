package curlutil

import (
	"net/http"
)

// IsSuccessCode .
func IsSuccessCode(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}
