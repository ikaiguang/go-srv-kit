package logpkg

import (
	"github.com/go-kratos/kratos/v2/log"
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
	logger, _ := NewDummyLogger()

	return log.NewHelper(logger)
}

// Log Print log by level and keyvals.
func Log(level log.Level, keyvals ...any) {
	helper.Log(level, keyvals...)
}

// Print logs a message at info level.
func Print(a ...any) {
	helper.Info(a...)
}

// Println logs a message at info level.
func Println(a ...any) {
	helper.Info(a...)
}

// Printf logs a message at info level.
func Printf(format string, a ...any) {
	helper.Infof(format, a...)
}

// Printw logs a message at info level.
func Printw(keyvals ...any) {
	helper.Infow(keyvals...)
}

// Debug logs a message at debug level.
func Debug(a ...any) {
	helper.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...any) {
	helper.Debugf(format, a...)
}

// Debugw logs a message at debug level.
func Debugw(keyvals ...any) {
	helper.Debugw(keyvals...)
}

// Info logs a message at info level.
func Info(a ...any) {
	helper.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...any) {
	helper.Infof(format, a...)
}

// Infow logs a message at info level.
func Infow(keyvals ...any) {
	helper.Infow(keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...any) {
	helper.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...any) {
	helper.Warnf(format, a...)
}

// Warnw logs a message at warnf level.
func Warnw(keyvals ...any) {
	helper.Warnw(keyvals...)
}

// Error logs a message at error level.
func Error(a ...any) {
	helper.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...any) {
	helper.Errorf(format, a...)
}

// Errorw logs a message at error level.
func Errorw(keyvals ...any) {
	helper.Errorw(keyvals...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...any) {
	helper.Fatal(a...)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...any) {
	helper.Fatalf(format, a...)
}

// Fatalw logs a message at fatal level.
func Fatalw(keyvals ...any) {
	helper.Fatalw(keyvals...)
}
