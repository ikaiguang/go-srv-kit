package errorutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// WithStack returns an error
func WithStack(err error) error {
	return pkgerrors.WithStack(err)
}

// New returns an error object for the code, reason, message.
func New(code int, reason, message string) error {
	return pkgerrors.WithStack(errors.New(code, reason, message))
}

// NewWithMetadata returns an error object for the code, reason, message.
// with an MD formed by the mapping of key, value.
func NewWithMetadata(code int, reason, message string, md map[string]string) error {
	err := errors.New(code, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// Wrap returns an error
// with an MD formed inError
func Wrap(err *errors.Error, eSlice ...error) error {
	if err == nil {
		return nil
	}
	err.Metadata = errorMetadata(eSlice)
	return pkgerrors.WithStack(err)
}

// WrapWithMetadata returns an error
// with an MD formed by the mapping of key, value.
func WrapWithMetadata(err *errors.Error, md map[string]string) error {
	if err == nil {
		return nil
	}
	err = err.WithMetadata(md)
	return pkgerrors.WithStack(err)
}
