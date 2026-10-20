package logpkg

import (
	"io"
	"log/slog"
	"time"

	"github.com/go-kratos/kratos/v3/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	writerpkg "github.com/ikaiguang/go-srv-kit/kit/v3/writer"
)

// 轮转日志参数
const (
	DefaultRotationStorageAge      = time.Hour * 24 * 30     // 30天
	DefaultRotationSize            = 100 << 20               // 100M
	_defaultRotationFilenameSuffix = "_app.%Y%m%d%H%M%S.log" // 文件名后缀
)

var _ Logger = (*File)(nil)

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
	// RotateSize 轮询规则：按文件大小RotateSize(默认：52428800 # 100<<20 = 100M)
	RotateSize int64

	// 存储规则：默认为：StorageAge
	// StorageAge 存储：n久(默认：30天)
	StorageAge time.Duration
	// StorageCounter 存储：n个(默认：10086个)
	StorageCounter uint
}

// File 输出到文件
type File struct {
	*slog.Logger
	loggerHandler *zap.Logger
}

// NewFileLogger 输出到文件
func NewFileLogger(conf *ConfigFile, opts ...Option) (*File, error) {
	if conf == nil {
		return nil, errNilConfig
	}
	handler := &File{}
	if err := handler.initLogger(conf, opts...); err != nil {
		return handler, err
	}
	return handler, nil
}

// sync zap.Logger.Sync
func (s *File) sync() error {
	return s.loggerHandler.Sync()
}

// Close zap.Logger.Sync
func (s *File) Close() error {
	return s.sync()
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
		if o != nil {
			o(&option)
		}
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
	s.Logger = slog.New(&zapHandler{logger: s.loggerHandler})
	return err
}

// getWriter log writer
func (s *File) getWriter(cfg *ConfigFile, opt *options) (writer io.Writer, err error) {
	writerConfig := &writerpkg.ConfigRotate{
		Dir:            cfg.Dir,
		Filename:       cfg.Filename,
		RotateTime:     cfg.RotateTime,
		RotateSize:     cfg.RotateSize,
		StorageCounter: cfg.StorageCounter,
		StorageAge:     cfg.StorageAge,
	}
	var opts []writerpkg.Option
	if opt.filenameSuffix != "" {
		opts = append(opts, writerpkg.WithFilenameSuffix(opt.filenameSuffix))
	}
	return writerpkg.NewRotateFile(writerConfig, opts...)
}
