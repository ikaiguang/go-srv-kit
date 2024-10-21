package rabbitmqutil

import (
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	rabbitmqpkg "github.com/ikaiguang/go-srv-kit/data/rabbitmq"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	stdlog "log"
	"sync"
)

type rabbitmqManager struct {
	conf          *configpb.Rabbitmq
	loggerManager loggerutil.LoggerManager

	rabbitmqOnce sync.Once
	rabbitmqConn *amqp.ConnectionWrapper
}

type RabbitmqManager interface {
	Enable() bool
	GetClient() (*amqp.ConnectionWrapper, error)
	Close() error
}

func NewRabbitmqManager(conf *configpb.Rabbitmq, loggerManager loggerutil.LoggerManager) (RabbitmqManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = rabbitmq")
		return nil, errorpkg.WithStack(e)
	}
	return &rabbitmqManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *rabbitmqManager) GetClient() (*amqp.ConnectionWrapper, error) {
	var err error
	s.rabbitmqOnce.Do(func() {
		s.rabbitmqConn, err = s.loadingRabbitmqClient()
	})
	if err != nil {
		s.rabbitmqOnce = sync.Once{}
	}
	return s.rabbitmqConn, err
}

func (s *rabbitmqManager) Close() error {
	if s.rabbitmqConn != nil {
		stdlog.Println("|*** STOP: close: rabbitmqConn")
		err := s.rabbitmqConn.Close()
		if err != nil {
			stdlog.Println("|*** STOP: close: rabbitmqConn failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *rabbitmqManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *rabbitmqManager) loadingRabbitmqClient() (*amqp.ConnectionWrapper, error) {
	stdlog.Println("|*** LOADING: Rabbitmq connection: ...")
	logger, err := s.loggerManager.GetLogger()
	if err != nil {
		return nil, err
	}
	opts := make([]rabbitmqpkg.Option, 0)
	opts = append(opts, rabbitmqpkg.WithLogger(rabbitmqpkg.NewLogger(logger)))
	uc, err := rabbitmqpkg.NewConnection(ToRabbitmqConfig(s.conf), opts...)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return uc, nil
}

// ToRabbitmqConfig ...
func ToRabbitmqConfig(cfg *configpb.Rabbitmq) *rabbitmqpkg.Config {
	return &rabbitmqpkg.Config{
		Url:        cfg.Url,
		TlsAddress: cfg.TlsAddress,
		TlsCaPem:   cfg.TlsCaPem,
		TlsCertPem: cfg.TlsCertPem,
		TlsKeyPem:  cfg.TlsKeyPem,
	}
}
