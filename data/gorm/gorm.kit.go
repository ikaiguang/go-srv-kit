package gormpkg

import (
	"context"
	"database/sql"
	stderrors "errors"
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
	Do(ctx context.Context, fc func(context.Context, *gorm.DB) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CommitAndErrRollback(ctx context.Context, resultErr error) (err error)
}

func NewTransaction(ctx context.Context, db *gorm.DB, opts ...*sql.TxOptions) TransactionInstance {
	tx := db.WithContext(ctx).Begin(opts...)

	return &transaction{tx: tx}
}

type transaction struct {
	tx *gorm.DB
}

func (s *transaction) Do(ctx context.Context, fc func(context.Context, *gorm.DB) error) error {
	return fc(ctx, s.tx)
}

func (s *transaction) Commit(ctx context.Context) error {
	return s.tx.WithContext(ctx).Commit().Error
}

func (s *transaction) Rollback(ctx context.Context) error {
	return s.tx.WithContext(ctx).Rollback().Error
}

func (s *transaction) CommitAndErrRollback(ctx context.Context, resultErr error) (err error) {
	defer func() {
		if err != nil {
			rollbackErr := s.Rollback(ctx)
			if rollbackErr != nil {
				err = stderrors.Join(err, rollbackErr)
			}
		}
	}()
	if resultErr != nil {
		err = resultErr
		return err
	}
	err = s.Commit(ctx)
	if err != nil {
		return err
	}
	return err
}
