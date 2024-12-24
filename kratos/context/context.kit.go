package contextpkg

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/trace"
)

// FromServerContext 等于 transport.FromServerContext
func FromServerContext(ctx context.Context) (tr transport.Transporter, ok bool) {
	return transport.FromServerContext(ctx)
}

// FromClientContext 等于 transport.FromClientContext
func FromClientContext(ctx context.Context) (tr transport.Transporter, ok bool) {
	return transport.FromClientContext(ctx)
}

// MatchHTTPContext 匹配
// Deprecated: 不建议使用此方法
// 建议使用： contextpkg.MatchHTTPServerContext or contextpkg.MatchGRPCServerContext
func MatchHTTPContext(ctx context.Context) (http.Context, bool) {
	httpCtx, ok := ctx.(http.Context)
	return httpCtx, ok
}

// MatchHTTPServerContext ...
func MatchHTTPServerContext(ctx context.Context) (tr http.Transporter, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return ToHTTPTransporter(kratosTr)
}

// MatchGRPCServerContext ...
func MatchGRPCServerContext(ctx context.Context) (tr *grpc.Transport, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return ToGRPCTransporter(kratosTr)
}

// MatchHTTPClientContext ...
func MatchHTTPClientContext(ctx context.Context) (tr http.Transporter, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return ToHTTPTransporter(kratosTr)
}

// MatchGRPCClientContext ...
func MatchGRPCClientContext(ctx context.Context) (tr *grpc.Transport, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return ToGRPCTransporter(kratosTr)
}

// ToHTTPTransporter ...
func ToHTTPTransporter(kratosTr transport.Transporter) (httpTr http.Transporter, ok bool) {
	httpTr, ok = kratosTr.(http.Transporter)
	return httpTr, ok
}

// ToGRPCTransporter ...
func ToGRPCTransporter(kratosTr transport.Transporter) (grpcTr *grpc.Transport, ok bool) {
	grpcTr, ok = kratosTr.(*grpc.Transport)
	return grpcTr, ok
}

func NewContext(ctx context.Context) context.Context {
	span := trace.SpanFromContext(ctx)
	return trace.ContextWithSpan(context.Background(), span)
}
