package rabbitmqutil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
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

func GetRabbitmqConn(rabbitmqManager RabbitmqManager) (*amqp.ConnectionWrapper, error) {
	return rabbitmqManager.GetClient()
}
