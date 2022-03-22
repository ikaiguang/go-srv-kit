package setup

import (
	stdlog "log"
	"sync"

	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	mysqlutil "github.com/ikaiguang/go-srv-kit/mysql"
)

// LoggerFileWriter 文件日志写手柄
func (s *modules) MysqlGormDB() (*gorm.DB, error) {
	var err error
	s.mysqlGormMutex.Do(func() {
		s.mysqlGormDB, err = s.loadingMysqlGormDB()
	})
	if err != nil {
		s.mysqlGormMutex = sync.Once{}
		return nil, err
	}
	if s.mysqlGormDB != nil {
		return s.mysqlGormDB, err
	}

	s.mysqlGormDB, err = s.loadingMysqlGormDB()
	if err != nil {
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
		opts    []mysqlutil.Option
	)
	if s.Config.EnableLoggingConsole() {
		writers = append(writers, mysqlutil.NewStdWriter())
	}
	if s.Config.EnableLoggingFile() {
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, mysqlutil.NewJSONWriter(writer))
	}
	if len(writers) > 0 {
		opts = append(opts, mysqlutil.WithWriters(writers...))
	}
	return mysqlutil.NewMysqlDB(s.Config.MySQLConfig(), opts...)
}
