package setup

import (
	stdlog "log"
	"sync"

	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"

	redisutil "github.com/ikaiguang/go-srv-kit/redis"
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
	if s.redisClient != nil {
		return s.redisClient, err
	}
	s.redisClient, err = s.loadingRedisClient()
	if err != nil {
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