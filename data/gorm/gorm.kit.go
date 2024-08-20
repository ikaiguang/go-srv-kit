package gormpkg

import (
	"database/sql"
	"gorm.io/gorm"
	"regexp"
)

var (
	// regColumn 正则表达式:列
	regColumn = regexp.MustCompile("^[A-Za-z-_]+$")
)

// IsValidField 判断是否为有效的字段名
func IsValidField(field string) bool {
	return regColumn.MatchString(field)
}

// Transaction 在事务中执行一系列操作; 无需手动开启事务
// DOCS: https://gorm.io/zh_CN/docs/transactions.html
func Transaction(dbConn *gorm.DB, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return dbConn.Transaction(fc, opts...)
}

type TransactionInstance interface {
	Rollback() error
	Commit() error
}

func NewTransaction(db *gorm.DB, opts ...*sql.TxOptions) TransactionInstance {
	tx := db.Begin(opts...)

	return &transaction{tx: tx}
}

type transaction struct {
	tx *gorm.DB
}

func (s *transaction) Rollback() error {
	return s.tx.Rollback().Error
}

func (s *transaction) Commit() error {
	return s.tx.Commit().Error
}
