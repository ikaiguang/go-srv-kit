package mysqlutil

import (
	"sync"

	configpb "github.com/ikaiguang/go-service-kit/api/config"
	loggerutil "github.com/ikaiguang/go-service-kit/logger"
	"gorm.io/gorm"
)

var (
	singletonMutex        sync.Once
	singletonMysqlManager MysqlManager
)

func NewSingletonMysqlManager(conf *configpb.MySQL, loggerManager loggerutil.LoggerManager) (MysqlManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonMysqlManager, err = NewMysqlManager(conf, loggerManager)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonMysqlManager, err
}

func GetDBConn(mysqlManager MysqlManager) (*gorm.DB, error) {
	return mysqlManager.GetDB()
}
