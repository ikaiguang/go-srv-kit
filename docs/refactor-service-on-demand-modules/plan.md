# service 按需模块引入重构计划

## 要做什么

将 `service` 启动链路从“import 一个核心包就编译进大量基础设施组件”调整为“业务服务在入口文件显式引入所需组件，并通过 all-in-one 启动选项注册”。

目标是让 `ping-service` 或真实业务服务可以在一个启动文件中声明实际需要的模块，例如：

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

基础启动包默认只负责配置、日志、生命周期和组件注册表，不再因为 import `service/setup` 或 `service/server` 自动拉入 Consul、Etcd、Redis、Mongo、RabbitMQ、Jaeger 等组件实现。

## 为什么这么做

- Go 的依赖是包级编译依赖。只要 `service/setup` 包内任意文件 import 了 Consul、Etcd、RabbitMQ 等依赖，业务服务即使运行时不用这些组件，也会在编译依赖图里拉入它们。
- 现有“懒加载”只能避免运行时初始化，不能避免编译期依赖进入模块。
- `AllInOneServer` 是业务服务最常用入口，组件选择应在业务入口显式表达，避免工具包默认替业务服务决定所有基础设施。
- Consul 配置加载、Consul 服务注册、cluster service API 的 Consul/Etcd discovery 都属于可选能力，应由对应模块提供选项，而不是由核心 server/setup 包硬编码。

## 如何做

- [x] 先阅读 `service/setup`、`service/server`、`service/config`、`service/cluster_service_api` 和 `testdata/ping-service` 的启动路径。
- [x] 识别当前硬依赖来源：`setup_launcher`、`interface`、`setup_provider`、`server_app`、`config_from_consul`、`cluster_service_api`。
- [x] 初步将 `setup` 核心收窄到配置、日志、生命周期、组件注册表。
- [x] 初步新增各组件子包的 `WithSetup()`，例如 `service/redis`、`service/postgres`、`service/consul` 等。
- [x] 初步将 Consul 配置加载器从 `service/config` 迁移到 `service/consul`。
- [x] 初步将 all-in-one 启动扩展为可接收 `serverutil.Option`，由入口文件传入组件注册和可选 provider。
- [x] 初步更新 `ping-service` 示例入口，让默认入口只注册 `cluster_service_api`。
- [ ] 继续修复上一轮中断后遗留的编译问题，重点是 import cycle、测试和旧 provider 调用点。
- [ ] 重新运行 `gofmt`、`wire ./testdata/ping-service/cmd/ping-service/export` 和核心包测试。
- [ ] 用 `go list -deps ./testdata/ping-service/cmd/ping-service` 检查默认入口是否不再隐式拉入 Consul/Etcd。
- [ ] 更新 `service/setup/README.md`、`service/server/README.md` 和 `docs/migration-guide.md`，说明新的推荐接入方式。

## 已执行的修改记录

- 2026-04-29：读取了 `go-srv-kit` skill、`service/setup`、`service/server`、`service/config`、`service/cluster_service_api` 和 `ping-service` 启动相关代码。
- 2026-04-29：将 `componentRegistry` 初步公开为 `ComponentRegistry`，并给 `Component[T]` 增加 `Init()`，用于按名称急切初始化。
- 2026-04-29：删除了 `service/setup/setup_launcher.util.go` 中集中 import 所有组件的实现，新增 `setup_component.util.go` 和 `setup_core.util.go` 承担核心能力。
- 2026-04-29：将 `LauncherManager` 初步收窄，不再直接组合数据库、Redis、Mongo、Consul、Jaeger、RabbitMQ、Auth、ServiceAPI provider。
- 2026-04-29：新增了多个组件子包的 `WithSetup()` 文件，让组件注册跟随对应模块 import。
- 2026-04-29：把 Consul 配置加载实现初步迁移到 `service/consul/consul_config_loader.util.go`。
- 2026-04-29：将 `cluster_service_api` 的 registry 依赖初步改为 discovery factory，避免核心 client 包硬 import Consul/Etcd registry。
- 2026-04-29：更新 `serverutil.AllInOneServer`，初步支持通过 `serverutil.Option` 注入 setup options、auth provider、jaeger provider 和 app option provider。
- 2026-04-29：更新 `testdata/ping-service/cmd/ping-service/main.go`、`cmd/all-in-one/main.go` 和 Wire provider，生成过一次 `wire_gen.go`。
- 2026-04-29：执行核心包测试时中断前已暴露问题：`service/setup` 测试 import cycle、`service/config` 中 Consul loader 测试引用已迁移函数、`launcherManager` 缺少核心方法实现。随后已开始补 `setup_core.util.go` 和调整部分测试，但尚未完成验证。

## 当前风险和注意事项

- 当前工作区存在未完成代码改动，不能视为可编译状态。
- `service/setup` 测试里仍可能保留旧组件 provider 思路，需要继续清理或迁移到各组件包测试。
- `service/config/config_from_consul.util.go` 已迁移后，相关测试应移动到 `service/consul` 或标记为 integration，并改用新入口。
- `NewLauncherManagerWithCleanup`、`NewWithCleanup` 的向后兼容策略需要明确：是保留极简核心行为，还是另建 `WithAllSetup()` 之类的兼容 helper。
- 不建议恢复 `WithAllComponents()` 到 `service/setup`，否则会重新把所有组件依赖拉回核心包。

## 确认状态

已确认。

## 继续执行记录

- 2026-04-29：用户确认本计划，继续执行代码修复、文档更新和验证命令。
