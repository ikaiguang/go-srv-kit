package logmiddle

import (
	"context"
	"fmt"
	"strconv"
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

// RequestMessage 请求信息
type RequestMessage struct {
	Kind      string
	Component string
	Method    string
	Operation string

	ExecTime time.Duration
	ClientIP string
}

// GetRequestInfo 获取服务信息
func (s *RequestMessage) GetRequestInfo() string {
	str := "kind=" + `"` + s.Kind + `"`
	str += " component=" + `"` + s.Component + `"`
	str += " latency=" + `"` + s.ExecTime.String() + `"`
	str += " clientIP=" + `"` + s.ClientIP + `"`
	return str
}

// GetOperationInfo .
func (s *RequestMessage) GetOperationInfo() string {
	str := "method=" + `"` + s.Method + `"`
	str += " operation=" + fmt.Sprintf("%q", s.Operation)
	return str
}

// ErrMessage 响应信息
type ErrMessage struct {
	Code   int32
	Reason string
	Msg    string
	Stack  string

	RequestArgs string
}

// GetErrorDetail ...
func (s *ErrMessage) GetErrorDetail() string {
	message := "code=" + `"` + strconv.FormatInt(int64(s.Code), 10) + `"`
	message += " reason=" + `"` + s.Reason + `"`
	message += " detail=" + fmt.Sprintf("%q", s.Msg)

	return message
}

// options ...
type options struct {
	withSkip      bool
	withSkipDepth int
}

// Option ...
type Option func(options *options)

// WithDefaultSkip ...
func WithDefaultSkip() Option {
	return func(o *options) {
		o.withSkip = true
		o.withSkipDepth = 1
	}
}

// WithCallerSkip ...
func WithCallerSkip(skip int) Option {
	return func(o *options) {
		o.withSkip = true
		o.withSkipDepth = skip
	}
}

// ServerLog 中间件日志
// 参考 logging.Server(logger)
func ServerLog(logger log.Logger, opts ...Option) middleware.Middleware {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				isWebsocket    = false
				loggingLevel   = log.LevelInfo
				requestMessage = &RequestMessage{
					Kind: "server",
				}
				errMessage = &ErrMessage{
					Code: 0,
				}
			)

			// 信息
			if info, ok := transport.FromServerContext(ctx); ok {
				requestMessage.Component = info.Kind().String()
				requestMessage.Operation = info.Operation()
			}

			// 时间
			startTime := time.Now()

			// 执行结果
			reply, err = handler(ctx, req)
			// 错误在最后判断
			//if err != nil {}

			// 执行时间
			requestMessage.ExecTime = time.Since(startTime)

			// request
			if httpContext, isHTTP := contextutil.MatchHTTPContext(ctx); isHTTP {
				requestMessage.Method = httpContext.Request().Method
				requestMessage.Operation = httpContext.Request().URL.String()
				if headerutil.GetIsWebsocket(httpContext.Request().Header) {
					isWebsocket = true
					requestMessage.Method = "WS"
				}
				requestMessage.ClientIP = contextutil.ClientIPFromHTTP(httpContext)
			} else {
				requestMessage.Method = "GRPC"
				requestMessage.ClientIP = contextutil.ClientIPFromGRPC(ctx)
			}

			// 打印日志
			var kv = []interface{}{
				"request", requestMessage.GetRequestInfo(),
				"operation", requestMessage.GetOperationInfo(),
			}

			// websocket 不输出错误
			if isWebsocket {
				_ = log.WithContext(ctx, logger).Log(loggingLevel, kv...)
				return
			}

			// 有错误的
			if err != nil {
				loggingLevel = log.LevelError
				// 错误信息
				errMessage.Msg = err.Error()
				if se := errors.FromError(err); se != nil {
					errMessage.Code = se.Code
					errMessage.Reason = se.Reason
				}
				// 错误调用
				var callers []string
				if opt.withSkip && opt.withSkipDepth > 0 {
					callers = errorutil.CallerWithSkip(err, opt.withSkipDepth)
				} else {
					callers = errorutil.Stack(err)
				}
				if len(callers) > 0 {
					errMessage.Stack = strings.Join(callers, "\n\t")
				}
				// 请求参数
				errMessage.RequestArgs = extractArgs(req)

				// 打印日志
				kv = append(kv,
					"error", errMessage.GetErrorDetail(),
					"args", errMessage.RequestArgs,
					"stack", errMessage.Stack,
				)
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
