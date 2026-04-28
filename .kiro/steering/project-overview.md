---
inclusion: always
---

# go-srv-kit 项目概述

## 核心信息

- 项目名称: go-srv-kit
- 框架: go-kratos v2.9.2
- 语言: Go 1.25+
- 架构: DDD 分层架构 + Wire 编译期依赖注入
- 通信协议: HTTP + gRPC 双协议支持
- 私有仓库: gitlab.uufff.com

## 架构关系

```
业务服务 (如 testdata/ping-service)
  ↓ Service Layer (internal/service/)  ← HTTP/gRPC handlers
  ↓ Business Layer (internal/biz/)     ← 业务逻辑
  ↓ Data Layer (internal/data/)        ← 数据访问
  ↓ go-srv-kit 基础设施
    ├── service/   - 服务启动、配置管理、LauncherManager
    ├── kratos/    - 框架扩展 (auth、middleware、error、log)
    ├── data/      - 数据组件 (MySQL、Redis、RabbitMQ、MongoDB)
    └── kit/       - 通用工具库 (加密、ID生成、文件操作)
```

## 数据转换流

```
Proto (API 定义) → DTO (service/dto/) → BO (biz/bo/) → PO (data/po/) → Database
```

## 中间件链（按顺序）

1. CORS → 2. Recovery → 3. Tracer (OpenTelemetry) → 4. Validator → 5. Header → 6. Auth (JWT 白名单) → 7. Rate Limiting

## 常用命令

```bash
make init                    # 初始化开发环境
make generate               # 生成 Wire 代码
make protoc-api-protobuf    # 生成所有 API
make run-service            # 运行示例服务
go test ./...               # 运行测试
gofmt -w .                  # 格式化代码
golangci-lint run           # 静态检查
```

## 目录结构

| 目录 | 用途 |
|------|------|
| `api/` | Proto 定义文件 |
| `data/` | 数据层组件 (MySQL、Redis、RabbitMQ、MongoDB) |
| `kit/` | 通用工具库 |
| `kratos/` | Kratos 框架扩展 (auth、middleware、error、log) |
| `service/` | 服务层工具 (server 初始化、config、logger) |
| `testdata/ping-service/` | 示例服务 |

## 核心开发原则

1. 编写代码前先描述方案并等待批准
2. 修改超过 3 个文件时，先分解为更小的任务
3. 编写代码后列出可能问题并建议测试用例
4. 发现 bug 时先编写重现测试再修复
5. Service 层只能调用 Business 层（禁止直接调用 Data 层）
6. Repository 接口定义在 biz/repo/，实现在 data/repo/
7. 使用 Wire 的 wire.Bind 进行接口绑定
