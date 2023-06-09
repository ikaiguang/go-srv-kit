package errorpkg

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
)

// StatusOK 200
func StatusOK(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusOK, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// BadRequest mapped to a 400 response.
func BadRequest(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusBadRequest, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// Unauthorized mapped to a 401 response.
func Unauthorized(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusUnauthorized, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// Forbidden mapped to a 403 response.
func Forbidden(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusForbidden, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// NotFound mapped to a 404 response.
func NotFound(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusNotFound, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// Conflict mapped to a 409 response.
func Conflict(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusConflict, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// TooManyRequests mapped to a 429 response.
func TooManyRequests(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusTooManyRequests, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// InternalServer mapped to a 500 response.
func InternalServer(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusInternalServerError, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// NotImplemented mapped to a HTTP 501 response.
func NotImplemented(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusNotImplemented, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// ServiceUnavailable mapped to a HTTP 503 response.
func ServiceUnavailable(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusServiceUnavailable, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// GatewayTimeout mapped to a HTTP 504 response.
func GatewayTimeout(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(http.StatusGatewayTimeout, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}

// ClientClosed mapped to a HTTP 499 response.
func ClientClosed(reason, message string, eSlice ...error) *errors.Error {
	e := errors.New(499, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return e
}
