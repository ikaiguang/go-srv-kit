# redis - Redis 客户端管理

`redis/` 提供 Redis 客户端的懒加载管理。

## 包名

```go
import redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
```

## 使用

```go
manager, err := redisutil.NewRedisManager(redisConfig)

if manager.Enable() {
    client, err := manager.GetClient()  // redis.UniversalClient
    defer manager.Close()
}
```

底层使用 `data/redis` 创建客户端，支持单机和集群模式。
