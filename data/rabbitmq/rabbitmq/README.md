# rabbitmq

`rabbitmq` 目录提供包 `rabbitmqpkg`，用于基于 Watermill AMQP 快速创建 RabbitMQ publisher、subscriber 和可复用连接。该包适合简单业务服务接入 RabbitMQ；需要更复杂队列、消费、发布配置时，可以继续使用底层 `watermill-amqp` API。

## 安装

```bash
go get github.com/ikaiguang/go-rabbitmq-kit
```

```go
import rabbitmqpkg "github.com/ikaiguang/go-rabbitmq-kit/rabbitmq"
```

## 核心能力

- `Config`：由 `config.proto` 生成的 RabbitMQ 配置，核心字段是 `Url`，并预留 TLS 相关字段。
- `NewConnection`：创建 `amqp.ConnectionWrapper` 的封装，便于后续复用连接。
- `NewPublisher` / `NewSubscriber`：基于配置直接创建 Watermill AMQP publisher 或 subscriber。
- `NewPublisherWithConnection` / `NewSubscriberWithConnection`：基于已有连接创建 publisher 或 subscriber。
- `NewPublisherAndSubscriber` / `NewPublisherAndSubscriberWithConnection`：同时创建发布者和订阅者。
- `WithLogger`：传入自定义 Watermill `LoggerAdapter`。
- `WithSlogLogger`：传入标准库 `*slog.Logger`。
- `WithNonDurable`：使用 non-durable queue 配置。
- `WithTLSConfig`：向 AMQP 连接配置注入 `*tls.Config`。
- `NewLogger` / `NewLoggerFromWriters`：创建 Watermill 日志适配器。

## 快速使用

创建发布者：

```go
conf := &rabbitmqpkg.Config{
	Url: "amqp://guest:guest@127.0.0.1:5672/",
}

publisher, err := rabbitmqpkg.NewPublisher(conf)
if err != nil {
	return err
}
defer publisher.Close()
```

创建订阅者：

```go
subscriber, err := rabbitmqpkg.NewSubscriber(conf)
if err != nil {
	return err
}
defer subscriber.Close()

messages, err := subscriber.Subscribe(ctx, "example.topic")
if err != nil {
	return err
}

for msg := range messages {
	// handle msg.Payload
	msg.Ack()
}
```

使用 `slog`：

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

publisher, err := rabbitmqpkg.NewPublisher(
	conf,
	rabbitmqpkg.WithSlogLogger(logger),
)
```

复用连接：

```go
conn, err := rabbitmqpkg.NewConnection(conf)
if err != nil {
	return err
}
defer conn.Close()

publisher, err := rabbitmqpkg.NewPublisherWithConnection(conn)
if err != nil {
	return err
}
defer publisher.Close()
```

启用 TLS：

```go
tlsConfig := &tls.Config{
	MinVersion: tls.VersionTLS12,
}

subscriber, err := rabbitmqpkg.NewSubscriber(
	conf,
	rabbitmqpkg.WithTLSConfig(tlsConfig),
)
```

## 测试

```bash
go test ./rabbitmq
```

当前测试文件包含需要 RabbitMQ 服务的集成示例。运行全量测试前，请确认本地 RabbitMQ 地址、账号和 topic 配置适合当前环境，避免连接到真实生产资源。

## 生成配置代码

`Config` 来自 `rabbitmq/config.proto`。修改 proto 后使用仓库 Makefile 生成：

```bash
make protoc-config-protobuf
```

不要手工修改 `config.pb.go` 或 `config.pb.validate.go`。

## 注意事项

- `Config.Url` 可能包含账号密码，示例、日志和提交记录中不要使用真实凭据。
- `WithNonDurable` 会创建 non-durable queue 配置，适合临时消息或示例，不适合需要持久化保障的场景。
- `WithTLSConfig` 只负责注入 TLS 配置；证书文件读取、密钥保护和服务端校验策略由调用方负责。
- `NewPublisherWithConnection`、`NewSubscriberWithConnection` 等函数复用传入连接，关闭 publisher/subscriber 时会同步关闭底层连接，使用前需要确认生命周期。
- 高级 AMQP 配置请直接参考 `watermill-amqp`，本包只封装常见入口。

## 参考

- [watermill-amqp](https://github.com/ThreeDotsLabs/watermill-amqp)
- [Watermill](https://watermill.io/)
