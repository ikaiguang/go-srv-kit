package errorutil

import (
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// AcceptedWithMetadata 202
func AcceptedWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusAccepted, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// NoContentWithMetadata 204
func NoContentWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusNoContent, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// BadRequestWithMetadata new BadRequest error that is mapped to a 400 response.
func BadRequestWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusBadRequest, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// UnauthorizedWithMetadata new Unauthorized error that is mapped to a 401 response.
func UnauthorizedWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusUnauthorized, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// ForbiddenWithMetadata new Forbidden error that is mapped to a 403 response.
func ForbiddenWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusForbidden, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// NotFoundWithMetadata new NotFound error that is mapped to a 404 response.
func NotFoundWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusNotFound, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// ConflictWithMetadata new Conflict error that is mapped to a 409 response.
func ConflictWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusConflict, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// ClientClosedWithMetadata new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosedWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(499, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// InternalServerWithMetadata new InternalServer error that is mapped to a 500 response.
func InternalServerWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusInternalServerError, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// NotImplementedWithMetadata new NotImplemented error that is mapped to a HTTP 501 response.
func NotImplementedWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusNotImplemented, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// ServiceUnavailableWithMetadata new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailableWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusServiceUnavailable, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// BadGatewayWithMetadata new BadGateway error that is mapped to a HTTP 502 response.
func BadGatewayWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusBadGateway, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// GatewayTimeoutWithMetadata new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeoutWithMetadata(reason, message string, md map[string]string) error {
	err := errors.New(http.StatusGatewayTimeout, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}
