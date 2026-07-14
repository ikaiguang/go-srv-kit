package rabbitmqutil

import (
	"sync"

	rabbitmqpkg "github.com/ikaiguang/go-rabbitmq-kit/rabbitmq"
	configpb "github.com/ikaiguang/go-service-kit/api/config"
	loggerutil "github.com/ikaiguang/go-service-kit/logger"
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
