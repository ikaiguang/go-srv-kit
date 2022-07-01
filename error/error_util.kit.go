package errorutil

import (
	"net/http"
)

// AcceptedWithMetadata 202
func AcceptedWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusAccepted, reason, message, md)
}

// NoContentWithMetadata 204
func NoContentWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusNoContent, reason, message, md)
}

// BadRequestWithMetadata new BadRequest error that is mapped to a 400 response.
func BadRequestWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusBadRequest, reason, message, md)
}

// UnauthorizedWithMetadata new Unauthorized error that is mapped to a 401 response.
func UnauthorizedWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusUnauthorized, reason, message, md)
}

// ForbiddenWithMetadata new Forbidden error that is mapped to a 403 response.
func ForbiddenWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusForbidden, reason, message, md)
}

// NotFoundWithMetadata new NotFound error that is mapped to a 404 response.
func NotFoundWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusNotFound, reason, message, md)
}

// ConflictWithMetadata new Conflict error that is mapped to a 409 response.
func ConflictWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusConflict, reason, message, md)
}

// ClientClosedWithMetadata new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosedWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(499, reason, message, md)
}

// InternalServerWithMetadata new InternalServer error that is mapped to a 500 response.
func InternalServerWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusInternalServerError, reason, message, md)
}

// NotImplementedWithMetadata new NotImplemented error that is mapped to a HTTP 501 response.
func NotImplementedWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusNotImplemented, reason, message, md)
}

// ServiceUnavailableWithMetadata new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailableWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusServiceUnavailable, reason, message, md)
}

// BadGatewayWithMetadata new BadGateway error that is mapped to a HTTP 502 response.
func BadGatewayWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusBadGateway, reason, message, md)
}

// GatewayTimeoutWithMetadata new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeoutWithMetadata(reason, message string, md map[string]string) error {
	return NewWithMetadata(http.StatusGatewayTimeout, reason, message, md)
}
