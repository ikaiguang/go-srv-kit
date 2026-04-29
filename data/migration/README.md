# migration - 数据库迁移

`migration/` 提供数据库迁移框架，支持建表、删表和自定义迁移操作。

## 包名

```go
import migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
```

## 迁移类型

### 建表

```go
migrator := migrationpkg.NewCreateTable(db.Migrator(), "v1.0.0", &UserPO{})
err := migrator.Up()   // 建表（已存在则跳过）
err := migrator.Down() // 删表
```

### 删表

```go
migrator := migrationpkg.NewDropTable(db.Migrator(), "v1.0.0", &UserPO{})
err := migrator.Up()   // 删表
err := migrator.Down() // 建表
```

### 自定义迁移

```go
migrator := migrationpkg.NewAnyMigrator("v1.0.0", "add_user_index",
    func() error { /* up */ return nil },
    func() error { /* down */ return nil },
)
```

## MigrationInterface

所有迁移实现统一接口：

```go
type MigrationInterface interface {
    Version() string
    Identifier() string
    Up() error
    Down() error
}
```

## 参考

- 迁移工具入口：`testdata/ping-service/cmd/database-migration/`
