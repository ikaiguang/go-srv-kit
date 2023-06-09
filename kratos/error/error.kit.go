package errorpkg

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
)

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
	e.Metadata = errorMetadata(eSlice)
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
	e.Metadata = md
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
	e = e.WithMetadata(md)
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
func errorMetadata(eSlice []error) map[string]string {
	var (
		metadata = make(map[string]string)
		errorKey = "error"
	)
	if len(eSlice) == 0 {
		return metadata
	}
	if len(eSlice) == 1 {
		if eSlice[0] != nil {
			metadata[errorKey] = eSlice[0].Error()
		}
		return metadata
	}
	for i := range eSlice {
		key := errorKey + "." + strconv.Itoa(i)
		if eSlice[i] == nil {
			metadata[key] = ""
		} else {
			metadata[key] = eSlice[i].Error()
		}
	}
	return metadata
}

type status struct {
	*errors.Error
}

// Error ...
type Error struct {
	*status
	stack *stack
}

func (e *Error) Error() string {
	if e.status == nil || e.status.Error == nil {
		return ""
	}
	return e.status.Error.Error()
}

func (e *Error) Message() string {
	return e.Error()
}

func (e *Error) Cause() error {
	return e.status.Error
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.Message())
			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Message())
	case 'q':
		fmt.Fprintf(s, "%q", e.Message())
	}
}

// StackTrace ...
func (e *Error) StackTrace() StackTrace {
	if e.stack == nil {
		return nil
	}
	return e.stack.StackTrace()
}
