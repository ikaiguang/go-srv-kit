package errorpkg

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const EnumNumberKey = "enum_number"

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

func (e *Error) GetCode() int32 {
	return e.status.Code
}

func (e *Error) GetReason() string {
	return e.status.Reason
}

func (e *Error) GetMessage() string {
	return e.status.Message
}

func (e *Error) GetMetadata() map[string]string {
	return e.status.Metadata
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

type Enum interface {
	String() string
	Number() protoreflect.EnumNumber
}

func Err(code int, e Enum, msg string) *Error {
	return NewWithMetadata(
		code,
		e.String(),
		msg,
		map[string]string{"enum": strconv.Itoa(int(e.Number()))},
	)
}

func Errf(code int, e Enum, format string, a ...interface{}) *Error {
	return NewWithMetadata(
		code,
		e.String(),
		fmt.Sprintf(format, a...),
		map[string]string{EnumNumberKey: strconv.Itoa(int(e.Number()))},
	)
}

// ErrWithMetadata ...
func ErrWithMetadata(code int, enum Enum, message string, md map[string]string) *Error {
	e := errors.New(code, enum.String(), message)
	e.Metadata = md
	if md != nil {
		if _, ok := md[EnumNumberKey]; !ok {
			md[EnumNumberKey] = strconv.Itoa(int(enum.Number()))
		}
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}

// WrapWithEnumNumber enum number
func WrapWithEnumNumber(e *errors.Error, enum Enum) error {
	if e == nil {
		return nil
	}
	if e.Metadata == nil {
		e.Metadata = map[string]string{}
	}
	if _, ok := e.Metadata[EnumNumberKey]; !ok {
		e.Metadata[EnumNumberKey] = strconv.Itoa(int(enum.Number()))
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}
