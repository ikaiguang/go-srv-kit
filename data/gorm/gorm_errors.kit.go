package gormpkg

import (
	stderrors "errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// IsErrRecordNotFound ...
func IsErrRecordNotFound(err error) bool {
	return stderrors.Is(err, gorm.ErrRecordNotFound)
}

// IsErrDuplicatedKey ...
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	var pgErr *pgconn.PgError
	if stderrors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	var mysqlErr *mysql.MySQLError
	if stderrors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	var sqliteErr *sqlite3.Error
	if stderrors.As(err, &sqliteErr) {
		return sqliteErr.ExtendedCode == 2067
	}
	return false
}
