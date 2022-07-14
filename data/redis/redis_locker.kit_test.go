package redisutil

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
	"github.com/stretchr/testify/require"
)

const (
	keyName = "test-redsync"
)

var (
	redisCC *redis.Client
)

// go test -v -count=1 ./pkg/locker/redis -test.run=TestNewRedisSimpleLocker
func TestNewRedisSimpleLocker(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name         string
		lockerStatus bool
		isLockFailed bool
		locker       lockerutil.Locker
	}{
		{
			name:         "#加锁成功",
			lockerStatus: true,
			isLockFailed: false,
		},
		{
			name:         "#加锁失败；锁时间过期后变加锁成功",
			lockerStatus: false,
			isLockFailed: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locker, err := NewSimpleLocker(redisCC)
			require.Nil(t, err)
			tests[i].locker = locker

			err = locker.Lock(ctx, keyName)
			t.Logf("我想要：加锁(%v)\n", tt.lockerStatus)
			if err != nil {
				t.Logf("我加锁失败啦！错误：%v\n", err)
			} else {
				t.Logf("===> 我加锁成功啦！\n")
			}
		})

		// 睡眠
		if i == len(tests)-1 {
			continue
		}
		sleepDuration := 10 * time.Second
		t.Logf("==> 睡眠%v,尝试加锁是否成功。V4简单加锁，默认时长为8s\n", sleepDuration)
		time.Sleep(sleepDuration)
	}

	// 解锁
	for i := range tests {
		ok, err := tests[i].locker.Unlock(context.Background())
		if err != nil {
			t.Logf("unlock error : %v\n", err)
		}
		if !ok {
			t.Logf("unlock status : %v\n", ok)
		}
	}
}

// go test -v -count=1 ./pkg/locker/redis -test.run=TestNewRedisLocker
func TestNewRedisLocker(t *testing.T) {

	ctx := context.Background()

	tests := []struct {
		name         string
		lockerStatus bool
		isLockFailed bool
		locker       lockerutil.Locker
	}{
		{
			name:         "#加锁成功",
			lockerStatus: true,
			isLockFailed: false,
		},
		{
			name:         "#加锁失败",
			lockerStatus: false,
			isLockFailed: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locker, err := NewSimpleLocker(redisCC)
			require.Nil(t, err)
			tests[i].locker = locker

			err = locker.Lock(ctx, keyName)
			t.Logf("我想要：加锁(%v)\n", tt.lockerStatus)
			if err != nil {
				t.Logf("我加锁失败啦！错误：%v\n", err)
			} else {
				t.Logf("===> 我加锁成功啦！\n")
			}
		})

		// 睡眠
		if i == len(tests)-1 {
			continue
		}
		sleepDuration := lockExpire + time.Second
		t.Logf("==> 睡眠%v,尝试加锁是否成功。V4简单加锁，默认时长为%v\n", sleepDuration, lockExpire)
		time.Sleep(sleepDuration)
	}

	// 解锁
	for i := range tests {
		ok, err := tests[i].locker.Unlock(context.Background())
		if err != nil {
			t.Logf("unlock error : %v\n", err)
		}
		if !ok {
			t.Logf("unlock status : %v\n", ok)
		}
	}
}
