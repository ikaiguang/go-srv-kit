# Wire 依赖注入模板

## Wire 定义文件模板
```go
//go:build wireinject
// +build wireinject

package exporter

import (
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-kratos/kratos/v2/transport/http"
    "github.com/google/wire"
    setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/biz/biz"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/data/data"
    "github.com/ikaiguang/go-srv-kit/testdata/{service_name}/internal/service/service"
)

// exportServices 导出服务
func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (cleanuputil.CleanupManager, error) {
    panic(wire.Build(
        // 基础设施
        setuputil.GetLogger,

        // Data 层
        data.New{Xxx}Data,

        // Business 层
        biz.New{Xxx}Biz,

        // Service 层
        service.New{Xxx}Service,

        // 注册服务
        service.RegisterServices,
    ))
}
```

## 接口绑定模板
```go
// 当接口和实现分离时，使用 wire.Bind 绑定
panic(wire.Build(
    wire.Bind(new(biz.{Xxx}BizRepo), new(*data.{xxx}Data)),
    data.New{Xxx}Data,
))
```

## Provider 函数模板
```go
// 提供数据库连接
func provideDatabase(launcher setuputil.LauncherManager) (*gorm.DB, error) {
    return launcher.GetMysqlDBConn()
}

// 提供缓存连接
func provideRedis(launcher setuputil.LauncherManager) (*redis.Client, error) {
    return launcher.GetRedisConn()
}

// 在 wire.Build 中使用
panic(wire.Build(
    provideDatabase,
    provideRedis,
    data.New{Xxx}Data,
))
```

## 依赖注入顺序规则
```
1. 基础设施（Logger, Config, DB, Redis）
2. Data 层（依赖基础设施）
3. Business 层（依赖 Data 层）
4. Service 层（依赖 Business 层）
5. 注册服务（依赖 Service 层）
```
