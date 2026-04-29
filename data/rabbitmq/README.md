# rabbitmq - RabbitMQ 消息队列

`rabbitmq/` 提供基于 [Watermill](https://github.com/ThreeDotsLabs/watermill) 的 RabbitMQ 客户端封装。

## 包名

```go
import rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
```

## 使用

```go
// 创建连接（内置重连机制）
conn, err := rabbitmqpkg.NewConnection(config, opts...)

// 基于连接创建发布者和订阅者
publisher, err := rabbitmqpkg.NewPublisherWithConnection(conn, opts...)
subscriber, err := rabbitmqpkg.NewSubscriberWithConnection(conn, opts...)

// 同时创建发布者和订阅者
publisher, subscriber, err := rabbitmqpkg.NewPublisherAndSubscriberWithConnection(conn, opts...)
```

## 选项

| 选项 | 说明 |
|------|------|
| `WithLogger(logger)` | 设置日志 |
| `WithNonDurable()` | 非持久化队列 |
| `WithTLSConfig(tlsConfig)` | TLS 配置 |

## 参考

- [watermill-amqp](https://github.com/ThreeDotsLabs/watermill-amqp)
- [go-kratos/examples: event](https://github.com/go-kratos/examples/blob/main/event/event/event.go)

通常通过 `service/rabbitmq` 管理器间接使用。
