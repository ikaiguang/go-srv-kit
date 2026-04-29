# 项目上下文

## 仓库标识

- 模块：`github.com/ikaiguang/go-srv-kit`
- Go：`1.25.9`
- 框架：`github.com/go-kratos/kratos/v2 v2.9.2`
- 依赖注入：`github.com/google/wire v0.7.0`
- 工作区：根目录 `go.work` 配合本地 `replace`，覆盖 `kit/`、`kratos/`、`service/` 和多个 `data/*` 子模块

## 仓库定位

这个仓库是一个微服务工具包，不是单个业务服务。

- `service/` 提供启动、配置、日志、数据库、中间件、认证和服务装配
- `kratos/` 提供认证、中间件、错误处理、日志、客户端等框架扩展
- `data/` 提供 MySQL、PostgreSQL、MongoDB、Redis、RabbitMQ、Consul、Etcd、Jaeger 等基础设施适配
- `kit/` 提供通用工具
- `testdata/ping-service/` 是业务服务如何接入该工具包的首选示例

## 分层模型

基于该工具包实现业务服务时，遵循：

1. `Service` 层负责 HTTP/gRPC handler 和 DTO 转换
2. `Biz` 层负责业务逻辑和仓储接口
3. `Data` 层负责仓储实现和外部访问

常见数据流：

`Proto -> DTO -> BO -> PO`

不要让 `Service` 层直接调用 `Data` 层。

## 示例入口

需要具体例子时，优先读这些文件：

- `testdata/ping-service/cmd/ping-service/main.go`
- `testdata/ping-service/cmd/ping-service/export/wire.go`
- `testdata/ping-service/internal/service/service/ping.service.go`
- `testdata/ping-service/internal/biz/biz/ping.biz.go`
- `testdata/ping-service/internal/data/data/ping.data.go`

## 工具包关键文件

- `service/setup/setup.util.go`：`LauncherManager` 和初始化流程
- `service/server/server_all_in_one.util.go`：服务端构建入口
- `api/config/config.proto`：全局配置结构
- `kratos/auth/`、`kratos/middleware/`、`kratos/error/`、`kratos/log/`：共享运行时能力

## 启动形态

典型启动流程：

1. 导出服务专属配置和认证白名单
2. 通过 Wire 装配业务服务
3. 使用工具包基础设施创建 all-in-one server
4. 启动 Kratos app
