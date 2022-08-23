package rabbitmqutil

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// NewConnection ...
func NewConnection() {
	_, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
}
