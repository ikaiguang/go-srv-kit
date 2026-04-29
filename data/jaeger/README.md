# jaeger - Jaeger 链路追踪

`jaeger/` 提供 Jaeger Exporter 的创建，支持 HTTP 和 gRPC 两种上报协议。

## 包名

```go
import jaegerpkg "github.com/ikaiguang/go-srv-kit/data/jaeger"
```

## 使用

```go
exporter, err := jaegerpkg.NewJaegerExporter(config)
```

## 上报协议

| Kind | 说明 |
|------|------|
| `KindHTTP` | 通过 HTTP 上报（otlptracehttp） |
| `KindGRPC` | 通过 gRPC 上报（otlptracegrpc） |

创建前会自动检测目标地址的连接可用性。

通常通过 `service/jaeger` 管理器间接使用。
