package redisutil

import (
	"context"
	stderrors "errors"
	"time"

	"github.com/go-redsync/redsync/v4"
	lockerutil "github.com/ikaiguang/go-srv-kit/kit/locker"
)

// mutexLock ...
type mutexLock struct {
	mutex         *redsync.Mutex
	extendChannel chan bool
}

// Unlock ...
func (s *mutexLock) Unlock(ctx context.Context) (ok bool, err error) {
	// 防止panic 信道
	if s.extendChannel != nil {
		s.extendChannel <- true
		close(s.extendChannel)
	}

	// 取锁的有效期。在获取锁之前，该值将为零值。
	if s.mutex.Until().IsZero() {
		return ok, err
	}
	return s.mutex.UnlockContext(ctx)
}

// extend ...
func (s *mutexLock) extend(ctx context.Context) {
	// 计时器
	timer := time.NewTimer(_lockExtendDelay)

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
		s.extend(ctx)
	case <-s.extendChannel: // 停止
		timer.Stop()
		// 调试
		//fmt.Println("redis mutex 停止续期")
		return
	}
}
