package mongoutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	singletonMutex        sync.Once
	singletonMongoManager MongoManager
)

func NewSingletonMongoManager(conf *configpb.Mongo, loggerManager loggerutil.LoggerManager) (MongoManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonMongoManager, err = NewMongoManager(conf, loggerManager)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonMongoManager, err
}

func GetMongoClient(mongoManager MongoManager) (*mongo.Client, error) {
	return mongoManager.GetMongoClient()
}
