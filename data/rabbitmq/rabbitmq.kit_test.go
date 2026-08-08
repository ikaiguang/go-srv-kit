package rabbitmqpkg

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
)

func TestRabbitMQConfigValidation(t *testing.T) {
	if _, err := NewConnection(nil); err == nil {
		t.Fatal("NewConnection(nil) error = nil")
	}
	if _, err := NewPublisher(&Config{}); err == nil {
		t.Fatal("NewPublisher() empty URL error = nil")
	}
	if _, err := NewSubscriberWithConnection(nil); err == nil {
		t.Fatal("NewSubscriberWithConnection(nil) error = nil")
	}
}

func TestNewOptionsHandlesNilValues(t *testing.T) {
	options := newOptions(nil, WithLogger(nil))
	if options.logger == nil {
		t.Fatal("newOptions() logger = nil")
	}

	nonDurable := newOptions(WithNonDurable())
	config := newQueueConfig(&Config{Url: "amqp://guest:guest@127.0.0.1:5672/"}, nonDurable)
	if config.Queue.Durable {
		t.Fatal("WithNonDurable() produced a durable queue")
	}
}

func TestNewLoggerFromWritersAddsSourceAndFields(t *testing.T) {
	var output bytes.Buffer
	logger := NewLoggerFromWriters(&output).With(watermill.LogFields{"module": "rabbitmq"})
	logger.Info("connected", watermill.LogFields{"attempt": 1})

	got := output.String()
	for _, want := range []string{"connected", "attempt=1", "module=rabbitmq", "source="} {
		if !strings.Contains(got, want) {
			t.Fatalf("log output = %q, want containing %q", got, want)
		}
	}
}
