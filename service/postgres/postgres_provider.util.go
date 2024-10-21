package postgresutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"gorm.io/gorm"
	"sync"
)

var (
	singletonMutex           sync.Once
	singletonPostgresManager PostgresManager
)

func NewSingletonPostgresManager(conf *configpb.PSQL, loggerManager loggerutil.LoggerManager) (PostgresManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonPostgresManager, err = NewPostgresManager(conf, loggerManager)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonPostgresManager, err
}

func GetDBConn(postgresManager PostgresManager) (*gorm.DB, error) {
	return postgresManager.GetDB()
}
