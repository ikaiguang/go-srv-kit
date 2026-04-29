# go-srv-kit 模块化迁移指南

## 概述

go-srv-kit 从单体模块重构为多模块架构，将基础设施组件（MySQL、Redis、MongoDB、RabbitMQ 等）拆分为独立的 Go 模块。这次重构的核心目标：

- **按需依赖**：业务服务只引入实际使用的组件，减少不必要的依赖
- **独立版本管理**：各组件可独立发布和升级
- **编译加速**：减少依赖图大小，加快编译速度
- **懒加载 + 按需注入**：通过 Registry 模式和 `WithXxx()` Option 实现组件的按需注册和懒加载初始化

## 模块化前后对比

### 模块结构变化

| 变更前（单体模块） | 变更后（多模块） |
|---|---|
| `github.com/ikaiguang/go-srv-kit/service` 包含所有组件 | `service/` 仍为核心模块，但各数据组件独立 |
| 所有数据组件在 `service/` 内部 | `data/mysql`、`data/redis` 等为独立模块 |
| 启动时初始化所有已启用的组件 | 通过 `WithXxx()` Option 按需注册，懒加载初始化 |
| `NewLauncherManager()` 自动初始化所有组件 | `New()` + `WithXxx()` 按需注册 |

### 目录结构变化

```
go-srv-kit/
├── service/              # 核心服务模块（go.mod）
│   ├── setup/            # LauncherManager + Registry 模式
│   ├── auth/             # 认证组件
│   ├── config/           # 配置加载
│   ├── server/           # HTTP/gRPC 服务器
│   └── ...
├── data/                 # 数据组件（各自独立 go.mod）
│   ├── mysql/            # MySQL 组件（独立模块）
│   ├── redis/            # Redis 组件（独立模块）
│   ├── postgres/         # PostgreSQL 组件（独立模块）
│   ├── mongo/            # MongoDB 组件（独立模块）
│   ├── consul/           # Consul 组件（独立模块）
│   ├── jaeger/           # Jaeger 组件（独立模块）
│   ├── rabbitmq/         # RabbitMQ 组件（独立模块）
│   └── gorm/             # GORM 通用组件（独立模块）
├── kit/                  # 通用工具库（独立模块）
└── kratos/               # Kratos 框架扩展（独立模块）
```

## 向后兼容说明

**重要：现有代码无需修改即可正常工作。**

`NewWithCleanup` 和 `NewLauncherManagerWithCleanup` 函数保持向后兼容，内部自动调用 `WithAllComponents()` 注册所有组件。如果你的服务使用 `AllInOneServer` 或 `NewWithCleanup`，升级后无需任何代码变更。

```go
// 这些调用方式仍然有效，无需修改
lm, cleanup, err := setuputil.NewWithCleanup(configPath)
lm, cleanup, err := setuputil.NewLauncherManagerWithCleanup(configPath)
app, cleanup, err := serverutil.AllInOneServer(flagconf, configOpts, services, whitelist)
```

## 导入路径变更

### go.mod 配置

如果你的项目使用 `go.work` 进行本地开发，需要将新的子模块加入 `go.work`：

```go
// go.work
go 1.25.9

use (
    .
    ./data/consul
    ./data/etcd
    ./data/gorm
    ./data/jaeger
    ./data/migration
    ./data/mongo
    ./data/mysql
    ./data/postgres
    ./data/rabbitmq
    ./data/redis
    ./kit
    ./kratos
    ./service
)
```

### 业务服务的 go.mod

业务服务的 `go.mod` 中，`service` 模块会自动传递引入所需的 `data/*` 子模块。如果你直接使用 `data/*` 包的类型，需要显式添加依赖：

```go
// go.mod
module your-service

require (
    github.com/ikaiguang/go-srv-kit/service v0.x.x
    // 仅在直接 import data/* 包时需要
    github.com/ikaiguang/go-srv-kit/data/mysql v0.x.x
    github.com/ikaiguang/go-srv-kit/data/redis v0.x.x
)
```

本地开发时使用 `replace` 指令：

```go
replace github.com/ikaiguang/go-srv-kit/service => ../go-srv-kit/service
replace github.com/ikaiguang/go-srv-kit/data/mysql => ../go-srv-kit/data/mysql
replace github.com/ikaiguang/go-srv-kit/data/redis => ../go-srv-kit/data/redis
```

## API 变更

### LauncherManager 创建方式

#### 方式一：向后兼容（推荐升级初期使用）

```go
import setuputil "github.com/ikaiguang/go-srv-kit/service/setup"

// 自动注册所有组件，与旧版行为一致
lm, cleanup, err := setuputil.NewWithCleanup(configPath)
```

#### 方式二：按需注入（推荐新项目使用）

```go
import (
    configutil "github.com/ikaiguang/go-srv-kit/service/config"
    setuputil  "github.com/ikaiguang/go-srv-kit/service/setup"
)

// 1. 加载配置
conf, err := configutil.Loading(configFilePath)

// 2. 按需注册组件
lm, err := setuputil.New(conf,
    setuputil.WithAllComponents(),  // 注册所有组件（向后兼容）
)
```

### WithXxx() Option 按需注入

`WithAllComponents()` 会注册所有组件。如果你只需要部分组件，可以使用 `WithComponentRegistrar` 自定义注册：

```go
// 注册所有组件（向后兼容）
lm, err := setuputil.New(conf, setuputil.WithAllComponents())
```

### WithEagerInit 急切初始化

默认情况下，组件采用懒加载模式（首次调用 `GetXxxClient()` 时才初始化）。如果需要在启动时立即初始化某些组件：

```go
lm, err := setuputil.New(conf,
    setuputil.WithAllComponents(),
    setuputil.WithEagerInit(
        setuputil.ComponentRedis,
        setuputil.ComponentMysql,
    ),
)
```

### 组件名称常量

```go
package setuputil

const (
    ComponentLogger     = "logger"      // 日志（始终注册）
    ComponentRedis      = "redis"       // Redis
    ComponentMysql      = "mysql"       // MySQL
    ComponentPostgres   = "postgres"    // PostgreSQL
    ComponentMongo      = "mongo"       // MongoDB
    ComponentConsul     = "consul"      // Consul
    ComponentJaeger     = "jaeger"      // Jaeger
    ComponentRabbitmq   = "rabbitmq"    // RabbitMQ
    ComponentAuth       = "auth"        // 认证
    ComponentServiceAPI = "serviceAPI"  // 集群服务 API
)
```

### 多实例支持

模块化后支持同类型组件的多实例（如多个 MySQL 数据库连接）：

```go
// 默认实例
db, err := lm.GetMysqlDBConn()

// 命名实例
orderDB, err := lm.GetNamedMysqlDBConn("order-db")
userDB, err := lm.GetNamedMysqlDBConn("user-db")

// Redis 命名实例
cacheRedis, err := lm.GetNamedRedisClient("cache")
sessionRedis, err := lm.GetNamedRedisClient("session")
```

多实例在配置文件中通过 `xxx_instances` 字段配置：

```yaml
mysql:
  enable: true
  dsn: "default-dsn"
mysql_instances:
  order-db:
    enable: true
    dsn: "order-db-dsn"
  user-db:
    enable: true
    dsn: "user-db-dsn"
```

### NewTokenManager 签名变更

`NewTokenManager` 函数签名保持不变，但通过 `LauncherManager` 获取 `TokenManager` 的方式有所调整：

```go
// 旧版：直接从 LauncherManager 获取（内部自动初始化）
tokenManager, err := lm.GetTokenManager()

// 新版：相同调用方式，但需要确保 Auth 组件已注册
// 使用 WithAllComponents() 时自动注册
// 未注册时调用会返回 "component not registered: auth" 错误
tokenManager, err := lm.GetTokenManager()
```

Auth 组件依赖 Redis，因此注册 Auth 时需要同时注册 Redis。

## go.work 配置说明

`go.work` 用于本地多模块联合开发。在 go-srv-kit 项目根目录下：

```go
// go.work
go 1.25.9

use (
    .                    // 根模块
    ./data/consul        // Consul 数据组件
    ./data/etcd          // Etcd 数据组件
    ./data/gorm          // GORM 通用组件
    ./data/jaeger        // Jaeger 数据组件
    ./data/migration     // 数据库迁移组件
    ./data/mongo         // MongoDB 数据组件
    ./data/mysql         // MySQL 数据组件
    ./data/postgres      // PostgreSQL 数据组件
    ./data/rabbitmq      // RabbitMQ 数据组件
    ./data/redis         // Redis 数据组件
    ./kit                // 通用工具库
    ./kratos             // Kratos 框架扩展
    ./service            // 核心服务模块
)
```

使用 `go.work` 后，所有模块间的 `replace` 指令由 workspace 自动处理，无需在各模块的 `go.mod` 中手动添加。

### 编译命令

```bash
# 联合编译所有模块
go build ./...

# 编译特定模块
go build ./service/...
go build ./data/mysql/...

# 编译示例服务
go build ./testdata/ping-service/cmd/ping-service/...
go build ./testdata/ping-service/cmd/all-in-one/...
```

## 常见问题和解决方案

### Q1: 升级后编译报错 "component not registered"

**原因**：使用了 `setuputil.New()` 但未注册所需组件。

**解决方案**：使用 `WithAllComponents()` 注册所有组件，或确保注册了所需的组件：

```go
lm, err := setuputil.New(conf, setuputil.WithAllComponents())
```

### Q2: 使用 NewWithCleanup 是否需要修改代码？

**不需要。** `NewWithCleanup` 内部自动调用 `WithAllComponents()`，保持完全向后兼容。

### Q3: go mod tidy 后根模块缺少依赖

**原因**：`testdata/` 目录被 Go 的 `./...` 模式排除，`go mod tidy` 会移除 testdata 所需的依赖。

**解决方案**：不要对根模块执行 `go mod tidy`，或执行后手动恢复 testdata 所需的依赖。根模块的 `go.mod` 文件头部有相关注释说明。

### Q4: 如何确认组件是否已注册？

调用 `GetXxxClient()` 时，如果组件未注册会返回明确的错误信息：

```
component not registered: redis; use corresponding WithXxx() option
```

根据错误信息添加对应的组件注册即可。

### Q5: 懒加载和急切初始化如何选择？

- **懒加载（默认）**：组件在首次使用时才初始化，适合大多数场景
- **急切初始化**：使用 `WithEagerInit()` 在启动时立即初始化，适合需要在启动时验证连接的场景（如数据库、缓存）

```go
// 启动时立即验证 Redis 和 MySQL 连接
lm, err := setuputil.New(conf,
    setuputil.WithAllComponents(),
    setuputil.WithEagerInit(setuputil.ComponentRedis, setuputil.ComponentMysql),
)
```

### Q6: 多实例配置不生效

**确认事项**：
1. 配置文件中使用了 `xxx_instances` 字段（如 `mysql_instances`）
2. 实例名称与代码中 `GetNamedXxxClient("name")` 的参数一致
3. 对应组件已注册（使用 `WithAllComponents()` 或单独注册）

### Q7: Wire 生成代码报错

**解决方案**：

```bash
# 删除旧的生成文件
rm ./cmd/*/export/wire_gen.go

# 重新生成
wire ./cmd/*/export
```

确保 Wire 定义文件中的 Provider 顺序正确：基础设施 → Data 层 → Business 层 → Service 层。
