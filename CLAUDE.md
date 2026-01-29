# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

`go-srv-kit` 是基于 **go-kratos v2** 框架的微服务开发工具包，提供开箱即用的工具和组件，用于快速构建微服务和业务系统。

- **框架**: go-kratos v2.9.1
- **语言**: Go 1.24.10
- **架构**: DDD 分层架构 + Wire 依赖注入
- **通信协议**: HTTP + gRPC 双协议支持

## 架构关系

```
┌─────────────────────────────────────────────────────────────────┐
│                    使用 go-srv-kit 的业务服务                      │
│                  (如 testdata/ping-service)                      │
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

**关键理解**: `testdata/ping-service` 是使用 `go-srv-kit` 的示例服务，展示如何基于这个工具包构建业务微服务。

## 常用命令

### 初始化开发环境
```bash
make init
```
安装必要的工具：protoc 相关插件、wire、kratos CLI 等。

### 生成代码
```bash
# 生成 Wire 依赖注入代码
make generate
# 或直接运行
wire ./testdata/ping-service/cmd/ping-service/export
```

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

### 运行测试服务
```bash
# 启动 ping-service（使用 cmd 或 git-bash）
make run-service
# 或
go run ./testdata/ping-service/cmd/ping-service/... -conf=./testdata/ping-service/configs

# 运行 HTTP 测试
make testing-service
```

### 访问 API 文档
```bash
# Swagger UI (OpenAPI v2)
http://127.0.0.1:10101/q/

# OpenAPI v3 JSON
http://127.0.0.1:10101/api/swagger/
```

### Protocol Buffers 生成
```bash
# 使用 Makefile 中定义的命令（具体见各子模块的 makefile）
make api-config
make api-ping-service
```

### 构建
```bash
make build
make deploy-on-docker
```

## go-srv-kit 目录结构

| 目录 | 用途 |
|------|------|
| `api/` | Proto 定义文件（配置、通用 API） |
| `cmd/` | 命令行工具 |
| `data/` | 数据层组件实现（MySQL、Redis、RabbitMQ、MongoDB 等） |
| `kit/` | 通用工具库（加密、ID 生成、文件操作等） |
| `kratos/` | Kratos 框架扩展（auth、middleware、error、log 等） |
| `service/` | 服务层工具（server 初始化、config、database、logger 等） |
| `testdata/ping-service/` | 示例服务，展示如何使用 go-srv-kit |
| `wire/` | Wire 依赖注入工具 |

## 业务服务开发模式

使用 go-srv-kit 开发业务服务时，遵循以下模式：

### 1. 定义 API（Proto）

在 `api/{service-name}/v1/` 下定义：
- `services/` - 服务定义
- `resources/` - 请求/响应消息
- `errors/` - 错误定义
- `enums/` - 枚举类型

### 2. 实现各层

**Service Layer** (`internal/service/service/`):
- 实现 gRPC/HTTP handler
- 处理 DTO ↔ BO 转换
- 调用业务逻辑

**Business Layer** (`internal/biz/biz/`):
- 实现业务逻辑
- 定义 Repository 接口（在 `biz/repo/`）
- 处理领域事件

**Data Layer** (`internal/data/data/`):
- 实现 Repository 接口（在 `data/repo/`）
- 使用 go-srv-kit 提供的数据组件（GORM、Redis 等）
- 返回 PO 模型

### 3. Wire 依赖注入

在 `cmd/{service}/export/wire.go` 中定义：

```go
//go:build wireinject
package exporter

import "github.com/google/wire"

func exportServices(launcher, hs, gs) (Cleanup, error) {
    panic(wire.Build(
        setuputil.GetLogger,
        data.NewPingData,      // Data 层
        biz.NewPingBiz,        // Business 层
        service.NewPingService, // Service 层
        service.RegisterServices,
    ))
}
```

运行 `wire` 生成 `wire_gen.go`。

### 4. 配置管理

**配置 Proto** (`internal/conf/config.conf.proto`):
```protobuf
message ServiceConfig {
  message PingService { string key = 1; }
  PingService ping_service = 1;
}
```

**导出配置** (`cmd/*/export/main.export.go`):
```go
func ExportServiceConfig() []configutil.Option {
    return []configutil.Option{
        configutil.WithOtherConfig(serviceConfig),
    }
}
```

**访问配置**:
```go
config := launcherManager.GetConfig()
```

## 服务启动流程

```
main.go (业务服务入口)
  ↓
ExportServiceConfig()    # 导出服务特定配置
ExportAuthWhitelist()    # 导出认证白名单
ExportServices()         # 导出服务注册函数
  ↓
AllInOneServer()         # [go-srv-kit] 创建服务器
  ↓
NewLauncherManager()     # [go-srv-kit] 初始化基础设施
  ├── 日志初始化
  ├── 数据库 (MySQL/PostgreSQL)
  ├── 缓存 (Redis)
  ├── 消息队列 (RabbitMQ)
  ├── 服务发现 (Consul/Etcd)
  ├── 链路追踪 (Jaeger)
  └── JWT 认证
  ↓
NewServerManager()       # [go-srv-kit] 创建 HTTP/gRPC 服务器
  ↓
Wire 依赖注入构建         # 业务服务的依赖关系
  ↓
kratos.App()             # 启动应用
```

## 数据转换流

```
Proto (API 定义) → DTO (service/dto/) → BO (biz/bo/) → PO (data/po/) → Database
```

## 重要参考文件

### go-srv-kit 核心文件

| 文件 | 说明 |
|------|------|
| `service/setup/setup.util.go` | LauncherManager - 基础设施统一入口 |
| `service/server/server_all_in_one.util.go` | AllInOneServer - 服务器初始化 |
| `api/config/config.proto` | 全局配置结构定义 |
| `kratos/auth/` | JWT 认证实现 |
| `kratos/middleware/` | 中间件集合 |
| `kratos/error/` | 错误处理工具 |
| `kratos/log/` | 日志工具 |

### 示例服务参考文件

| 文件 | 说明 |
|------|------|
| `testdata/ping-service/cmd/ping-service/main.go` | 服务入口模板 |
| `testdata/ping-service/cmd/ping-service/export/wire.go` | Wire 依赖注入模板 |
| `testdata/ping-service/internal/service/service/ping.service.go` | Service 层示例 |
| `testdata/ping-service/internal/biz/biz/ping.biz.go` | Business 层示例 |
| `testdata/ping-service/internal/data/data/ping.data.go` | Data 层示例 |

## 中间件链

默认中间件（按顺序）：
1. CORS - 跨域资源共享
2. Recovery - Panic 恢复
3. Tracer - OpenTelemetry 链路追踪
4. Validator - 请求验证
5. Header - HTTP 头管理
6. Auth - JWT 认证（基于白名单）
7. Rate Limiting - 请求限流

## 认证机制

- **JWT Token**: 支持 Access Token / Refresh Token
- **白名单模式**: 在 `ExportAuthWhitelist()` 中定义无需认证的路径
- **Token 类型**: USER, ADMIN, EMPLOYEE, DEFAULT
- **Claims**: UserID, UserUuid, LoginPlatform, TokenType

## 错误处理

使用 `kratos/error/` 包：
```go
errorpkg.ErrorBadRequest("message")
errorpkg.WrapWithMetadata(err, metadata)
errorpkg.FormatError(err) // 带堆栈信息
```

## 日志系统

基于 Zap 的结构化日志：
- Console 输出（开发环境）
- 文件轮转（生产环境）
- 分类日志：app, gorm, rabbitmq

## 基础设施组件

### 数据库
- **MySQL/PostgreSQL**: GORM ORM，连接池
- **MongoDB**: 官方驱动封装

### 缓存
- **Redis**: 集群支持，连接池
- **本地缓存**: go-cache

### 消息队列
- **RabbitMQ**: AMQP 协议支持
- **Watermill**: 消息处理框架

### 服务发现
- **Consul**: 服务注册、健康检查
- **Etcd**: 替代方案

### 可观测性
- **日志**: Zap 结构化日志
- **追踪**: OpenTelemetry + Jaeger
- **指标**: Prometheus 支持
