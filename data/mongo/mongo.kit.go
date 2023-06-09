package mongopkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Config ...
type Config struct {
	Hosts             []string
	Addr              string
	AppName           string
	MaxPoolSize       uint64
	MinPoolSize       uint64
	MaxConnecting     uint64
	ConnectTimeout    time.Duration
	HeartbeatInterval time.Duration
	MaxConnIdleTime   time.Duration
	Timeout           time.Duration
	Debug             bool
}

// NewMongoClient ...
func NewMongoClient(config *Config, logger log.Logger) (*mongo.Client, error) {
	clientOpt := options.Client()
	clientOpt.SetHosts(config.Hosts)
	if config.Addr != "" {
		clientOpt.ApplyURI(config.Addr)
	}
	if config.AppName != "" {
		clientOpt.SetAppName(config.AppName)
	}
	if config.ConnectTimeout > 0 {
		clientOpt.SetConnectTimeout(config.ConnectTimeout)
	}
	if config.HeartbeatInterval > 0 {
		clientOpt.SetHeartbeatInterval(config.HeartbeatInterval)
	}
	if config.MaxConnIdleTime > 0 {
		clientOpt.SetMaxConnIdleTime(config.MaxConnIdleTime)
	}
	if config.Timeout > 0 {
		clientOpt.SetTimeout(config.Timeout)
	}
	if config.MaxPoolSize > 0 {
		clientOpt.SetMaxPoolSize(config.MaxPoolSize)
	}
	if config.MinPoolSize > 0 {
		clientOpt.SetMinPoolSize(config.MinPoolSize)
	}
	if config.MaxConnecting > 0 {
		clientOpt.SetMaxConnecting(config.MaxConnecting)
	}
	if config.Debug {
		clientOpt.SetMonitor(NewMonitor(logger))
	}

	client, err := mongo.Connect(context.Background(), clientOpt)
	if err != nil {
		err = fmt.Errorf("mongo connect failed: %w", err)
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		err = fmt.Errorf("mongo ping failed: %w", err)
		return nil, err
	}
	return client, err
}
