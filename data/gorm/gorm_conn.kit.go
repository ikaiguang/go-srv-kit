package gormutil

import (
	"gorm.io/gorm"
)

// NewDB creates a new DB instance
func NewDB(dialect gorm.Dialector, connOption *ConnOption) (db *gorm.DB, err error) {
	// 日志
	loggerHandler := NewLoggerForConn(connOption)

	// 选项
	gormConfig := &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   loggerHandler,
	}

	// 数据库链接
	db, err = gorm.Open(dialect, gormConfig)
	if err != nil {
		return
	}

	// 连接池
	connPool, err := db.DB()
	if err != nil {
		return
	}
	//  连接的最大数量
	if connOption.ConnMaxActive > 0 {
		connPool.SetMaxOpenConns(connOption.ConnMaxActive)
	}
	// 连接可复用的最大时间
	if connOption.ConnMaxLifetime > 0 {
		connPool.SetConnMaxLifetime(connOption.ConnMaxLifetime)
	}
	// 连接池中空闲连接的最大数量
	if connOption.ConnMaxIdle > 0 {
		connPool.SetMaxIdleConns(connOption.ConnMaxIdle)
	}
	// 设置连接空闲的最长时间
	if connOption.ConnMaxIdleTime > 0 {
		connPool.SetConnMaxIdleTime(connOption.ConnMaxIdleTime)
	}
	return db, err
}
