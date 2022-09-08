package setuppkg

import (
	stdlog "log"
	"sync"

	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	psqlutil "github.com/ikaiguang/go-srv-kit/data/postgres"

	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetPostgresGormDB 数据库
func (s *engines) GetPostgresGormDB() (*gorm.DB, error) {
	var err error
	s.postgresGormMutex.Do(func() {
		s.postgresGormDB, err = s.loadingPostgresGormDB()
	})
	if err != nil {
		s.postgresGormMutex = sync.Once{}
	}
	return s.postgresGormDB, err
}

// reloadPostgresGormDB 重新加载 postgres gorm 数据库
func (s *engines) reloadPostgresGormDB() error {
	if s.Config.PostgresConfig() == nil {
		return nil
	}
	dbConn, err := s.loadingPostgresGormDB()
	if err != nil {
		return err
	}
	*s.postgresGormDB = *dbConn
	return nil
}

// loadingPostgresGormDB postgres gorm 数据库
func (s *engines) loadingPostgresGormDB() (*gorm.DB, error) {
	if s.Config.PostgresConfig() == nil {
		stdlog.Println("|*** 加载：Postgres-GORM：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载：Postgres-GORM：...")

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
	return psqlutil.NewDB(s.Config.PostgresConfig(), opts...)
}
