package headermiddle

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"go.opentelemetry.io/otel/trace"

	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// RequestHeader 响应头
func RequestHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var traceID string
			if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
				traceID = span.TraceID().String()
			}
			//traceID = "testdata"

			// 设置请求头
			if httpContext, ok := contextutil.MatchHTTPContext(ctx); ok {
				// request id
				headerutil.SetRequestID(httpContext.Request().Header, traceID)
				// 是否websocket
				// 在升级为websocket链接后设置：websocketutil.UpgradeConn
				//if connectionutil.IsWebSocketConn(httpContext.Request()) {
				//	headerutil.SetIsWebsocket(httpContext.Request().Header)
				//}
			}

			return handler(ctx, req)
		}
	}
}
