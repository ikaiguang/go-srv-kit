package setup

import (
	stdlog "log"
	"sync"

	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlutil "github.com/ikaiguang/go-srv-kit/data/mysql"

	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetMySQLGormDB 数据库
func (s *engines) GetMySQLGormDB() (*gorm.DB, error) {
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
func (s *engines) loadingMysqlGormDB() (*gorm.DB, error) {
	if s.Config.MySQLConfig() == nil {
		stdlog.Println("|*** 加载MySQL-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载MySQL-GORM：...")

	// logger writer
	var (
		writers []logger.Writer
		opts    []gormutil.Option
	)
	if s.Config.EnableLoggingConsole() {
		writers = append(writers, gormutil.NewStdWriter())
	}
	if s.Config.EnableLoggingFile() {
		writer, err := s.LoggerFileWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, gormutil.NewJSONWriter(writer))
	}
	if len(writers) > 0 {
		opts = append(opts, gormutil.WithWriters(writers...))
	}
	return mysqlutil.NewMysqlDB(s.Config.MySQLConfig(), opts...)
}
