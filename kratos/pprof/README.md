# pprof

`pprof` 包的包名为 `pprofpkg`，用于给 Kratos HTTP server 注册 Go 标准 pprof 调试路由。适合在需要本地或受控环境性能诊断时启用。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/kratos/v3
```

```go
import pprofpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/pprof"
```

## 核心能力

- `RegisterPprof`：注册 `/debug/pprof`、`/debug/pprof/profile`、`/debug/pprof/trace`、`/debug/heap`、`/debug/goroutine` 等调试端点。

## 快速使用

```go
srv := http.NewServer()
pprofpkg.RegisterPprof(srv)
```

## 测试

```bash
go test ./pprof
```

## 注意事项

- pprof 端点可能暴露性能、goroutine、内存和路径信息，生产环境应通过鉴权、网络隔离或独立调试端口限制访问。
- `RegisterPprof` 需要传入 Kratos HTTP server，不适用于标准库 `net/http.Server`。
