package logutil

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap/zapcore"

	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"
)

// LoggerKey 日志消息key；例：time=2022-03-17T20:11:32.031, msg=testing
type LoggerKey string

// Value 值
func (l LoggerKey) Value() string {
	return string(l)
}

const (
	// DefaultCallerSkip 日志 runtime caller skip
	// 使用 单个Logger 		ConfigStd.CallerSkip = DefaultCallerSkip
	// 使用 log.MultiLogger	ConfigStd.CallerSkip = DefaultCallerSkip + 1
	//
	// 使用 logutil.Setup 单个Logger 			ConfigStd.CallerSkip = DefaultCallerSkip + 1
	// 使用 logutil.Setup + log.With 		ConfigStd.CallerSkip = DefaultCallerSkip + 2
	// 使用 logutil.Setup log.MultiLogger 	ConfigStd.CallerSkip = DefaultCallerSkip + 2
	// log.With 会把 单个Logger 转换为 MultiLogger
	DefaultCallerSkip = 3

	// DefaultCallerValuer log.With 之 log.Caller
	//
	// 使用 单个Logger 		CallerValuer = DefaultCallerValuer
	// 使用 log.MultiLogger 	CallerValuer = DefaultCallerValuer
	// 使用 logutil.Setup 	CallerValuer = DefaultCallerValuer + 2
	DefaultCallerValuer = 4

	// DefaultTimeFormat
	DefaultTimeFormat = timeutil.YmdHmsMLogger

	// zapcore.EncoderConfig keys
	LoggerKeyMessage    LoggerKey = "msg"
	LoggerKeyLevel      LoggerKey = "level"
	LoggerKeyTime       LoggerKey = "time"
	LoggerKeyName       LoggerKey = "name"
	LoggerKeyCaller     LoggerKey = "caller"
	LoggerKeyFunction   LoggerKey = "func"
	LoggerKeyStacktrace LoggerKey = "stack"
)

// NewMultiLogger wraps multi logger.
func NewMultiLogger(logs ...log.Logger) log.Logger {
	return log.MultiLogger(logs...)
}

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
