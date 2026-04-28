package headerpkg

import (
	stdhttp "net/http"
)

// SetRequestID 设置RequestID
func SetRequestID(header stdhttp.Header, requestID string) {
	header.Set(RequestID, requestID)
}

// GetRequestID 获取RequestID
func GetRequestID(header stdhttp.Header) string {
	return header.Get(RequestID)
}
