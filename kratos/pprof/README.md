# pprof - 性能分析

`pprof/` 提供 Go pprof 性能分析路由的快捷注册。

## 包名

```go
import pprofpkg "github.com/ikaiguang/go-srv-kit/kratos/pprof"
```

## 使用

```go
pprofpkg.RegisterPprof(httpServer)
```

注册后可访问以下路由：

| 路由 | 说明 |
|------|------|
| `/debug/pprof` | pprof 首页 |
| `/debug/pprof/profile` | CPU 分析 |
| `/debug/heap` | 堆内存分析 |
| `/debug/goroutine` | Goroutine 分析 |
| `/debug/pprof/trace` | 执行追踪 |
| `/debug/allocs` | 内存分配分析 |
| `/debug/block` | 阻塞分析 |
| `/debug/mutex` | 互斥锁分析 |

> 生产环境建议通过配置控制是否启用。
