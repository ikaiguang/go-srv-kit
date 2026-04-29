# CLAUDE.md

本文件为 Claude Code 在 `go-srv-kit` 仓库中工作时提供项目级说明。

## 核心规则

1. 默认使用中文回复、解释和说明。
2. 编码前先确认需求和改动范围；若需求不清楚，先澄清。
3. 发现 bug 时，优先补充可复现的测试或最小验证步骤，再修复。
4. 代码修改后，说明潜在风险和建议验证方式。
5. 超过少量文件的大改动，优先拆分成可验证的小步提交。

## 当前仓库事实

- 模块：`github.com/ikaiguang/go-srv-kit`
- Go：`1.25.9`
- 框架：`go-kratos v2.9.2`
- 依赖注入：`google/wire v0.7.0`
- 架构：DDD 分层 + Wire + HTTP/gRPC 双协议

## 仓库定位

`go-srv-kit` 是一个微服务工具包，不是单个业务服务仓库。

- `service/`：服务启动、配置、`LauncherManager`、服务器装配
- `kratos/`：认证、中间件、错误处理、日志等框架扩展
- `data/`：MySQL、Redis、RabbitMQ、MongoDB、Consul、Etcd、Jaeger 等组件
- `kit/`：通用工具
- `testdata/ping-service/`：示例服务，也是最重要的接入参考

## 开发约定

- 分层保持 `Service -> Biz -> Data`
- `Service` 层只负责 handler、参数校验和 DTO 转换，不直接访问 `Data`
- `Biz` 层定义仓储接口，`Data` 层提供实现
- 常见数据流：`Proto -> DTO -> BO -> PO`
- 不要手工修改生成文件：
  - `*.pb.go`
  - `*_grpc.pb.go`
  - `*_http.pb.go`
  - `*.validate.go`
  - `wire_gen.go`

## 常用命令

```bash
make init
make generate
make run-service
make testing-service
go test ./...
go vet ./...
wire ./testdata/ping-service/cmd/ping-service/export
```

## 重要注意事项

- 不要在根模块执行 `go mod tidy`
  - `testdata/` 不会被 Go 的 `./...` 包模式纳入，可能导致示例服务依赖被移除
- Windows 下部分 `make` 目标更适合在 `cmd` 或 `git-bash` 中运行
- 优先参考 `testdata/ping-service/` 的目录组织和 Wire 接入方式

## 优先阅读顺序

1. `README.md`
2. `AGENTS.md`
3. `testdata/ping-service/README.md`
4. `service/setup/README.md`
5. `.claude/rules/` 下与你当前任务最相关的规则文件

## 相关文档

- `README_STYLE.md`：编码风格规范
- `docs/migration-guide.md`：模块化迁移说明
- `docs/codex-resume-guide.md`：Codex 会话续接说明
- `.agents/skills/go-srv-kit/`：Repo-local skill，上下文拆分到 `references/`
