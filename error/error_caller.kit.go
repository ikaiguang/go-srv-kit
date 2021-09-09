package errorutil

import (
	"fmt"
	"runtime"
	"strconv"

	pkgerrors "github.com/pkg/errors"
)

const (
	// caller stack trace depth
	_stackTracerDepth = 7
)

// stackTracer errors.StackTrace
type stackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

// Caller serializes a caller in file:line format
func Caller(err error) (callers []string) {
	trace, ok := err.(stackTracer)
	if !ok {
		pc, file, line, _ := runtime.Caller(1)
		fn := runtime.FuncForPC(pc)
		funcName := "unknown"
		if fn != nil {
			funcName = fn.Name()
		}
		callers = []string{
			funcName + "\n" + file + ":" + strconv.Itoa(line),
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
