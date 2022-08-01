package headermiddle

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"go.opentelemetry.io/otel/trace"

	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// RequestHeader 响应头
func RequestHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if httpTr, ok := contextutil.MatchHTTPServerContext(ctx); ok {
				var traceID string
				if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
					traceID = span.TraceID().String()
				}
				if traceID == "" {
					traceID = uuidutil.New()
				}
				headerutil.SetRequestID(httpTr.Request().Header, traceID)
			}
			return handler(ctx, req)
		}
	}
}
