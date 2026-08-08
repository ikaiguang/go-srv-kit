package mongopkg

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const defaultProbeTimeout = 5 * time.Second

// NewMongoClient ...
func NewMongoClient(config *Config, logger *slog.Logger) (*mongo.Client, error) {
	if config == nil {
		return nil, errors.New("mongo config is nil")
	}

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
	clientOpt.SetMonitor(NewMonitor(
		logger,
		WithSlowThreshold(config.SlowThreshold.AsDuration()),
		WithCommandLogging(config.Debug),
	))

	client, err := mongo.Connect(clientOpt)
	if err != nil {
		err = fmt.Errorf("mongo connect failed: %w", err)
		return nil, err
	}
	probeTimeout := config.ConnectTimeout.AsDuration()
	if probeTimeout <= 0 {
		probeTimeout = defaultProbeTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), probeTimeout)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		disconnectCtx, disconnectCancel := context.WithTimeout(context.Background(), defaultProbeTimeout)
		defer disconnectCancel()
		if disconnectErr := client.Disconnect(disconnectCtx); disconnectErr != nil {
			return nil, fmt.Errorf("mongo ping failed: %w; disconnect client: %v", err, disconnectErr)
		}
		return nil, fmt.Errorf("mongo ping failed: %w", err)
	}
	return client, nil
}
