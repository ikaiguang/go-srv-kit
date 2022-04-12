package setup

import (
	stdlog "log"
	"sync"

	mysqlutil2 "github.com/ikaiguang/go-srv-kit/data/mysql"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MysqlGormDB 数据库
func (s *modules) MysqlGormDB() (*gorm.DB, error) {
	var err error
	s.mysqlGormMutex.Do(func() {
		s.mysqlGormDB, err = s.loadingMysqlGormDB()
	})
	if err != nil {
		s.mysqlGormMutex = sync.Once{}
		return nil, err
	}
	return s.mysqlGormDB, err
}

// loadingMysqlGormDB mysql gorm 数据库
func (s *modules) loadingMysqlGormDB() (*gorm.DB, error) {
	if s.Config.MySQLConfig() == nil {
		stdlog.Println("|*** 加载MySQL-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载MySQL-GORM：...")

	// logger writer
	var (
		writers []logger.Writer
		opts    []mysqlutil2.Option
	)
	if s.Config.EnableLoggingConsole() {
		writers = append(writers, mysqlutil2.NewStdWriter())
	}
	if s.Config.EnableLoggingFile() {
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, mysqlutil2.NewJSONWriter(writer))
	}
	if len(writers) > 0 {
		opts = append(opts, mysqlutil2.WithWriters(writers...))
	}
	return mysqlutil2.NewMysqlDB(s.Config.MySQLConfig(), opts...)
}
