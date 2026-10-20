package logpkg

import (
	"log/slog"
	"time"

	"github.com/go-kratos/kratos/v3/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*Std)(nil)

// ConfigStd 标准输出
type ConfigStd struct {
	// Level 日志级别
	Level log.Level
	// CallerSkip 日志 runtime caller skips
	CallerSkip     int
	UseJSONEncoder bool
}

// Std 标准输出
type Std struct {
	*slog.Logger
	conf          *ConfigStd
	loggerHandler *zap.Logger
}

// NewStdLogger 输出到控制台
func NewStdLogger(conf *ConfigStd) (*Std, error) {
	if conf == nil {
		return nil, errNilConfig
	}
	handler := &Std{
		conf: conf,
	}
	if err := handler.InitLogger(conf); err != nil {
		return handler, err
	}
	return handler, nil
}

// sync zap.Logger.Sync
func (s *Std) sync() error {
	return s.loggerHandler.Sync()
}

// Close zap.Logger.Sync
func (s *Std) Close() error {
	return s.sync()
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
		if o != nil {
			o(&option)
		}
	}

	var (
		encoding    = "console"
		encodeLevel = zapcore.CapitalColorLevelEncoder
	)
	if conf.UseJSONEncoder {
		encoding = "json"
		encodeLevel = zapcore.CapitalLevelEncoder
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
		EncodeLevel: encodeLevel,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(option.timeFormat))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		//EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
	SetZapLoggerKeys(&encoderConf, option.loggerKeys)

	// logger
	callerSkip := DefaultCallerSkip
	if conf.CallerSkip > 0 {
		callerSkip = conf.CallerSkip
	}
	stacktraceLevel := zapcore.DPanicLevel
	var zOpts = []zap.Option{
		zap.AddCallerSkip(callerSkip),
		zap.AddStacktrace(stacktraceLevel),
	}

	// 参考 zap.NewDevelopmentConfig()
	loggerConf := &zap.Config{
		Level:            zap.NewAtomicLevelAt(ToZapLevel(conf.Level)),
		Development:      true,
		Sampling:         nil,
		Encoding:         encoding,
		EncoderConfig:    encoderConf,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	s.loggerHandler, err = loggerConf.Build(zOpts...)
	if err != nil {
		return
	}
	s.Logger = slog.New(&zapHandler{logger: s.loggerHandler})
	return err
}
