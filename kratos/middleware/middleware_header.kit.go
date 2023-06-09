package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"go.opentelemetry.io/otel/trace"

	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	contextpkg "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// RequestHeader 响应头
func RequestHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var traceID string
			if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
				traceID = span.TraceID().String()
			}
			if traceID == "" {
				traceID = uuidpkg.New()
				if httpTr, ok := contextpkg.MatchHTTPServerContext(ctx); ok {
					httpTr.RequestHeader().Set(headerpkg.RequestID, traceID)
				}
				if grpcTr, ok := contextpkg.MatchGRPCServerContext(ctx); ok {
					grpcTr.ReplyHeader().Set(headerpkg.RequestID, traceID)
				}
			}
			return handler(ctx, req)
		}
	}
}
