---
inclusion: fileMatch
fileMatchPattern: "**/wire*.go"
---

# Wire 依赖注入规范

## 文件结构

```
cmd/{service}/export/
├── wire.go           # Wire 定义文件（手写）
├── wire_gen.go       # 自动生成（不要修改）
└── main.export.go    # 导出函数
```

## Wire 定义模板

```go
//go:build wireinject
// +build wireinject

package exporter

func exportServices(launcherManager setupv2.LauncherManager, hs *http.Server, gs *grpc.Server) (Cleanup, error) {
    panic(wire.Build(
        // 1. 基础设施
        setupv2.GetLogger,

        // 2. Data 层（从底层到上层）
        data.NewXxxData,

        // 3. Business 层
        biz.NewXxxBiz,

        // 4. Service 层
        service.NewXxxService,

        // 5. 注册服务
        service.RegisterServices,
    ))
}
```

## 接口绑定

```go
wire.Bind(new(biz.XxxBizRepo), new(*data.xxxData))
```

## 依赖注入顺序

基础设施 (Logger, DB, Redis) → Data 层 → Business 层 → Service 层

## 常见错误

- `cycle detected`：循环依赖，引入中间层或接口解耦
- `no provider found`：检查返回类型和 wire.Bind
- 生成失败：`rm ./cmd/*/export/wire_gen.go` 后重新 `wire ./cmd/*/export`

## 生成命令

```bash
wire ./cmd/{service}/export
# 或
make generate
```
