# jaeger

`jaeger` 目录提供包 `jaegerpkg`，导入路径为 `github.com/ikaiguang/go-jaeger-kit/jaeger`。该包负责根据配置创建 OpenTelemetry OTLP trace exporter，用于把服务链路追踪数据发送到 Jaeger Collector 或兼容 OTLP 的后端。

## 安装

```bash
go get github.com/ikaiguang/go-jaeger-kit
```

```go
import jaegerpkg "github.com/ikaiguang/go-jaeger-kit/jaeger"
```

## 核心能力

- `Config`：由 `jaeger/config.proto` 生成的配置结构，包含 `Kind`、`Addr`、`IsInsecure`、`Timeout` 等字段。
- `KindHTTP` / `KindGRPC`：用于选择 OTLP HTTP 或 OTLP gRPC exporter。
- `NewExporter`：根据 `Config.Kind` 创建 exporter；`KindHTTP` 使用 HTTP exporter，其他值默认走 gRPC exporter。
- `NewJaegerExporter`：`NewExporter` 的兼容入口。
- `NewHTTPExporter`：直接创建 OTLP HTTP trace exporter。
- `NewGRPCExporter`：直接创建 OTLP gRPC trace exporter。
- `Validate` / `ValidateAll`：由 `protoc-gen-validate` 生成的配置校验方法；当前 proto 未声明额外字段规则。

## 快速使用

```go
package main

import (
	"context"
	"log"
	"time"

	jaegerpkg "github.com/ikaiguang/go-jaeger-kit/jaeger"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	exp, err := jaegerpkg.NewExporter(&jaegerpkg.Config{
		Kind:       string(jaegerpkg.KindGRPC),
		Addr:       "localhost:4317",
		IsInsecure: true,
		Timeout:    durationpb.New(30 * time.Second),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = exp.Shutdown(context.Background())
	}()

	// 将 exp 交给 go.opentelemetry.io/otel/sdk/trace.NewBatchSpanProcessor
	// 或项目现有 tracing 初始化流程。
}
```

HTTP exporter 示例：

```go
exp, err := jaegerpkg.NewExporter(&jaegerpkg.Config{
	Kind:       string(jaegerpkg.KindHTTP),
	Addr:       "localhost:4318",
	IsInsecure: true,
	Timeout:    durationpb.New(30 * time.Second),
})
```

## 测试

```bash
go test ./jaeger
```

如果默认 Go 构建缓存目录不可写，可使用：

```bash
env GOCACHE=/tmp/go-build-cache go test ./jaeger
```

## 生成文件

本目录包含由 `jaeger/config.proto` 生成的文件：

- `config.pb.go`
- `config.pb.validate.go`
- `config.swagger.json`

不要手动修改生成文件。修改 `config.proto` 后，按仓库 Makefile 重新生成：

```bash
make protoc-config-protobuf
```

## 注意事项

- `NewExporter` 会先调用地址有效性检查，`Config.Addr` 应为可解析的 `host:port`。
- `Config.Kind` 只有值为 `http` 时走 HTTP exporter；其他值会走 gRPC exporter。
- `Config.IsInsecure` 会启用非 TLS 连接，生产环境应根据链路追踪后端和网络边界谨慎配置。
- `Config.WithHttpBasicAuth`、`Username` 和 `Password` 当前未被 `NewHTTPExporter` 或 `NewGRPCExporter` 使用。
- `Option` 和 `WithWriter` 已定义，但当前 exporter 创建流程未消费这些选项。
