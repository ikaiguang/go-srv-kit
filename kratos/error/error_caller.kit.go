package errorpkg

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	// DefaultStackTracerDepth 错误追踪最深层数
	DefaultStackTracerDepth = 10
)

// Println 输出错误
func Println(err error) {
	callers := Caller(err)
	fmt.Println(strings.Join(callers, "\n"))
}

// stackTracer errors.StackTrace
type stackTracer interface {
	StackTrace() StackTrace
}

// Stack 序列化调用者：file:line
func Stack(err error) (callers []string) {
	trace, ok := err.(stackTracer)
	if !ok {
		pc, _, _, _ := runtime.Caller(1)
		callers = []string{
			fmt.Sprintf("%+v", Frame(pc)),
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

// Caller 序列化调用者：file:line
func Caller(err error) (callers []string) {
	trace, ok := err.(stackTracer)
	if !ok {
		pc, _, _, _ := runtime.Caller(1)
		callers = []string{
			fmt.Sprintf("%+v", Frame(pc)),
		}
		return callers
	}

	// stack trace
	st := trace.StackTrace()
	if len(st) > DefaultStackTracerDepth {
		st = st[:DefaultStackTracerDepth]
	}
	callers = make([]string, len(st))
	for i := range st {
		callers[i] = fmt.Sprintf("%+v", st[i])
	}
	return callers
}

// CallerWithSkip 序列化调用者：file:line
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
