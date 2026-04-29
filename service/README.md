# service - 服务层工具

`service/` 是 go-srv-kit 的服务层，负责基础设施的统一初始化、配置管理和服务器生命周期管理。

业务服务通过 `LauncherManager` 按需获取各类基础设施组件（数据库、缓存、消息队列等），无需关心底层初始化细节。

## 模块导入

```go
import "github.com/ikaiguang/go-srv-kit/service"
```

## 目录结构

| 目录 | 包名 | 说明 |
|------|------|------|
| `setup/` | `setuputil` | LauncherManager 核心入口，统一管理所有基础设施组件的生命周期 |
| `server/` | `serverutil` | HTTP/gRPC 服务器创建、启动和 All-In-One 模式 |
| `config/` | `configutil` | 配置加载（文件 / Consul），配置项便捷访问 |
| `logger/` | `loggerutil` | 日志管理器，支持控制台、文件、GORM、RabbitMQ 分类日志 |
| `middleware/` | `middlewareutil` | 中间件设置，JWT 白名单管理 |
| `auth/` | `authutil` | 认证实例管理，Token 和 Auth 的懒加载初始化 |
| `app/` | `apputil` | 应用标识、环境信息、HTTP 编解码器 |
| `cleanup/` | `cleanuputil` | 资源清理管理器（后进先出） |
| `database/` | `dbutil` | 数据库迁移辅助工具 |
| `tracer/` | `tracerutil` | OpenTelemetry 链路追踪初始化 |
| `store/` | `storeutil` | 配置文件存储（Consul） |
| `cluster_service_api/` | `clientutil` | 集群服务间 API 调用客户端 |
| `consul/` | `consulutil` | Consul 客户端管理器 |
| `mysql/` | `mysqlutil` | MySQL 连接管理器 |
| `postgres/` | `postgresutil` | PostgreSQL 连接管理器 |
| `mongo/` | `mongoutil` | MongoDB 客户端管理器 |
| `redis/` | `redisutil` | Redis 客户端管理器 |
| `rabbitmq/` | `rabbitmqutil` | RabbitMQ 连接管理器 |
| `jaeger/` | `jaegerutil` | Jaeger 链路追踪管理器 |

## 核心概念：LauncherManager

`LauncherManager` 是所有基础设施组件的统一入口，采用懒加载模式，按需初始化：

```go
// 创建 LauncherManager（推荐方式）
lm, cleanup, err := setuputil.NewWithCleanup(configFilePath)
defer cleanup()

// 按需获取组件
logger, err := lm.GetLogger()
db, err := lm.GetMysqlDBConn()
redisClient, err := lm.GetRedisClient()
```

## 快速启动服务

```go
// All-In-One 模式：一行代码启动多服务
app, cleanup, err := serverutil.AllInOneServer(
    configFilePath,
    configOpts,
    services,
    authWhitelist,
)
defer cleanup()
serverutil.RunServer(app, nil)
```

## 与其他模块的关系

```
service/setup (LauncherManager)
  ├── 使用 data/* 组件创建底层连接
  ├── 使用 kratos/log 初始化日志
  ├── 使用 kratos/auth 初始化认证
  └── 使用 kratos/middleware 配置中间件

service/server
  ├── 使用 service/setup 获取基础设施
  ├── 使用 service/middleware 配置中间件链
  └── 创建 kratos.App 启动服务
```

## 参考

- 示例服务：[testdata/ping-service](../testdata/ping-service/)
- Wire 装配：`testdata/ping-service/cmd/ping-service/export/wire.go`
