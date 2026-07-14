package postgresutil

import (
	"sync"

	configpb "github.com/ikaiguang/go-service-kit/api/config"
	loggerutil "github.com/ikaiguang/go-service-kit/logger"
	"gorm.io/gorm"
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
