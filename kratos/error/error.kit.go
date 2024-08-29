package errorpkg

import (
	"fmt"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
)

const (
	EnumMetadataKey = "reason"
	EnumErrorKey    = "with_error"
)

func Kvs(e *errors.Error, kvs ...string) {
	if e.Metadata == nil {
		e.Metadata = make(map[string]string)
	}
	if len(kvs)%2 != 0 {
		kvs = append(kvs, "KEYVALS UNPAIRED")
	}
	for i := 0; i < len(kvs); i += 2 {
		e.Metadata[kvs[i]] = kvs[i+1]
	}
}

// WithStack returns an error
func WithStack(e *errors.Error) *Error {
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

// Wrap returns an error
func Wrap(e *errors.Error, eSlice ...error) *Error {
	if e == nil {
		return nil
	}
	e.Metadata = errorMetadata(e.Metadata, eSlice)
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

func WrapKvs(e *errors.Error, kvs ...string) *Error {
	if e == nil {
		return nil
	}
	if e.Metadata == nil {
		e.Metadata = make(map[string]string)
	}
	if len(kvs)%2 != 0 {
		kvs = append(kvs, "KEYVALS UNPAIRED")
	}
	for i := 0; i < len(kvs); i += 2 {
		e.Metadata[kvs[i]] = kvs[i+1]
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

// New ...
func New(code int, reason, message string) *Error {
	return &Error{
		status: &status{Error: newError(code, reason, message)},
		stack:  callers(),
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return &Error{
		status: &status{Error: newError(code, reason, fmt.Sprintf(format, a...))},
		stack:  callers(),
	}
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code int, reason, format string, a ...interface{}) *Error {
	return &Error{
		status: &status{Error: newError(code, reason, fmt.Sprintf(format, a...))},
		stack:  callers(),
	}
}

// NewWithMetadata ...
func NewWithMetadata(code int, reason, message string, md map[string]string) *Error {
	e := errors.New(code, reason, message)
	if e.Metadata == nil {
		e.Metadata = md
	} else {
		for k, v := range md {
			e.Metadata[k] = v
		}
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

// WrapWithMetadata returns an error
func WrapWithMetadata(e *errors.Error, md map[string]string) error {
	if e == nil {
		return nil
	}
	if e.Metadata == nil {
		e.Metadata = md
	} else {
		for k, v := range md {
			e.Metadata[k] = v
		}
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

// newError ...
func newError(code int, reason, message string) *errors.Error {
	e := errors.New(code, reason, message)
	e.Metadata = map[string]string{}
	return e
}

// errorMetadata .
func errorMetadata(metadata map[string]string, eSlice []error) map[string]string {
	if metadata == nil {
		metadata = make(map[string]string)
	}
	if len(eSlice) == 0 {
		return metadata
	}

	var (
		errorKey = EnumErrorKey
	)
	for i := range eSlice {
		key := errorKey + "." + strconv.Itoa(i+1)
		if eSlice[i] == nil {
			metadata[key] = "nil"
		} else {
			metadata[key] = eSlice[i].Error()
		}
	}
	return metadata
}
