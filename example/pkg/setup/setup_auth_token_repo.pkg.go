package setuppkg

import (
	"github.com/redis/go-redis/v9"
	stdlog "log"
	"sync"

	tokenutil "github.com/ikaiguang/go-srv-kit/kratos/token"
)

// GetAuthTokenRepo 验证Token工具
func (s *engines) GetAuthTokenRepo(redisCC *redis.Client) tokenutil.AuthTokenRepo {
	var err error
	s.authTokenRepoMutex.Do(func() {
		s.authTokenRepo = s.loadingAuthTokenRepo(redisCC)
	})
	if err != nil {
		s.authTokenRepoMutex = sync.Once{}
	}
	return s.authTokenRepo
}

// loadingAuthTokenRepo 验证Token工具
func (s *engines) loadingAuthTokenRepo(redisCC *redis.Client) tokenutil.AuthTokenRepo {
	stdlog.Println("|*** 加载：验证Token工具：...")
	return tokenutil.NewRedisTokenRepo(
		redisCC,
		tokenutil.WithAuthConfig(s.BusinessAuthConfig()),
	)
}
