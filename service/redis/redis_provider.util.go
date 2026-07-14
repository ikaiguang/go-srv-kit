package redisutil

import (
	"sync"

	configpb "github.com/ikaiguang/go-service-kit/api/config"
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
