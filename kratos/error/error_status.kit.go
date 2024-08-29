package errorpkg

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"io"
)

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

func (e *Error) GetMetadataReason() string {
	if e.status.Metadata == nil {
		return ""
	}
	if v, ok := e.status.Metadata[EnumMetadataKey]; ok {
		return v
	}
	return ""
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, e.Message())
			if e.stack != nil {
				e.stack.Format(s, verb)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Message())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Message())
	}
}

// StackTrace ...
func (e *Error) StackTrace() StackTrace {
	if e.stack == nil {
		return nil
	}
	return e.stack.StackTrace()
}
