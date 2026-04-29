# tracer - 链路追踪初始化

`tracer/` 提供 OpenTelemetry 链路追踪的初始化工具。

## 包名

```go
import tracerutil "github.com/ikaiguang/go-srv-kit/service/tracer"
```

## 使用

```go
// 使用 Jaeger Exporter 初始化
err := tracerutil.InitTracerWithJaegerExporter(appConfig, jaegerExporter)

// 使用默认配置初始化（无 Exporter）
err := tracerutil.InitTracer(appConfig)
```

初始化后，`kratos/middleware` 中的 `tracing.Server()` 中间件会自动采集链路信息。
