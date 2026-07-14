package gormpkg

import (
	"gorm.io/gorm"
	"gorm.io/hints"
)

// UseIndex 使用索引
func UseIndex(db *gorm.DB, indexNameList ...string) *gorm.DB {
	return db.Clauses(hints.UseIndex(indexNameList...))
}

// ForceIndex 强制使用索引
func ForceIndex(db *gorm.DB, indexNameList ...string) *gorm.DB {
	return db.Clauses(hints.ForceIndex(indexNameList...))
}

// IgnoreIndex 忽略使用索引
func IgnoreIndex(db *gorm.DB, indexNameList ...string) *gorm.DB {
	return db.Clauses(hints.IgnoreIndex(indexNameList...))
}
