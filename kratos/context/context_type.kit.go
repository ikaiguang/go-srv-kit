package contextutil

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// MatchHTTPContext 匹配
func MatchHTTPContext(ctx context.Context) (http.Context, bool) {
	httpCtx, ok := ctx.(http.Context)
	return httpCtx, ok
}
