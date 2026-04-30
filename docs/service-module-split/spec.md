# Service Module Split

## 背景

用户目标不是拆分 `ping-service` 内部业务 handler，而是让 `ping-service` 作为独立服务模块运行，并在编译依赖和运行初始化两层都只引入实际需要的能力。

当前已有 `testdata/ping-service/go.mod`，但它还没有完整依赖，也没有加入 `go.work`。同时 `service/go.mod` 仍是一个大模块，直接声明了 Redis、MySQL、PostgreSQL、MongoDB、RabbitMQ、Consul、Etcd、Jaeger、GORM 等依赖。只要业务服务 require `github.com/ikaiguang/go-srv-kit/service`，这些可选基础设施就会进入模块图。

已有 `docs/service-module-split/README.md` 描述了总体方向。本规格作为执行入口，并补充源码阅读后发现的边界问题：

- `service/auth` 实际依赖 `data/redis` 和 `go-redis`，不应留在核心 `service` 模块。
- `service/database` 实际暴露 `gorm.DB`，会带入 GORM，不应由普通 `ping-service` 入口间接引入。
- `service/server` 当前通过 `service/tracer` 和 `server_option` 引入 `otlptrace` 类型，需要评估是否保留为核心轻量依赖，或进一步用接口/Option 解耦 Jaeger exporter 类型。
- `testdata/ping-service/cmd/ping-service/export/main.export.go` 同时导出普通服务和数据库迁移函数，导致普通服务入口 import `service/database` 与 `cmd/database-migration/migrate`，会污染最小服务模块依赖。
- 继续阅读后发现 `kratos/go.mod` 也存在同类聚合问题：核心 `kratos` 模块直接 require Consul registry、Etcd registry、Redis、`otlptrace` 等依赖。即使 `service/*` 拆分完成，只要默认 `ping-service` 仍 require `github.com/ikaiguang/go-srv-kit/kratos`，这些依赖仍会进入模块图。
- `service/middleware` 当前只需要 `kratos/auth` 的认证接口和 JWT middleware，但 `kratos/auth` 同包内包含 Redis token manager 实现，导致认证接口包与 Redis 存储实现没有模块边界。

## 目标

- 让 `testdata/ping-service` 成为可独立 `go test` / `go build` 的 Go module。
- `ping-service` 默认服务入口只依赖配置文件加载、日志、server、cluster service api、当前业务代码等实际使用能力。
- `ping-service` 默认模块图不包含 Redis、MySQL、PostgreSQL、MongoDB、RabbitMQ、Consul、Etcd、GORM、数据库迁移等未使用依赖。
- 将 `service/*` 中可选基础设施能力拆出核心 `service` 模块，保留显式 `WithSetup()` / provider 注入方式。
- 将 `kratos` 中可选实现依赖从核心模块图中移走，至少避免默认服务因 `kratos` 核心模块引入 Redis、Consul、Etcd 等依赖。
- 后续业务如果需要 Redis、Mongo、Consul、数据库等，只需 require 对应 service 子模块并在入口显式注入。

## 非目标

- 不拆分 `ping-service` 的 `PingService`、`HomeService`、`WebsocketService` 等业务 handler。
- 不引入运行时自动扫描、反射插件系统或隐藏式全量注册。
- 不执行根模块 `go mod tidy`。
- 不修改 Proto 生成文件，除非后续明确需要并通过约定生成命令产生。
- 不处理真实部署、发布、数据库迁移执行或线上资源访问。
- 不大规模重写认证、注册中心或 tracing 业务语义；本轮只调整模块边界和必要的 provider 适配。

## 影响范围

- `go.work`
- `service/go.mod`
- `service/{consul,redis,mysql,postgres,mongo,rabbitmq,jaeger,auth,database,store}` 及其可能新增的子模块 `go.mod`
- `service/server`、`service/tracer` 中 Jaeger exporter 类型边界
- `kratos/go.mod`
- `kratos/auth` 的 Redis token manager 边界
- `kratos/registry` 的 Consul/Etcd 适配边界
- 必要时新增 `kratos/auth/redis`、`kratos/registry/consul`、`kratos/registry/etcd` 等可选子模块或等价拆分路径
- `testdata/ping-service/go.mod`
- `testdata/ping-service/cmd/ping-service/export`
- `testdata/ping-service/cmd/database-migration`
- `docs/service-module-split/README.md`
- 必要时更新 `docs/migration-guide.md` 或相关 README

## 方案

1. 保留 `service/` 作为核心服务模块，只包含普通业务服务默认会用到的轻量包，例如 `setup`、`server`、`config`、`logger`、`app`、`middleware`、`cleanup`、`cluster_service_api`。
2. 将重型或可选基础设施包拆为独立 service 子模块，例如 `service/redis`、`service/mongo`、`service/consul`、`service/postgres` 等。各子模块继续 import 核心 `service` 模块并提供 `WithSetup()`。
3. 将 `service/auth` 按 Redis 依赖视为可选模块；如果要保留无 Redis auth 抽象，另行设计轻量接口，本轮不默认把 Redis auth 放入核心。
4. 将 `service/database` 按 GORM 依赖视为可选模块，避免普通服务入口因为迁移 helper 引入 GORM。
5. 处理 `service/server` 的 Jaeger provider 类型边界：优先评估是否可以把 `otlptrace.Exporter` 类型从核心 Option 中移走，或确认它只属于轻量 OTel 依赖且不带 `data/jaeger`。若需要拆分，新增更小接口或由 `service/jaeger` 提供适配 Option。
6. 处理 `kratos` 模块的可选实现边界：
   - 将 Redis token manager 从 `kratos/auth` 核心接口包中拆出，或通过文件/包拆分让默认 `kratos` 模块不 require `go-redis`。
   - 将 Consul/Etcd registry 适配从默认 `kratos` 模块图中拆出，保留业务显式引入注册中心时的适配能力。
   - 如 tracing exporter 类型会污染默认依赖，按同样原则改为可选适配。
7. 拆开 `ping-service` 普通服务导出和数据库迁移导出：普通 `cmd/ping-service/export` 不再 import `service/database` 或 `cmd/database-migration/migrate`；迁移入口自己持有迁移相关依赖。
8. 补全 `testdata/ping-service/go.mod` 和 `go.work`，验证从 `testdata/ping-service` 模块视角构建默认服务入口。

## 任务列表

- [x] 梳理 `service/*` 当前 import 和模块依赖，确认核心包与可选包边界。
- [x] 为可选基础设施 service 子包创建独立 `go.mod`。
- [x] 清理 `service/go.mod`，移除已拆出可选模块的 require/replace。
- [x] 调整 `kratos` 模块边界，移除默认模块图中的 Redis、Consul、Etcd 等可选实现依赖。
- [x] 更新 `go.work`，加入 `testdata/ping-service` 和新增 service 子模块。
- [x] 调整 `service/server` / `service/tracer` 的 Jaeger exporter 类型边界，避免核心模块不必要依赖。
- [x] 调整 `ping-service` export 包，移除普通服务入口的迁移/GORM 依赖污染。
- [x] 补全 `testdata/ping-service/go.mod` 的 require/replace。
- [x] 更新 README / migration guide，说明独立服务模块和按需 service 子模块接入方式。
- [x] 运行 `gofmt`。
- [x] 运行 `wire ./testdata/ping-service/cmd/ping-service/export`。
- [x] 运行定向 `go test` / `go build` 验证核心 service 和 ping-service。
- [x] 使用 `go list -m all` 或 `go list -deps` 验证默认 ping-service 模块图不含未使用基础设施依赖。

## 验收标准

- `testdata/ping-service` 能作为独立 Go module 构建默认服务入口。
- 默认 `ping-service` 依赖图不出现 Redis、Mongo、MySQL、PostgreSQL、RabbitMQ、Consul、Etcd、GORM、数据库迁移相关依赖。
- `service/go.mod` 不再直接 require 已拆出的可选基础设施依赖。
- `kratos/go.mod` 不再直接 require 默认服务未使用的 Redis、Consul、Etcd 适配依赖。
- 需要 Redis/Mongo/Consul/数据库等能力时，可以通过对应 service 子模块 require + `WithSetup()` 显式接入。
- 现有示例入口、Wire 生成和定向测试通过。

## 风险与回滚

本任务属于中高风险重构，影响 Go module 边界、工作区、可选组件 import path 和示例服务依赖。

控制方式：

- 先按模块边界做最小拆分，保持现有 import path 和 `WithSetup()` 使用方式尽量不变。
- 不执行根模块 `go mod tidy`。
- 每次发现 README 假设与源码不一致，先更新规格记录再继续。
- 使用 `go list` 和定向测试验证模块图，而不是只看运行时日志。

回滚方式：

- 如拆分不可行，按文件级别回退本任务新增/修改的 go.mod、go.work、示例入口和文档。
- 不使用 `git reset --hard` 或丢弃用户未确认修改。

## 执行记录

- 2026-04-30：读取 `mytest/temp/service.md`，确认用户目标是独立服务模块和按需编译依赖，不是拆分业务 handler。
- 2026-04-30：读取 `go.work`、根 `go.mod`、`service/go.mod`、`testdata/ping-service/go.mod` 和 `docs/service-module-split/README.md`。
- 2026-04-30：确认 `testdata/ping-service/go.mod` 已存在但未补全，`go.work` 尚未纳入 `testdata/ping-service`。
- 2026-04-30：确认 `service/go.mod` 仍直接声明多种可选基础设施依赖，单纯运行时懒加载不足以解决模块图膨胀。
- 2026-04-30：确认 README 中 `service/auth`、`service/database` 的轻量性判断与源码不一致，需要在执行中一并处理。
- 2026-04-30：创建本规格文档，等待用户确认后再执行代码和模块文件修改。
- 2026-04-30：用户确认后继续读取 `kratos/middleware`、`kratos/auth`、`kratos/go.mod`。发现 `kratos` 核心模块也聚合了 Redis、Consul、Etcd 等可选依赖；仅拆 `service/*` 无法满足默认 `ping-service` 模块图最小化。已将方案范围扩大到 `kratos` 模块边界，需用户再次确认后再修改代码。
- 2026-04-30：用户确认扩大范围后，已将 `kratos/auth` 的 Redis token manager 拆到 `kratos/auth/redis` 可选子模块，核心 `kratos/auth` 仅保留 `TokenManager` 接口。
- 2026-04-30：已将 `kratos/registry`、`kratos/registry_etcd` 作为可选子模块纳入 `go.work`，核心 `kratos/go.mod` 不再直接 require Redis、Consul、Etcd、`otlptrace`。
- 2026-04-30：已为 `service/auth`、`service/consul`、`service/database`、`service/jaeger`、`service/mongo`、`service/mysql`、`service/postgres`、`service/rabbitmq`、`service/redis`、`service/store` 创建独立模块，并清理核心 `service/go.mod`。
- 2026-04-30：已将 Jaeger tracing 注入从 `serverutil.WithJaegerExporterProvider` 改为通用 `serverutil.WithTracerOptionProvider`，由 `service/jaeger` 提供 `WithTracerOptions()` 适配。
- 2026-04-30：已移除默认 `cmd/ping-service/export` 中的 `ExportDatabaseMigration`，避免普通服务入口引入 `service/database` 和迁移依赖。
- 2026-04-30：已将 `cmd/database-migration`、`cmd/store-configuration` 拆成 `ping-service` 下的独立命令模块，避免默认服务模块扫描到数据库和 Consul 配置存储依赖。
- 2026-04-30：已运行 `wire ./testdata/ping-service/cmd/ping-service/export`，成功生成 `wire_gen.go`。
- 2026-04-30：已运行核心验证：`go test ./kratos/auth ./kratos/middleware ./service/setup ./service/server ./service/cluster_service_api`，结果通过。
- 2026-04-30：已运行可选模块验证：`service/redis`、`service/auth`、`service/consul`、`service/postgres`、`service/mysql`、`service/mongo`、`service/rabbitmq`、`service/jaeger`、`service/database`、`service/store`、`kratos/auth/redis`、`kratos/registry`、`kratos/registry_etcd` 的 `go test ./...` 均通过。
- 2026-04-30：已运行命令模块验证：`testdata/ping-service/cmd/database-migration` 和 `cmd/store-configuration` 的 `go test ./...` 均通过。
- 2026-04-30：已关闭 workspace 验证独立模块：`GOWORK=off go test ./cmd/ping-service/...`、`GOWORK=off go test ./setup ./server ./cluster_service_api`、`GOWORK=off go test ./auth ./middleware` 均通过。
- 2026-04-30：已关闭 workspace 验证默认 `ping-service` 模块图和包依赖图，未匹配到 Redis、DB、Consul、Etcd、Mongo、RabbitMQ、Watermill、GORM 或拆出的可选 service/kratos 子模块。
