package middlewareutil

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"go.opentelemetry.io/otel/trace"

	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// ResponseHeader 响应头
func ResponseHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var traceID string
			if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
				traceID = span.TraceID().String()
			}
			//traceID = "testdata"

			// 设置请求头
			if httpContext, ok := contextutil.MatchHTTPContext(ctx); ok {
				// http
				httpContext.Response().Header().Set(headerutil.RequestID, traceID)
			}
			return handler(ctx, req)
		}
	}
}
