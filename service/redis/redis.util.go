package redisutil

import (
	stdlog "log"
	"sync"

	errorpkg "github.com/ikaiguang/go-kratos-kit/error"
	redispkg "github.com/ikaiguang/go-redis-kit/redis"
	configpb "github.com/ikaiguang/go-service-kit/api/config"
)

type redisManager struct {
	conf *configpb.Redis

	redisOnce   sync.Once
	redisClient redis.UniversalClient
}

type RedisManager interface {
	Enable() bool
	GetClient() (redis.UniversalClient, error)
	Close() error
}

func NewRedisManager(conf *configpb.Redis) (RedisManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = redis")
		return nil, errorpkg.WithStack(e)
	}
	return &redisManager{
		conf: conf,
	}, nil
}

func (s *redisManager) GetClient() (redis.UniversalClient, error) {
	var err error
	s.redisOnce.Do(func() {
		s.redisClient, err = s.loadingRedisClient()
	})
	if err != nil {
		s.redisOnce = sync.Once{}
	}
	return s.redisClient, err
}

func (s *redisManager) Close() error {
	if s.redisClient != nil {
		stdlog.Println("|*** STOP: close: redisClient")
		err := s.redisClient.Close()
		if err != nil {
			stdlog.Println("|*** STOP: close: redisClient failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *redisManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *redisManager) loadingRedisClient() (redis.UniversalClient, error) {
	stdlog.Println("|*** LOADING: Redis client: ...")
	uc, err := redispkg.NewDB(ToRedisConfig(s.conf))
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return uc, nil
}

// ToRedisConfig ...
func ToRedisConfig(cfg *configpb.Redis) *redispkg.Config {
	return &redispkg.Config{
		Addresses:       cfg.Addresses,
		Username:        cfg.Username,
		Password:        cfg.Password,
		Db:              cfg.Db,
		DialTimeout:     cfg.DialTimeout,
		ReadTimeout:     cfg.ReadTimeout,
		WriteTimeout:    cfg.WriteTimeout,
		ConnMaxActive:   cfg.ConnMaxActive,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdle:     cfg.ConnMaxIdle,
		ConnMinIdle:     cfg.ConnMinIdle,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
	}
}
