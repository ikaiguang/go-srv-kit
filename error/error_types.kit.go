package errorutil

import (
	"net/http"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// errorMetadata .
func errorMetadata(eSlice []error) map[string]string {
	if len(eSlice) == 0 {
		return nil
	}

	var (
		metadata = make(map[string]string)
		errorKey = "error"
	)
	if len(eSlice) == 1 {
		metadata[errorKey] = eSlice[0].Error()
		return metadata
	}
	for i := range eSlice {
		if eSlice[i] == nil {
			continue
		}
		key := errorKey + "_" + strconv.Itoa(i)
		metadata[key] = eSlice[i].Error()
	}
	return metadata
}

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusBadRequest, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusUnauthorized, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusForbidden, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusNotFound, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusConflict, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// ClientClosed new ClientClosed error that is mapped to a HTTP 499 response.
func ClientClosed(reason, message string, eSlice ...error) error {
	e := errors.New(499, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusInternalServerError, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// NotImplemented new NotImplemented error that is mapped to a HTTP 501 response.
func NotImplemented(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusNotImplemented, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func ServiceUnavailable(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusServiceUnavailable, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// BadGateway new BadGateway error that is mapped to a HTTP 502 response.
func BadGateway(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusBadGateway, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}

// GatewayTimeout new GatewayTimeout error that is mapped to a HTTP 504 response.
func GatewayTimeout(reason, message string, eSlice ...error) error {
	e := errors.New(http.StatusGatewayTimeout, reason, message)
	e.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(e)
}
