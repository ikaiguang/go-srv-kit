package rabbitmqutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"sync"
)

var (
	singletonMutex           sync.Once
	singletonRabbitmqManager RabbitmqManager
)

func NewSingletonRabbitmqManager(conf *configpb.Rabbitmq, loggerManager loggerutil.LoggerManager) (RabbitmqManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonRabbitmqManager, err = NewRabbitmqManager(conf, loggerManager)
	})
	if err != nil {
		singletonMutex = sync.Once{}
	}
	return singletonRabbitmqManager, err
}

func GetRabbitmqConn(rabbitmqManager RabbitmqManager) (*rabbitmqpkg.ConnectionWrapper, error) {
	return rabbitmqManager.GetClient()
}
