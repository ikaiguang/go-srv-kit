package gormpkg

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ForUpdate FOR UPDATE
func ForUpdate(db *gorm.DB) *gorm.DB {
	return db.Clauses(clause.Locking{Strength: "UPDATE"})
}

// ForUpdateNowait FOR UPDATE NOWAIT
func ForUpdateNowait(db *gorm.DB) *gorm.DB {
	return db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  "NOWAIT",
	})
}

// ForShareOfTable FOR SHARE OF `table_name`
func ForShareOfTable(db *gorm.DB) *gorm.DB {
	return db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	})
}
