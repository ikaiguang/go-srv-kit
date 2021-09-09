package errorutil

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusBadRequest, reason, message))
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusUnauthorized, reason, message))
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusForbidden, reason, message))
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusNotFound, reason, message))
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusConflict, reason, message))
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusInternalServerError, reason, message))
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusServiceUnavailable, reason, message))
}

// BadGateway new BadGateway error that is mapped to a HTTP 502 response.
func BadGateway(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusBadGateway, reason, message))
}

// GatewayTimeout new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeout(reason, message string) error {
	return pkgerrors.WithStack(errors.New(http.StatusGatewayTimeout, reason, message))
}

// ClientClosed new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosed(reason, message string) error {
	return pkgerrors.WithStack(errors.New(499, reason, message))
}
