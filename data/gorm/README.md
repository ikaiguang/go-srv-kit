# gorm - ORM 通用工具

`gorm/` 提供基于 GORM 的通用数据库操作工具，被 `data/mysql` 和 `data/postgres` 共同使用。

## 包名

```go
import gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
```

## 功能列表

| 文件 | 功能 | 说明 |
|------|------|------|
| `gorm_conn.kit.go` | 连接管理 | 数据库连接创建和连接池配置 |
| `gorm_logger.kit.go` | 日志 | GORM 日志集成（控制台 + JSON 文件） |
| `gorm_page.kit.go` | 分页 | 标准分页查询 |
| `gorm_order.kit.go` | 排序 | 安全排序（防 SQL 注入，校验字段名） |
| `gorm_batch_insert.kit.go` | 批量插入 | 分批写入 |
| `gorm_where.kit.go` | 条件构建 | WHERE 条件辅助 |
| `gorm_locking.kit.go` | 锁 | 悲观锁 `ForUpdate`、共享锁 `ForShare` |
| `gorm_hint.kit.go` | Hint | 查询提示 |
| `gorm_helper.kit.go` | 辅助工具 | 通用辅助函数 |
| `gorm_errors.kit.go` | 错误处理 | GORM 错误判断 |
| `gorm_callback.kit.go` | 回调 | GORM 回调注册 |

## 事务

```go
// 简单事务
gormpkg.ExecWithTransaction(db, func(tx *gorm.DB) error {
    // 在事务中操作
    return nil
})

// 手动事务
tx := gormpkg.NewTransaction(ctx, db)
tx.Do(ctx, func(ctx context.Context, tx *gorm.DB) error {
    return nil
})
tx.CommitAndErrRollback(ctx, resultErr)
```

## 分页

```go
gormpkg.PageQuery(db, page, pageSize)
```

## 排序

```go
gormpkg.OrderBy(db, "created_at", "desc")
// 自动校验字段名合法性，防止 SQL 注入
```
