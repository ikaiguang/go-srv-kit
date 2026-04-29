# database - 数据库迁移辅助

`database/` 提供数据库迁移的辅助工具和选项。

## 包名

```go
import dbutil "github.com/ikaiguang/go-srv-kit/service/database"
```

## 使用

```go
// 定义迁移函数
var migrationFunc dbutil.MigrationFunc = func(dbConn *gorm.DB, opts ...dbutil.MigrationOption) {
    // 执行迁移
}

// 带选项执行
migrationFunc(dbConn, dbutil.WithLogger(logger), dbutil.WithClose())
```

## 选项

| 选项 | 说明 |
|------|------|
| `WithLogger(logger)` | 设置迁移日志 |
| `WithClose()` | 迁移完成后关闭数据库连接 |
