# data - 数据层组件

`data/` 提供各类数据基础设施的底层客户端封装，包括数据库、缓存、消息队列、服务发现和链路追踪。

每个子目录是一个独立的 Go 模块（有自己的 `go.mod`），业务服务按需引入，避免不必要的依赖。

## 目录结构

| 目录 | 包名 | 说明 | 核心依赖 |
|------|------|------|----------|
| `gorm/` | `gormpkg` | GORM ORM 通用工具（连接、日志、分页、排序、批量插入等） | `gorm.io/gorm` |
| `mysql/` | `mysqlpkg` | MySQL 数据库连接，基于 GORM | `gorm.io/driver/mysql` |
| `postgres/` | `psqlpkg` | PostgreSQL 数据库连接，基于 GORM | `gorm.io/driver/postgres` |
| `mongo/` | `mongopkg` | MongoDB 客户端封装，含慢查询监控 | `go.mongodb.org/mongo-driver/v2` |
| `redis/` | `redispkg` | Redis 客户端封装，含分布式锁 | `github.com/redis/go-redis/v9` |
| `rabbitmq/` | `rabbitmqpkg` | RabbitMQ 消息队列，基于 Watermill | `github.com/ThreeDotsLabs/watermill-amqp/v3` |
| `consul/` | `consulpkg` | Consul 服务发现与配置中心客户端 | `github.com/hashicorp/consul/api` |
| `etcd/` | `etcdpkg` | Etcd 服务发现客户端 | `go.etcd.io/etcd/client/v3` |
| `jaeger/` | `jaegerpkg` | Jaeger 分布式追踪 Exporter（HTTP/gRPC） | `go.opentelemetry.io/otel` |
| `migration/` | `migrationpkg` | 数据库迁移框架（建表、删表、自定义迁移） | `gorm.io/gorm` |

## 使用方式

`data/` 中的组件通常不直接使用，而是通过 `service/` 层的管理器间接调用：

```go
// 通过 LauncherManager 获取（推荐）
db, err := launcherManager.GetMysqlDBConn()
redisClient, err := launcherManager.GetRedisClient()

// 直接使用（适用于独立场景）
db, err := mysqlpkg.NewMysqlDB(config, opts...)
redisClient, err := redispkg.NewDB(config)
```

## gorm/ 工具集

`gorm/` 提供了丰富的 ORM 辅助功能：

- **连接管理** - 连接池配置、日志集成
- **分页查询** - `PageQuery` 标准分页
- **排序** - `OrderBy` 安全排序（防 SQL 注入）
- **批量插入** - `BatchInsert` 分批写入
- **事务** - `ExecWithTransaction`、`NewTransaction`
- **锁** - 悲观锁 `ForUpdate`、共享锁 `ForShare`
- **Hint** - 查询提示

## redis/ 分布式锁

```go
locker := redispkg.NewLocker(redisClient)

// 一次性锁（不续期）
unlocker, err := locker.Once(ctx, "lock:order:123")
defer unlocker.Unlock(ctx)

// 互斥锁（自动续期，防止锁过期）
unlocker, err := locker.Mutex(ctx, "lock:payment:456")
defer unlocker.Unlock(ctx)
```

## migration/ 数据库迁移

```go
// 建表迁移
migrator := migrationpkg.NewCreateTable(db.Migrator(), "v1.0.0", &UserPO{})

// 自定义迁移
migrator := migrationpkg.NewAnyMigrator("v1.0.0", "add_index", upFunc, downFunc)
```

## 与 service/ 层的关系

```
service/mysql   → 使用 data/mysql + data/gorm 创建连接
service/redis   → 使用 data/redis 创建连接
service/mongo   → 使用 data/mongo 创建连接
service/rabbitmq → 使用 data/rabbitmq 创建连接
service/consul  → 使用 data/consul 创建连接
service/jaeger  → 使用 data/jaeger 创建 Exporter
```
