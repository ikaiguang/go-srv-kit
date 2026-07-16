package rabbitmq

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	threadpkg "github.com/ikaiguang/go-srv-kit/kit/v3/thread"
	uuidpkg "github.com/ikaiguang/go-srv-kit/kit/v3/uuid"
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
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	return NewLogger(slog.New(handler))
}
