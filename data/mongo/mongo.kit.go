package mongopkg

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Config ...
type Config struct {
	Debug             bool
	AppName           string
	Hosts             []string
	Addr              string
	MaxPoolSize       uint32
	MinPoolSize       uint32
	MaxConnecting     uint32
	ConnectTimeout    *durationpb.Duration
	Timeout           *durationpb.Duration
	HeartbeatInterval *durationpb.Duration
	MaxConnIdleTime   *durationpb.Duration
	SlowThreshold     *durationpb.Duration
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
	if config.ConnectTimeout.AsDuration() > 0 {
		clientOpt.SetConnectTimeout(config.ConnectTimeout.AsDuration())
	}
	if config.HeartbeatInterval.AsDuration() > 0 {
		clientOpt.SetHeartbeatInterval(config.HeartbeatInterval.AsDuration())
	}
	if config.MaxConnIdleTime.AsDuration() > 0 {
		clientOpt.SetMaxConnIdleTime(config.MaxConnIdleTime.AsDuration())
	}
	if config.Timeout.AsDuration() > 0 {
		clientOpt.SetTimeout(config.Timeout.AsDuration())
	}
	if config.MaxPoolSize > 0 {
		clientOpt.SetMaxPoolSize(uint64(config.MaxPoolSize))
	}
	if config.MinPoolSize > 0 {
		clientOpt.SetMinPoolSize(uint64(config.MinPoolSize))
	}
	if config.MaxConnecting > 0 {
		clientOpt.SetMaxConnecting(uint64(config.MaxConnecting))
	}

	// logger
	if config.Debug {
		clientOpt.SetMonitor(NewMonitor(logger, WithSlowThreshold(config.SlowThreshold.AsDuration())))
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
