package redisutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	singletonMutex        sync.Once
	singletonRedisManager RedisManager
)

func NewSingletonRedisManager(conf *configpb.Redis) (RedisManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonRedisManager, err = NewRedisManager(conf)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonRedisManager, err
}

func GetRedisClient(redisManager RedisManager) (redis.UniversalClient, error) {
	return redisManager.GetClient()
}
