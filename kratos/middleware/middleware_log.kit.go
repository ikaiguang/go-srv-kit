package middlewareutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	errorutil "github.com/ikaiguang/go-srv-kit/error"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// ServerLog is an server logging middleware. 日志
// logging.Server(logger)
func ServerLog(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)

			// 时间
			startTime := time.Now()

			// 信息
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}

			// 执行结果
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}

			// logging
			var (
				loggingLevel = log.LevelInfo
				kv           = []interface{}{
					"kind", "server",
					"component", kind,
					"latency", time.Since(startTime).Seconds(),
				}
			)

			// request
			var isWebsocket = false
			if httpContext, isHTTP := contextutil.MatchHTTPContext(ctx); isHTTP {
				method := httpContext.Request().Method
				if headerutil.GetIsWebsocket(httpContext.Request().Header) {
					isWebsocket = true
					method = "WS"
				}
				kv = append(kv, "operation", method+" "+httpContext.Request().URL.String())
			} else {
				kv = append(kv, "operation", operation)
			}

			// websocket 不输出错误
			if isWebsocket {
				_ = log.WithContext(ctx, logger).Log(loggingLevel, kv...)
				return
			}

			// args
			if err != nil {
				loggingLevel = log.LevelError
				kv = append(kv,
					"code", code,
					"reason", reason,
					"error", err.Error(),
					"args", extractArgs(req),
				)
				callers := errorutil.CallerWithSkip(err, 1)
				if len(callers) > 1 {
					kv = append(kv, "stack", strings.Join(callers, "\n\t"))
				} else {
					kv = append(kv, "stack", "")
				}
			}

			// 输出日志
			_ = log.WithContext(ctx, logger).Log(loggingLevel, kv...)
			return
		}
	}
}

// ClientLog is an client logging middleware.
// logging.Client(logger)
func ClientLog(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				reason    string
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				reason = se.Reason
			}
			level, stack := extractError(err)
			_ = log.WithContext(ctx, logger).Log(level,
				"kind", "client",
				"component", kind,
				"operation", operation,
				"args", extractArgs(req),
				"code", code,
				"reason", reason,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
