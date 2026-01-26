# Wire 依赖注入规范

## Wire 基础

### Wire 文件结构

```
cmd/{service}/export/
├── wire.go           # Wire 定义文件（手写）
├── wire_gen.go       # Wire 生成的代码（自动生成，不要修改）
└── main.export.go    # 导出函数
```

### Wire 定义文件模板

```go
//go:build wireinject
// +build wireinject

package exporter

import (
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-kratos/kratos/v2/transport/http"
    "github.com/google/wire"
    setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
)

func exportServices(launcherManager setuputil.LauncherManager, hs *http.Server, gs *grpc.Server) (Cleanup, error) {
    panic(wire.Build(
        // 基础设施
        setuputil.GetLogger,
        setuputil.GetServiceAPIManager,

        // Data 层 - 顺序：从底层到上层
        data.NewXxxData,
        data.NewYyyData,

        // Business 层
        biz.NewXxxBiz,
        biz.NewYyyBiz,

        // Service 层
        service.NewXxxService,
        service.NewYyyService,

        // 注册服务
        service.RegisterServices,
    ))
}
```

### 生成 Wire 代码

```bash
# 单个服务
wire ./cmd/ping-service/export

# 使用 Makefile
make generate
```

## Wire 高级用法

### 接口绑定

```go
// 在 biz/repo/ 定义接口
type XxxBizRepo interface {
    GetUser(ctx context.Context, id uint) (*bo.User, error)
}

// 在 wire.go 中绑定实现
func exportServices(...) {
    panic(wire.Build(
        wire.Bind(new(biz.XxxBizRepo), new(*data.xxxData)),
        data.NewXxxData,
    ))
}
```

### 结构体注入

```go
// 使用 wire.StructField 指定注入字段
panic(wire.Build(
    wire.Struct(new(service.XxxService), "*"),
    data.NewXxxData,
))
```

### Provider 函数

```go
// 提供额外的依赖
func provideGormDB(launcher setuputil.LauncherManager) (*gorm.DB, error) {
    return launcher.GetMysqlDBConn()
}

func exportServices(...) {
    panic(wire.Build(
        provideGormDB,
        data.NewXxxData,
    ))
}
```

### Provider Set

```go
// 创建可复用的 Provider Set
var DataSet = wire.NewSet(
    data.NewXxxData,
    data.NewYyyData,
)

var BizSet = wire.NewSet(
    biz.NewXxxBiz,
    biz.NewYyyBiz,
)

func exportServices(...) {
    panic(wire.Build(
        DataSet,
        BizSet,
        ServiceSet,
    ))
}
```

## 常见错误处理

### 循环依赖

```
错误: cycle detected
```

**解决方案**：
- 重构代码，引入中间层
- 使用接口解耦
- 检查依赖关系是否合理

### 类型不匹配

```
错误: no provider found for *biz.XxxBiz
```

**解决方案**：
- 检查返回类型是否匹配
- 检查是否需要使用 wire.Bind
- 检查函数命名是否正确

## 依赖注入顺序

**原则：从底层到上层**

```go
panic(wire.Build(
    // 1. 基础设施（Logger, Config, DB, Redis）
    setuputil.GetLogger,
    provideDatabase,
    provideRedis,

    // 2. Data 层（依赖基础设施）
    data.NewUserData,
    data.NewOrderData,

    // 3. Business 层（依赖 Data 层）
    biz.NewUserBiz,
    biz.NewOrderBiz,

    // 4. Service 层（依赖 Business 层）
    service.NewUserService,
    service.NewOrderService,
))
```

## 清理资源

```go
// Cleanup 函数
func (d *Data) Cleanup() {
    if d.db != nil {
        sqlDB, _ := d.db.DB()
        sqlDB.Close()
    }
    if d.redis != nil {
        d.redis.Close()
    }
}

// 在 Wire 中注册清理
func exportServices(...) (cleanuputil.CleanupManager, error) {
    panic(wire.Build(
        wire.Bind(new(cleanuputil.Cleanup, new(*Data)),
        data.NewData,
    ))
}
```
