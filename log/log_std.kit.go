package logutil

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ log.Logger = &Std{}

// ConfigStd 标准输出
type ConfigStd struct {
	// Level 日志级别
	Level log.Level
	// CallerSkip 日志 runtime caller skips
	CallerSkip int
}

// Std 标准输出
type Std struct {
	loggerHandler *zap.Logger
}

// NewStdLogger 输出到控制台
func NewStdLogger(conf *ConfigStd) (*Std, error) {
	handler := &Std{}
	if err := handler.InitLogger(conf); err != nil {
		return handler, err
	}
	return handler, nil
}

// Sync zap.Logger.Sync
func (s *Std) Sync() error {
	return s.loggerHandler.Sync()
}

// Log .
func (s *Std) Log(level log.Level, keyvals ...interface{}) (err error) {
	if len(keyvals) == 0 {
		return err
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}

	// field
	var (
		msg = "\n"
	)
	for i := 0; i < len(keyvals); i += 2 {
		msg += "*** ｜ " + fmt.Sprint(keyvals[i]) + "\n"
		msg += "\t" + fmt.Sprint(keyvals[i+1]) + "\n"
	}

	switch level {
	case log.LevelDebug:
		s.loggerHandler.Debug(msg)
	case log.LevelInfo:
		s.loggerHandler.Info(msg)
	case log.LevelWarn:
		s.loggerHandler.Warn(msg)
	case log.LevelError:
		s.loggerHandler.Error(msg)
	case log.LevelFatal:
		s.loggerHandler.Fatal(msg)
	}
	return err
}

// InitLogger .
func (s *Std) InitLogger(conf *ConfigStd, opts ...Option) (err error) {
	// 可选项
	option := options{
		writer:     nil,
		loggerKeys: DefaultLoggerKey(),
		timeFormat: DefaultTimeFormat,
	}
	for _, o := range opts {
		o(&option)
	}

	// 参考 zap.NewDevelopmentEncoderConfig()
	encoderConf := zapcore.EncoderConfig{
		MessageKey:    LoggerKeyMessage.Value(),
		LevelKey:      LoggerKeyLevel.Value(),
		TimeKey:       LoggerKeyTime.Value(),
		NameKey:       LoggerKeyName.Value(),
		CallerKey:     LoggerKeyCaller.Value(),
		FunctionKey:   LoggerKeyFunction.Value(),
		StacktraceKey: LoggerKeyStacktrace.Value(),

		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(option.timeFormat))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
	SetZapLoggerKeys(&encoderConf, option.loggerKeys)

	// 参考 zap.NewDevelopmentConfig()
	loggerConf := &zap.Config{
		Level:            zap.NewAtomicLevelAt(ToZapLevel(conf.Level)),
		Development:      true,
		Sampling:         nil,
		Encoding:         "console",
		EncoderConfig:    encoderConf,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// logger
	callerSkip := DefaultCallerSkip
	if conf.CallerSkip > 0 {
		callerSkip = conf.CallerSkip
	}
	stacktraceLevel := zapcore.DPanicLevel
	s.loggerHandler, err = loggerConf.Build(
		zap.AddCallerSkip(callerSkip),
		zap.AddStacktrace(stacktraceLevel),
	)
	if err != nil {
		return
	}
	return err
}
