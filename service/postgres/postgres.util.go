package postgresutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	psqlpkg "github.com/ikaiguang/go-srv-kit/data/postgres"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlog "log"
	"sync"
)

type postgresManager struct {
	conf          *configpb.PSQL
	loggerManager loggerutil.LoggerManager

	postgresOnce   sync.Once
	postgresClient *gorm.DB
}

type PostgresManager interface {
	Enable() bool
	GetDB() (*gorm.DB, error)
	Close() error
}

func NewPostgresManager(conf *configpb.PSQL, loggerManager loggerutil.LoggerManager) (PostgresManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = psql")
		return nil, errorpkg.WithStack(e)
	}
	return &postgresManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *postgresManager) GetDB() (*gorm.DB, error) {
	var err error
	s.postgresOnce.Do(func() {
		s.postgresClient, err = s.loadingPostgresDB()
	})
	if err != nil {
		s.postgresOnce = sync.Once{}
	}
	return s.postgresClient, err
}

func (s *postgresManager) Close() error {
	if s.postgresClient != nil {
		stdlog.Println("|*** STOP: close: postgresClient")
		db, err := s.postgresClient.DB()
		if err != nil {
			stdlog.Println("|*** STOP: close: postgresClient failed: ", err.Error())
			return err
		}
		err = db.Close()
		if err != nil {
			stdlog.Println("|*** STOP: close: postgresClient failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *postgresManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *postgresManager) loadingPostgresDB() (*gorm.DB, error) {
	stdlog.Println("|*** LOADING: PostgresDB: ...")
	// logger
	var (
		writers = make([]logger.Writer, 0, 2)
	)
	if s.loggerManager.EnableConsole() {
		writers = append(writers, gormpkg.NewStdWriter())
	}
	if s.loggerManager.EnableFile() {
		writer, err := s.loggerManager.GetWriterForGORM()
		if err != nil {
			return nil, err
		}
		writers = append(writers, gormpkg.NewJSONWriter(writer))
	}

	var opts = make([]gormpkg.Option, 0, 1)
	if len(writers) > 0 {
		opts = append(opts, gormpkg.WithWriters(writers...))
	}
	return psqlpkg.NewDB(ToPSQLConfig(s.conf), opts...)
}

// ToPSQLConfig ...
func ToPSQLConfig(cfg *configpb.PSQL) *psqlpkg.Config {
	return &psqlpkg.Config{
		Dsn:             cfg.Dsn,
		SlowThreshold:   cfg.SlowThreshold,
		LoggerEnable:    cfg.LoggerEnable,
		LoggerColorful:  cfg.LoggerColorful,
		LoggerLevel:     cfg.LoggerLevel,
		ConnMaxActive:   cfg.ConnMaxActive,
		ConnMaxLifetime: cfg.ConnMaxLifetime,
		ConnMaxIdle:     cfg.ConnMaxIdle,
		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
	}
}
