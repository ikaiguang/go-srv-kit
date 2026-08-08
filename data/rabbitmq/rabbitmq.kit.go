package rabbitmqpkg

// 仅适用于简单例子使用，高级使用请配置后再实例化
// 仅做例子参考，实例化 amqp.NewPublisher
// 仅做例子参考，实例化 amqp.NewPublisherWithConnection
// 仅做例子参考，实例化 amqp.NewSubscriber
// 仅做例子参考，实例化 amqp.NewSubscriberWithConnection
// 仅做例子参考，按需配置 amqp.Config
// 仅做例子参考，按需配置 amqp.ConnectionConfig
// 仅做例子参考，按需配置 amqp.PublishConfig
// 仅做例子参考，按需配置 amqp.ConsumeConfig
// 仅做例子参考，按需配置 amqp.QueueConfig
import (
	"errors"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

type ConnectionWrapper struct {
	*amqp.ConnectionWrapper
	config *Config
}

// NewConnection 链接
// 已默认支持重连机制： amqp.DefaultReconnectConfig
func NewConnection(conf *Config, opts ...Option) (*ConnectionWrapper, error) {
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)
	conn, err := amqp.NewConnection(amqpConfig.Connection, op.logger)
	if err != nil {
		return nil, err
	}
	return &ConnectionWrapper{
		ConnectionWrapper: conn,
		config:            conf,
	}, nil
}

// NewSubscriberWithConnection 发布者
// 注意：Close 同步调用了 conn.Close
func NewSubscriberWithConnection(conn *ConnectionWrapper, opts ...Option) (*amqp.Subscriber, error) {
	if err := validateConnection(conn); err != nil {
		return nil, err
	}
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conn.config, op)
	)
	return amqp.NewSubscriberWithConnection(amqpConfig, op.logger, conn.ConnectionWrapper)
}

// NewPublisherWithConnection 发布者
// 注意：Close 同步调用了 conn.Close
func NewPublisherWithConnection(conn *ConnectionWrapper, opts ...Option) (*amqp.Publisher, error) {
	if err := validateConnection(conn); err != nil {
		return nil, err
	}
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conn.config, op)
	)
	return amqp.NewPublisherWithConnection(amqpConfig, op.logger, conn.ConnectionWrapper)
}

// NewPublisherAndSubscriberWithConnection 发布订阅
// 注意：Close 同步调用了 conn.Close
func NewPublisherAndSubscriberWithConnection(conn *ConnectionWrapper, opts ...Option) (publisher message.Publisher, subscriber message.Subscriber, err error) {
	if err = validateConnection(conn); err != nil {
		return nil, nil, err
	}
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conn.config, op)
	)
	publisher, err = amqp.NewPublisherWithConnection(amqpConfig, op.logger, conn.ConnectionWrapper)
	if err != nil {
		return publisher, subscriber, err
	}
	subscriber, err = amqp.NewSubscriberWithConnection(amqpConfig, op.logger, conn.ConnectionWrapper)
	if err != nil {
		if closeErr := publisher.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close publisher after subscriber failure: %w", closeErr))
		}
		return nil, nil, err
	}
	return publisher, subscriber, err
}

// NewSubscriber 订阅者
// 注意：Close 同步调用了 conn.Close
func NewSubscriber(conf *Config, opts ...Option) (*amqp.Subscriber, error) {
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)
	return amqp.NewSubscriber(amqpConfig, op.logger)
}

// NewPublisher 发布者
// 注意：Close 同步调用了 conn.Close
func NewPublisher(conf *Config, opts ...Option) (*amqp.Publisher, error) {
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)
	return amqp.NewPublisher(amqpConfig, op.logger)
}

// NewPublisherAndSubscriber 发布订阅
// 注意：Close 同步调用了 conn.Close
func NewPublisherAndSubscriber(conf *Config, opts ...Option) (publisher message.Publisher, subscriber message.Subscriber, err error) {
	if err = validateConfig(conf); err != nil {
		return nil, nil, err
	}
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)
	publisher, err = amqp.NewPublisher(amqpConfig, op.logger)
	if err != nil {
		return publisher, subscriber, err
	}
	subscriber, err = amqp.NewSubscriber(amqpConfig, op.logger)
	if err != nil {
		if closeErr := publisher.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close publisher after subscriber failure: %w", closeErr))
		}
		return nil, nil, err
	}
	return publisher, subscriber, err
}

// newOptions ...
func newOptions(opts ...Option) *options {
	op := options{
		isNonDurable: false,
		logger:       &watermill.NopLogger{},
	}
	for i := range opts {
		if opts[i] != nil {
			opts[i](&op)
		}
	}
	if op.logger == nil {
		op.logger = &watermill.NopLogger{}
	}
	return &op
}

// newQueueConfig ...
func newQueueConfig(conf *Config, op *options) amqp.Config {
	// 配置
	var amqpConfig amqp.Config
	if op.isNonDurable {
		amqpConfig = amqp.NewNonDurableQueueConfig(conf.Url)
	} else {
		amqpConfig = amqp.NewDurableQueueConfig(conf.Url)
	}
	if op.tlsConfig != nil {
		amqpConfig.Connection.TLSConfig = op.tlsConfig
	}
	return amqpConfig
}

func validateConfig(conf *Config) error {
	if conf == nil {
		return errors.New("rabbitmq config is nil")
	}
	if conf.Url == "" {
		return errors.New("rabbitmq url is empty")
	}
	return nil
}

func validateConnection(conn *ConnectionWrapper) error {
	if conn == nil || conn.ConnectionWrapper == nil {
		return errors.New("rabbitmq connection is nil")
	}
	return validateConfig(conn.config)
}
