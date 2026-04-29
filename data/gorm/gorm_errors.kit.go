package gormpkg

import (
	stderrors "errors"

	"gorm.io/gorm"
)

// IsErrRecordNotFound ...
func IsErrRecordNotFound(err error) bool {
	return stderrors.Is(err, gorm.ErrRecordNotFound)
}

// IsErrDuplicatedKey 检查是否为唯一键冲突错误（仅检查 GORM 层面）
// 如需检查特定数据库驱动的错误码，请使用：
//   - mysqlpkg.IsErrDuplicatedKey（MySQL 1062）
//   - psqlpkg.IsErrDuplicatedKey（PostgreSQL 23505）
func IsErrDuplicatedKey(err error) bool {
	if err == nil {
		return false
	}
	return stderrors.Is(err, gorm.ErrDuplicatedKey)
}
