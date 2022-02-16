package logutil

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"
)

var _ log.Logger = &Std{}

// Std 标准输出
type Std struct {
	loggerHandler *zap.Logger
}

// NewStdLogger 输出到控制台
func NewStdLogger(conf *ConfigStd) (*Std, error) {
	handler := &Std{}
	if err := handler.InitLogger(conf); err != nil {
		err = errors.WithStack(err)
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
		msg += "*** | " + fmt.Sprint(keyvals[i]) + "\n"
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
func (s *Std) InitLogger(conf *ConfigStd) (err error) {
	// 参考 zap.NewDevelopmentEncoderConfig()
	encoderConf := zapcore.EncoderConfig{
		MessageKey:    ZapMessageKey,
		LevelKey:      ZapLevelKey,
		TimeKey:       ZapTimeKey,
		NameKey:       ZapNameKey,
		CallerKey:     ZapCallerKey,
		FunctionKey:   ZapFunctionKey,
		StacktraceKey: ZapStacktraceKey,

		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(timeutil.YmdHmsMLogger))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
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
		err = errors.WithStack(err)
		return
	}
	return err
}
