package gormutil

import (
	"gorm.io/gorm"
)

const (
	OptionKeyTableOptions = "gorm:table_options"
	OptionKeyCreate       = "gorm:create"
	OptionKeyUpdate       = "gorm:update"
	OptionKeyQuery        = "gorm:query"
	OptionKeyDelete       = "gorm:delete"

	OptionValueEngineInnoDB = "ENGINE=InnoDB CHARSET=utf8mb4"
)

// SetOption set option
func SetOption(db *gorm.DB, key, value string) *gorm.DB {
	return db.Set(key, value)
}
