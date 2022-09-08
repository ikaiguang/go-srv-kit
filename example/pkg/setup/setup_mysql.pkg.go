package setuppkg

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
	}
	return s.mysqlGormDB, err
}

// reloadMysqlGormDB 重新加载 mysql gorm 数据库
func (s *engines) reloadMysqlGormDB() error {
	if s.Config.MySQLConfig() == nil {
		return nil
	}
	dbConn, err := s.loadingMysqlGormDB()
	if err != nil {
		return err
	}
	*s.mysqlGormDB = *dbConn
	return nil
}

// loadingMysqlGormDB mysql gorm 数据库
func (s *engines) loadingMysqlGormDB() (*gorm.DB, error) {
	if s.Config.MySQLConfig() == nil {
		stdlog.Println("|*** 加载：MySQL-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载：MySQL-GORM：...")

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
