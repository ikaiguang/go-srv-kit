# mongo - MongoDB 客户端

`mongo/` 提供 MongoDB 客户端的创建，含慢查询监控和命令日志。

## 包名

```go
import mongopkg "github.com/ikaiguang/go-srv-kit/data/mongo"
```

## 使用

```go
client, err := mongopkg.NewMongoClient(config, logger)
```

## 功能

| 文件 | 说明 |
|------|------|
| `mongo.kit.go` | 客户端创建（连接池、超时、心跳配置） |
| `mongo_monitor.kit.go` | 命令监控（慢查询日志） |
| `mongo_log.kit.go` | 日志适配器 |
| `mongo_const.kit.go` | 常量定义 |

通常通过 `service/mongo` 管理器间接使用。
