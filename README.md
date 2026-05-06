# go-srv-kit

[![Go Report Card](https://goreportcard.com/badge/github.com/ikaiguang/go-srv-kit/kit)](https://goreportcard.com/report/github.com/ikaiguang/go-srv-kit/kit)
[![GoDoc](https://godoc.org/github.com/ikaiguang/go-srv-kit/kit?status.svg)](https://godoc.org/github.com/ikaiguang/go-srv-kit/kit)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

`go-srv-kit` 为微服务、业务系统开发提供开箱即用的工具。

- 按需配置启动基础组件，如：数据库、缓存、消息队列等。
- 提供基础的工具，如：日志、配置、HTTP、GRPC、JWT、SnowflakeId、...

## 目录

- [特性](#特性)
- [架构](#架构)
- [快速开始](#快速开始)
- [目录结构](#目录结构)
- [核心组件](#核心组件)
- [配置](#配置)
- [开发指南](#开发指南)
- [常见问题](#常见问题)
- [贡献](#贡献)
- [许可证](#许可证)

## 特性

- **🏗️ DDD 分层架构** - Service → Business → Data 清晰分层
- **🔧 Wire 依赖注入** - 编译期依赖注入，类型安全
- **🔄 双协议支持** - HTTP + gRPC 从同一 Proto 定义生成
- **🔐 JWT 认证** - 白名单模式，支持多 Token 类型
- **📊 可观测性** - 结构化日志、链路追踪、指标监控
- **🗄️ 多数据库** - MySQL、PostgreSQL、MongoDB
- **⚡ 缓存支持** - Redis 集群、本地缓存
- **📬 消息队列** - RabbitMQ、Watermill
- **🔍 服务发现** - Consul、Etcd
- **🛡️ 中间件** - CORS、Recovery、Tracer、Auth、限流

## 架构

```
┌─────────────────────────────────────────────────────────────────┐
│                    使用 go-srv-kit 的业务服务                      │
├─────────────────────────────────────────────────────────────────┤
│  Service Layer (internal/service/)  ← HTTP/gRPC handlers         │
├─────────────────────────────────────────────────────────────────┤
│  Business Layer (internal/biz/)     ← 业务逻辑                   │
├─────────────────────────────────────────────────────────────────┤
│  Data Layer (internal/data/)        ← 数据访问                   │
├─────────────────────────────────────────────────────────────────┤
│                   ↓ 调用 go-srv-kit 基础设施 ↓                    │
├─────────────────────────────────────────────────────────────────┤
│  service/        - 服务启动、配置管理、LauncherManager            │
│  kratos/         - 框架扩展 (auth、middleware、error、log)        │
│  data/           - 数据组件 (MySQL、Redis、RabbitMQ、MongoDB)     │
│  kit/            - 通用工具库 (加密、ID生成、文件操作)            │
└─────────────────────────────────────────────────────────────────┘
```

## 快速开始

### 前置要求

- Go 1.24+
- protoc + protoc-gen-*
- wire

### 安装

```bash
# 克隆仓库
git clone https://github.com/ikaiguang/go-srv-kit.git
cd go-srv-kit

# 初始化开发环境
make init
```

### 运行示例服务

**Windows** 系统，请使用 `cmd` 或 `git-bash` 运行。

```bash
# 启动项目
make run-service
# 或
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs

# 运行测试
make testing-service
# 或
curl http://127.0.0.1:10101/api/v1/ping/logger && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/error && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/panic && echo "\n"
curl http://127.0.0.1:10101/api/v1/ping/say_hello && echo "\n"
```

### 访问 API 文档

```bash
# Swagger UI (OpenAPI v2)
http://127.0.0.1:10101/q/

# OpenAPI v3 JSON
http://127.0.0.1:10101/api/swagger/
```

### 创建新服务

参考：[service-layout](https://github.com/ikaiguang/service-layout)

## 目录结构

```
go-srv-kit/
├── api/              # Proto 定义文件
├── cmd/              # 命令行工具
├── data/             # 数据层组件实现
├── kit/              # 通用工具库
├── kratos/           # Kratos 框架扩展
├── service/          # 服务层工具
├── testdata/         # 测试数据和示例服务
├── websocket/        # WebSocket 支持
├── wire/             # Wire 依赖注入工具
├── .claude/          # Claude Code 智能开发配置
├── CLAUDE.md         # 项目架构和开发指南
└── README.md         # 本文件
```

## 核心组件

### 服务层 (service/)

| 组件 | 说明 |
|------|------|
| `setup/` | LauncherManager - 基础设施统一初始化入口 |
| `server/` | HTTP/gRPC 服务器创建和管理 |
| `config/` | 配置加载（文件/Consul） |
| `database/` | 数据库连接管理 |
| `logger/` | 日志初始化（Zap） |
| `middleware/` | 中间件设置 |
| `auth/` | 认证提供者 |

### Kratos 扩展 (kratos/)

| 组件 | 说明 |
|------|------|
| `auth/` | JWT Token 管理和验证 |
| `middleware/` | CORS、Recovery、Tracer、Auth、限流等 |
| `error/` | 统一错误处理 |
| `log/` | 结构化日志（Zap） |
| `client/` | gRPC/HTTP 客户端封装 |

### 数据组件 (data/)

| 组件 | 说明 |
|------|------|
| `gorm/` | GORM ORM 工具 |
| `mysql/` | MySQL 专用工具 |
| `postgres/` | PostgreSQL 专用工具 |
| `mongo/` | MongoDB 客户端封装 |
| `redis/` | Redis 客户端封装 |
| `rabbitmq/` | RabbitMQ 客户端封装 |
| `consul/` | Consul 客户端封装 |
| `etcd/` | Etcd 客户端封装 |
| `jaeger/` | 分布式追踪 |

### 工具库 (kit/)

加密、ID 生成、文件操作、网络请求等 70+ 工具函数。

## 配置

### 配置文件

配置采用 Proto 定义，支持 YAML 格式：

```yaml
# configs/config.yaml
server:
  http:
    addr: 0.0.0.0:10101
  grpc:
    addr: 0.0.0.0:10102

data:
  mysql:
    host: localhost
    port: 3306
    database: mydb

data:
  redis:
    addr: localhost:6379
```

### 环境变量

支持通过环境变量覆盖配置：

```bash
export SERVER_HTTP_ADDR=0.0.0.0:8080
```

## 开发指南

### 文档

- [CLAUDE.md](CLAUDE.md) - 项目架构和开发指南
- [.claude/rules/](.claude/rules/) - 编码规范和开发流程

### API 开发流程

1. 定义 Proto (`api/{service}/v1/`)
2. 生成代码 (`make api-{service}`)
3. 实现 Service 层
4. 实现 Business 层
5. 实现 Data 层
6. 配置 Wire 依赖注入
7. 运行 `wire` 生成代码

### Proto 代码生成

```bash
# 生成所有 API
make protoc-api-protobuf

# 生成配置
make protoc-config-protobuf

# 生成指定服务
make protoc-specified-api service=ping-service

# 生成 ping-service v1
make protoc-ping-v1-protobuf
```

### 生成的代码

每个 `.proto` 文件会生成：
- `{file}.pb.go` - Proto 消息定义
- `{file}_grpc.pb.go` - gRPC 服务接口
- `{file}_http.pb.go` - HTTP 服务接口
- `{file}_errors.pb.go` - 错误定义
- `{file}.validate.go` - 验证代码
- `{file}.swagger.json` - OpenAPI v2 文档
- `{file}.openapi.yaml` - OpenAPI v3 文档

### 常用命令

```bash
# 初始化
make init

# 生成 Wire 代码
make generate

# 生成 Proto 代码
make api-{service}

# 运行测试
cd kit && go test ./...

# 代码检查
go vet ./...
golangci-lint run

# 构建
make build
```

## 概述

- 使用服务框架：[go-kratos](https://github.com/go-kratos/kratos)

**参考链接**

- [github.com/go-kratos/kratos](https://github.com/go-kratos/kratos)
- [github.com/uber-go/guide](https://github.com/uber-go/guide)
- [Go Package names](https://blog.golang.org/package-names)

## 常见问题

### Q: 如何添加新的中间件？

在 `service/middleware/` 或 `kratos/middleware/` 中添加，然后在 `ExportAuthWhitelist()` 配置白名单。

### Q: 如何切换数据库？

修改配置文件中的数据库连接信息，Go-srv-kit 支持 MySQL、PostgreSQL、MongoDB。

### Q: 如何禁用认证？

在配置中设置 `setting.enable_auth_middleware: false`。

### Q: Wire 生成失败怎么办？

```bash
# 清理后重新生成
rm ./cmd/*/export/wire_gen.go
wire ./cmd/*/export
```

## 贡献

欢迎贡献！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'feat: add some amazing feature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## Give a star! ⭐

如果您觉得这个项目有趣，或者对您有帮助，请给个 star 吧！

If you think this project is interesting, or helpful to you, please give a star!

## 许可证

本项目采用 Apache License 2.0 许可证 - 详见 [LICENSE](LICENSE) 文件

Copyright [2020] [ckaiguang@outlook.com]
