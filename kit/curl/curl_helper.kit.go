package curlutil

// IsSuccessCode .
func IsSuccessCode(code int) bool {
	return code >= 200 && code < 300
}
