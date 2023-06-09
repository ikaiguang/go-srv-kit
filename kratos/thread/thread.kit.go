package threadpkg

import (
	"context"
	"runtime/debug"

	"github.com/ikaiguang/go-srv-kit/kratos/log"
	"go.opentelemetry.io/otel/trace"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	go func() {
		defer Recover()
		fn()
	}()
}

// GoWithCtx 新开启协程执行fn，并将trace-id传递到新的协程
func GoWithCtx(ctx context.Context, fn func(ctx context.Context)) {
	span := trace.SpanFromContext(ctx)
	newCtx := trace.ContextWithSpan(context.Background(), span)
	go func() {
		defer Recover()
		fn(newCtx)
	}()
}

// Recover is used with defer to do cleanup on panics.
// Use it like:
// defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logpkg.Error(p)
		logpkg.Error(string(debug.Stack()))
	}
}
