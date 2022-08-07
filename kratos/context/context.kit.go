package contextutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// MatchHTTPServerContext ...
// 然后参考 http.Context 实现的 Vars、Query
func MatchHTTPServerContext(ctx context.Context) (tr http.Transporter, ok bool) {
	stdTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	tr, ok = stdTr.(http.Transporter)
	return tr, ok
}

// MatchGRPCServerContext ...
func MatchGRPCServerContext(ctx context.Context) (tr *grpc.Transport, ok bool) {
	stdTr, ok := transport.FromServerContext(ctx)
	if !ok {
		return tr, ok
	}
	tr, ok = stdTr.(*grpc.Transport)
	return tr, ok
}

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
// 建议使用： contextutil.MatchHTTPServerContext or contextutil.MatchGRPCServerContext
func MatchHTTPContext(ctx context.Context) (http.Context, bool) {
	httpCtx, ok := ctx.(http.Context)
	return httpCtx, ok
}
