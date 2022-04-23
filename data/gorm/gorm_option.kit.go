package gormutil

import (
	"gorm.io/gorm"
)

// SetOption set option
func SetOption(db *gorm.DB, key, value string) *gorm.DB {
	return db.Set(key, value)
}

// SetTableEngine 设置表引擎
func SetTableEngine(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
}
