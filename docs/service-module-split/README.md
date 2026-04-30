# service 模块拆分：按需编译依赖

## 要做什么

将 `service/` 从一个大模块拆分为多个独立子模块，使业务服务（如 ping-service）在编译层面只引入实际需要的依赖，而不是被迫拉入 Redis、MySQL、Mongo、Consul 等全部基础设施。

## 为什么这么做

当前 `service/go.mod` 是一个整体模块，包含了所有基础设施子包（consul、redis、mysql、postgres、mongo、rabbitmq、jaeger）。即使 ping-service 只用了 `service/setup`、`service/server`、`service/cluster_service_api`，但因为它们都在同一个 `service/` 模块下，`go.mod` 必须声明所有依赖。

这导致：
- ping-service 的 `go.mod` 一旦 require `go-srv-kit/service`，就会传递引入 consul、redis、mysql、mongo、rabbitmq 等全部间接依赖
- 编译依赖图膨胀，与"按需引入"的设计意图不符
- 运行时虽然已经是按需的（WithSetup 模式），但编译时不是

## 当前问题分析

### service/go.mod 的依赖全景

`service/go.mod` 直接 require 了：
- `data/consul` → consul API
- `data/gorm` → gorm ORM
- `data/jaeger` → jaeger tracing
- `data/mongo` → mongo driver
- `data/mysql` → mysql driver
- `data/postgres` → postgres driver
- `data/rabbitmq` → rabbitmq (watermill)
- `data/redis` → redis client
- `go-redis/v9` → redis
- `etcd/client/v3` → etcd
- `mongo-driver/v2` → mongo
- `gorm.io/gorm` → gorm
- `consul/api` → consul

### 哪些 service 子包引入了重型依赖

| 子包 | 引入的重型依赖 | 被 ping-service 使用 |
|------|---------------|---------------------|
| `service/setup` | 无（只依赖 logger、config、kratos） | ✅ |
| `service/server` | 无（只依赖 setup、middleware、tracer、app） | ✅ |
| `service/config` | 无（纯配置加载） | ✅ |
| `service/logger` | 无（纯日志） | ✅ |
| `service/app` | 无 | ✅ |
| `service/middleware` | 无（依赖 kratos auth） | ✅ |
| `service/cleanup` | 无 | ✅ |
| `service/tracer` | 无（otel） | ✅ |
| `service/cluster_service_api` | 无（gRPC/HTTP client） | ✅ |
| `service/database` | gorm（MigrationFunc 签名含 *gorm.DB） | ❌（通过 export 间接引入） |
| `service/auth` | go-redis、data/redis、service/redis | ❌ |
| `service/store` | consul API | ❌ |
| `service/consul` | consul API、consul registry | ❌ |
| `service/redis` | go-redis、data/redis | ❌ |
| `service/mysql` | gorm、data/mysql、data/gorm | ❌ |
| `service/postgres` | gorm、data/postgres、data/gorm | ❌ |
| `service/mongo` | mongo-driver、data/mongo | ❌ |
| `service/rabbitmq` | watermill、data/rabbitmq | ❌ |
| `service/jaeger` | jaeger exporter、data/jaeger | ❌ |

### 关键发现

`service/setup`（核心）本身**不直接 import** 任何重型子包。各基础设施子包通过 `WithSetup()` Option 模式注册，import 关系是**调用方**（如 main.go）决定的。

**问题的本质**：这些子包虽然代码上解耦了，但因为在同一个 `go.mod` 下，编译依赖无法分离。

## 方案

### 拆分策略

将 `service/` 拆分为：

1. **`service/` 核心模块**（保留现有 `service/go.mod`）
   - 包含：setup、server、config、logger、app、middleware、cleanup、tracer、cluster_service_api、database
   - 这些是任何业务服务都需要的基础能力
   - `go.mod` 只依赖 kratos、kit、kratos 扩展，不依赖任何 data/* 子模块

2. **各基础设施子模块**（各自独立 `go.mod`）
   - `service/consul/go.mod` → 依赖 data/consul
   - `service/redis/go.mod` → 依赖 data/redis
   - `service/mysql/go.mod` → 依赖 data/mysql、data/gorm
   - `service/postgres/go.mod` → 依赖 data/postgres、data/gorm
   - `service/mongo/go.mod` → 依赖 data/mongo
   - `service/rabbitmq/go.mod` → 依赖 data/rabbitmq
   - `service/jaeger/go.mod` → 依赖 data/jaeger
   - `service/auth/go.mod` → 依赖 service/redis、data/redis、go-redis
   - `service/store/go.mod` → 依赖 consul（或合并到 service/consul）
   - `service/database/go.mod` → 依赖 gorm（MigrationFunc 签名含 *gorm.DB）

### ping-service 的 go.mod

拆分后，ping-service 的 `go.mod` 只需要：

```
require (
    github.com/ikaiguang/go-srv-kit/service v0.0.0  // 核心：setup、server、config、logger
    github.com/ikaiguang/go-srv-kit/kratos v0.0.0   // 框架扩展
    github.com/ikaiguang/go-srv-kit/kit v0.0.0       // 工具库
    github.com/ikaiguang/go-srv-kit v0.0.0           // api/config proto
)
```

不会引入 Redis、MySQL、Mongo、Consul、RabbitMQ 等任何重型依赖。

### 后续扩展示例

如果 ping-service 以后需要 Redis：

```go
// main.go
import redisutil "github.com/ikaiguang/go-srv-kit/service/redis"

setupOpts := []serverutil.Option{
    serverutil.WithSetupOptions(
        clientutil.WithSetup(),
        redisutil.WithSetup(),  // 新增这一行
    ),
}
```

`go.mod` 新增：
```
require github.com/ikaiguang/go-srv-kit/service/redis v0.0.0
```

## 任务列表

### 阶段一：拆分 service 子模块

1. [ ] 为 `service/consul` 创建独立 `go.mod`
2. [ ] 为 `service/redis` 创建独立 `go.mod`
3. [ ] 为 `service/mysql` 创建独立 `go.mod`
4. [ ] 为 `service/postgres` 创建独立 `go.mod`
5. [ ] 为 `service/mongo` 创建独立 `go.mod`
6. [ ] 为 `service/rabbitmq` 创建独立 `go.mod`
7. [ ] 为 `service/jaeger` 创建独立 `go.mod`
8. [ ] 处理 `service/store`（依赖 consul，合并到 service/consul 或独立）
9. [ ] 清理 `service/go.mod`，移除已拆出子模块的依赖
10. [ ] 更新 `go.work`，添加新的子模块路径

### 阶段二：完善 ping-service 独立 go.mod

11. [ ] 补全 `testdata/ping-service/go.mod` 的 require 和 replace
12. [ ] 将 `testdata/ping-service` 加入 `go.work`
13. [ ] 验证 ping-service 可以独立编译

### 阶段三：验证

14. [ ] 确认 ping-service 的依赖图不包含 redis/mysql/mongo/consul 等
15. [ ] 确认现有的 all-in-one 入口仍然正常工作
16. [ ] 确认 Wire 生成正常

## 风险点

1. `service/store` 包依赖 consul，需要确认是否有其他包引用它
2. `service/go.mod` 清理后需要确认没有遗漏的交叉引用
3. `go.work` 更新后需要确认所有模块能正确解析
4. 现有使用全量 `service` 模块的代码（如 all-in-one）需要同步更新 import
