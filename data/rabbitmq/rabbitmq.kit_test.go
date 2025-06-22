package rabbitmqpkg

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/uuid"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./data/rabbitmq/ -count=1 -run TestNewSubscriber
func TestNewSubscriber(t *testing.T) {
	var (
		//amqpURI = "amqp://guest:guest@127.0.0.1:5672/"
		amqpURI = "amqp://rabbitmq:Rabbitmq.123456@my-rabbitmq:5672/"
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

	go doSub(sub, topic)
	go doPub(pub, topic)

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
	return NewLogger(logImpl)
}
