package apppkg

import (
	"context"
	"fmt"
	"log/slog"
	stdhttp "net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	headerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/header"
	ippkg "github.com/ikaiguang/go-srv-kit/kit/v3/ip"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/context"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/error"
)

const (
	defaultMaxRequestArgs = 1024 * 1024
	_minInfoLevelCode     = 1000
)

var maxRequestArgs atomic.Uint64

func init() {
	maxRequestArgs.Store(defaultMaxRequestArgs)
}

// SetMaxRequestArgSize 设置最大请求参数
func SetMaxRequestArgSize(size uint) {
	maxRequestArgs.Store(uint64(size))
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

func (s *RequestInfoForServer) Kvs() []any {
	return []any{
		"kind", s.Kind,
		"component", s.Component,
		"latency", s.Latency.String(),
		"client_ip", s.ClientIP,
	}
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

func (s *OperationInfo) Kvs() []any {
	return []any{
		"method", s.Method,
		"operation", s.Operation,
		"args", s.Args,
	}
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

func (s *ErrMessage) Kvs() []any {
	return []any{
		"code", s.Code,
		"reason", s.Reason,
		"message", s.Message,
	}
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

func ServerLoggingKvs(requestInfo *RequestInfoForServer, operationInfo *OperationInfo) []any {
	return []any{
		"client_ip", requestInfo.ClientIP,
		"latency", requestInfo.Latency,
		"kind", requestInfo.Kind,
		"component", requestInfo.Component,
		"method", operationInfo.Method,
		"operation", operationInfo.Operation,
		"args", operationInfo.Args,
	}
}

// ServerLog 中间件日志
// 参考 logging.Server(logger)
func ServerLog(logger *slog.Logger, opts ...Option) middleware.Middleware {
	if logger == nil {
		logger = slog.Default()
	}
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			var (
				isWebsocket  = false
				loggingLevel = slog.LevelInfo
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
				request := httpTr.Request()
				if request != nil {
					operationInfo.Method = request.Method
					if request.URL != nil {
						operationInfo.Operation = request.URL.Path
					}
					if headerpkg.GetIsWebSocket(request.Header) {
						isWebsocket = true
						operationInfo.Method = "WS"
					}
					requestInfo.ClientIP = contextpkg.ClientIPFromHTTP(ctx, request)
				}
			} else if _, isGRPC := tr.(*grpc.Transport); isGRPC {
				operationInfo.Method = "GRPC"
				requestInfo.ClientIP = contextpkg.ClientIPFromGRPC(ctx)
			}

			// websocket 不输出错误
			if isWebsocket {
				var kv = operationInfo.Kvs()
				kv = append(kv, requestInfo.Kvs()...)
				logger.Log(ctx, loggingLevel, "server request", kv...)
				return
			}

			// 请求参数
			operationInfo.Args = truncateRequestArgs(extractArgs(req))

			// 打印日志
			var kv = operationInfo.Kvs()
			kv = append(kv, requestInfo.Kvs()...)

			// 有错误的
			if err != nil {
				loggingLevel = slog.LevelError
				// 错误信息
				errMessage.Message = err.Error()
				se := errorpkg.FromError(err)
				if se != nil {
					errMessage.Code = se.Code
					errMessage.Reason = se.Reason
				}
				if se != nil && (se.GetCode() < stdhttp.StatusInternalServerError || se.GetCode() >= _minInfoLevelCode) {
					loggingLevel = slog.LevelInfo
				}
				// 错误调用
				var stackMessage string
				var callers = errorpkg.CallStackWithSkipAndDepth(err, opt.withSkipDepth, opt.withTracerDepth)
				if len(callers) > 0 {
					stackMessage = strings.Join(callers, "\n\t")
				}

				// 打印日志
				kv = append(kv, errMessage.Kvs()...)
				kv = append(kv, "stack", stackMessage)
			}

			// 输出日志
			logger.Log(ctx, loggingLevel, "server request", kv...)
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

func (s *RequestInfoForClient) Kvs() []any {
	return []any{
		"kind", s.Kind,
		"component", s.Component,
		"latency", s.Latency.String(),
		"client_ip", s.ClientIP,
	}
}

func ClientLoggingKvs(requestInfo *RequestInfoForClient, operationInfo *OperationInfo) []any {
	return []any{
		"client_ip", requestInfo.ClientIP,
		"latency", requestInfo.Latency,
		"kind", requestInfo.Kind,
		"component", requestInfo.Component,
		"method", operationInfo.Method,
		"operation", operationInfo.Operation,
		"args", operationInfo.Args,
	}
}

// ClientLog is an client logging middleware.
// logging.Client(logger)
func ClientLog(logger *slog.Logger) middleware.Middleware {
	if logger == nil {
		logger = slog.Default()
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			var (
				startTime = time.Now()
				level     = slog.LevelInfo
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
			operationInfo.Args = truncateRequestArgs(extractArgs(req))

			// 请求信息
			requestInfo := &RequestInfoForClient{
				Kind:      "client",
				Component: kind,
				ClientIP:  ippkg.LocalIP(),
			}

			// request
			if httpTr, isHTTP := tr.(http.Transporter); isHTTP {
				request := httpTr.Request()
				if request != nil {
					operationInfo.Method = request.Method
					if request.URL != nil {
						operationInfo.Operation = request.URL.Path
					}
					if headerpkg.GetIsWebSocket(request.Header) {
						operationInfo.Method = "WS"
					}
				}
			} else if _, isGRPC := tr.(*grpc.Transport); isGRPC {
				operationInfo.Method = "GRPC"
			}

			reply, err = handler(ctx, req)
			requestInfo.Latency = time.Since(startTime)

			var kv = operationInfo.Kvs()
			kv = append(kv, requestInfo.Kvs()...)
			if err != nil {
				se := errorpkg.FromError(err)
				if se != nil {
					code = se.Code
					reason = se.Reason
				}
				var stack string
				level, stack = extractError(err)
				if se != nil && (se.GetCode() < stdhttp.StatusInternalServerError || se.GetCode() >= _minInfoLevelCode) {
					level = slog.LevelInfo
				}
				if level == slog.LevelError {
					level = slog.LevelWarn // no error level in client
				}
				errMessage := &ErrMessage{
					Code:   code,
					Reason: reason,
				}
				kv = append(kv,
					"error", errMessage.String(),
					"stack", stack,
				)
			}
			logger.Log(ctx, level, "client request", kv...)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req any) string {
	if redacter, ok := req.(logging.Redacter); ok {
		return redacter.Redact()
	}
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

func truncateRequestArgs(args string) string {
	limit := maxRequestArgs.Load()
	if uint64(len(args)) > limit {
		return args[:int(limit)]
	}
	return args
}

// extractError returns the string of the error
func extractError(err error) (slog.Level, string) {
	if err != nil {
		return slog.LevelError, fmt.Sprintf("%+v", err)
	}
	return slog.LevelInfo, ""
}
