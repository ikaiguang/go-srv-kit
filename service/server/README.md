# server - 服务器管理

`server/` 负责 HTTP/gRPC 服务器创建、中间件配置和 Kratos 应用启动。

`server/` 不替业务服务隐式注册所有基础设施组件。All-In-One 启动时，业务入口应通过 `serverutil.WithSetupOptions(...)` 显式传入所需组件。

## 包名

```go
import serverutil "github.com/ikaiguang/go-srv-kit/service/server"
```

## ServerManager

`ServerManager` 管理 HTTP、gRPC 和 Kratos App 的单例创建：

```go
srvManager, err := serverutil.NewServerManager(launcherManager, authWhitelist)
httpServer, err := srvManager.GetHTTPServer()
grpcServer, err := srvManager.GetGRPCServer()
app, err := srvManager.GetApp()
```

## All-In-One 模式

多个服务合并到一个进程中启动时，入口文件显式声明组件：

```go
import (
    clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
    postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
    redisutil "github.com/ikaiguang/go-srv-kit/service/redis"
    serverutil "github.com/ikaiguang/go-srv-kit/service/server"
)

runOpts := []serverutil.Option{
    serverutil.WithSetupOptions(
        clientutil.WithSetup(),
        postgresutil.WithSetup(),
        redisutil.WithSetup(),
    ),
}

app, cleanup, err := serverutil.AllInOneServer(
    configFilePath,
    configOpts,
    []serverutil.ServiceExporter{service1, service2},
    authWhitelist,
    runOpts...,
)
defer cleanup()
serverutil.RunServer(app, cleanup)
```

## 可选扩展点

`serverutil.Option` 当前支持：

| Option | 说明 |
|---|---|
| `WithSetupOptions(opts ...setuputil.Option)` | 注入 LauncherManager 组件注册选项 |
| `WithAppOptionProvider(providers ...AppOptionProvider)` | 注入 Kratos App 选项 |
| `WithAuthManagerProvider(provider AuthManagerProvider)` | 注入认证管理器提供者 |
| `WithJaegerExporterProvider(provider JaegerExporterProvider)` | 注入 Jaeger exporter 提供者 |

## 独立启动

已有 `LauncherManager`、HTTP server 和 gRPC server 时，可以直接创建 App：

```go
app, err := serverutil.NewApp(launcherManager, httpServer, grpcServer)
serverutil.RunServer(app, cleanup)
```

## 文件说明

| 文件 | 说明 |
|---|---|
| `server.util.go` | ServerManager 接口和实现 |
| `server_all_in_one.util.go` | All-In-One 多服务启动 |
| `server_app.util.go` | Kratos App 创建 |
| `server_http.util.go` | HTTP 服务器创建和中间件配置 |
| `server_grpc.util.go` | gRPC 服务器创建和中间件配置 |
| `server_option.util.go` | 服务器和 All-In-One 选项 |
| `server_provider.util.go` | Wire Provider |
