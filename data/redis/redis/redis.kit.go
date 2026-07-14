package redispkg

import (
	"context"
	stderrors "errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewDB redis db
func NewDB(conf *Config) (db redis.UniversalClient, err error) {
	redisOpt := &redis.UniversalOptions{
		Addrs:           conf.Addresses,
		Username:        conf.Username,
		Password:        conf.Password,
		DB:              int(conf.Db),
		DialTimeout:     conf.DialTimeout.AsDuration(),
		ReadTimeout:     conf.ReadTimeout.AsDuration(),
		WriteTimeout:    conf.WriteTimeout.AsDuration(),
		PoolSize:        int(conf.ConnMaxActive),
		ConnMaxLifetime: conf.ConnMaxLifetime.AsDuration(),
		MinIdleConns:    int(conf.ConnMinIdle),
		MaxIdleConns:    int(conf.ConnMaxIdle),
		ConnMaxIdleTime: conf.ConnMaxIdleTime.AsDuration(),
	}
	db = redis.NewUniversalClient(redisOpt)

	// ping 测试连接
	err = db.Ping(context.Background()).Err()
	if err != nil {
		err = fmt.Errorf("redis connection ping failed : %w", err)
		return db, err
	}
	return db, err
}

func IsNilErr(err error) bool {
	return stderrors.Is(err, redis.Nil)
}
