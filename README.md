# go-srv-kit

[![Go Report Card](https://goreportcard.com/badge/github.com/ikaiguang/go-srv-kit)](https://goreportcard.com/report/github.com/ikaiguang/go-srv-kit)
[![GoDoc](https://godoc.org/github.com/ikaiguang/go-srv-kit?status.svg)](https://godoc.org/github.com/ikaiguang/go-srv-kit)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

`go-srv-kit` 是一个基于 `go-kratos v2.9.2` 的微服务工具包，用于快速搭建带有 DDD 分层、Wire 依赖注入、HTTP/gRPC 双协议和常见基础设施能力的业务服务。

## 适合做什么

- 快速搭建 `Service -> Biz -> Data` 分层的业务服务
- 统一接入配置、日志、数据库、缓存、消息队列、认证和链路追踪
- 通过 Proto 同时生成 gRPC 与 HTTP 接口代码
- 基于 `testdata/ping-service` 作为完整示例服务进行二次开发

## 当前技术事实

- Go 版本：`1.25.9`
- go-kratos：`v2.9.2`
- Wire：`v0.7.0`
- 工作区：根模块配合 `go.work`，联动多个 `data/*`、`kit/`、`kratos/`、`service/` 子模块

## 仓库结构

| 路径 | 说明 |
|------|------|
| `api/` | 项目级 Proto 定义，例如全局配置 |
| `cmd/` | 项目级命令行工具 |
| `data/` | MySQL、Redis、RabbitMQ、MongoDB、Consul、Etcd、Jaeger 等数据组件 |
| `kit/` | 与业务无关的通用工具库 |
| `kratos/` | 认证、中间件、错误处理、日志、客户端等 Kratos 扩展 |
| `service/` | `LauncherManager`、配置、服务端、中间件、基础设施管理 |
| `testdata/` | 示例服务、测试配置与辅助工具 |
| `wire/` | Wire 依赖注入辅助 |
| `docs/` | 补充文档与迁移说明 |

## 快速开始

### 前置要求

- Go `1.25.9` 或更高兼容版本
- `protoc`
- `wire`

可通过以下命令安装本仓库常用开发工具：

```bash
make init
```

### 重要注意事项

- 不要在仓库根目录随手执行 `go mod tidy`
  - 根模块的 `go.mod` 已明确说明：`testdata/` 不会被 Go 的 `./...` 包模式纳入，`tidy` 可能移除示例服务依赖
- Windows 下部分 `make` 目标更适合在 `cmd` 或 `git-bash` 执行
- 不要手工修改 `*.pb.go`、`*_grpc.pb.go`、`*_http.pb.go`、`*.validate.go`、`wire_gen.go` 等生成文件

### 运行示例服务

```bash
# 推荐
make run-service

# 或直接运行
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs
```

接口测试：

```bash
make testing-service
```

常用访问地址：

- Swagger UI（OpenAPI v2）：`http://127.0.0.1:10101/q/`
- OpenAPI v3：`http://127.0.0.1:10101/api/swagger/`

## 常用开发命令

```bash
# 安装工具
make init

# 生成示例服务的 Wire 代码
make generate

# 生成所有 API Proto
make protoc-api-protobuf

# 生成配置 Proto
make protoc-config-protobuf

# 运行测试
go test ./...

# 静态检查
go vet ./...
```

如果只需要重新生成示例服务的 Wire 代码：

```bash
wire ./testdata/ping-service/cmd/ping-service/export
```

## 推荐开发路径

如果你要基于本仓库开发业务服务，优先参考：

1. `testdata/ping-service/`
2. `service/setup/`
3. `service/server/`
4. `kratos/error/`、`kratos/log/`、`kratos/middleware/`

典型业务服务分层：

```text
Proto -> DTO -> BO -> PO
Service -> Biz -> Data
```

## 文档索引

### 面向开发者

- [docs/README.md](docs/README.md) - 文档目录
- [docs/migration-guide.md](docs/migration-guide.md) - 模块化迁移指南
- [README_STYLE.md](README_STYLE.md) - 代码风格规范
- [README_GO_PRIVATE.md](README_GO_PRIVATE.md) - Go 私有仓库配置
- [README_GIT_COLLABORATOR.md](README_GIT_COLLABORATOR.md) - Git 协作流程
- [README_SUBMODULE.md](README_SUBMODULE.md) - Git 子模块使用指南

### 面向 AI 辅助开发

- [CLAUDE.md](CLAUDE.md) - Claude Code 项目指南
- [AGENTS.md](AGENTS.md) - Agent 常驻规则
- [docs/codex-resume-guide.md](docs/codex-resume-guide.md) - Codex 会话续接说明

## 参考项目

- 业务服务模板：[service-layout](https://github.com/ikaiguang/service-layout)
- go-kratos：https://github.com/go-kratos/kratos
- Uber Go Style Guide：https://github.com/uber-go/guide

## 许可证

本项目采用 Apache License 2.0，详见 [LICENSE](LICENSE)。
