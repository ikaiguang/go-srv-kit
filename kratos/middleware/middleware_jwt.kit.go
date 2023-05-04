package middlewareutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher(whiteList map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		//if tr, ok := contextutil.MatchHTTPServerContext(ctx); ok {
		//	if _, ok := whiteList[tr.Request().URL.Path]; ok {
		//		return false
		//	}
		//}

		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
