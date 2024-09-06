package errorpkg

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type Enum interface {
	String() string
	Number() protoreflect.EnumNumber
	HTTPCode() int
}

func Err(e Enum, msg string) *Error {
	return NewWithMetadata(
		e.HTTPCode(),
		e.String(),
		msg,
		map[string]string{EnumMDReasonKey: strconv.Itoa(int(e.Number()))},
	)
}

func Errf(e Enum, format string, a ...interface{}) *Error {
	return NewWithMetadata(
		e.HTTPCode(),
		e.String(),
		fmt.Sprintf(format, a...),
		map[string]string{EnumMDReasonKey: strconv.Itoa(int(e.Number()))},
	)
}

// ErrWithMetadata ...
func ErrWithMetadata(enum Enum, message string, md map[string]string) *Error {
	e := errors.New(enum.HTTPCode(), enum.String(), message)
	e.Metadata = md
	if md != nil {
		if _, ok := md[EnumMDReasonKey]; !ok {
			md[EnumMDReasonKey] = strconv.Itoa(int(enum.Number()))
		}
	}
	return &Error{
		status: &status{Error: e},
		stack:  callers(),
	}
}
