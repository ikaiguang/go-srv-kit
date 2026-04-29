# Refactor Service On Demand Modules

## 背景

当前 `service` 启动链路存在包级硬依赖问题。只要业务服务 import 核心启动包，就可能在编译依赖图中拉入 Consul、Etcd、Redis、MongoDB、RabbitMQ、Jaeger 等基础设施组件，即使业务运行时并不使用这些组件。

已有 `docs/refactor-service-on-demand-modules/plan.md` 记录了本轮重构的背景、已执行内容和中断点。该计划已确认，但执行过程中低风险操作被反复询问确认，没有完全遵循“用户确认规格后低风险操作可直接执行”的约定。

本文件作为后续执行的 `spec-workflow` 入口，后续继续执行时以本文件为准，并保留 `plan.md` 作为历史计划和上下文记录。

## 目标

- 将 `service/setup` 和 `service/server` 核心启动链路收窄到配置、日志、生命周期和组件注册表。
- 让 Redis、PostgreSQL、MySQL、MongoDB、RabbitMQ、Consul、Etcd、Jaeger、Auth、cluster service API 等组件由业务入口显式引入和注册。
- 修复上一轮中断后遗留的编译问题、测试问题和旧 provider 调用点。
- 更新示例服务和文档，说明新的按需接入方式。
- 验证默认 `ping-service` 入口不再隐式拉入 Consul/Etcd 相关依赖。

## 非目标

- 不恢复 `service/setup` 中集中注册所有组件的 `WithAllComponents()` 式设计。
- 不为了兼容旧行为而重新让核心包 import 所有基础设施组件。
- 不执行 `go mod tidy`。
- 不修改生成文件以外的 Proto 生成产物，除非通过约定生成命令产生。
- 不处理真实环境发布、部署、数据库迁移或线上 API 调用。

## 影响范围

重点检查和修改范围：

- `service/setup`
- `service/server`
- `service/config`
- `service/cluster_service_api`
- `service/redis`
- `service/postgres`
- `service/mysql`
- `service/mongo`
- `service/consul`
- `service/jaeger`
- `service/rabbitmq`
- `service/auth`
- `testdata/ping-service/cmd/ping-service`
- `testdata/ping-service/cmd/all-in-one`
- `testdata/ping-service/cmd/database-migration`
- `testdata/ping-service/configs`
- `docs/migration-guide.md`
- `service/setup/README.md`
- `service/server/README.md`

## 方案

继续沿用已确认的重构方向：

- 核心启动包只保留配置、日志、生命周期、组件注册表和 option 聚合能力。
- 各基础设施组件在自己的 `service/<component>` 包中提供 `WithSetup()` 或等价 option。
- Consul 配置加载能力归属 `service/consul`，避免 `service/config` 直接硬依赖 Consul。
- cluster service API 通过 discovery factory 或显式 option 接入 Consul/Etcd，避免核心 client 包硬 import registry 实现。
- `AllInOneServer` 通过 `serverutil.Option` 接收 setup options、auth provider、jaeger provider、app option provider 等扩展点。
- 示例服务入口显式声明实际需要的组件。

执行阶段遇到编译错误时，优先修复 import cycle、旧 provider 引用、测试中已迁移函数的调用点和缺失的核心方法实现。

## 任务列表

- [x] 将既有 `plan.md` 转换为本 `spec.md` 执行入口
- [x] 读取 go-srv-kit 相关技术参考和当前工作区状态
- [x] 定位当前编译错误、import cycle 和旧 provider 调用点
- [x] 修复 `service/setup`、`service/config`、`service/server`、`service/cluster_service_api` 等核心包问题
- [x] 修复示例服务入口、Wire 装配和生成结果
- [x] 更新 README 和 migration guide
- [x] 运行 `gofmt`
- [x] 运行 `wire ./testdata/ping-service/cmd/ping-service/export`
- [x] 运行核心包测试
- [x] 运行 `go list -deps ./testdata/ping-service/cmd/ping-service` 检查默认入口依赖
- [x] 更新执行记录和最终验证结果

## 验收标准

- 核心测试包能够通过必要的 `go test` 验证。
- `wire ./testdata/ping-service/cmd/ping-service/export` 能够成功执行。
- 默认 `ping-service` 入口的依赖检查不再隐式拉入 Consul/Etcd 相关 registry 实现。
- 文档说明新的按需组件注册方式。
- `service/setup` 或 `service/server` 不再因为核心 import 自动拉入所有基础设施组件。

## 风险与回滚

本任务属于中风险重构，影响核心启动路径、Wire 装配、示例服务和文档。

控制方式：

- 优先做最小修复，不引入新的全量注册模式。
- 每次发现方案偏离时先更新本文档。
- 通过 `go test`、`wire` 和 `go list -deps` 验证结果。

回滚方式：

- 如方案不可行，可回退本轮相关源码和文档改动。
- 不通过 `git reset --hard` 或丢弃用户未确认修改来回滚。

## 执行记录

- 已根据用户要求启用 `spec-workflow` 和 `go-srv-kit` 协同流程。
- 已确认现有 `plan.md` 标记为“已确认”，后续低风险操作不再逐项询问。
- 已创建本 `spec.md`，作为后续继续执行的规格入口。
- 继续执行时读取了 `service/setup`、`service/server`、`service` README、`docs/migration-guide.md`、当前 API 和示例入口。
- 已确认当前核心编译问题在定向测试中不再复现。
- 已重写 `service/setup/README.md`、`service/server/README.md`、`service/README.md` 和 `docs/migration-guide.md`，移除过期的 `WithAllComponents()` / 核心包 `WithXxx()` 推荐用法，改为各组件包 `WithSetup()` 和入口显式注册。
- 已修正 `service/setup` 中少量注释，将旧的 `WithXxx` 表述调整为组件 `WithSetup`。
- 已执行 `gofmt -w service\setup`。
- 已执行 `wire ./testdata/ping-service/cmd/ping-service/export`，成功生成 `wire_gen.go`。
- 已执行定向测试：`go test ./service/setup ./service/config ./service/cluster_service_api ./service/server ./service/redis ./service/postgres ./service/consul ./testdata/ping-service/cmd/ping-service ./testdata/ping-service/cmd/all-in-one ./testdata/ping-service/cmd/database-migration`，结果通过。
- 已执行依赖检查：`go list -deps ./testdata/ping-service/cmd/ping-service | Select-String -Pattern "go-srv-kit/(data|kratos)/(consul|etcd|registry_etcd)|hashicorp/consul|go.etcd.io/etcd"`，无输出，默认入口未匹配到 Consul/Etcd 相关依赖。
