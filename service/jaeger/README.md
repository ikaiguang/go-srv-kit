# jaeger - Jaeger 链路追踪管理

`jaeger/` 提供 Jaeger Exporter 的懒加载管理。

## 包名

```go
import jaegerutil "github.com/ikaiguang/go-srv-kit/service/jaeger"
```

## 使用

```go
manager, err := jaegerutil.NewJaegerManager(jaegerConfig)

if manager.Enable() {
    exporter, err := manager.GetExporter()  // *otlptrace.Exporter
    defer manager.Close()
}
```

底层使用 `data/jaeger` 创建 Exporter，支持 HTTP 和 gRPC 两种上报协议。
