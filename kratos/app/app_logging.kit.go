package apppkg

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	stdhttp "net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
	"github.com/ikaiguang/go-srv-kit/kratos/error"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

var (
	// _maxRequestArgs 设置最大请求参数
	_maxRequestArgs   uint  = 1024 * 1024
	_minInfoLevelCode int32 = 1000
)

// SetMaxRequestArgSize 设置最大请求参数
func SetMaxRequestArgSize(size uint) {
	_maxRequestArgs = size
}

type RequestInfoForServer struct {
	Kind      string        `json:"kind"`
	Component string        `json:"component"`
	Latency   time.Duration `json:"latency"`
	ClientIP  string        `json:"client_ip"`
}

func (s *RequestInfoForServer) String() string {
	res := `{`
	res += `"kind":"` + s.Kind + `",`
	res += `"component":"` + s.Component + `",`
	res += `"latency":"` + s.Latency.String() + `",`
	res += `"client_ip":"` + s.ClientIP + `"`
	res += `}`

	return res
}

type OperationInfo struct {
	Method    string `json:"method"`
	Operation string `json:"operation"`
	Args      string `json:"args"`
}

func (s *OperationInfo) String() string {
	res := `{`
	res += `"method":"` + s.Method + `",`
	res += `"operation":` + fmt.Sprintf("%q", s.Operation) + `,`
	res += `"args":` + fmt.Sprintf("%q", s.Args)
	res += `}`

	return res
}

// ErrMessage 响应信息
type ErrMessage struct {
	Code    int32  `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (s *ErrMessage) String() string {
	res := `{`
	res += `"code":"` + strconv.Itoa(int(s.Code)) + `",`
	res += `"reason":"` + s.Reason + `",`
	res += `"message":` + fmt.Sprintf("%q", s.Message)
	res += `}`

	return res
}

// options ...
type options struct {
	withSkipDepth   int
	withTracerDepth int
}

// Option ...
type Option func(options *options)

// WithDefaultSkip ...
func WithDefaultSkip() Option {
	return func(o *options) {
		o.withSkipDepth = 1
	}
}

// WithCallerSkip ...
func WithCallerSkip(skip int) Option {
	return func(o *options) {
		o.withSkipDepth = skip
	}
}

// WithDefaultDepth ...
func WithDefaultDepth() Option {
	return func(o *options) {
		o.withTracerDepth = errorpkg.DefaultStackTracerDepth
	}
}

// WithCallerDepth ...
func WithCallerDepth(depth int) Option {
	return func(o *options) {
		o.withTracerDepth = depth
	}
}

// ServerLog 中间件日志
// 参考 logging.Server(logger)
func ServerLog(logHelper *log.Helper, opts ...Option) middleware.Middleware {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				isWebsocket  = false
				loggingLevel = log.LevelInfo
				requestInfo  = &RequestInfoForServer{
					Kind: "server",
				}
				operationInfo = &OperationInfo{}
				errMessage    = &ErrMessage{
					Code: 0,
				}
			)

			// 信息
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				requestInfo.Component = tr.Kind().String()
				operationInfo.Method = tr.Kind().String()
				operationInfo.Operation = tr.Operation()
			}

			// 时间
			startTime := time.Now()

			// 执行结果
			reply, err = handler(ctx, req)
			// 错误在最后判断
			//if err != nil {}

			// 执行时间
			requestInfo.Latency = time.Since(startTime)

			// request
			if httpTr, isHTTP := tr.(http.Transporter); isHTTP {
				operationInfo.Method = httpTr.Request().Method
				operationInfo.Operation = httpTr.Request().URL.String()
				if headerpkg.GetIsWebsocket(httpTr.Request().Header) {
					isWebsocket = true
					operationInfo.Method = "WS"
				}
				requestInfo.ClientIP = contextpkg.ClientIPFromHTTP(ctx, httpTr.Request())
			} else if _, isGRPC := tr.(*grpc.Transport); isGRPC {
				operationInfo.Method = "GRPC"
				requestInfo.ClientIP = contextpkg.ClientIPFromGRPC(ctx)
			}

			// websocket 不输出错误
			if isWebsocket {
				var kv = []interface{}{
					"operation", operationInfo.String(),
					"request", requestInfo.String(),
				}
				logHelper.WithContext(ctx).Log(loggingLevel, kv...)
				return
			}

			// 请求参数
			args := extractArgs(req)
			if len(args) > int(_maxRequestArgs) {
				args = args[:_maxRequestArgs]
			}
			operationInfo.Args = args

			// 打印日志
			var kv = []interface{}{
				"operation", operationInfo.String(),
				"request", requestInfo.String(),
			}

			// 有错误的
			if err != nil {
				loggingLevel = log.LevelError
				// 错误信息
				errMessage.Message = err.Error()
				se := errorpkg.FromError(err)
				if se != nil {
					errMessage.Code = se.Code
					errMessage.Reason = se.Reason
				}
				if se.GetCode() < stdhttp.StatusInternalServerError || se.GetCode() >= _minInfoLevelCode {
					loggingLevel = log.LevelInfo
				}
				// 错误调用
				var stackMessage string
				var callers = errorpkg.CallStackWithSkipAndDepth(err, opt.withSkipDepth, opt.withTracerDepth)
				if len(callers) > 0 {
					stackMessage = strings.Join(callers, "\n\t")
				}

				// 打印日志
				kv = append(kv,
					"error", errMessage.String(),
					"stack", stackMessage,
				)
			}

			// 输出日志
			logHelper.WithContext(ctx).Log(loggingLevel, kv...)
			return
		}

	}
}

type RequestInfoForClient struct {
	Kind      string        `json:"kind"`
	Component string        `json:"component"`
	Latency   time.Duration `json:"latency"`
	ClientIP  string        `json:"client_ip"`
}

func (s *RequestInfoForClient) String() string {
	res := `{`
	res += `"kind":"` + s.Kind + `",`
	res += `"component":"` + s.Component + `",`
	res += `"latency":"` + s.Latency.String() + `",`
	res += `"client_ip":"` + s.ClientIP + `"`
	res += `}`

	return res
}

// ClientLog is an client logging middleware.
// logging.Client(logger)
func ClientLog(logHelper *log.Helper) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				startTime = time.Now()
				level     = log.LevelInfo
				code      int32
				reason    string
				kind      string
			)

			// 请求信息
			operationInfo := &OperationInfo{}
			tr, ok := transport.FromClientContext(ctx)
			if ok {
				kind = tr.Kind().String()
				operationInfo.Method = tr.Kind().String()
				operationInfo.Operation = tr.Operation()
			}

			// 请求参数
			args := extractArgs(req)
			if len(args) > int(_maxRequestArgs) {
				args = args[:_maxRequestArgs]
			}
			operationInfo.Args = args

			// 请求信息
			requestInfo := &RequestInfoForClient{
				Kind:      "client",
				Component: kind,
				Latency:   time.Since(startTime),
			}

			// request
			if httpTr, isHTTP := tr.(http.Transporter); isHTTP {
				operationInfo.Method = httpTr.Request().Method
				operationInfo.Operation = httpTr.Request().URL.String()
				if headerpkg.GetIsWebsocket(httpTr.Request().Header) {
					operationInfo.Method = "WS"
				}
				//requestInfo.ClientIP = contextpkg.ClientIPFromHTTP(ctx, httpTr.Request())
			} else if _, isGRPC := tr.(*grpc.Transport); isGRPC {
				operationInfo.Method = "GRPC"
				//requestInfo.ClientIP = contextpkg.ClientIPFromGRPC(ctx)
			}

			// log
			var (
				kv = []interface{}{
					"operation", operationInfo.String(),
					"request", requestInfo.String(),
				}
			)

			reply, err = handler(ctx, req)
			if err != nil {
				se := errorpkg.FromError(err)
				if se != nil {
					code = se.Code
					reason = se.Reason
				}
				var stack string
				level, stack = extractError(err)
				if se.GetCode() < stdhttp.StatusInternalServerError || se.GetCode() >= _minInfoLevelCode {
					level = log.LevelInfo
				}
				errMessage := &ErrMessage{
					Code:   code,
					Reason: reason,
					//Message: err.Error(),
				}
				kv = append(kv,
					"error", errMessage.String(),
					"stack", stack,
				)
			}
			logHelper.WithContext(ctx).Log(level, kv...)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if redacter, ok := req.(logging.Redacter); ok {
		return redacter.Redact()
	}
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
