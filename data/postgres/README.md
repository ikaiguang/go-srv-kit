# postgres - PostgreSQL 数据库

`postgres/` 提供 PostgreSQL 数据库连接的创建，基于 GORM + PostgreSQL 驱动。

## 包名

```go
import psqlpkg "github.com/ikaiguang/go-srv-kit/data/postgres"
```

## 使用

```go
db, err := psqlpkg.NewPostgresDB(config, opts...)
```

配置项包括：DSN、慢查询阈值、日志级别、连接池参数等。

通常通过 `service/postgres` 管理器间接使用。
