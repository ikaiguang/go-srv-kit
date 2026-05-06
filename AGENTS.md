# AGENTS.md

本文件只保留本项目的常驻规则。更细的架构、编码、命令和流程说明放在 `.agents/skills/` 中按需读取。

## 项目事实

- Go 版本：`1.25.9`
- 默认使用中文交流、解释和补充文档；代码标识符保持英文。

## 架构硬约束

- 业务分层遵循 `Service -> Biz -> Data`。
- `Service` 只负责 HTTP/gRPC handler、入参校验和 DTO 转换，不直接访问 `Data`。
- `Biz` 负责业务逻辑和仓储接口；`Data` 负责仓储实现和外部访问。
- 常见数据流：`Proto -> DTO -> BO -> PO`。

## 生成代码

- 不要手改生成文件，如 `*.pb.go`、`*_grpc.pb.go`、`*_http.pb.go`、`*.validate.go`、`wire_gen.go`。
- Proto 生成命令以 `Makefile` 为准。
- 修改 Wire 装配后，运行：

```bash
make generate
wire ./testdata/ping-service/cmd/ping-service/export
```

## 常用命令

```bash
make init
make generate
go test ./...
go vet ./...
make run-service
```

Windows 下如果 `make` 在 PowerShell 中表现异常，优先使用等价的 `go`、`wire` 或在 `git-bash` 中运行。

## Skill 入口

- 本仓库 Go/Proto/Wire/配置/测试修改：先用 `.agents/skills/my-project/`。
- 需要先文档后执行、规格确认或风险分级：用 `.agents/skills/spec-workflow/`。
- 全库审计、补测试、补文档或综合风险扫描：如可用，使用 `.agents/skills/code-audit-repair/`。

常用参考文件：

- `.agents/skills/my-project/references/project-context.md`
- `.agents/skills/my-project/references/service-workflow.md`
- `.agents/skills/my-project/references/coding-rules.md`
- `.agents/skills/my-project/references/commands-and-generation.md`

## 工作方式

- 修改文件、执行有副作用命令或推进实现前，除非用户明确跳过，先在 `.agents/specs/<task-name>/spec.md` 写清目标、方案、任务、验收和风险，等待确认。
- 任务文档只作临时过程记录；长期规则沉淀到 `AGENTS.md`、repo-local skill、模块 README 或稳定 docs。
- 用户确认规格后，读文件、搜索、修改相关普通文件、`gofmt`、`go test`、`wire`、`go list` 等低风险本地操作默认直接执行。
- 涉及删除大量文件、改写 Git 历史、丢弃未确认修改、系统级依赖、发布部署、线上数据、敏感信息或风险不确定的操作，必须先确认。
- 先读目标模块和相邻实现，再做与现有结构一致的最小改动；新增或修改行为时同步考虑测试、文档、安全、稳定性和性能风险。
- 当用户指出规则应沉淀、回答应复用或行为需纠正时，主动更新合适的规则或文档；边界不清时提出 1-3 个具体问题。
