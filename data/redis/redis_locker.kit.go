package redisutil

import (
	"context"
	stderrors "errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
	uuid "github.com/satori/go.uuid"
)

// NewLocker 获取锁
// 此加锁，锁默认时间：8s；每过5s续期一次；除了解锁，否则永不过期；
func NewLocker(client *redis.Client) (locker lockerutil.Locker, err error) {
	handler := new(redisLocker)
	handler.Init(client)

	return handler, err
}

const (
	lockExpire      = 8 * time.Second        // 锁过期时间
	lockResetDelay  = 5 * time.Second        // 重置锁时间，防止锁自动过期
	lockTries       = 1                      // 尝试次数
	lockRetryDelay  = 500 * time.Millisecond // 尝试间隔
	lockDriftFactor = 0.01                   // 时钟漂移系数
)

var (
	// locker opts
	lockerOpts []redsync.Option

	// locker func
	lockerValueFunc = func() (string, error) {
		id := uuid.NewV4()
		return id.String(), nil
	}
)

func init() {
	// 锁过期时间
	lockerOpts = append(lockerOpts, redsync.WithExpiry(lockExpire))
	// try && 尝试间隔
	lockerOpts = append(lockerOpts, redsync.WithTries(lockTries))
	lockerOpts = append(lockerOpts, redsync.WithRetryDelay(lockRetryDelay))
	// 设置时钟漂移系数。
	lockerOpts = append(lockerOpts, redsync.WithDriftFactor(lockDriftFactor))
	// 随机值
	lockerOpts = append(lockerOpts, redsync.WithGenValueFunc(lockerValueFunc))
}

// redisLocker 锁默认时间：8s；每过5s续期一次；除了解锁，否则永不过期；
type redisLocker struct {
	rs        *redsync.Redsync
	mutex     *redsync.Mutex
	resetChan chan bool
}

// Init .
func (s *redisLocker) Init(client *redis.Client) {
	s.rs = redsync.New(goredis.NewPool(client))
}

// Lock .
func (s *redisLocker) Lock(ctx context.Context, name string) (err error) {
	s.mutex = s.rs.NewMutex(name, lockerOpts...)
	if err = s.mutex.LockContext(ctx); err != nil {
		err = lockerutil.NewErrLockerFailed(true, name, err)
		return err
	}

	// 续期锁，防止锁自动过期
	s.resetChan = make(chan bool)
	go s.resetExpire(ctx)

	return err
}

// Unlock 解锁
func (s *redisLocker) Unlock(ctx context.Context) (ok bool, err error) {
	// 防止panic
	if s.resetChan != nil {
		s.resetChan <- true
		close(s.resetChan)
	}

	// 取锁的有效期。在获取锁之前，该值将为零值。
	if s.mutex.Until().IsZero() {
		return ok, err
	}
	return s.mutex.UnlockContext(ctx)
}

// resetExpire 重置锁时间，防止自动过期而解锁
func (s *redisLocker) resetExpire(ctx context.Context) {
	// 计时器
	timer := time.NewTimer(lockResetDelay)

	select {
	case <-timer.C: // 续期
		// 结束计时
		timer.Stop()
		// 续期
		if ok, err := s.mutex.ExtendContext(ctx); err != nil || !ok {
			if stderrors.Is(err, redsync.ErrExtendFailed) {
				err = lockerutil.NewErrExtendFailed(true, s.mutex.Name(), err)
			}
			// 调试
			//fmt.Println("redis mutex 续期失败")
			return
		}
		// 调试
		//fmt.Println("redis mutex 续期成功")
		// 再次续期
		s.resetExpire(ctx)
	case <-s.resetChan: // 停止
		timer.Stop()
		// 调试
		//fmt.Println("redis mutex 停止续期")
		return
	}
}
