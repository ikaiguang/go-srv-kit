# mysql - MySQL 连接管理

`mysql/` 提供 MySQL 数据库连接的懒加载管理。

## 包名

```go
import mysqlutil "github.com/ikaiguang/go-srv-kit/service/mysql"
```

## 使用

```go
manager, err := mysqlutil.NewMysqlManager(mysqlConfig, loggerManager)

if manager.Enable() {
    db, err := manager.GetDB()  // *gorm.DB
    defer manager.Close()
}
```

底层使用 `data/mysql` + `data/gorm` 创建连接，自动配置日志（控制台 + 文件）。
