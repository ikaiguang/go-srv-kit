package psqlpkg

import (
	stderrors "errors"

	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsErrDuplicatedKey 检查是否为唯一键冲突错误（PostgreSQL）
// 同时检查 GORM 层面的 ErrDuplicatedKey 和 PostgreSQL 驱动的 23505 错误码
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	if gormpkg.IsErrDuplicatedKey(err) {
		return true
	}
	var pgErr *pgconn.PgError
	if stderrors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
