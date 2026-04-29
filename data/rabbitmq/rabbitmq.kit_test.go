//go:build ignore

package rabbitmqpkg

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/go-kratos/kratos/v2/log"
	threadpkg "github.com/ikaiguang/go-srv-kit/kit/thread"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"github.com/stretchr/testify/require"
)

func getTestRabbitMQURL() string {
	if url := os.Getenv("DB_RABBITMQ_URL"); url != "" {
		return url
	}
	return "amqp://rabbitmq:Rabbitmq.123456@my-rabbitmq:5672/"
}

// kratosToLogger 将 kratos log.Logger 适配为本包的 Logger 接口
func kratosToLogger(l log.Logger) Logger {
	return LogAdapter(func(level Level, keyvals ...any) error {
		return l.Log(log.Level(level), keyvals...)
	})
}

// go test -v ./data/rabbitmq/ -count=1 -run TestNewSubscriber
func TestNewSubscriber(t *testing.T) {
	var (
		//amqpURI = "amqp://guest:guest@127.0.0.1:5672/"
		amqpURI = getTestRabbitMQURL()
		topic   = "example.topic"
	)
	logger := newMultiLogger()

	conf := &Config{
		Url: amqpURI,
	}

	sub, err := NewSubscriber(conf, WithLogger(logger))
	require.NoError(t, err)
	pub, err := NewPublisher(conf, WithLogger(logger))
	require.NoError(t, err)

	defer func() { _ = sub.Close() }()
	defer func() { _ = pub.Close() }()

	threadpkg.GoSafe(func() {
		doSub(sub, topic)
	})
	threadpkg.GoSafe(func() {
		doPub(pub, topic)
	})

	time.Sleep(3 * time.Second)
}

func doSub(subscriber *amqp.Subscriber, topic string) {
	messages, err := subscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}

	fmt.Println("==> sub...")
	var i = 0
	for msg := range messages {
		i++
		fmt.Printf("==> subscriber.Subscribe : index : %d\n", i)
		fmt.Printf("==> subscriber.Subscribe : messages UUID : %v\n", msg.UUID)
		fmt.Printf("==> subscriber.Subscribe : messages Payload : %v\n", string(msg.Payload))

		msg.Ack()
	}
}

func doPub(publisher *amqp.Publisher, topic string) {
	fmt.Println("==> pub...")
	var i = 0
	for {
		i++
		msg := message.NewMessage(uuidpkg.NewUUID(), []byte(fmt.Sprintf("%d : Hello, world!", i)))

		if err := publisher.Publish(topic, msg); err != nil {
			//panic(err)
			fmt.Println("==> publisher.Publish error : ", err)
			fmt.Println("==> publisher.Publish error : ", err)
			fmt.Println("==> publisher.Publish error : ", err)
		}

		time.Sleep(time.Second)
	}
}

func newMultiLogger() watermill.LoggerAdapter {
	stdLoggerConfig := &logpkg.ConfigStd{
		Level:      logpkg.ParseLevel("DEBUG"),
		CallerSkip: logpkg.DefaultCallerSkip,
	}
	stdLoggerImpl, err := logpkg.NewStdLogger(stdLoggerConfig)
	if err != nil {
		panic(err)
	}

	cfg := &logpkg.ConfigFile{
		//Level:      log.LevelDebug,
		//Level:      log.LevelInfo,
		//CallerSkip: logutil.DefaultCallerSkip,
		Level:      logpkg.ParseLevel("DEBUG"),
		CallerSkip: logpkg.DefaultCallerSkip,

		Dir:      "./runtime/logs",
		Filename: "rotation",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		//StorageCounter: 2,
		StorageAge: time.Hour,
	}
	fileLogger, err := logpkg.NewFileLogger(
		cfg,
		//WithFilenameSuffix("_xxx"),
		logpkg.WithFilenameSuffix("_xxx.%Y%m%d%H%M%S.log"),
		logpkg.WithLoggerKey(map[logpkg.LoggerKey]string{logpkg.LoggerKeyTime: "date"}),
		logpkg.WithTimeFormat(timepkg.YmdHmsMillisecond),
	)
	if err != nil {
		panic(err)
	}
	logImpl := logpkg.NewMultiLogger(stdLoggerImpl, fileLogger)
	return NewLogger(kratosToLogger(logImpl))
}
