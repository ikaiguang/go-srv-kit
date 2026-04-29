# AGENTS.md

本文件为进入 `go-srv-kit` 仓库工作的代理提供最小必要上下文。

## 项目事实

- 模块名：`github.com/ikaiguang/go-srv-kit`
- Go 版本：`1.25.9`
- 框架：`go-kratos v2.9.2`
- 依赖注入：`google/wire v0.7.0`
- 工作区：根模块配合 `go.work`，本地 `replace` 到 `kit/`、`kratos/`、`service/` 和多个 `data/*` 子模块

## 语言规则

- 默认使用中文进行交流、解释、说明和文档补充
- 代码中的变量名、函数名、类型名保持英文，遵循既有编程规范
- 技术术语可以保留英文，但解释优先使用中文
- 错误信息和日志说明优先使用中文

## 仓库定位

`go-srv-kit` 是一个微服务工具包，不只是单个业务服务。

- `service/` 提供服务启动、配置、日志、数据库和 `LauncherManager`
- `kratos/` 提供认证、中间件、错误处理、日志等框架扩展
- `data/` 提供 MySQL、PostgreSQL、MongoDB、Redis、RabbitMQ、Consul、Etcd、Jaeger 等组件
- `kit/` 提供通用工具
- `testdata/ping-service/` 是示例服务，也是理解推荐接入方式的首选参考

## 架构约束

- 业务分层遵循 `Service -> Biz -> Data`
- `Service` 层负责 HTTP/gRPC handler 和 DTO 转换
- `Biz` 层负责业务逻辑，仓储接口放在 `biz/repo/`
- `Data` 层负责仓储实现，实现在 `data/repo/`
- 常见数据流：`Proto -> DTO -> BO -> PO`
- 不要让 `Service` 直接访问 `Data`

## 生成代码与禁止事项

- 不要手改 Proto 生成文件，如 `*.pb.go`、`*_grpc.pb.go`、`*_http.pb.go`、`*.validate.go`
- 修改 Wire 装配后，重新生成：
  - `wire ./testdata/ping-service/cmd/ping-service/export`
- Proto 相关生成命令以 `Makefile` 为准，例如：
  - `make protoc-api-protobuf`
  - `make protoc-config-protobuf`
  - `make protoc-ping-v1-protobuf`

## 重要约束

- 根目录 `go.mod` 明确说明：不要随手对根模块执行 `go mod tidy`
- 原因：`testdata/` 不会被 `./...` 包模式纳入，`tidy` 会移除示例服务依赖
- 如果确实需要整理依赖，先确认是否会影响 `testdata/ping-service`

## 常用命令

```bash
make init
make generate
go test ./...
go vet ./...
make run-service
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs
```

Windows 下部分命令更适合在 `cmd` 或 `git-bash` 运行；如果 `make` 目标在 PowerShell 中表现异常，优先退回到等价的 `go` / `wire` 命令。

## 修改代码时的优先参考

- 示例服务入口：`testdata/ping-service/cmd/ping-service/main.go`
- Wire 装配：`testdata/ping-service/cmd/ping-service/export/wire.go`
- 基础设施装配：`service/setup/setup.util.go`
- 服务器初始化：`service/server/server_all_in_one.util.go`

## 深层上下文

- Repo-local Codex skill 位于 `.agents/skills/go-srv-kit/`
- `AGENTS.md` 保留常驻规则；更深的架构、流程、生成命令和编码约定放在该 skill 的 `references/` 中按需读取

## 编码习惯

- 使用 `gofmt`，必要时使用 `goimports`
- 函数长度不超过 150 行；嵌套层级不超过 3 层，超过就拆分
- 参数数量不超过 4 个；超过时改为参数结构体或 `Options` 模式
- `context.Context` 作为第一个参数，命名 `ctx`；不要塞进结构体
- 参数顺序固定为：`ctx` -> 主参数/请求结构体 -> 其他参数 -> `...Option`
- `error` 作为最后一个返回值
- 返回值不超过 3 个（不含 `error`）；超过优先封装结构体
- 命名返回值只在 `defer` 需要修改返回值时使用；其他场景使用匿名返回值
- 新增业务逻辑时优先沿用现有命名：`NewXxxService`、`NewXxxBiz`、`NewXxxData`
- 接收器命名使用类型名首字母小写，禁止 `me`、`this`、`self`
- 业务逻辑中禁止使用 `panic`；异步 goroutine 使用 `threadpkg.GoSafe()` 包装
- 类型断言始终使用 `comma ok`
- 禁止硬编码配置值、密码、Token、连接串等敏感信息
- 重复出现的魔法数字必须提取为常量；多处复用的字符串标识符同样必须常量化，禁止直接硬编码
- Go 常量遵循驼峰命名；私有常量用小写开头；编辑 `.proto` 时枚举值保持全大写
- 导出常量和变量应补注释；成组常量可写总注释并给关键项加行尾说明
- goroutine 安全封装、错误处理风格、分层边界优先遵循现有实现，而不是另起一套

## 错误与日志

- 优先使用 `github.com/ikaiguang/go-srv-kit/kratos/error` 提供的错误构造与包装：
  - `errorpkg.ErrorBadRequest(...)`
  - `errorpkg.ErrorUnauthorized(...)`
  - `errorpkg.ErrorForbidden(...)`
  - `errorpkg.ErrorNotFound(...)`
  - `errorpkg.ErrorConflict(...)`
  - `errorpkg.ErrorInternal(...)`
  - `errorpkg.WrapWithMetadata(err, metadata)`
  - `errorpkg.FormatError(err)`
- 分层处理保持一致：
  - `Service` 层做参数校验、必要错误日志、返回业务错误
  - `Biz` 层做业务校验和错误包装
  - `Data` 层做底层错误转换，例如 `gorm.ErrRecordNotFound` -> `ErrorNotFound`
- 记录日志优先使用 `logpkg.WithContext(ctx)`，保留 TraceID
- 使用结构化日志 `Infow/Warnw/Errorw`，不要优先写 `Infof/Errorf` 这类格式化日志
- 错误日志至少带上 `"error", err`；需要排查堆栈时再补 `"stack", stringutil.GetStackTrace(2)`
- 日志里禁止输出未脱敏的密码、手机号、Token、密钥等敏感信息；按现有工具先脱敏再记录

## 工作方式

- 先文档后执行：除非用户明确要求跳过，收到任何需要修改文件、执行有副作用命令或推进实现的任务时，先在 `docs/<task-name>/*.md` 创建任务说明文档，写清楚“要做什么、为什么这么做、如何做（任务列表）”，并等待用户确认后再开始实际执行。
- 执行过程中如果发现原方案不合适、约束不成立或需要改变方向，先回到对应文档记录“发现的问题、原方案为什么不合适、准备怎么改、为什么这样改”，必要时再次等待用户确认，然后再继续执行。
- 用户确认任务文档后，执行阶段的低风险操作可以静默执行，不需要逐项询问确认；低风险操作包括读取和搜索文件、修改当前仓库内与任务相关的普通文件、运行本地格式化、测试、构建和约定内代码生成。
- 涉及危险或高影响操作时必须先确认，包括改写数据库、请求第三方或线上 API 修改数据、删除大量文件、改写 Git 历史、丢弃未确认修改、安装或修改系统级依赖、发布部署、发送外部消息、处理未脱敏敏感信息，以及任何风险不确定的操作。
- 当用户指出某个回答应沉淀为文档、规则或使用指导时，应主动更新对应文件；如果不能更新，先说明原因，不要只停留在聊天回答。
- 当用户纠正代理行为、指出应沉淀文档或要求以后不要反复提醒时，应主动判断并更新 `AGENTS.md`、skill 或 docs；适合常驻的规则应进入 `AGENTS.md`，细节放入对应 skill 或 docs。
- 已确认任务文档后，读文件、搜索、`gofmt`、`go test`、`wire`、`go list` 等低风险本地操作默认直接执行；只有越出当前 workspace、破坏性操作、线上或系统级影响、处理敏感信息、风险不确定时才询问确认。
- 如果规则边界、执行风险或文档分层不清楚，不要默不作声或只聊天回答；应主动提出 1-3 个具体问题让用户决策。
- 先读将要修改的模块和相邻实现，再决定改法
- 优先做与当前结构一致的最小改动
- 涉及新服务或新模块时，优先对照 `testdata/ping-service` 的目录组织和 Wire 接入方式
