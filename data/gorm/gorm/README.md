# gorm

`gorm` 包提供围绕 `gorm.io/gorm` 的常用辅助能力，包名为 `gormpkg`。它适合在业务项目的数据访问层复用，用来统一数据库连接、分页、排序、条件拼装、事务、锁、索引 hint、批量插入和常见错误判断。

## 安装

```bash
go get github.com/ikaiguang/go-gorm-kit
```

导入时建议使用别名，避免和官方 GORM 包名混淆：

```go
import gormpkg "github.com/ikaiguang/go-gorm-kit/gorm"
```

## 核心能力

- `NewDB`：基于 `gorm.Dialector` 创建 `*gorm.DB`，同时设置日志、预编译、默认事务策略和连接池参数。
- `ExecWithTransaction`、`NewTransaction`：事务执行辅助。
- `Paginator`、`InitPaginatorArgs`：把页码参数转换为 GORM `Limit` 与 `Offset`。
- `AssembleWheres`、`AssembleOrders`：拼装查询条件与排序，并对字段名做基础校验。
- `UseIndex`、`ForceIndex`、`IgnoreIndex`：封装 GORM index hint。
- `ForUpdate`、`ForUpdateNowait`、`ForShareOfTable`：封装常用行锁语义。
- `BatchInsert`、`BatchInsertWithContext`：分批执行批量插入，支持 `INSERT IGNORE` 和冲突处理。
- `Model`、`ModelForMysql`、`ModelForPostgres`：常用字段模型。
- `QueryUndeletedData`、`SoftDelete`、`Deleted`：软删除和物理删除辅助。
- `IsErrRecordNotFound`、`IsErrDuplicatedKey`：常见 GORM 错误判断。

## 创建数据库连接

`NewDB` 接收具体数据库驱动的 `gorm.Dialector`。调用方仍然需要在业务项目中引入对应驱动，例如 MySQL、Postgres 或 SQLite 的 GORM driver。

```go
package data

import (
	"time"

	gormpkg "github.com/ikaiguang/go-gorm-kit/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	return gormpkg.NewDB(mysql.Open(dsn), &gormpkg.ConnOption{
		LoggerEnable:              true,
		LoggerLevel:               logger.Warn,
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: true,
		ConnMaxActive:             50,
		ConnMaxIdle:               10,
		ConnMaxLifetime:           time.Hour,
		ConnMaxIdleTime:           10 * time.Minute,
	})
}
```

## 分页、条件和排序

`AssembleWheres` 与 `AssembleOrders` 会校验字段名，只允许字母、数字、下划线和点号。字段名、操作符和表名仍应来自服务端可信配置，不要把用户输入直接拼进字段名或操作符。

```go
args := gormpkg.InitPaginatorArgs(1, 20)
args.PageWheres = []*gormpkg.Where{
	gormpkg.NewWhere("users.status", "=", "enabled"),
}
args.PageOrders = []*gormpkg.Order{
	gormpkg.NewOrder("users.id", gormpkg.DefaultOrderDesc),
}

query := db.Model(&User{})
query = gormpkg.AssembleWheres(query, args.PageWheres)
query = gormpkg.AssembleOrders(query, args.PageOrders)
query = gormpkg.Paginator(query, args.PageOption)

var users []*User
if err := query.Find(&users).Error; err != nil {
	return err
}
```

仅当字段、操作符和表达式完全来自可信代码时，才使用 `UnsafeAssembleWheres` 或 `UnsafeAssembleOrders`。

## 事务

简单场景优先使用 `ExecWithTransaction`：

```go
err := gormpkg.ExecWithTransaction(db, func(tx *gorm.DB) error {
	if err := tx.Create(&user).Error; err != nil {
		return err
	}
	return tx.Create(&profile).Error
})
```

需要显式控制提交和回滚时，可以使用 `NewTransaction`：

```go
tx := gormpkg.NewTransaction(ctx, db)
err := tx.Do(ctx, func(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&user).Error
})
return tx.CommitAndErrRollback(ctx, err)
```

## 批量插入

批量插入通过 `BatchInsertRepo` 由调用方提供表名、列名、占位符和值。`TableName`、`InsertColumns` 和冲突 SQL 必须由可信代码生成，不能拼接未校验的外部输入。

```go
type UserRows []*User

func (r *UserRows) TableName() string { return "users" }
func (r *UserRows) Len() int          { return len(*r) }

func (r *UserRows) InsertColumns() ([]string, string) {
	return []string{"name", "age"}, "?, ?"
}

func (r *UserRows) InsertValues(args *gormpkg.BatchInsertValueArgs) ([]any, []string) {
	rows := (*r)[args.StepStart:args.StepEnd]
	values := make([]any, 0, len(rows)*2)
	placeholders := make([]string, 0, len(rows))
	for _, row := range rows {
		placeholders = append(placeholders, "("+args.InsertPlaceholder+")")
		values = append(values, row.Name, row.Age)
	}
	return values, placeholders
}

rows := UserRows{
	{Name: "alice", Age: 18},
	{Name: "bob", Age: 20},
}
if err := gormpkg.BatchInsertWithContext(ctx, db, &rows); err != nil {
	return err
}
```

## 锁和索引 Hint

```go
err := gormpkg.ForUpdate(db.WithContext(ctx)).
	Where("id = ?", id).
	First(&user).Error

err = gormpkg.ForceIndex(db.WithContext(ctx), "idx_users_status").
	Where("status = ?", "enabled").
	Find(&users).Error
```

索引 hint 依赖数据库方言支持，使用前请确认目标数据库行为。

## 测试

```bash
go test ./gorm
```

仓库内部分数据库集成测试当前以注释形式保留，需要真实数据库环境时再按测试文件中的示例配置连接。

## 注意事项

- 本包不管理具体数据库驱动依赖，业务项目需要自行引入 `gorm.io/driver/...`。
- `BatchInsert` 的表名、列名和冲突 SQL 由调用方实现，务必保持可信来源。
- 安全版本的条件和排序 helper 只校验字段名，不校验操作符语义。
- `RemoveCallback` 当前是占位实现，不会移除 GORM callback。
