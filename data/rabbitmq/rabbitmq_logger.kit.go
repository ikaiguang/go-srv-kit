package rabbitmqpkg

import (
	"context"
	"io"
	"log/slog"
	"sort"

	"github.com/ThreeDotsLabs/watermill"
)

var _ watermill.LoggerAdapter = (*logger)(nil)

// logger ...
type logger struct {
	handler *slog.Logger
	fields  watermill.LogFields
}

// NewLogger ...
func NewLogger(handler *slog.Logger) watermill.LoggerAdapter {
	if handler == nil {
		handler = slog.Default()
	}

	return &logger{
		handler: handler,
		fields:  watermill.LogFields{},
	}
}

func NewLoggerFromWriters(writers ...io.Writer) watermill.LoggerAdapter {
	handler := slog.NewTextHandler(io.MultiWriter(writers...), &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	return NewLogger(slog.New(handler))
}

func attrs(fields watermill.LogFields) []slog.Attr {
	var attributes = make([]slog.Attr, 0, len(fields))
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		attributes = append(attributes, slog.Any(k, fields[k]))
	}
	return attributes
}

func (s *logger) log(level slog.Level, msg string, fields watermill.LogFields) {
	fields = s.fields.Add(fields)
	s.handler.LogAttrs(context.Background(), level, msg, attrs(fields)...)
}

func (s *logger) Error(msg string, err error, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"error": err})
	s.log(slog.LevelError, msg, fields)
}

func (s *logger) Info(msg string, fields watermill.LogFields) {
	s.log(slog.LevelInfo, msg, fields)
}

func (s *logger) Debug(msg string, fields watermill.LogFields) {
	s.log(slog.LevelDebug, msg, fields)
}

func (s *logger) Trace(msg string, fields watermill.LogFields) {
	fields = fields.Add(watermill.LogFields{"isTrace": true})
	s.log(slog.LevelDebug, msg, fields)
}

func (s *logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	handler := &logger{
		handler: s.handler,
		fields:  s.fields.Add(fields),
	}
	return handler
}
