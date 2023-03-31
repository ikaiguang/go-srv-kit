package dbv1_0_0_example

import (
	"gorm.io/gorm"

	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	schemas "github.com/ikaiguang/go-srv-kit/example/internal/infra/mysql/schema/example"
)

// Upgrade .
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	upgradeHandler := NewMigrateHandler(dbConn, migrateRepo)

	// 创建表 example
	err = upgradeHandler.CreateTableExample()
	if err != nil {
		return err
	}
	return err
}

// migrate 数据库迁移
type migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationutil.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) *migrate {
	return &migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// CreateTableExample ...
func (s *migrate) CreateTableExample() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.ExampleSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.ExampleSchema)
}
