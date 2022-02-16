package middlewareutil

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"

	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

// ErrorStack 错误跟踪
func ErrorStack(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err == nil {
				return reply, err
			}

			//errorutil.Println(err)
			callers := errorutil.CallerWithSkip(err, 1)
			_ = log.WithContext(ctx, logger).Log(
				log.LevelError,
				"msg", err.Error(),
				"error stack", strings.Join(callers, "\n\t"),
			)
			return
		}
	}
}
