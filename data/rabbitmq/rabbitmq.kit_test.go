package rabbitmqutil

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"
	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// go test -v ./data/rabbitmq/ -count=1 -test.run=TestNewSubscriber
func TestNewSubscriber(t *testing.T) {
	var (
		amqpURI = "amqp://guest:guest@127.0.0.1:5672/"
		topic   = "example.topic"
	)
	logger := newMultiLogger()

	conf := &confv1.Base_Rabbitmq{
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
		msg := message.NewMessage(uuidutil.NewUUID(), []byte(
			fmt.Sprintf("%d : Hello, world!", i)))

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
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      logutil.ParseLevel("DEBUG"),
		CallerSkip: logutil.DefaultCallerSkip,
	}
	stdLoggerImpl, err := logutil.NewStdLogger(stdLoggerConfig)
	if err != nil {
		panic(err)
	}

	cfg := &logutil.ConfigFile{
		//Level:      log.LevelDebug,
		//Level:      log.LevelInfo,
		//CallerSkip: logutil.DefaultCallerSkip,
		Level:      logutil.ParseLevel("DEBUG"),
		CallerSkip: logutil.DefaultCallerSkip,

		Dir:      "./runtime/logs",
		Filename: "rotation",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		//StorageCounter: 2,
		StorageAge: time.Hour,
	}
	fileLogger, err := logutil.NewFileLogger(
		cfg,
		//WithFilenameSuffix("_xxx"),
		logutil.WithFilenameSuffix("_xxx.%Y%m%d%H%M%S.log"),
		logutil.WithLoggerKey(map[logutil.LoggerKey]string{logutil.LoggerKeyTime: "date"}),
		logutil.WithTimeFormat(timeutil.YmdHmsMillisecond),
	)
	if err != nil {
		panic(err)
	}
	logImpl := logutil.NewMultiLogger(stdLoggerImpl, fileLogger)
	return NewLogger(logImpl)
}
