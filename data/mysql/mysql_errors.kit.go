package mysqlpkg

import (
	stderrors "errors"

	"github.com/go-sql-driver/mysql"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
)

// IsErrDuplicatedKey 检查是否为唯一键冲突错误（MySQL）
// 同时检查 GORM 层面的 ErrDuplicatedKey 和 MySQL 驱动的 1062 错误码
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	if gormpkg.IsErrDuplicatedKey(err) {
		return true
	}
	var mysqlErr *mysql.MySQLError
	if stderrors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	return false
}
