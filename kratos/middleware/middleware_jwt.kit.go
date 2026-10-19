package middlewarepkg

import (
	"context"

	"github.com/go-kratos/kratos/v3/middleware/selector"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher(whiteList map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
