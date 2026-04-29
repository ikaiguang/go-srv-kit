# go-srv-kit 模块化迁移指南

## 概述

go-srv-kit 的服务启动链路已经从“核心包隐式引入大量基础设施组件”调整为“业务入口显式声明所需组件”。

这次重构的核心目标：

- **按需依赖**：业务服务只 import 实际使用的组件，减少编译依赖图。
- **显式注册**：Redis、MySQL、PostgreSQL、MongoDB、Consul、Jaeger、RabbitMQ、Auth、cluster service API 等组件由业务入口通过 `WithSetup()` 注册。
- **核心收窄**：`service/setup` 只保留配置、日志、组件注册表、生命周期和关闭能力。
- **懒加载**：组件注册后仍在首次使用时初始化；需要启动时校验连接时再配合 `setuputil.WithEagerInit(...)`。

## 迁移前后对比

| 迁移前 | 迁移后 |
|---|---|
| import 核心启动包可能拉入所有基础设施依赖 | 业务入口 import 哪个组件包，才引入对应依赖 |
| `LauncherManager` 直接组合数据库、Redis、Consul 等 Provider | `LauncherManager` 只提供配置、日志、注册表和生命周期 |
| 通过核心包里的 `WithXxx()` 或全量注册注册组件 | 通过各 `service/<component>.WithSetup()` 注册组件 |
| 在 `lm.GetXxx()` 上直接取基础设施 | 在对应组件包调用 `GetXxx(lm)` |
| Consul 配置加载在 `service/config` 中 | Consul 相关能力归属 `service/consul` |

## 目录职责

```text
service/
├── setup/              # LauncherManager 核心：配置、日志、注册表、生命周期
├── server/             # HTTP/gRPC server 和 All-In-One 启动
├── config/             # 文件配置加载和基础配置工具
├── cluster_service_api/# 集群服务 API 客户端组件
├── auth/               # 认证组件
├── consul/             # Consul 客户端、配置加载、registry 接入
├── mysql/              # MySQL 组件
├── postgres/           # PostgreSQL 组件
├── mongo/              # MongoDB 组件
├── redis/              # Redis 组件
├── rabbitmq/           # RabbitMQ 组件
└── jaeger/             # Jaeger 组件
```

## 基础迁移方式

### 只需要核心启动能力

如果服务只需要配置、日志、生命周期，不需要数据库、Redis 等组件：

```go
import setuputil "github.com/ikaiguang/go-srv-kit/service/setup"

lm, cleanup, err := setuputil.NewWithCleanup(configPath)
if err != nil {
    return err
}
defer cleanup()
```

`NewWithCleanup` 不会自动注册所有基础设施组件。

### 需要基础设施组件

业务入口显式 import 对应组件包，并传入 `WithSetup()`：

```go
import (
    postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
    redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
    setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

lm, cleanup, err := setuputil.NewWithCleanupOptions(
    configPath,
    nil,
    postgresutil.WithSetup(),
    redisutil.WithSetup(),
)
if err != nil {
    return err
}
defer cleanup()

db, err := postgresutil.GetDB(lm)
redisClient, err := redisutil.GetClient(lm)
```

### All-In-One 服务入口

`serverutil.AllInOneServer` 通过 `serverutil.WithSetupOptions(...)` 接收组件注册选项：

```go
import (
    clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
    serverutil "github.com/ikaiguang/go-srv-kit/service/server"
)

runOpts := []serverutil.Option{
    serverutil.WithSetupOptions(
        clientutil.WithSetup(),
    ),
}

app, cleanup, err := serverutil.AllInOneServer(
    flagconf,
    configOpts,
    services,
    whitelist,
    runOpts...,
)
if err != nil {
    return err
}
serverutil.RunServer(app, cleanup)
```

如果业务服务还需要 PostgreSQL、Redis 或 Auth，在同一个 `WithSetupOptions` 中继续加入：

```go
runOpts := []serverutil.Option{
    serverutil.WithSetupOptions(
        clientutil.WithSetup(),
        postgresutil.WithSetup(),
        redisutil.WithSetup(),
        authutil.WithSetup(),
    ),
}
```

Auth 依赖 Redis，注册 Auth 时应同时注册 Redis。

## 组件注册和获取

| 组件 | 注册方式 | 获取方式 |
|---|---|---|
| Redis | `redisutil.WithSetup()` | `redisutil.GetClient(lm)` / `redisutil.GetNamedClient(lm, name)` |
| MySQL | `mysqlutil.WithSetup()` | `mysqlutil.GetDB(lm)` / `mysqlutil.GetNamedDB(lm, name)` |
| PostgreSQL | `postgresutil.WithSetup()` | `postgresutil.GetDB(lm)` / `postgresutil.GetNamedDB(lm, name)` |
| MongoDB | `mongoutil.WithSetup()` | `mongoutil.GetClient(lm)` / `mongoutil.GetNamedClient(lm, name)` |
| Consul | `consulutil.WithSetup()` | `consulutil.GetClient(lm)` / `consulutil.GetNamedClient(lm, name)` |
| Jaeger | `jaegerutil.WithSetup()` | `jaegerutil.GetExporter(lm)` / `jaegerutil.GetNamedExporter(lm, name)` |
| RabbitMQ | `rabbitmqutil.WithSetup()` | `rabbitmqutil.GetConn(lm)` / `rabbitmqutil.GetNamedConn(lm, name)` |
| Auth | `authutil.WithSetup()` | `authutil.GetTokenManager(lm)` / `authutil.GetAuthManager(lm)` |
| Cluster service API | `clientutil.WithSetup()` | `clientutil.GetManager(lm)` |

`setuputil.WithComponentRegistrar` 是组件包内部注册扩展点。业务入口优先使用组件包的 `WithSetup()`，不要直接写注册表细节。

## 急切初始化

默认情况下，组件采用懒加载模式。需要启动时立即验证连接时，配合 `setuputil.WithEagerInit(...)`：

```go
lm, err := setuputil.New(conf,
    postgresutil.WithSetup(),
    redisutil.WithSetup(),
    setuputil.WithEagerInit(
        setuputil.ComponentPostgres,
        setuputil.ComponentRedis,
    ),
)
```

如果对未注册组件执行急切初始化，会返回 `component not registered` 错误。

## 多实例支持

同类型组件支持多实例，例如：

```go
db, err := postgresutil.GetDB(lm)
orderDB, err := postgresutil.GetNamedDB(lm, "order-db")

redisClient, err := redisutil.GetClient(lm)
sessionRedis, err := redisutil.GetNamedClient(lm, "session")
```

多实例在配置文件中通过 `xxx_instances` 字段配置：

```yaml
psql:
  enable: true
  dsn: "default-dsn"
psql_instances:
  order-db:
    enable: true
    dsn: "order-db-dsn"
```

## Consul 配置加载

Consul 相关能力不再放在 `service/config` 核心包中。

需要从 Consul 读取配置或接入 Consul registry 时，应显式使用 `service/consul` 对应能力，并在入口中 import `consulutil`。

这样默认服务入口不会因为配置包而隐式拉入 Consul 依赖。

## go.work 配置说明

本仓库使用 `go.work` 进行本地多模块联合开发。根目录 `go.work` 包含根模块、`service`、`kit`、`kratos` 和多个 `data/*` 子模块。

不要在根模块随手执行 `go mod tidy`。根目录 `go.mod` 已说明：`testdata/` 不会被 Go 的 `./...` 包模式纳入，`tidy` 可能移除示例服务依赖。

## 常见问题

### Q1: 报错 "component not registered"

**原因**：调用了某个组件包的 `GetXxx(lm)`，但入口没有传入对应组件包的 `WithSetup()`。

**解决方式**：在入口注册所需组件。

```go
lm, cleanup, err := setuputil.NewWithCleanupOptions(
    configPath,
    nil,
    redisutil.WithSetup(),
)
```

All-In-One 模式：

```go
runOpts := []serverutil.Option{
    serverutil.WithSetupOptions(redisutil.WithSetup()),
}
```

### Q2: 使用 NewWithCleanup 是否会注册所有组件？

不会。`NewWithCleanup` 只加载配置并创建核心 `LauncherManager`。需要基础设施组件时，使用 `NewWithCleanupOptions` 并显式传入组件包的 `WithSetup()`。

### Q3: 默认 ping-service 入口为什么只注册 cluster service API？

示例服务默认入口只演示服务启动和集群服务 API 调用，因此只注册：

```go
serverutil.WithSetupOptions(
    clientutil.WithSetup(),
)
```

这样默认入口不会因为未使用的组件而拉入 Consul、Etcd、数据库或消息队列依赖。

### Q4: 数据库迁移入口怎么注册数据库？

数据库迁移入口显式注册 PostgreSQL：

```go
launcher, cleanup, err := setuputil.NewWithCleanupOptions(
    flagconf,
    nil,
    postgresutil.WithSetup(),
)
```

然后通过：

```go
db, err := postgresutil.GetDB(launcher)
```

### Q5: Wire 生成代码报错怎么办？

修改 provider 或构造函数后，运行：

```bash
wire ./testdata/ping-service/cmd/ping-service/export
```

不要手工修改生成的 `wire_gen.go`。
