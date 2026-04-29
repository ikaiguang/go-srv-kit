# redis - Redis 客户端

`redis/` 提供 Redis 客户端创建和分布式锁实现。

## 包名

```go
import redispkg "github.com/ikaiguang/go-srv-kit/data/redis"
```

## 客户端

```go
client, err := redispkg.NewDB(config)
```

支持单机和集群模式（`redis.UniversalClient`）。

## 分布式锁

基于 [go-redsync](https://github.com/go-redsync/redsync)：

```go
locker := redispkg.NewLocker(redisClient)

// 一次性锁（不续期，8 秒过期）
unlocker, err := locker.Once(ctx, "lock:key")
defer unlocker.Unlock(ctx)

// 互斥锁（自动续期，防止锁过期）
unlocker, err := locker.Mutex(ctx, "lock:key")
defer unlocker.Unlock(ctx)
```

| 文件 | 说明 |
|------|------|
| `redis.kit.go` | 客户端创建 |
| `redis_lock.kit.go` | 分布式锁（Locker） |
| `redis_lock_once.kit.go` | 一次性锁实现 |
| `redis_lock_mutex.kit.go` | 互斥锁实现（自动续期） |

通常通过 `service/redis` 管理器间接使用。
