package migrationuitl

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	DefaultMigrationTableName = "srv_migration"
)

// MigrationRepo 数据库迁移
type MigrationRepo interface {
	// Identifier 迁移标识
	Identifier() string
	// Up 运行迁移
	Up() error
	// Down 回滚迁移
	Down() error
}

// ===== 创建表 =====

// CreateTable 创建表
type CreateTable struct {
	migrator gorm.Migrator

	table schema.Tabler
}

// NewCreateTable 创建表
func NewCreateTable(migrator gorm.Migrator, table schema.Tabler) *CreateTable {
	return &CreateTable{
		migrator: migrator,
		table:    table,
	}
}

// Identifier implements the Migration interface.
func (s *CreateTable) Identifier() string {
	return "create_table_" + s.table.TableName()
}

// Up implements the Migration interface.
func (s *CreateTable) Up() error {
	if s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.CreateTable(s.table)
}

// Down implements the Migration interface.
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

	table schema.Tabler
}

// NewDropTable 删除表
func NewDropTable(migrator gorm.Migrator, table schema.Tabler) *DropTable {
	return &DropTable{
		migrator: migrator,
		table:    table,
	}
}

// Identifier implements the Migration interface.
func (s *DropTable) Identifier() string {
	return "drop_table_" + s.table.TableName()
}

// Up implements the Migration interface.
func (s *DropTable) Up() error {
	if !s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.DropTable(s.table)
}

// Down implements the Migration interface.
func (s *DropTable) Down() error {
	if s.migrator.HasTable(s.table) {
		return nil
	}
	return s.migrator.CreateTable(s.table)
}
