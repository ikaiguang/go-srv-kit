package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"go.opentelemetry.io/otel/trace"

	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	headerpkg "github.com/ikaiguang/go-srv-kit/kratos/header"
)

// RequestAndResponseHeader 请求头 and 响应头
func RequestAndResponseHeader() middleware.Middleware {
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
				if tr.RequestHeader().Get(headerpkg.RequestID) == "" {
					tr.RequestHeader().Set(headerpkg.RequestID, traceID)
				}
				if tr.ReplyHeader().Get(headerpkg.RequestID) == "" {
					tr.ReplyHeader().Set(headerpkg.RequestID, traceID)
				}
			}
			return handler(ctx, req)
		}
	}
}
