package migrationutil

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	Version = "v0.0.1"
)

var (
	_ MigrationInterface = (*AnyMigrator)(nil)
	_ MigrationInterface = (*CreateTable)(nil)
	_ MigrationInterface = (*DropTable)(nil)
)

// MigrationInterface 数据库迁移
type MigrationInterface interface {
	// Version 版本
	Version() string
	// Identifier 迁移标识
	Identifier() string
	// Up 运行迁移
	Up() error
	// Down 回滚迁移
	Down() error
}

// ===== Any =====

// AnyMigrator 任意
type AnyMigrator struct {
	version    string
	identifier string
	up         AnyMigratorExec
	down       AnyMigratorExec
}

// AnyMigratorExec ...
type AnyMigratorExec func() error

// NewAnyMigrator ...
func NewAnyMigrator(version, identifier string, up, down AnyMigratorExec) MigrationInterface {
	return &AnyMigrator{
		version:    version,
		identifier: identifier,
		up:         up,
		down:       down,
	}
}

// Version implements the Migration.
func (s *AnyMigrator) Version() string {
	return s.version
}

// Identifier implements the Migration.
func (s *AnyMigrator) Identifier() string {
	return s.identifier
}

// Up implements the Migration
func (s *AnyMigrator) Up() error {
	if s.up == nil {
		return nil
	}
	return s.up()
}

// Down implements the Migration
func (s *AnyMigrator) Down() error {
	if s.down == nil {
		return nil
	}
	return s.down()
}

// ===== 创建表 =====

// CreateTable 创建表
type CreateTable struct {
	migrator gorm.Migrator

	version string
	table   schema.Tabler
}

// NewCreateTable 创建表
func NewCreateTable(migrator gorm.Migrator, version string, table schema.Tabler) MigrationInterface {
	return &CreateTable{
		migrator: migrator,
		version:  version,
		table:    table,
	}
}

// Version implements the Migration.
func (s *CreateTable) Version() string {
	return s.version
}

// Identifier implements the Migration
func (s *CreateTable) Identifier() string {
	return "create_table_" + s.table.TableName()
}

// Up implements the Migration
func (s *CreateTable) Up() error {
	if s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.CreateTable(s.table)
}

// Down implements the Migration
func (s *CreateTable) Down() error {
	if !s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.DropTable(s.table)
}

// ===== 删除表 =====

// DropTable 删除表
type DropTable struct {
	migrator gorm.Migrator

	version string
	table   schema.Tabler
}

// NewDropTable 删除表
func NewDropTable(migrator gorm.Migrator, version string, table schema.Tabler) MigrationInterface {
	return &DropTable{
		version:  version,
		migrator: migrator,
		table:    table,
	}
}

// Version implements the Migration.
func (s *DropTable) Version() string {
	return s.version
}

// Identifier implements the Migration
func (s *DropTable) Identifier() string {
	return "drop_table_" + s.table.TableName()
}

// Up implements the Migration
func (s *DropTable) Up() error {
	if !s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.DropTable(s.table)
}

// Down implements the Migration
func (s *DropTable) Down() error {
	if s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.CreateTable(s.table)
}
