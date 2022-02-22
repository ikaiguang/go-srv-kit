package loghelper

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/ikaiguang/go-srv-kit/log"
)

// debug
var (
	// helper log helper
	helper = defaultHandler()
)

// Setup 启动
// @Param loggers 请注意 ConfigStd.CallerSkip 的值
// @Param loggers 请注意 ConfigFile.CallerSkip 的值
//
// 此处 CallerSkip = logutil.DefaultCallerSkip + 2
func Setup(logger log.Logger) {
	helper = log.NewHelper(logger)
}

// defaultHandler .
func defaultHandler() *log.Helper {
	logger, _ := logutil.NewDummyLogger()

	return log.NewHelper(logger)
}

// WithContext returns a shallow copy of h with its context changed
// to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context) *log.Helper {
	return helper.WithContext(ctx)
}

// Log Print log by level and keyvals.
func Log(level log.Level, keyvals ...interface{}) {
	helper.Log(level, keyvals...)
}

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	helper.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	helper.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func Debugw(keyvals ...interface{}) {
	helper.Debugw(keyvals...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	helper.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	helper.Infof(format, a...)
}

// Infow logs a message at info level.
func Infow(keyvals ...interface{}) {
	helper.Infow(keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	helper.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	helper.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(keyvals ...interface{}) {
	helper.Warnw(keyvals...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	helper.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	helper.Errorf(format, a...)
}

// Errorw logs a message at error level.
func Errorw(keyvals ...interface{}) {
	helper.Errorw(keyvals...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	helper.Fatal(a...)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	helper.Fatalf(format, a...)
}

// Fatalw logs a message at fatal level.
func Fatalw(keyvals ...interface{}) {
	helper.Fatalw(keyvals...)
}