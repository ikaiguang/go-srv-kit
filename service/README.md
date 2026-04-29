# service - 服务层工具

`service/` 是 go-srv-kit 的服务层，负责配置加载、日志、组件注册、服务器创建和生命周期管理。

业务服务通过 `LauncherManager` 管理核心启动上下文，并通过各 `service/<component>` 子包按需注册和获取基础设施组件。

## 目录结构

| 目录 | 包名 | 说明 |
|---|---|---|
| `setup/` | `setuputil` | LauncherManager 核心入口，负责配置、日志、组件注册表和生命周期 |
| `server/` | `serverutil` | HTTP/gRPC 服务器创建、启动和 All-In-One 模式 |
| `config/` | `configutil` | 配置加载和配置项便捷访问 |
| `logger/` | `loggerutil` | 日志管理器，支持控制台、文件、GORM、RabbitMQ 分类日志 |
| `middleware/` | `middlewareutil` | 中间件设置，JWT 白名单管理 |
| `auth/` | `authutil` | 认证组件注册和 Token/Auth 管理器获取 |
| `app/` | `apputil` | 应用标识、环境信息、HTTP 编解码器 |
| `cleanup/` | `cleanuputil` | 资源清理管理器（后进先出） |
| `database/` | `dbutil` | 数据库迁移辅助工具 |
| `tracer/` | `tracerutil` | OpenTelemetry 链路追踪初始化 |
| `store/` | `storeutil` | 配置文件存储 |
| `cluster_service_api/` | `clientutil` | 集群服务间 API 调用客户端 |
| `consul/` | `consulutil` | Consul 客户端、配置加载和 registry 接入 |
| `mysql/` | `mysqlutil` | MySQL 组件注册和连接获取 |
| `postgres/` | `postgresutil` | PostgreSQL 组件注册和连接获取 |
| `mongo/` | `mongoutil` | MongoDB 组件注册和客户端获取 |
| `redis/` | `redisutil` | Redis 组件注册和客户端获取 |
| `rabbitmq/` | `rabbitmqutil` | RabbitMQ 组件注册和连接获取 |
| `jaeger/` | `jaegerutil` | Jaeger 组件注册和 exporter 获取 |

## 核心概念

`LauncherManager` 只包含配置、日志、组件注册表、生命周期和关闭能力。具体基础设施由业务入口显式注册：

```go
import (
    postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
    redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
    setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

lm, cleanup, err := setuputil.NewWithCleanupOptions(
    configFilePath,
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

## 快速启动服务

All-In-One 模式下，通过 `serverutil.WithSetupOptions(...)` 声明所需组件：

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
    configFilePath,
    configOpts,
    services,
    authWhitelist,
    runOpts...,
)
if err != nil {
    return err
}
serverutil.RunServer(app, cleanup)
```

## 与其他模块的关系

```text
service/setup
  ├── 保持核心能力：配置、日志、组件注册表、生命周期
  └── 不隐式 import 所有基础设施组件

service/<component>
  ├── 提供 WithSetup() 注册组件
  └── 提供 GetXxx(launcherManager) 获取组件实例

service/server
  ├── 使用 service/setup 创建 LauncherManager
  ├── 通过 WithSetupOptions 接收按需组件注册
  └── 创建 kratos.App 启动服务
```

## 参考

- 示例服务：[testdata/ping-service](../testdata/ping-service/)
- Wire 装配：`testdata/ping-service/cmd/ping-service/export/wire.go`
- LauncherManager 说明：[setup/README.md](setup/README.md)
- All-In-One 说明：[server/README.md](server/README.md)
