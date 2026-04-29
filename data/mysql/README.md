# mysql - MySQL 数据库

`mysql/` 提供 MySQL 数据库连接的创建，基于 GORM + MySQL 驱动。

## 包名

```go
import mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql"
```

## 使用

```go
db, err := mysqlpkg.NewMysqlDB(config, opts...)
```

配置项包括：DSN、慢查询阈值、日志级别、连接池参数等。

通常通过 `service/mysql` 管理器间接使用。
