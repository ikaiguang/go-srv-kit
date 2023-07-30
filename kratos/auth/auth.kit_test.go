package authpkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

func ExampleServer() {
	var (
		redisCC   = &redis.Client{}
		signKey   = ""
		logger    = log.DefaultLogger
		whiteList = map[string]struct{}{}
	)
	authConfig := Config{
		SigningMethod:      jwt.SigningMethodHS256,
		SignKey:            signKey,
		RefreshCrypto:      NewCBCCipher(signKey),
		AuthCacheKeyPrefix: CheckAuthCacheKeyPrefix(nil),
	}
	repo, err := NewAuthRepo(redisCC, logger, authConfig)
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
			WithTokenValidator(repo.VerifyToken),
		),
	).
		Match(ExampleWhiteListMatcher(whiteList)).
		Build()

	return
}
