package redispkg

import (
	"context"
	stderrors "errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const defaultProbeTimeout = 5 * time.Second

// NewDB redis db
func NewDB(conf *Config) (db redis.UniversalClient, err error) {
	if conf == nil {
		return nil, stderrors.New("redis config is nil")
	}
	if len(conf.Addresses) == 0 {
		return nil, stderrors.New("redis addresses are empty")
	}
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
	probeTimeout := conf.DialTimeout.AsDuration()
	if probeTimeout <= 0 {
		probeTimeout = defaultProbeTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), probeTimeout)
	defer cancel()
	err = db.Ping(ctx).Err()
	if err != nil {
		pingErr := fmt.Errorf("redis connection ping failed: %w", err)
		if closeErr := db.Close(); closeErr != nil {
			return nil, stderrors.Join(pingErr, fmt.Errorf("close redis client: %w", closeErr))
		}
		return nil, pingErr
	}
	return db, nil
}

func IsNilErr(err error) bool {
	return stderrors.Is(err, redis.Nil)
}
