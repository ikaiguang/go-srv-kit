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
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Config rabbitmq config
type Config struct {
	Url        string
	TlsAddress string
	TlsCaPem   string
	TlsCertPem string
	TlsKeyPem  string
}

// NewConnection 链接
// 已默认支持重连机制： amqp.DefaultReconnectConfig
func NewConnection(conf *Config, opts ...Option) (*amqp.ConnectionWrapper, error) {
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)
	return amqp.NewConnection(amqpConfig.Connection, op.logger)
}

// NewSubscriber 订阅者
// 注意：Close 同步调用了 conn.Close
func NewSubscriber(conf *Config, opts ...Option) (*amqp.Subscriber, error) {
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)

	return amqp.NewSubscriber(
		amqpConfig,
		op.logger,
	)
}

// NewPublisher 发布者
// 注意：Close 同步调用了 conn.Close
func NewPublisher(conf *Config, opts ...Option) (*amqp.Publisher, error) {
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newQueueConfig(conf, op)
	)

	return amqp.NewPublisher(
		amqpConfig,
		op.logger,
	)
}

// NewPubSub 发布订阅
// 注意：Close 同步调用了 conn.Close
func NewPubSub(conf *Config, opts ...Option) (publisher message.Publisher, subscriber message.Subscriber, err error) {
	// 配置
	var (
		op         = newOptions(opts...)
		amqpConfig = newPubSubConfig(conf, op)
	)
	publisher, err = amqp.NewPublisher(
		amqpConfig,
		op.logger,
	)
	if err != nil {
		return publisher, subscriber, err
	}

	subscriber, err = amqp.NewSubscriber(
		amqpConfig,
		op.logger,
	)
	if err != nil {
		return publisher, subscriber, err
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
		opts[i](&op)
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

// newPubSubConfig ...
func newPubSubConfig(conf *Config, op *options) amqp.Config {
	// 配置
	var amqpConfig amqp.Config
	if op.isNonDurable {
		amqpConfig = amqp.NewNonDurablePubSubConfig(conf.Url, amqp.GenerateQueueNameTopicName)
	} else {
		amqpConfig = amqp.NewDurablePubSubConfig(conf.Url, amqp.GenerateQueueNameTopicName)
	}
	if op.tlsConfig != nil {
		amqpConfig.Connection.TLSConfig = op.tlsConfig
	}
	return amqpConfig
}
