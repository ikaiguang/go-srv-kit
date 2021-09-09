package errorutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// New returns an error object for the code, message.
// @param code == http.Code
func New(code int, reason, message string) error {
	return pkgerrors.WithStack(errors.New(code, reason, message))
}

// Error returns an error object for the code, message.
// @param code == http.Code
func Error(code int, reason, message string) error {
	return pkgerrors.WithStack(errors.New(code, reason, message))
}

// NewWithMetadata returns an error object for the code, message.
// with an MD formed by the mapping of key, value.
// @param code == http.Code
func NewWithMetadata(code int, reason, message string, md map[string]string) error {
	err := errors.New(code, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// FromError try to convert an error to *errors.Error
func FromError(err error) *errors.Error {
	return errors.FromError(Cause(err))
}

// Cause returns the underlying cause of the error
func Cause(err error) error {
	return pkgerrors.Cause(err)
}

// Code returns the http code for a error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200
	}
	if se := FromError(err); err != nil {
		return int(se.Code)
	}
	return errors.UnknownCode
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if se := FromError(err); err != nil {
		return se.Reason
	}
	return errors.UnknownReason
}
