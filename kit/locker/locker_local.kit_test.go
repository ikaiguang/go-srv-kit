package lockerpkg

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	localKeyName = "test-redsync"
)

// go test -v -count 1 ./locker -run TestLocal_LockOnce
func TestLocal_LockOnce(t *testing.T) {
	locker := NewLocalLocker()
	ctx := context.Background()

	t.Run("首次加锁成功", func(t *testing.T) {
		unlock, err := locker.Once(ctx, localKeyName)
		require.Nil(t, err, "首次加锁应成功")
		require.NotNil(t, unlock)

		// 解锁
		ok, err := unlock.Unlock(ctx)
		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("锁过期后再次加锁成功", func(t *testing.T) {
		unlock1, err := locker.Once(ctx, localKeyName+"_expire")
		require.Nil(t, err)
		require.NotNil(t, unlock1)

		// 等待锁过期（Once 会在 _cacheLockerExpire 后自动解锁）
		time.Sleep(_cacheLockerExpire + 2*time.Second)

		// 过期后再次加锁应成功
		unlock2, err := locker.Once(ctx, localKeyName+"_expire")
		require.Nil(t, err, "锁过期后再次加锁应成功")
		require.NotNil(t, unlock2)

		_, _ = unlock2.Unlock(ctx)
	})
}

// go test -v -count 1 ./locker -run TestLocal_LockMutex
func TestLocal_LockMutex(t *testing.T) {
	locker := NewLocalLocker()
	ctx := context.Background()

	t.Run("首次加锁成功", func(t *testing.T) {
		unlock, err := locker.Mutex(ctx, localKeyName+"_mutex")
		require.Nil(t, err, "首次 Mutex 加锁应成功")
		require.NotNil(t, unlock)

		// 持有锁期间再次加锁应失败
		t.Run("持有锁期间再次加锁失败", func(t *testing.T) {
			_, err := locker.Mutex(ctx, localKeyName+"_mutex")
			assert.NotNil(t, err, "持有锁期间再次加锁应失败")
			assert.True(t, IsErrorLockFailed(err), "应返回 LockFailed 错误")
		})

		// 解锁
		ok, err := unlock.Unlock(ctx)
		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("解锁后再次加锁成功", func(t *testing.T) {
		unlock1, err := locker.Mutex(ctx, localKeyName+"_mutex2")
		require.Nil(t, err)
		require.NotNil(t, unlock1)

		_, _ = unlock1.Unlock(ctx)

		// 解锁后应能再次加锁
		unlock2, err := locker.Mutex(ctx, localKeyName+"_mutex2")
		require.Nil(t, err, "解锁后再次加锁应成功")
		require.NotNil(t, unlock2)

		_, _ = unlock2.Unlock(ctx)
	})
}
