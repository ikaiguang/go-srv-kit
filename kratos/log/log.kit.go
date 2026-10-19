package logpkg

import (
	"context"
	"log/slog"
	"runtime"
	"strings"

	"github.com/go-kratos/kratos/v3/log"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/v3/time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerKey 日志消息key；例：time=2022-03-17T20:11:32.031, msg=testing
type LoggerKey string

// Value 值
func (l LoggerKey) Value() string {
	return string(l)
}

const (
	// DefaultCallerSkip 日志 runtime caller skip
	DefaultCallerSkip = 3
	// DefaultCallerValuer log.With 之 log.Caller
	DefaultCallerValuer = 4
)

var (
	// CallerSkipForMiddleware 中间件日志
	CallerSkipForMiddleware = DefaultCallerSkip
	// CallerSkipForLogger 基础日志
	CallerSkipForLogger = DefaultCallerSkip + 1
	// CallerSkipForHelper 日志助手
	CallerSkipForHelper = DefaultCallerSkip + 2
)

// Logger is implemented by slog.Logger and the package logger wrappers.
type Logger interface {
	Handler() slog.Handler
}

// ToSlog returns logger as a standard library slog logger.
func ToSlog(logger Logger) *slog.Logger {
	if logger == nil {
		return NewNopLogger()
	}
	if slogLogger, ok := logger.(*slog.Logger); ok {
		return slogLogger
	}
	return slog.New(logger.Handler())
}

// NewMultiLogger wraps multiple loggers.
func NewMultiLogger(loggers ...Logger) *slog.Logger {
	handlers := make([]slog.Handler, 0, len(loggers))
	for _, logger := range loggers {
		if logger != nil {
			handlers = append(handlers, logger.Handler())
		}
	}
	return slog.New(&multiHandler{handlers: handlers})
}

const (
	// LoggerKeyMessage zapcore.EncoderConfig keys
	LoggerKeyMessage    LoggerKey = "msg"
	LoggerKeyLevel      LoggerKey = "level"
	LoggerKeyTime       LoggerKey = "time"
	LoggerKeyName       LoggerKey = "name"
	LoggerKeyCaller     LoggerKey = "caller"
	LoggerKeyFunction   LoggerKey = "func"
	LoggerKeyStacktrace LoggerKey = "stack"

	// DefaultTimeFormat 日志时间格式
	DefaultTimeFormat = timepkg.YmdHmsMLogger
)

// DefaultLoggerKey 日志消息key
func DefaultLoggerKey() map[LoggerKey]string {
	return map[LoggerKey]string{
		LoggerKeyMessage:    LoggerKeyMessage.Value(),
		LoggerKeyLevel:      LoggerKeyLevel.Value(),
		LoggerKeyTime:       LoggerKeyTime.Value(),
		LoggerKeyName:       LoggerKeyName.Value(),
		LoggerKeyCaller:     LoggerKeyCaller.Value(),
		LoggerKeyFunction:   LoggerKeyFunction.Value(),
		LoggerKeyStacktrace: LoggerKeyStacktrace.Value(),
	}
}

// ParseLevel 日志级别
func ParseLevel(s string) log.Level {
	return log.ParseLevel(s)
}

// ToZapLevel .
func ToZapLevel(lv log.Level) zapcore.Level {
	switch lv {
	case log.LevelDebug:
		return zapcore.DebugLevel
	case log.LevelInfo:
		return zapcore.InfoLevel
	case log.LevelWarn:
		return zapcore.WarnLevel
	case log.LevelError:
		return zapcore.ErrorLevel
	case log.LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// SetZapLoggerKeys 设置日志key
func SetZapLoggerKeys(encoderConfig *zapcore.EncoderConfig, loggerKeys map[LoggerKey]string) {
	if data, ok := loggerKeys[LoggerKeyMessage]; ok {
		encoderConfig.MessageKey = data
	}
	if data, ok := loggerKeys[LoggerKeyLevel]; ok {
		encoderConfig.LevelKey = data
	}
	if data, ok := loggerKeys[LoggerKeyTime]; ok {
		encoderConfig.TimeKey = data
	}
	if data, ok := loggerKeys[LoggerKeyName]; ok {
		encoderConfig.NameKey = data
	}
	if data, ok := loggerKeys[LoggerKeyCaller]; ok {
		encoderConfig.CallerKey = data
	}
	if data, ok := loggerKeys[LoggerKeyFunction]; ok {
		encoderConfig.FunctionKey = data
	}
	if data, ok := loggerKeys[LoggerKeyStacktrace]; ok {
		encoderConfig.StacktraceKey = data
	}
}

type multiHandler struct {
	handlers []slog.Handler
}

func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *multiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, record.Level) {
			if err := handler.Handle(ctx, record.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithAttrs(attrs)
	}
	return &multiHandler{handlers: handlers}
}

func (h *multiHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithGroup(name)
	}
	return &multiHandler{handlers: handlers}
}

type zapHandler struct {
	logger *zap.Logger
	attrs  []zap.Field
	groups []string
}

func (h *zapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return h.logger.Core().Enabled(ToZapLevel(level))
}

func (h *zapHandler) Handle(_ context.Context, record slog.Record) error {
	fields := append([]zap.Field(nil), h.attrs...)
	record.Attrs(func(attr slog.Attr) bool {
		fields = appendSlogAttr(fields, h.groups, attr)
		return true
	})
	entry := zapcore.Entry{
		Level:   ToZapLevel(record.Level),
		Time:    record.Time,
		Message: record.Message,
	}
	if record.PC != 0 {
		frame, _ := runtime.CallersFrames([]uintptr{record.PC}).Next()
		entry.Caller = zapcore.EntryCaller{
			Defined:  true,
			PC:       record.PC,
			File:     frame.File,
			Line:     frame.Line,
			Function: frame.Function,
		}
	}
	if checked := h.logger.Core().Check(entry, nil); checked != nil {
		checked.Write(fields...)
	}
	return nil
}

func (h *zapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	clone := &zapHandler{
		logger: h.logger,
		attrs:  append([]zap.Field(nil), h.attrs...),
		groups: append([]string(nil), h.groups...),
	}
	for _, attr := range attrs {
		clone.attrs = appendSlogAttr(clone.attrs, clone.groups, attr)
	}
	return clone
}

func (h *zapHandler) WithGroup(name string) slog.Handler {
	clone := &zapHandler{
		logger: h.logger,
		attrs:  append([]zap.Field(nil), h.attrs...),
		groups: append([]string(nil), h.groups...),
	}
	if name != "" {
		clone.groups = append(clone.groups, name)
	}
	return clone
}

func appendSlogAttr(fields []zap.Field, groups []string, attr slog.Attr) []zap.Field {
	attr.Value = attr.Value.Resolve()
	if attr.Value.Kind() == slog.KindGroup {
		groups = append(groups, attr.Key)
		for _, child := range attr.Value.Group() {
			fields = appendSlogAttr(fields, groups, child)
		}
		return fields
	}
	keyParts := append(append([]string(nil), groups...), attr.Key)
	return append(fields, zap.Any(strings.Join(keyParts, "."), attr.Value.Any()))
}
