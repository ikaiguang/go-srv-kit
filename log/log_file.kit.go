package logutil

import (
	"fmt"
	"io"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	writerutil "github.com/ikaiguang/go-srv-kit/writer"
)

// 轮转日志参数
const (
	DefaultRotationStorageAge      = time.Hour * 24 * 30     // 30天
	DefaultRotationSize            = 50 << 20                // 50M
	_defaultRotationFilenameSuffix = "_app.%Y%m%d%H%M%S.log" // 文件名后缀
)

// file 输出到文件
type file struct {
	loggerHandler *zap.Logger
}

// NewFileLogger 输出到文件
func NewFileLogger(conf *ConfigFile, opts ...Option) (log.Logger, error) {
	handler := &file{}
	if err := handler.InitLogger(conf, opts...); err != nil {
		err = errors.WithStack(err)
		return handler, err
	}
	return handler, nil
}

// Log .
func (s *file) Log(level log.Level, keyvals ...interface{}) (err error) {
	if len(keyvals) == 0 {
		return err
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}

	// field
	var (
		msg  = "\n"
		data []zap.Field
	)
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}

	switch level {
	case log.LevelDebug:
		s.loggerHandler.Debug(msg, data...)
	case log.LevelInfo:
		s.loggerHandler.Info(msg, data...)
	case log.LevelWarn:
		s.loggerHandler.Warn(msg, data...)
	case log.LevelError:
		s.loggerHandler.Error(msg, data...)
	case log.LevelFatal:
		s.loggerHandler.Fatal(msg, data...)
	}
	return err
}

// InitLogger .
func (s *file) InitLogger(conf *ConfigFile, opts ...Option) (err error) {
	// 可选项
	options := options{
		writer: nil,
	}
	for _, o := range opts {
		o(&options)
	}

	// 参考 zap.NewProductionEncoderConfig()
	encoderConf := zapcore.EncoderConfig{
		MessageKey:    ZapMessageKey,
		LevelKey:      ZapLevelKey,
		TimeKey:       ZapTimeKey,
		NameKey:       ZapNameKey,
		CallerKey:     ZapCallerKey,
		FunctionKey:   ZapFunctionKey,
		StacktraceKey: ZapStacktraceKey,

		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02T15:04:05.999"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		//EncodeCaller: zapcore.FullCallerEncoder,
	}

	// writer
	if options.writer == nil {
		options.writer, err = s.getWriter(conf)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}

	encoder := zapcore.NewJSONEncoder(encoderConf)
	zapCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(options.writer),
		zap.NewAtomicLevelAt(ToZapLevel(conf.Level)),
	)

	// logger
	callerSkip := DefaultCallerSkip
	if conf.CallerSkip > 0 {
		callerSkip = conf.CallerSkip
	}
	stacktraceLevel := zapcore.DPanicLevel
	s.loggerHandler = zap.New(zapCore,
		zap.WithCaller(true),
		zap.AddCallerSkip(callerSkip),
		zap.AddStacktrace(stacktraceLevel),
	)
	return err
}

// getWriter log writer
func (s *file) getWriter(cfg *ConfigFile) (writer io.Writer, err error) {
	writerConfig := &writerutil.ConfigRotate{
		Dir:            cfg.Dir,
		Filename:       cfg.Filename,
		RotateTime:     cfg.RotateTime,
		RotateSize:     cfg.RotateSize,
		StorageCounter: cfg.StorageCounter,
		StorageAge:     cfg.StorageAge,
	}
	return writerutil.NewRotateFile(writerConfig)
}
