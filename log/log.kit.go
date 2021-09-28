package logutil

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap/zapcore"
)

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

	// zapcore.EncoderConfig keys
	ZapMessageKey    = "z_Msg"
	ZapLevelKey      = "z_Lv"
	ZapTimeKey       = "z_Time"
	ZapNameKey       = "z_Name"
	ZapCallerKey     = "z_Caller"
	ZapFunctionKey   = "z_Fn"
	ZapStacktraceKey = "z_ST"
)

// NewMultiLogger wraps multi logger.
func NewMultiLogger(logs ...log.Logger) log.Logger {
	return log.MultiLogger(logs...)
}

// ConfigStd 标准输出
type ConfigStd struct {
	// Level 日志级别
	Level log.Level
	// CallerSkip 日志 runtime caller skips
	CallerSkip int
}

// ConfigFile 输出到文件
type ConfigFile struct {
	// Level 日志级别
	Level log.Level
	// CallerSkip 日志 runtime caller skips
	CallerSkip int

	// 存储位置
	// Dir 文件夹
	Dir string
	// Filename 文件名(默认：${filename}_app.%Y%m%d.log)
	Filename string

	// RotateTime 轮询规则：n久 或 文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateTime time.Duration
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateSize int64

	// StorageCounter 存储：n个 或 有效期StorageAge(默认：30天)
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
