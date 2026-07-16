# migration

`migration` 包提供轻量的数据库迁移记录能力，包名为 `migrationpkg`。它基于 GORM 保存迁移执行记录，帮助业务项目按迁移标识避免重复执行同一个 `Up` 操作，并提供对应的 `Down` 回滚入口。

## 安装

```bash
go get github.com/ikaiguang/go-srv-kit/data/gorm/v3
```

```go
import migrationpkg "github.com/ikaiguang/go-srv-kit/data/gorm/v3/migration"
```

## 核心概念

- `MigrationInterface`：单个迁移任务接口，包含 `Version`、`Identifier`、`Up` 和 `Down`。
- `MigrateRepo`：迁移执行入口，负责初始化迁移记录表、查询记录、创建记录、运行 `Up` 和 `Down`。
- `Migration` / `MigrationEntity`：默认迁移记录表模型。
- `DefaultTableName`：默认迁移记录表名，当前为 `srv_migration_record`。
- `NewAnyMigrator`：用函数快速创建任意迁移。
- `NewCreateTable`、`NewDropTable`：用 GORM migrator 创建或删除表的迁移封装。

## 基本流程

1. 业务项目创建好 `*gorm.DB`。
2. 使用 `NewMigrateRepo(db)` 创建迁移仓储。
3. 调用 `InitializeSchema(ctx)` 创建迁移记录表。
4. 为每个业务变更定义一个 `MigrationInterface`。
5. 调用 `RunMigratorUp(ctx, migrator)` 执行迁移。
6. 如需回滚，调用 `RunMigratorDown(ctx, migrator)`。

## 快速使用

```go
ctx := context.Background()
migrateRepo := migrationpkg.NewMigrateRepo(db)

if err := migrateRepo.InitializeSchema(ctx); err != nil {
	return err
}

migrator := migrationpkg.NewAnyMigrator(
	"v1.0.0",
	"v1.0.0:create:users",
	func() error {
		return db.WithContext(ctx).AutoMigrate(&User{})
	},
	func() error {
		return db.WithContext(ctx).Migrator().DropTable(&User{})
	},
)

if err := migrateRepo.RunMigratorUp(ctx, migrator); err != nil {
	return err
}
```

`RunMigratorUp` 会先按 `Identifier()` 查询迁移记录；如果记录已存在，则跳过 `Up`。`RunMigratorDown` 会执行 `Down`，随后删除对应迁移记录。

## 创建表迁移

实现了 `schema.Tabler` 的模型可以直接使用 `NewCreateTable` 或模型上的 `CreateTableMigrator`。

```go
type User struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;size:255;not null"`
}

func (User) TableName() string {
	return "users"
}

migrator := migrationpkg.NewCreateTable(
	db.WithContext(ctx).Migrator(),
	"v1.0.0",
	User{},
)

if err := migrateRepo.RunMigratorUp(ctx, migrator); err != nil {
	return err
}
```

也可以使用内置迁移记录表模型：

```go
schemaMigrator := migrationpkg.MigrationSchema.CreateTableMigrator(db.Migrator())
if err := migrateRepo.RunMigratorUp(ctx, schemaMigrator); err != nil {
	return err
}
```

通常初始化迁移记录表时直接调用 `InitializeSchema` 即可，不需要手动执行上面的内置 schema migrator。

## 迁移记录

迁移记录写入 `srv_migration_record`，主要字段包括：

- `server_version`：服务版本。
- `migration_identifier`：迁移唯一标识。
- `migration_batch`：迁移批次。
- `migration_desc`：迁移描述。
- `migration_extra_info`：额外信息。
- `created_time`、`updated_time`：记录时间。

可以通过 `QueryRecord` 查询某个迁移是否已执行，也可以通过 `CreateRecordByEntity` 写入自定义记录。

## 测试

```bash
go test ./migration
```

MySQL 和 Postgres 相关测试文件当前主要用于记录模型和方言差异，真实数据库验证需要准备对应数据库实例。

## 注意事项

- `Identifier()` 应保持全局唯一且稳定，推荐包含版本、动作和业务对象，例如 `v1.2.0:create:users`。
- `RunMigratorUp` 只在迁移记录不存在时执行 `Up`；如果 `Up` 成功但创建记录失败，下一次仍会再次尝试执行。
- `RunMigratorDown` 会删除迁移记录，调用方需要确认回滚操作的业务风险。
- 迁移函数内部应使用传入业务上下文创建的 `*gorm.DB`，并自行处理事务需求。
