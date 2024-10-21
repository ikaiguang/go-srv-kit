package mysqlutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlpkg "github.com/ikaiguang/go-srv-kit/data/mysql"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlog "log"
	"sync"
)

type mysqlManager struct {
	conf          *configpb.MySQL
	loggerManager loggerutil.LoggerManager

	mysqlOnce   sync.Once
	mysqlClient *gorm.DB
}

type MysqlManager interface {
	Enable() bool
	GetDB() (*gorm.DB, error)
	Close() error
}

func NewMysqlManager(conf *configpb.MySQL, loggerManager loggerutil.LoggerManager) (MysqlManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = mysql")
		return nil, errorpkg.WithStack(e)
	}
	return &mysqlManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *mysqlManager) GetDB() (*gorm.DB, error) {
	var err error
	s.mysqlOnce.Do(func() {
		s.mysqlClient, err = s.loadingMysqlDB()
	})
	if err != nil {
		s.mysqlOnce = sync.Once{}
	}
	return s.mysqlClient, err
}

func (s *mysqlManager) Close() error {
	if s.mysqlClient != nil {
		stdlog.Println("|*** STOP: close: mysqlClient")
		db, err := s.mysqlClient.DB()
		if err != nil {
			stdlog.Println("|*** STOP: close: mysqlClient failed: ", err.Error())
			return err
		}
		err = db.Close()
		if err != nil {
			stdlog.Println("|*** STOP: close: mysqlClient failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *mysqlManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *mysqlManager) loadingMysqlDB() (*gorm.DB, error) {
	stdlog.Println("|*** LOADING: MysqlDB: ...")
	// logger
	var (
		writers = make([]logger.Writer, 0, 2)
	)
	if s.loggerManager.EnableConsole() {
		writers = append(writers, gormpkg.NewStdWriter())
	}
	if s.loggerManager.EnableFile() {
		writer, err := s.loggerManager.GetWriter()
		if err != nil {
			return nil, err
		}
		writers = append(writers, gormpkg.NewJSONWriter(writer))
	}

	var opts = make([]gormpkg.Option, 0, 1)
	if len(writers) > 0 {
		opts = append(opts, gormpkg.WithWriters(writers...))
	}
	db, err := mysqlpkg.NewMysqlDB(ToMysqlConfig(s.conf), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return db, nil
}

// ToMysqlConfig ...
func ToMysqlConfig(cfg *configpb.MySQL) *mysqlpkg.Config {
	return &mysqlpkg.Config{
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
