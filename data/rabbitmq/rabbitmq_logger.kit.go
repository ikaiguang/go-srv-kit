package rabbitmqpkg

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-kratos/kratos/v2/log"
	"io"
)

var _ watermill.LoggerAdapter = (*logger)(nil)

// logger ...
type logger struct {
	handler log.Logger
	fields  watermill.LogFields
}

// NewLogger ...
func NewLogger(handler log.Logger) watermill.LoggerAdapter {
	return &logger{
		handler: handler,
		fields:  watermill.LogFields{},
	}
}

func NewLoggerFromWriters(writers ...io.Writer) watermill.LoggerAdapter {
	stdLogger := log.NewStdLogger(io.MultiWriter(writers...))
	return NewLogger(stdLogger)
}

func kvs(msg string, fields watermill.LogFields) []interface{} {
	var kv = make([]interface{}, 0, 2+len(fields)*2)
	kv = append(kv, "msg", msg)
	for k := range fields {
		kv = append(kv, k, fields[k])
	}
	return kv
}

func (s *logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"error": err})
	_ = s.handler.Log(log.LevelError, kvs(msg, fields)...)
}

func (s *logger) Info(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelInfo, kvs(msg, fields)...)
}

func (s *logger) Debug(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(log.LevelDebug, kvs(msg, fields)...)
}

func (s *logger) Trace(msg string, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"isTrace": true})
	_ = s.handler.Log(log.LevelDebug, kvs(msg, fields)...)
}

func (s *logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	handler := &logger{
		handler: s.handler,
		fields:  s.fields.Add(fields),
	}
	return handler
}
