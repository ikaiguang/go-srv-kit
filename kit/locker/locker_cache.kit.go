package lockerpkg

import (
	"context"
	stderrors "errors"
	"sync"
	"time"

	cachepkg "github.com/patrickmn/go-cache"
)

const (
	_cacheLockerExpire    = 8 * time.Second // 锁过期时间
	_cacheLockExtendDelay = 3 * time.Second // 重置锁时间，防止锁自动过期
	_cacheLockerClean     = 5 * time.Minute // 清除缓存
)

var _ LocalLocker = (*cache)(nil)

type cache struct {
	sm    sync.Map
	cache *cachepkg.Cache
}

// NewCacheLocker ...
func NewCacheLocker() Lock {
	return &cache{
		cache: cachepkg.New(_cacheLockerExpire, _cacheLockerClean),
	}
}

func (s *cache) Mutex(ctx context.Context, lockName string) (Unlock, error) {
	mu := &sync.Mutex{}
	muInterface, _ := s.sm.LoadOrStore(lockName, mu)
	mu = muInterface.(*sync.Mutex)
	mu.Lock()
	defer mu.Unlock()

	lockerInstance, ok := s.cache.Get(lockName)
	if ok {
		locker := lockerInstance.(*cacheLock)
		if !locker.mu.TryLock() {
			err := ErrorLockerFailed(lockName, stderrors.New("try lock failed"))
			return locker, err
		}
		return locker, nil
	}
	locker := newCacheLock(lockName, s.cache, &sync.Mutex{}, false)
	locker.mu.Lock()

	s.cache.Set(lockName, locker, _cacheLockerExpire)
	go locker.extend(ctx)

	return locker, nil
}

func (s *cache) Once(ctx context.Context, lockName string) (Unlock, error) {
	mu := &sync.Mutex{}
	muInterface, _ := s.sm.LoadOrStore(lockName, mu)
	mu = muInterface.(*sync.Mutex)
	mu.Lock()
	defer mu.Unlock()

	lockerInstance, ok := s.cache.Get(lockName)
	if ok {
		locker := lockerInstance.(*cacheLock)
		if !locker.mu.TryLock() {
			err := ErrorLockerFailed(lockName, stderrors.New("try lock failed"))
			return locker, err
		}
		return locker, nil
	}
	locker := newCacheLock(lockName, s.cache, &sync.Mutex{}, false)
	locker.mu.Lock()

	s.cache.Set(lockName, locker, _cacheLockerExpire)

	return locker, nil
}

func (s *cache) Unlock(ctx context.Context, lockName string) {
	lockerInstance, ok := s.cache.Get(lockName)
	if !ok {
		return
	}
	locker := lockerInstance.(*cacheLock)
	locker.mu.TryLock()
	_, _ = locker.Unlock(ctx)
	return
}

// cacheLock ...
type cacheLock struct {
	lockName   string
	cache      *cachepkg.Cache
	mu         *sync.Mutex
	isMutex    bool
	stopExtend chan bool
}

func newCacheLock(lockName string, cache *cachepkg.Cache, mu *sync.Mutex, isMutex bool) *cacheLock {
	l := &cacheLock{
		lockName:   lockName,
		cache:      cache,
		mu:         mu,
		isMutex:    isMutex,
		stopExtend: nil,
	}
	if isMutex {
		l.stopExtend = make(chan bool)
	}
	return l
}

// Unlock ...
func (s *cacheLock) Unlock(ctx context.Context) (bool, error) {
	// 防止panic 信道
	if s.stopExtend != nil {
		s.stopExtend <- true
		close(s.stopExtend)
	}

	_ = s.mu.TryLock()
	s.mu.Unlock()
	return true, nil
}

// extend ...
func (s *cacheLock) extend(ctx context.Context) {
	// 计时器
	timer := time.NewTimer(_cacheLockExtendDelay)
	defer timer.Stop()

	select {
	case <-timer.C: // 续期
		// 续期
		lockerInstance, ok := s.cache.Get(s.lockName)
		if !ok {
			return
		}
		locker := lockerInstance.(*cacheLock)
		if locker.mu.TryLock() {
			_, _ = locker.Unlock(ctx)
			return
		}
		s.cache.Set(s.lockName, locker, _cacheLockerExpire)
		// 调试
		// fmt.Println("cache mutex 续期成功")
		// 再次续期
		s.extend(ctx)
	case <-s.stopExtend: // 停止
		// fmt.Println("cache mutex 停止续期")
		return
	}
}
