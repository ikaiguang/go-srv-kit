# server - 服务器管理

`server/` 负责 HTTP/gRPC 服务器的创建、中间件配置和应用启动。

## 包名

```go
import serverutil "github.com/ikaiguang/go-srv-kit/service/server"
```

## 核心功能

### ServerManager

管理 HTTP 和 gRPC 服务器的单例创建：

```go
srvManager, err := serverutil.NewServerManager(launcherManager, authWhitelist)
httpServer, err := srvManager.GetHTTPServer()
grpcServer, err := srvManager.GetGRPCServer()
app, err := srvManager.GetApp()
```

### All-In-One 模式

支持多个服务合并到一个进程中启动：

```go
app, cleanup, err := serverutil.AllInOneServer(
    configFilePath,
    configOpts,
    []serverutil.ServiceExporter{service1, service2},
    authWhitelist,
)
defer cleanup()
serverutil.RunServer(app, nil)
```

### 独立启动

```go
app, err := serverutil.NewApp(launcherManager, httpServer, grpcServer)
serverutil.RunServer(app, cleanup)
```

## 文件说明

| 文件 | 说明 |
|------|------|
| `server.util.go` | ServerManager 接口和实现 |
| `server_all_in_one.util.go` | All-In-One 多服务启动 |
| `server_app.util.go` | kratos.App 创建 |
| `server_http.util.go` | HTTP 服务器创建和中间件配置 |
| `server_grpc.util.go` | gRPC 服务器创建和中间件配置 |
| `server_init.util.go` | 服务器初始化（Tracer、Pprof 等） |
| `server_option.util.go` | 服务器选项 |
| `server_provider.util.go` | Wire Provider |
