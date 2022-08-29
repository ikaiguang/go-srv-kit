package migrationutil

import (
	"gorm.io/gorm"
	"time"
)

// MigrateRepo ...
type MigrateRepo interface {
	QueryOneByIdentifier(identifier string) (dataModel *MigrationEntity, isNotFound bool, err error)
	CreateRecord(dataModel *MigrationEntity) (err error)
	CreateDefaultRecord(serverVersion, migrateIdentifier string) (err error)
	InitializeMigrationSchema() (err error)
}

// migrate ...
type migrate struct {
	dbConn        *gorm.DB
	migrationRepo MigrationDataRepo
}

// NewMigrateRepo ...
func NewMigrateRepo(dbConn *gorm.DB) MigrateRepo {
	return &migrate{
		dbConn:        dbConn,
		migrationRepo: NewMigrationDataRepo(dbConn),
	}
}

// QueryOneByIdentifier 查询迁移记录
func (s *migrate) QueryOneByIdentifier(identifier string) (dataModel *MigrationEntity, isNotFound bool, err error) {
	return s.migrationRepo.QueryOneByIdentifier(identifier)
}

// CreateRecord 默认的迁移记录
func (s *migrate) CreateRecord(dataModel *MigrationEntity) (err error) {
	return s.migrationRepo.Create(dataModel)
}

// CreateDefaultRecord 默认的迁移记录
func (s *migrate) CreateDefaultRecord(serverVersion, migrateIdentifier string) (err error) {
	var (
		now = time.Now()
	)
	dataModel := &MigrationEntity{
		ServerVersion:       serverVersion,
		MigrationIdentifier: migrateIdentifier,
		MigrationBatch:      1,
		MigrationDesc:       "数据库迁移：" + migrateIdentifier,
		MigrationExtraInfo:  "{}",
		CreatedTime:         now,
		UpdatedTime:         now,
	}
	return s.dbConn.Create(dataModel).Error
}

// InitializeMigrationSchema 初始化创建迁移表
func (s *migrate) InitializeMigrationSchema() (err error) {
	if s.dbConn.Migrator().HasTable(MigrationSchema) {
		return err
	}
	err = s.dbConn.Migrator().CreateTable(MigrationSchema)
	if err != nil {
		return err
	}

	// 记录
	var (
		serverVersion     = "v0.0.1"
		migrateIdentifier = serverVersion + ":init:create:migration_schema"
	)
	return s.CreateDefaultRecord(serverVersion, migrateIdentifier)
}
