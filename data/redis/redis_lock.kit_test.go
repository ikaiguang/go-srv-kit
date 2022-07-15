package redisutil

import (
	"context"
	"testing"
	"time"

	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
	"github.com/stretchr/testify/require"
)

const (
	keyName = "test-redsync"
)

// go test -v -count=1 ./data/redis -test.run=TestLockOnce
func TestLockOnce(t *testing.T) {

	redisCC, err := NewDB(redisConfig)
	require.Nil(t, err)
	locker := NewLocker(redisCC)

	ctx := context.Background()

	tests := []struct {
		name         string
		lockerStatus bool
		isLockFailed bool
		unlock       lockerutil.Unlock
	}{
		{
			name:         "#加锁成功",
			lockerStatus: true,
			isLockFailed: false,
		},
		{
			name:         "#加锁一定的时间后",
			lockerStatus: true,
			isLockFailed: false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unlock, err := locker.Once(ctx, keyName)
			t.Logf("想要：加锁(%v)\n", tt.lockerStatus)
			if err != nil {
				t.Logf("加锁失败！错误：%v\n", err)
				t.Logf("==> IsLockFailedError : %v\n", lockerutil.IsLockFailedError(err))
			} else {
				t.Logf("===> 加锁成功！\n")
			}
			tests[i].unlock = unlock
		})

		// 睡眠
		if i == len(tests)-1 {
			continue
		}
		sleepDuration := _lockExpire + 2*time.Second
		t.Logf("==> 睡眠%v,尝试加锁是否成功。设置的加锁时长为%v\n", sleepDuration, _lockExpire)
		time.Sleep(sleepDuration)
	}

	// 解锁
	for i := range tests {
		ok, err := tests[i].unlock.Unlock(context.Background())
		if err != nil {
			t.Logf("==> 解锁 error : %v\n", err)
		}
		if !ok {
			t.Logf("==> 解锁 status : %v\n", ok)
		}
	}
}

// go test -v -count=1 ./data/redis -test.run=TestLockMutex
func TestLockMutex(t *testing.T) {

	redisCC, err := NewDB(redisConfig)
	require.Nil(t, err)
	locker := NewLocker(redisCC)

	ctx := context.Background()

	tests := []struct {
		name         string
		lockerStatus bool
		isLockFailed bool
		unlock       lockerutil.Unlock
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
			unlock, err := locker.Mutex(ctx, keyName)
			t.Logf("想要：加锁(%v)\n", tt.lockerStatus)
			if err != nil {
				t.Logf("加锁失败啦！错误：%v\n", err)
				t.Logf("==> IsLockFailedError : %v\n", lockerutil.IsLockFailedError(err))
			} else {
				t.Logf("===> 加锁成功啦！\n")
			}
			tests[i].unlock = unlock
		})

		// 睡眠
		if i == len(tests)-1 {
			continue
		}
		// lockExpire 默认8秒
		sleepDuration := _lockExpire + time.Second
		t.Logf("==> 睡眠:%v,尝试加锁是否成功。设置的加锁时长为:%v\n", sleepDuration, _lockExpire)
		time.Sleep(sleepDuration)
	}

	// 解锁
	for i := range tests {
		ok, err := tests[i].unlock.Unlock(context.Background())
		if err != nil {
			t.Logf("unlock error : %v\n", err)
		}
		if !ok {
			t.Logf("unlock status : %v\n", ok)
		}
	}
}
