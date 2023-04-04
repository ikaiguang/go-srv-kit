package redisutil

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewDB redis db
func NewDB(conf *confv1.Data_Redis) (db *redis.Client, err error) {
	redisOpt := &redis.Options{
		Addr:            conf.Addr,
		Username:        conf.Username,
		Password:        conf.Password,
		DB:              int(conf.Db),
		DialTimeout:     conf.DialTimeout.AsDuration(),
		ReadTimeout:     conf.ReadTimeout.AsDuration(),
		WriteTimeout:    conf.WriteTimeout.AsDuration(),
		PoolSize:        int(conf.ConnMaxActive),
		ConnMaxLifetime: conf.ConnMaxLifetime.AsDuration(),
		MinIdleConns:    0,
		MaxIdleConns:    int(conf.ConnMaxIdle),
		ConnMaxIdleTime: conf.ConnMaxIdleTime.AsDuration(),
	}
	db = redis.NewClient(redisOpt)

	// ping 测试连接
	err = db.Ping(context.Background()).Err()
	if err != nil {
		err = fmt.Errorf("redis connection ping failed : %w", err)
		return db, err
	}
	return db, err
}
