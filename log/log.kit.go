package logutil

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap/zapcore"
)

// logger config
const (
	// DefaultCallerSkip 日志 runtime caller skip
	DefaultCallerSkip = 3
)

// NewMultiLogger wraps multi logger.
func NewMultiLogger(logs ...log.Logger) log.Logger {
	return log.MultiLogger(logs...)
}

// ConfigStd 标准输出
type ConfigStd struct {
	// Level 日志级别
	Level log.Level
	// 日志 runtime caller skips
	// log.NewHelper callerSkip = DefaultCallerSkip
	// log.MultiLogger callerSkip = DefaultCallerSkip + 1
	CallerSkip int
}

// ConfigFile 输出到文件
type ConfigFile struct {
	// Level 日志级别
	Level log.Level
	// CallerSkip 日志 runtime caller skips
	// log.NewHelper callerSkip = DefaultCallerSkip
	// log.MultiLogger callerSkip = DefaultCallerSkip + 1
	CallerSkip int

	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// 轮询：n久 或 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	// RotateTime n久
	RotateTime time.Duration
	// RotateSize 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateSize int64

	// 存储：n个 或 有效期StorageAge(默认：30天)
	// StorageCounter n个
	StorageCounter uint
	// StorageAge 存储n久(默认：30天)
	StorageAge time.Duration
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
