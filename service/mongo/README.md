# mongo - MongoDB 客户端管理

`mongo/` 提供 MongoDB 客户端的懒加载管理。

## 包名

```go
import mongoutil "github.com/ikaiguang/go-srv-kit/service/mongo"
```

## 使用

```go
manager, err := mongoutil.NewMongoManager(mongoConfig, loggerManager)

if manager.Enable() {
    client, err := manager.GetMongoClient()  // *mongo.Client
    defer manager.Close()
}
```

底层使用 `data/mongo` 创建客户端，自动配置慢查询监控和日志。
