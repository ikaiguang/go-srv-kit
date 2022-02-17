package errorutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	pkgerrors "github.com/pkg/errors"
)

// WithStack returns an error
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	return pkgerrors.WithStack(err)
}

// New returns an error object for the code, reason, message.
func New(code int, reason, message string) error {
	return pkgerrors.WithStack(errors.New(code, reason, message))
}

// NewWithError returns an error object for the code, reason, message.
// with an MD formed inError
func NewWithError(code int, reason, message string, inError error) error {
	err := errors.New(code, reason, message)
	if inError != nil {
		err.Metadata = map[string]string{
			"error": inError.Error(),
		}
	}
	return pkgerrors.WithStack(err)
}

// NewWithMetadata returns an error object for the code, reason, message.
// with an MD formed by the mapping of key, value.
func NewWithMetadata(code int, reason, message string, md map[string]string) error {
	err := errors.New(code, reason, message)
	err.Metadata = md
	return pkgerrors.WithStack(err)
}

// Wrap returns an error
func Wrap(err *errors.Error) error {
	if err == nil {
		return nil
	}
	return pkgerrors.WithStack(err)
}

// WrapWithMetadata returns an error
// with an MD formed inError
func WrapWithError(err *errors.Error, inError error) error {
	if err == nil {
		return nil
	}
	if inError != nil {
		err = err.WithMetadata(map[string]string{
			"error": inError.Error(),
		})
	}
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
