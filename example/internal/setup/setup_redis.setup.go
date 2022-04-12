package setup

import (
	stdlog "log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/ikaiguang/go-srv-kit/data/redis"
	pkgerrors "github.com/pkg/errors"
)

// RedisClient redis 客户端
func (s *modules) RedisClient() (*redis.Client, error) {
	var err error
	s.redisClientMutex.Do(func() {
		s.redisClient, err = s.loadingRedisClient()
	})
	if err != nil {
		s.redisClientMutex = sync.Once{}
		return nil, err
	}
	return s.redisClient, err
}

// loadingRedisClient redis 客户端
func (s *modules) loadingRedisClient() (*redis.Client, error) {
	if s.Config.RedisConfig() == nil {
		stdlog.Println("|*** 加载Redis客户端：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载Redis客户端：...")

	return redisutil.NewDB(s.Config.RedisConfig())
}
