# mongo

`mongo` module 提供包 `mongo`，用于在 Go 服务中创建 MongoDB v2 驱动客户端、挂载 Mongo command monitor、记录慢查询日志，并复用常见 Mongo 操作常量。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/mongo/v3
```

```go
import "github.com/ikaiguang/go-srv-kit/data/mongo/v3"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 Mongo 配置类型，包含连接地址、hosts、连接池、超时、心跳、空闲连接和慢查询阈值等字段。
- `NewMongoClient(config *Config, logger *slog.Logger)`：根据配置创建 `*mongo.Client`，设置 command monitor，并在返回前执行 `Ping`。
- `NewMonitor(logger *slog.Logger, opts ...MonitorOption)`：创建 MongoDB command monitor，记录 started、succeeded 和 failed 事件。
- `WithSlowThreshold(time.Duration)`：设置慢查询阈值；成功命令耗时达到阈值时输出 warn 日志，否则输出 debug 日志。
- `OperationSet`、`OperationIn`、`OperationLookup` 等常量：常用 Mongo 更新、查询和聚合操作符。
- `FieldFrom`、`FieldLocalField`、`FieldForeignField` 等常量：常用 `$lookup` 字段名。

## 快速使用

```go
package example

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/ikaiguang/go-srv-kit/data/mongo/v3"
	"google.golang.org/protobuf/types/known/durationpb"
)

func openMongo(ctx context.Context) error {
	cfg := &mongo.Config{
		AppName:           "example-service",
		Addr:              os.Getenv("MONGO_URI"),
		ConnectTimeout:    durationpb.New(3 * time.Second),
		Timeout:           durationpb.New(3 * time.Second),
		HeartbeatInterval: durationpb.New(10 * time.Second),
		MaxConnIdleTime:   durationpb.New(60 * time.Second),
		SlowThreshold:     durationpb.New(100 * time.Millisecond),
		MaxPoolSize:       100,
		MinPoolSize:       2,
		MaxConnecting:     10,
	}

	client, err := mongo.NewMongoClient(cfg, slog.Default())
	if err != nil {
		return err
	}
	defer func() { _ = client.Disconnect(ctx) }()

	return nil
}
```

使用操作常量组织更新条件：

```go
update := map[string]any{
	mongo.OperationSet: map[string]any{
		"updated_by": "system",
	},
	mongo.OperationInc: map[string]any{
		"version": 1,
	},
}
_ = update
```

## 配置生成

`Config` 的源文件是 `config.proto`，生成文件包括：

- `config.pb.go`
- `config.pb.validate.go`

不要手工修改生成文件。修改 `config.proto` 后在仓库根目录运行：

```bash
protoc --proto_path=. --proto_path="$(go env GOPATH)/src" --proto_path=./third_party \
  --go_out=paths=source_relative:. \
  --validate_out=paths=source_relative,lang=go:. \
  data/mongo/config.proto
```

## 测试

```bash
go test ./...
```

当前 `TestNewMongoClient` 会连接测试配置中的 MongoDB 地址并执行 `Ping`。如果本地没有对应 MongoDB 实例，测试会因连接失败而失败；这属于集成测试环境依赖。

只做编译级验证时可以运行：

```bash
go test ./... -run '^$'
```

## 注意事项

- `NewMongoClient` 会立即 `Ping` MongoDB；连接串、hosts、认证信息和网络环境必须可用。
- 示例中请使用环境变量或占位值传递连接串，不要把真实账号、密码、Token 或内网地址写入文档和代码。
- command monitor 会记录 Mongo 命令内容，业务侧应避免把敏感明文放入会被日志记录的命令字段。
- `logger` 可以传 `nil`，包内会回退到 `slog.Default()`。
- 当前实现直接读取配置中的 duration 字段，调用方应为超时和慢查询字段提供合适的 `durationpb.Duration` 值。
