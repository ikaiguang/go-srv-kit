package migrationuitl

import (
	"time"
)

// Migration 数据库迁移
// 文档地址：https://gorm.io/docs/models.html
// MySQL 支持 unsigned
// Postgres 不支持 unsigned
// MySQL 支持 auto_increment
// Postgres : serial or bigserial
type Migration struct {
	Id                 uint64    `gorm:"COLUMN:id;primaryKey;type:uint;autoIncrement;comment:ID"`
	MigrationKey       string    `gorm:"COLUMN:migration_key;uniqueIndex;type:string;size:255;not null;default:'';comment:迁移key：唯一"`
	MigrationBatch     uint      `gorm:"COLUMN:migration_batch;type:uint;not null;default:0;comment:迁移批次"`
	MigrationDesc      string    `gorm:"COLUMN:migration_desc;type:text;not null;comment:迁移描述"`
	MigrationExtraInfo string    `gorm:"COLUMN:migration_extra_info;type:json;not null;comment:迁移：额外信息"`
	CreatedTime        time.Time `gorm:"COLUMN:created_time;type:time;not null;autoCreateTime:milli;comment:创建时间"`
	UpdatedTime        time.Time `gorm:"COLUMN:updated_time;type:time;not null;autoUpdateTime:milli;comment:更新时间"`
}

// TableName 表名
func (s *Migration) TableName() string {
	return DefaultMigrationTableName
}

// TestMigration test
// 文档地址：https://gorm.io/docs/models.html
// MySQL 支持 unsigned
// Postgres 不支持 unsigned
// MySQL 支持 auto_increment
// Postgres : serial or bigserial
type TestMigration struct {
	Id                uint64     `gorm:"COLUMN:id;primaryKey;type:uint;autoIncrement;comment:ID"`
	ColumnUniqueIndex string     `gorm:"COLUMN:column_unique_index;uniqueIndex;type:string;size:255;not null;default:'';comment:唯一索引"`
	ColumnIndex       string     `gorm:"COLUMN:column_index;index;type:string;size:255;not null;default:0;comment:普通索引"`
	ColumnBool        bool       `gorm:"COLUMN:column_bool;type:bool;not null;default:0;comment:布尔值"`
	ColumnInt         int32      `gorm:"COLUMN:column_int;type:int;not null;default:0;comment:整型"`
	ColumnUint        uint32     `gorm:"COLUMN:column_uint;type:uint;not null;default:0;comment:整型：无符号"`
	ColumnInt64       int64      `gorm:"COLUMN:column_int64;type:int;not null;default:0;comment:整型：64位"`
	ColumnUint64      uint64     `gorm:"COLUMN:column_uint64;type:uint;not null;default:0;comment:整型：无符号64位"`
	ColumnFloat64     float64    `gorm:"COLUMN:column_float64;type:decimal(30,10);not null;default:0;comment:浮点型：64位"`
	ColumnString      string     `gorm:"COLUMN:column_string;type:string;size:255;not null;default:'';comment:字符串"`
	ColumnText        string     `gorm:"COLUMN:column_text;type:text;not null;comment:文本"`
	ColumnJSON        string     `gorm:"COLUMN:column_json;type:json;not null;comment:JSON"`
	ColumnBytes       string     `gorm:"COLUMN:column_bytes;type:bytes;not null;comment:字节"`
	CreatedTime       time.Time  `gorm:"COLUMN:column_created_time;type:time;not null;autoCreateTime:milli;comment:创建时间"`
	ColumnUpdatedTime time.Time  `gorm:"COLUMN:column_updated_time;type:time;not null;autoUpdateTime:milli;comment:更新时间"`
	ColumnDeletedTime *time.Time `gorm:"COLUMN:column_deleted_time;type:time;default:null;comment:删除时间"`

	// IgnoreReadWrite ignore this field when write and read with struct
	IgnoreReadWrite string `gorm:"-"`
	// IgnoreAll ignore this field when write/read and migrate with struct
	IgnoreAll string `gorm:"-:all"`
	// IgnoreMigration ignore this field when migrate with struct
	IgnoreMigration string `gorm:"-:migration"`
}

// TableName 表名
func (s *TestMigration) TableName() string {
	return "srv_test_migration"
}
