package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel/trace"

	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
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
			}
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				if tr.ReplyHeader().Get(headerpkg.TraceID) == "" {
					tr.ReplyHeader().Set(headerpkg.TraceID, traceID)
				}
				if tr.RequestHeader().Get(headerpkg.RequestID) == "" {
					tr.RequestHeader().Set(headerpkg.RequestID, traceID)
				}
			}
			return handler(ctx, req)
		}
	}
}
