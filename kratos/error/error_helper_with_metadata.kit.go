package errorpkg

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// StatusOKWithMetadata 200
func StatusOKWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := StatusOK(reason, message)
	e.Metadata = md
	return e
}

// BadRequestWithMetadata mapped to a 400 response.
func BadRequestWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := BadRequest(reason, message)
	e.Metadata = md
	return e
}

// UnauthorizedWithMetadata mapped to a 401 response.
func UnauthorizedWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := Unauthorized(reason, message)
	e.Metadata = md
	return e
}

// ForbiddenWithMetadata mapped to a 403 response.
func ForbiddenWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := Forbidden(reason, message)
	e.Metadata = md
	return e
}

// NotFoundWithMetadata mapped to a 404 response.
func NotFoundWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := NotFound(reason, message)
	e.Metadata = md
	return e
}

// ConflictWithMetadata mapped to a 409 response.
func ConflictWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := Conflict(reason, message)
	e.Metadata = md
	return e
}

// TooManyRequestsWithMetadata mapped to a 409 response.
func TooManyRequestsWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := TooManyRequests(reason, message)
	e.Metadata = md
	return e
}

// InternalServerWithMetadata mapped to a 500 response.
func InternalServerWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := InternalServer(reason, message)
	e.Metadata = md
	return e
}

// NotImplementedWithMetadata mapped to a HTTP 501 response.
func NotImplementedWithMetadata(reason, message string, md map[string]string) error {
	e := NotImplemented(reason, message)
	e.Metadata = md
	return e
}

// ServiceUnavailableWithMetadata mapped to a HTTP 503 response.
func ServiceUnavailableWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := ServiceUnavailable(reason, message)
	e.Metadata = md
	return e
}

// GatewayTimeoutWithMetadata mapped to a HTTP 504 response.
func GatewayTimeoutWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := GatewayTimeout(reason, message)
	e.Metadata = md
	return e
}

// ClientClosedWithMetadata mapped to a HTTP 499 response.
func ClientClosedWithMetadata(reason, message string, md map[string]string) *errors.Error {
	e := ClientClosed(reason, message)
	e.Metadata = md
	return e
}
