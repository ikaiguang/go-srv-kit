# redis

`redis` 目录提供包 `redispkg`，用于创建 Redis 客户端、判断 `redis.Nil` 错误，并提供基于 `go-redsync/redsync` 的 Redis 分布式锁封装。该包适合在服务端项目的数据访问层或基础设施层中使用。

## 安装

当前 `go.mod` 模块路径为：

```bash
go get github.com/ikaiguang/go-srv-kit/data/redis
```

导入本包：

```go
import redispkg "github.com/ikaiguang/go-srv-kit/data/redis/redis"
```

## 核心能力

- `Config`：由 `redis/config.proto` 生成的 Redis 配置结构，包含地址、认证、DB、超时和连接池参数。
- `NewDB(conf *Config)`：根据 `Config` 创建 `redis.UniversalClient`，并通过 `Ping` 验证连接。
- `IsNilErr(err error)`：判断错误是否为 `redis.Nil`。
- `NewLocker(redisCC redis.UniversalClient, opts ...redsync.Option)`：创建实现 `github.com/ikaiguang/go-srv-kit/kit/locker.Locker` 的分布式锁。
- `Locker.Once(ctx, lockName)`：获取一次性锁，不启动自动续期。
- `Locker.Mutex(ctx, lockName)`：获取互斥锁，并按内部间隔续期，直到调用 `Unlock`。

## 快速使用

创建客户端并写入键值：

```go
conf := &redispkg.Config{
	Addresses:       []string{"127.0.0.1:6379"},
	DialTimeout:     durationpb.New(3 * time.Second),
	ReadTimeout:     durationpb.New(3 * time.Second),
	WriteTimeout:    durationpb.New(3 * time.Second),
	ConnMaxActive:   100,
	ConnMaxLifetime: durationpb.New(30 * time.Minute),
	ConnMaxIdle:     10,
	ConnMinIdle:     1,
	ConnMaxIdleTime: durationpb.New(time.Hour),
}

db, err := redispkg.NewDB(conf)
if err != nil {
	return err
}
defer db.Close()

err = db.Set(context.Background(), "example:key", "value", time.Minute).Err()
if err != nil {
	return err
}
```

判断空值错误：

```go
value, err := db.Get(ctx, "missing:key").Result()
if redispkg.IsNilErr(err) {
	return "", nil
}
if err != nil {
	return "", err
}
return value, nil
```

使用分布式锁：

```go
locker := redispkg.NewLocker(db)

unlock, err := locker.Mutex(ctx, "task:sync")
if err != nil {
	return err
}
defer unlock.Unlock(context.Background())
```

## 测试

本包测试会连接默认 Redis 地址 `127.0.0.1:6379`，属于集成测试。运行前请先启动本地 Redis，或按测试代码调整环境。

```bash
go test ./redis
```

也可以运行指定用例：

```bash
go test -v ./redis -count=1 -run TestNewDB_Xxx
go test -v ./redis -count=1 -run TestLockOnce
go test -v ./redis -count=1 -run TestLockMutex
```

## 生成代码

`config.pb.go` 和 `config.pb.validate.go` 是生成文件，不要手动修改。修改 `redis/config.proto` 后，按 `Makefile` 目标重新生成：

```bash
make protoc-config-protobuf
```

## 注意事项

- `NewDB` 会立即执行 `Ping`；Redis 不可达时会返回连接错误。
- `Config` 中的 duration 字段会被直接调用 `AsDuration`，使用前应显式赋值。
- `Password`、Redis 地址和锁名可能包含敏感业务信息，示例和日志中不要输出真实凭据。
- 默认锁过期时间为 8 秒，默认加锁尝试次数为 1；需要改变 redsync 行为时可通过 `NewLocker` 的 `redsync.Option` 参数传入。
