package authpkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/redis/go-redis/v9"
)

func ExampleServerMiddleware() {
	var (
		redisCC   = &redis.Client{}
		signKey   = ""
		logger    = log.DefaultLogger
		whiteList = map[string]struct{}{}
	)
	authConfig := Config{
		RefreshCrypto: NewCBCCipher(signKey),
	}
	tokenM := NewTokenManger(logger, redisCC, nil)
	repo, err := NewAuthRepo(authConfig, logger, tokenM)
	if err != nil {
		return
	}

	// ExampleWhiteListMatcher 路由白名单
	var ExampleWhiteListMatcher = func(whiteList map[string]struct{}) selector.MatchFunc {
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

	_ = selector.Server(
		Server(
			repo.JWTSigningKeyFunc,
			WithSigningMethod(repo.JWTSigningMethod()),
			WithClaims(repo.JWTSigningClaims),
			WithAccessTokenValidator(repo.VerifyAccessToken),
		),
	).
		Match(ExampleWhiteListMatcher(whiteList)).
		Build()

	return
}
