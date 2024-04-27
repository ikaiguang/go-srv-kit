package redispkg

import (
	"context"
	stderrors "errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Config redis config
type Config struct {
	Addresses    []string
	Username     string
	Password     string
	Db           uint32
	DialTimeout  *durationpb.Duration
	ReadTimeout  *durationpb.Duration
	WriteTimeout *durationpb.Duration
	// ConnMaxActive 连接的最大数量
	ConnMaxActive uint32
	// ConnMaxLifetime 连接可复用的最大时间
	ConnMaxLifetime *durationpb.Duration
	// ConnMaxIdle 连接池中空闲连接的最大数量
	ConnMaxIdle uint32
	ConnMinIdle uint32
	// ConnMaxIdleTime 设置连接空闲的最长时间
	ConnMaxIdleTime *durationpb.Duration
}

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
