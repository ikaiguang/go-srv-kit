package redisutil

import (
	"context"

	"github.com/go-redis/redis/v8"
	pkgerrors "github.com/pkg/errors"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewDB redis db
func NewDB(conf *confv1.Data_Redis) (db *redis.Client, err error) {
	redisOpt := &redis.Options{
		Addr:         conf.Addr,
		Username:     conf.Username,
		Password:     conf.Password,
		DB:           int(conf.Db),
		DialTimeout:  conf.DialTimeout.AsDuration(),
		ReadTimeout:  conf.ReadTimeout.AsDuration(),
		WriteTimeout: conf.WriteTimeout.AsDuration(),
		PoolSize:     int(conf.ConnMaxActive),
		MaxConnAge:   conf.ConnMaxLifetime.AsDuration(),
		MinIdleConns: int(conf.ConnMaxIdle),
		IdleTimeout:  conf.ConnMaxIdleTime.AsDuration(),
	}
	db = redis.NewClient(redisOpt)

	// ping 测试连接
	err = db.Ping(context.Background()).Err()
	if err != nil {
		err = pkgerrors.Errorf("redis connection ping failed : %s", err.Error())
		return db, err
	}
	return db, err
}
