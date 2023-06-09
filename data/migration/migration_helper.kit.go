package migrationpkg

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	DefaultTableName         = "srv_migration_record"
	DefaultVersion           = "v0.0.1"
	DefaultInitIdentifier    = DefaultVersion + ":init:create:migration_schema"
	FieldMigrationIdentifier = "migration_identifier"
)

// MigrateRepo ...
type MigrateRepo interface {
	InitializeSchema(ctx context.Context) (err error)
	QueryRecord(ctx context.Context, identifier string) (dataModel *MigrationEntity, isNotFound bool, err error)
	CreateRecord(ctx context.Context, serverVersion, migrateIdentifier string) (err error)
	CreateRecordByEntity(ctx context.Context, dataModel *MigrationEntity) (err error)
	RunMigratorUp(ctx context.Context, migratorRepo MigrationInterface) error
	RunMigratorDown(ctx context.Context, migratorRepo MigrationInterface) error
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

// InitializeSchema 初始化创建迁移表
func (s *migrate) InitializeSchema(ctx context.Context) (err error) {
	migrator := s.dbConn.WithContext(ctx).Migrator()
	if migrator.HasTable(MigrationSchema) {
		return err
	}
	err = migrator.CreateTable(MigrationSchema)
	if err != nil {
		return err
	}

	// 记录
	var (
		serverVersion     = DefaultVersion
		migrateIdentifier = DefaultInitIdentifier
	)
	serverVersion, migrateIdentifier = s.initIdentifier(serverVersion, migrateIdentifier)
	return s.CreateRecord(ctx, serverVersion, migrateIdentifier)
}

// RunMigratorUp 运行迁移
func (s *migrate) RunMigratorUp(ctx context.Context, migratorRepo MigrationInterface) error {
	dataModel, isNotFound, err := s.QueryRecord(ctx, migratorRepo.Identifier())
	if err != nil {
		return err
	}
	if !isNotFound && dataModel.Id > 0 {
		return nil
	}

	// up
	if err = migratorRepo.Up(); err != nil {
		return err
	}
	return s.CreateRecord(ctx, migratorRepo.Version(), migratorRepo.Identifier())
}

// RunMigratorDown 运行迁移
func (s *migrate) RunMigratorDown(ctx context.Context, migratorRepo MigrationInterface) error {
	// down
	if err := migratorRepo.Down(); err != nil {
		return err
	}
	return s.migrationRepo.DeleteOneByIdentifier(ctx, migratorRepo.Identifier())
}

// QueryRecord 查询迁移记录
func (s *migrate) QueryRecord(ctx context.Context, identifier string) (dataModel *MigrationEntity, isNotFound bool, err error) {
	return s.migrationRepo.QueryOneByIdentifier(ctx, identifier)
}

// CreateRecordByEntity 迁移记录
func (s *migrate) CreateRecordByEntity(ctx context.Context, dataModel *MigrationEntity) (err error) {
	return s.migrationRepo.Create(ctx, dataModel)
}

// CreateRecord 迁移记录
func (s *migrate) CreateRecord(ctx context.Context, serverVersion, migrateIdentifier string) (err error) {
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
	return s.migrationRepo.Create(ctx, dataModel)
}

// initIdentifier ...
func (s *migrate) initIdentifier(serverVersion, migrateIdentifier string) (string, string) {
	if serverVersion == "" {
		serverVersion = Version
	}
	if migrateIdentifier == "" {
		migrateIdentifier = serverVersion + ":init:create:migration_schema"
	}
	return serverVersion, migrateIdentifier
}
