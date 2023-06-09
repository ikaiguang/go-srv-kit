package contextpkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// FromServerContext 等于 transport.FromServerContext
func FromServerContext(ctx context.Context) (tr transport.Transporter, ok bool) {
	return transport.FromServerContext(ctx)
}

// FromClientContext 等于 transport.FromClientContext
func FromClientContext(ctx context.Context) (tr transport.Transporter, ok bool) {
	return transport.FromClientContext(ctx)
}

// MatchHTTPServerContext ...
// 然后参考 http.Context 实现的 Vars、Query
func MatchHTTPServerContext(ctx context.Context) (tr http.Transporter, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return IsHTTPTransporter(kratosTr)
}

// MatchGRPCServerContext ...
func MatchGRPCServerContext(ctx context.Context) (tr *grpc.Transport, ok bool) {
	kratosTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	return IsGRPCTransporter(kratosTr)
}

// IsHTTPTransporter ...
func IsHTTPTransporter(kratosTr transport.Transporter) (httpTr http.Transporter, ok bool) {
	httpTr, ok = kratosTr.(http.Transporter)
	return httpTr, ok
}

// IsGRPCTransporter ...
func IsGRPCTransporter(kratosTr transport.Transporter) (grpcTr *grpc.Transport, ok bool) {
	grpcTr, ok = kratosTr.(*grpc.Transport)
	return grpcTr, ok
}

// MatchHTTPContext 匹配
// Deprecated: 不建议使用此方法
// 建议使用： contextpkg.MatchHTTPServerContext or contextpkg.MatchGRPCServerContext
func MatchHTTPContext(ctx context.Context) (http.Context, bool) {
	httpCtx, ok := ctx.(http.Context)
	return httpCtx, ok
}
