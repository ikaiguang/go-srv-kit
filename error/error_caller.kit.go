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
	if len(callers) <= skip {
		callers = []string{}
		return callers
	}
	return callers[skip:]
}
