package rabbitmqpkg

import (
	"io"
	"log"

	"github.com/ThreeDotsLabs/watermill"
)

var _ watermill.LoggerAdapter = (*logger)(nil)

// logger ...
type logger struct {
	handler Logger
	fields  watermill.LogFields
}

// NewLogger ...
func NewLogger(handler Logger) watermill.LoggerAdapter {
	return &logger{
		handler: handler,
		fields:  watermill.LogFields{},
	}
}

// NewLoggerFromWriters 基于 io.Writer 创建 watermill.LoggerAdapter。
// 内部使用标准库 log 实现，不依赖 kratos。
func NewLoggerFromWriters(writers ...io.Writer) watermill.LoggerAdapter {
	stdLogger := log.New(io.MultiWriter(writers...), "", log.LstdFlags)
	return NewLogger(&stdWriterLogger{logger: stdLogger})
}

// stdWriterLogger 基于标准库 log 的 Logger 实现
type stdWriterLogger struct {
	logger *log.Logger
}

func (s *stdWriterLogger) Log(level Level, keyvals ...any) error {
	s.logger.Println(keyvals...)
	return nil
}

func kvs(msg string, fields watermill.LogFields) []any {
	var kv = make([]any, 0, 2+len(fields)*2)
	kv = append(kv, "msg", msg)
	for k := range fields {
		kv = append(kv, k, fields[k])
	}
	return kv
}

func (s *logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"error": err})
	_ = s.handler.Log(LevelError, kvs(msg, fields)...)
}

func (s *logger) Info(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(LevelInfo, kvs(msg, fields)...)
}

func (s *logger) Debug(msg string, fields watermill.LogFields) {
	_ = s.handler.Log(LevelDebug, kvs(msg, fields)...)
}

func (s *logger) Trace(msg string, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"isTrace": true})
	_ = s.handler.Log(LevelDebug, kvs(msg, fields)...)
}

func (s *logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	handler := &logger{
		handler: s.handler,
		fields:  s.fields.Add(fields),
	}
	return handler
}
