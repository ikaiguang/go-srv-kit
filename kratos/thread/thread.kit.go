package threadpkg

import (
	"context"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
	"github.com/ikaiguang/go-srv-kit/kratos/log"
	"runtime"
)

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn func()) {
	go func() {
		defer Recover(context.Background())
		fn()
	}()
}

// GoSafeWithContext ...
func GoSafeWithContext(ctx context.Context, fn func(ctx context.Context)) {
	newCtx := contextpkg.NewContext(ctx)
	go func() {
		defer Recover(ctx)
		fn(newCtx)
	}()
}

// Recover is used with defer to do cleanup on panics.
// Use it like: defer Recover(func() {})
func Recover(ctx context.Context) {
	if rerr := recover(); rerr != nil {
		buf := make([]byte, 64<<10) //nolint:mnd
		n := runtime.Stack(buf, false)
		buf = buf[:n]
		logpkg.ErrorfWithContext(ctx, "threadpkg.Recover: %+v\n%s\n", rerr, buf)
	}
}
