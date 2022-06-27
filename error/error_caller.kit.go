package errorutil

import (
	"fmt"
	"runtime"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

const (
	// caller stack trace depth
	_stackTracerDepth = 7
)

// Println 输出错误
func Println(err error) {
	callers := Caller(err)
	fmt.Println(strings.Join(callers, "\n"))
}

// stackTracer errors.StackTrace
type stackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

// Stack ...
func Stack(err error) (callers []string) {
	trace, ok := err.(stackTracer)
	if !ok {
		pc, _, _, _ := runtime.Caller(1)
		callers = []string{
			fmt.Sprintf("%+v", pkgerrors.Frame(pc)),
		}
		return callers
	}

	// stack trace
	st := trace.StackTrace()
	callers = make([]string, len(st))
	for i := range st {
		callers[i] = fmt.Sprintf("%+v", st[i])
	}
	return callers
}

// Caller serializes a caller in file:line format
func Caller(err error) (callers []string) {
	trace, ok := err.(stackTracer)
	if !ok {
		pc, _, _, _ := runtime.Caller(1)
		callers = []string{
			fmt.Sprintf("%+v", pkgerrors.Frame(pc)),
		}
		return callers
	}

	// stack trace
	st := trace.StackTrace()
	if len(st) > _stackTracerDepth {
		st = st[:_stackTracerDepth]
	}
	callers = make([]string, len(st))
	for i := range st {
		callers[i] = fmt.Sprintf("%+v", st[i])
	}
	return callers
}

// CallerWithSkip .
func CallerWithSkip(err error, skip int) (callers []string) {
	callers = Caller(err)
	callCounter := len(callers)
	switch callCounter {
	case 0, 1:
		return callers
	default:

	}
	if skip >= callCounter {
		return callers[callCounter-1:]
	}
	return callers[skip:]
}
