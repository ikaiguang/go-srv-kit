package headerpkg

import "net/http"

// SetContentType 设置 Content-Type。
func SetContentType(header http.Header, contentType string) {
	header.Set(ContentType, contentType)
}

// GetContentType 获取 Content-Type。
func GetContentType(header http.Header) string {
	return header.Get(ContentType)
}

// SetAuthorization 设置 Authorization。
func SetAuthorization(header http.Header, authorization string) {
	header.Set(AuthorizationKey, authorization)
}

// GetAuthorization 获取 Authorization。
func GetAuthorization(header http.Header) string {
	return header.Get(AuthorizationKey)
}

// SetTraceID 设置 TraceID。
func SetTraceID(header http.Header, traceID string) {
	header.Set(TraceID, traceID)
}

// GetTraceID 获取 TraceID。
func GetTraceID(header http.Header) string {
	return header.Get(TraceID)
}
