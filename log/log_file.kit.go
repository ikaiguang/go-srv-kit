package logutil

import (
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 轮转日志参数
const (
	_defaultRotationMaxAge         = time.Hour * 24 * 30     // 30天
	_defaultRotationSize           = 50 << 20                // 50M
	_defaultRotationFilenameSuffix = "_app.%Y%m%d%H%M%S.log" // 文件名后缀
)

// file 输出到文件
type file struct {
	loggerHandler *zap.Logger
}

// NewFileLogger 输出到文件
func NewFileLogger(conf *ConfigFile) (log.Logger, error) {
	handler := &file{}
	if err := handler.InitLogger(conf); err != nil {
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
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}

	logPrefix := "\n"
	switch level {
	case log.LevelDebug:
		s.loggerHandler.Debug(logPrefix, data...)
	case log.LevelInfo:
		s.loggerHandler.Info(logPrefix, data...)
	case log.LevelWarn:
		s.loggerHandler.Warn(logPrefix, data...)
	case log.LevelError:
		s.loggerHandler.Error(logPrefix, data...)
	case log.LevelFatal:
		s.loggerHandler.Fatal(logPrefix, data...)
	}
	return err
}

// InitLogger .
func (s *file) InitLogger(conf *ConfigFile) (err error) {
	// 参考 zap.NewProductionEncoderConfig()
	encoderConf := zapcore.EncoderConfig{
		MessageKey:    "xxx",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02T15:04:05.999"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		//EncodeCaller: zapcore.FullCallerEncoder,
	}

	// writer
	writer, err := s.getWriter(conf)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	encoder := zapcore.NewJSONEncoder(encoderConf)
	zapCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(writer),
		zap.NewAtomicLevelAt(ToZapLevel(conf.Level)),
	)

	// logger
	callerSkip := _defaultCallerSkipFile
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

	var opts []rotatelogs.Option

	// 轮询 时间 或 文件大小
	switch {
	case cfg.RotateTime > 0:
		opts = append(opts, rotatelogs.WithRotationTime(cfg.RotateTime))
	case cfg.RotateSize > 0:
		opts = append(opts, rotatelogs.WithRotationSize(cfg.RotateSize))
	default:
		opts = append(opts, rotatelogs.WithRotationSize(_defaultRotationSize))
	}

	// 存储 n个 或 n久
	switch {
	case cfg.StorageCounter > 0:
		opts = append(opts, rotatelogs.WithRotationCount(cfg.StorageCounter))
	case cfg.StorageAge > 0:
		opts = append(opts, rotatelogs.WithMaxAge(cfg.StorageAge))
	default:
		opts = append(opts, rotatelogs.WithMaxAge(_defaultRotationMaxAge))
	}

	// 写
	writer, err = rotatelogs.New(
		filepath.Join(cfg.Dir, cfg.Filename+_defaultRotationFilenameSuffix),
		opts...,
	)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
