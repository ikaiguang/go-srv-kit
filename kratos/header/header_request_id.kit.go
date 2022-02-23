package headerutil

import (
	"net/http"
)

// SetRequestID 设置RequestID
func SetRequestID(header http.Header, requestID string) {
	header.Set(RequestID, requestID)
}

// GetRequestID 获取RequestID
func GetRequestID(header http.Header) string {
	return header.Get(RequestID)
}
