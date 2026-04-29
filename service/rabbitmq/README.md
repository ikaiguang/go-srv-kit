# rabbitmq - RabbitMQ 连接管理

`rabbitmq/` 提供 RabbitMQ 连接的懒加载管理。

## 包名

```go
import rabbitmqutil "github.com/ikaiguang/go-srv-kit/service/rabbitmq"
```

## 使用

```go
manager, err := rabbitmqutil.NewRabbitmqManager(rabbitmqConfig, loggerManager)

if manager.Enable() {
    conn, err := manager.GetClient()  // *rabbitmqpkg.ConnectionWrapper
    defer manager.Close()
}
```

底层使用 `data/rabbitmq`（基于 Watermill）创建连接，自动配置日志，内置重连机制。
