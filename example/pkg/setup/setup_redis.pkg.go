package setuppkg

import (
	pkgerrors "github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	stdlog "log"
	"sync"

	"github.com/ikaiguang/go-srv-kit/data/redis"
)

// GetRedisClient redis 客户端
func (s *engines) GetRedisClient() (*redis.Client, error) {
	var err error
	s.redisClientMutex.Do(func() {
		s.redisClient, err = s.loadingRedisClient()
	})
	if err != nil {
		s.redisClientMutex = sync.Once{}
	}
	return s.redisClient, err
}

// reloadRedisClient 重新加载 redis 客户端
func (s *engines) reloadRedisClient() error {
	if s.Config.RedisConfig() == nil {
		return nil
	}
	redisClient, err := s.loadingRedisClient()
	if err != nil {
		return err
	}
	*s.redisClient = *redisClient
	return nil
}

// loadingRedisClient redis 客户端
func (s *engines) loadingRedisClient() (*redis.Client, error) {
	if s.Config.RedisConfig() == nil {
		stdlog.Println("|*** 加载：Redis客户端：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载：Redis客户端：...")

	return redisutil.NewDB(s.Config.RedisConfig())
}
