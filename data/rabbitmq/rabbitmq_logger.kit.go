package rabbitmqutil

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-kratos/kratos/v2/log"
)

var _ watermill.LoggerAdapter = (*logger)(nil)

// logger ...
type logger struct {
	handler log.Logger
	fields  watermill.LogFields
}

// NewLogger ...
func NewLogger(handler log.Logger) *logger {
	return &logger{
		handler: handler,
		fields:  watermill.LogFields{},
	}
}

func (s *logger) Error(msg string, err error, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelError,
		"msg", msg,
		"error", err,
		"fields", fields,
	)
}

func (s *logger) Info(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelInfo,
		"msg", msg,
		"fields", fields,
	)
}

func (s *logger) Debug(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelDebug,
		"msg", msg,
		"fields", fields,
	)
}

func (s *logger) Trace(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelDebug,
		"isTrace", true,
		"msg", msg,
		"fields", fields,
	)
}

func (s *logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	handler := &logger{
		handler: s.handler,
		fields:  s.fields.Add(fields),
	}
	return handler
}
