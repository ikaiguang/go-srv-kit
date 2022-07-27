// Package migrationutil
// Package entities
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package migrationutil

import (
	"time"
)

var _ = time.Time{}

// MigrationEntity srv_migration
// Migration 数据库迁移
// 文档地址：https://gorm.io/docs/models.html
// MySQL 支持 unsigned
// Postgres 不支持 unsigned
// MySQL 支持 auto_increment
// Postgres : serial or bigserial
type MigrationEntity struct {
	Id                  uint64    `gorm:"column:id;primaryKey" json:"id"`                          // ID
	ServerVersion       string    `gorm:"column:server_version" json:"server_version"`             // 服务版本
	MigrationIdentifier string    `gorm:"column:migration_identifier" json:"migration_identifier"` // 迁移key：唯一
	MigrationBatch      uint64    `gorm:"column:migration_batch" json:"migration_batch"`           // 迁移批次
	MigrationDesc       string    `gorm:"column:migration_desc" json:"migration_desc"`             // 迁移描述
	MigrationExtraInfo  string    `gorm:"column:migration_extra_info" json:"migration_extra_info"` // 迁移：额外信息
	CreatedTime         time.Time `gorm:"column:created_time" json:"created_time"`                 // 创建时间
	UpdatedTime         time.Time `gorm:"column:updated_time" json:"updated_time"`                 // 更新时间
}

// TableName table name
func (s *MigrationEntity) TableName() string {
	return DefaultMigrationTableName
}
