# postgres - PostgreSQL 连接管理

`postgres/` 提供 PostgreSQL 数据库连接的懒加载管理。

## 包名

```go
import postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
```

## 使用

```go
manager, err := postgresutil.NewPostgresManager(psqlConfig, loggerManager)

if manager.Enable() {
    db, err := manager.GetDB()  // *gorm.DB
    defer manager.Close()
}
```

底层使用 `data/postgres` + `data/gorm` 创建连接，自动配置日志。
