package logutil

import (
	"fmt"
	"io"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
)

// 轮转日志参数
const (
	DefaultRotationStorageAge      = time.Hour * 24 * 30     // 30天
	DefaultRotationSize            = 50 << 20                // 50M
	_defaultRotationFilenameSuffix = "_app.%Y%m%d%H%M%S.log" // 文件名后缀
)

var _ log.Logger = &File{}

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

	// 轮询规则：默认为：RotateTime
	// RotateTime 轮询规则：n久(默认：86400s # 86400s = 1天)
	RotateTime time.Duration
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 50<<20 = 50M)
	RotateSize int64

	// 存储规则：默认为：StorageAge
	// StorageAge 存储：n久(默认：30天)
	StorageAge time.Duration
	// StorageCounter 存储：n个(默认：10086个)
	StorageCounter uint
}

// File 输出到文件
type File struct {
	loggerHandler *zap.Logger
}

// NewFileLogger 输出到文件
func NewFileLogger(conf *ConfigFile, opts ...Option) (*File, error) {
	handler := &File{}
	if err := handler.initLogger(conf, opts...); err != nil {
		err = errors.WithStack(err)
		return handler, err
	}
	return handler, nil
}

// Sync zap.Logger.Sync
func (s *File) Sync() error {
	return s.loggerHandler.Sync()
}

// Log .
func (s *File) Log(level log.Level, keyvals ...interface{}) (err error) {
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

// initLogger .
func (s *File) initLogger(conf *ConfigFile, opts ...Option) (err error) {
	// 可选项
	option := options{
		writer:     nil,
		loggerKeys: DefaultLoggerKey(),
		timeFormat: DefaultTimeFormat,
	}
	for _, o := range opts {
		o(&option)
	}

	// 参考 zap.NewProductionEncoderConfig()
	encoderConf := zapcore.EncoderConfig{
		MessageKey:    LoggerKeyMessage.Value(),
		LevelKey:      LoggerKeyLevel.Value(),
		TimeKey:       LoggerKeyTime.Value(),
		NameKey:       LoggerKeyName.Value(),
		CallerKey:     LoggerKeyCaller.Value(),
		FunctionKey:   LoggerKeyFunction.Value(),
		StacktraceKey: LoggerKeyStacktrace.Value(),

		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(option.timeFormat))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		//EncodeCaller: zapcore.FullCallerEncoder,
	}
	SetZapLoggerKeys(&encoderConf, option.loggerKeys)

	// writer
	if option.writer == nil {
		option.writer, err = s.getWriter(conf, &option)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}

	encoder := zapcore.NewJSONEncoder(encoderConf)
	zapCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(option.writer),
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
func (s *File) getWriter(cfg *ConfigFile, opt *options) (writer io.Writer, err error) {
	writerConfig := &writerutil.ConfigRotate{
		Dir:            cfg.Dir,
		Filename:       cfg.Filename,
		RotateTime:     cfg.RotateTime,
		RotateSize:     cfg.RotateSize,
		StorageCounter: cfg.StorageCounter,
		StorageAge:     cfg.StorageAge,
	}
	var opts []writerutil.Option
	if opt.filenameSuffix != "" {
		opts = append(opts, writerutil.WithFilenameSuffix(opt.filenameSuffix))
	}
	return writerutil.NewRotateFile(writerConfig, opts...)
}
