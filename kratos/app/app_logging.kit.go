package apppkg

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
	"github.com/go-kratos/kratos/v2/transport/http"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
	"github.com/ikaiguang/go-srv-kit/kratos/error"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

var (
	// _maxRequestArgs 设置最大请求参数
	_maxRequestArgs uint = 1024 * 1024
)

// SetMaxRequestArgSize 设置最大请求参数
func SetMaxRequestArgSize(size uint) {
	_maxRequestArgs = size
}

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
	str += " ip=" + `"` + s.ClientIP + `"`
	str += " latency=" + `"` + s.ExecTime.String() + `"`
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
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				requestMessage.Component = tr.Kind().String()
				requestMessage.Operation = tr.Operation()
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
			if httpTr, isHTTP := tr.(http.Transporter); isHTTP {
				requestMessage.Method = httpTr.Request().Method
				requestMessage.Operation = httpTr.Request().URL.String()
				if headerpkg.GetIsWebsocket(httpTr.Request().Header) {
					isWebsocket = true
					requestMessage.Method = "WS"
				}
				requestMessage.ClientIP = contextpkg.ClientIPFromHTTP(ctx, httpTr.Request())
			} else {
				requestMessage.Method = "GRPC"
				requestMessage.ClientIP = contextpkg.ClientIPFromGRPC(ctx)
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
					callers = errorpkg.CallerWithSkip(err, opt.withSkipDepth)
				} else {
					callers = errorpkg.Stack(err)
				}
				if len(callers) > 0 {
					errMessage.Stack = strings.Join(callers, "\n\t")
				}
				// 请求参数
				errMessage.RequestArgs = extractArgs(req)
				if len(errMessage.RequestArgs) > int(_maxRequestArgs) {
					errMessage.RequestArgs = errMessage.RequestArgs[:_maxRequestArgs]
				}

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
