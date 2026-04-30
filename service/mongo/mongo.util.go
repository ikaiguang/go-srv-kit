package mongoutil

import (
	"context"
	stdlog "log"
	"sync"

	"github.com/go-kratos/kratos/v2/log"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	mongopkg "github.com/ikaiguang/go-srv-kit/data/mongo"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	loggerutil "github.com/ikaiguang/go-srv-kit/service/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoManager struct {
	conf          *configpb.Mongo
	loggerManager loggerutil.LoggerManager

	mongoOnce   sync.Once
	mongoClient *mongo.Client
}

type MongoManager interface {
	Enable() bool
	GetMongoClient() (*mongo.Client, error)
	Close() error
}

func NewMongoManager(conf *configpb.Mongo, loggerManager loggerutil.LoggerManager) (MongoManager, error) {
	if conf == nil {
		e := errorpkg.ErrorBadRequest("[CONFIGURATION] config error, key = mongo")
		return nil, errorpkg.WithStack(e)
	}
	return &mongoManager{
		conf:          conf,
		loggerManager: loggerManager,
	}, nil
}

func (s *mongoManager) GetMongoClient() (*mongo.Client, error) {
	var err error
	s.mongoOnce.Do(func() {
		s.mongoClient, err = s.loadingMongoDB()
	})
	if err != nil {
		s.mongoOnce = sync.Once{}
	}
	return s.mongoClient, err
}

func (s *mongoManager) Close() error {
	if s.mongoClient != nil {
		stdlog.Println("|*** STOP: close: mongoClient")
		err := s.mongoClient.Disconnect(context.Background())
		if err != nil {
			stdlog.Println("|*** STOP: close: mongoClient failed: ", err.Error())
			return err
		}
	}
	return nil
}

func (s *mongoManager) Enable() bool {
	return s.conf.GetEnable()
}

func (s *mongoManager) loadingMongoDB() (*mongo.Client, error) {
	stdlog.Println("|*** LOADING: MongoDB: ...")
	// logger
	logger, err := s.loggerManager.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := mongopkg.NewMongoClient(ToMongoConfig(s.conf), toMongoLogger(logger))
	if err != nil {
		e := errorpkg.ErrorInternalError("%s", err.Error())
		return nil, errorpkg.WithStack(e)
	}
	return db, nil
}

// toMongoLogger 将 kratos log.Logger 适配为 mongopkg.Logger
func toMongoLogger(l log.Logger) mongopkg.Logger {
	return mongopkg.LogAdapter(func(level mongopkg.Level, keyvals ...any) error {
		return l.Log(log.Level(level), keyvals...)
	})
}

// ToMongoConfig ...
func ToMongoConfig(cfg *configpb.Mongo) *mongopkg.Config {
	return &mongopkg.Config{
		Debug:             cfg.GetDebug(),
		AppName:           cfg.GetAppName(),
		Hosts:             cfg.GetHosts(),
		Addr:              cfg.GetAddr(),
		MaxPoolSize:       cfg.GetMaxPoolSize(),
		MinPoolSize:       cfg.GetMinPoolSize(),
		MaxConnecting:     cfg.GetMaxConnecting(),
		ConnectTimeout:    cfg.GetConnectTimeout(),
		Timeout:           cfg.GetTimeout(),
		HeartbeatInterval: cfg.GetHeartbeatInterval(),
		MaxConnIdleTime:   cfg.GetMaxConnIdleTime(),
		SlowThreshold:     cfg.GetSlowThreshold(),
	}
}
