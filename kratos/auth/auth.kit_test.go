package authpkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	lockerpkg "github.com/ikaiguang/go-srv-kit/kit/locker"
	"github.com/redis/go-redis/v9"
)

// testLocker 测试用的 Locker mock 实现
type testLocker struct{}

func (t *testLocker) Mutex(_ context.Context, lockName string) (lockerpkg.Unlocker, error) {
	return &testUnlocker{name: lockName}, nil
}

func (t *testLocker) Once(_ context.Context, lockName string) (lockerpkg.Unlocker, error) {
	return &testUnlocker{name: lockName}, nil
}

// testUnlocker 测试用的 Unlocker mock 实现
type testUnlocker struct {
	name string
}

func (t *testUnlocker) Unlock(_ context.Context) (bool, error) { return true, nil }
func (t *testUnlocker) Name() string                           { return t.name }

func ExampleServer() {
	var (
		redisCC   = &redis.Client{}
		signKey   = ""
		logger    = log.DefaultLogger
		whiteList = map[string]struct{}{}
	)
	authConfig := Config{
		RefreshCrypto: NewCBCCipher(signKey),
	}
	tokenM := NewTokenManager(logger, redisCC, nil, &testLocker{})
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
