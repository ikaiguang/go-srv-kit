package logpkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-kratos/kratos/v3/log"
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
func Setup(logger Logger) {
	helper = ToSlog(logger)
}

// defaultHandler .
func defaultHandler() *slog.Logger {
	return NewNopLogger()
}

// Log Print log by level and keyvals.
func Log(level log.Level, keyvals ...any) {
	helper.Log(context.Background(), level, "", normalizeKeyValues(keyvals)...)
}

// Print logs a message at info level.
func Print(a ...any) {
	helper.Info(fmt.Sprint(a...))
}

// Println logs a message at info level.
func Println(a ...any) {
	helper.Info(fmt.Sprintln(a...))
}

// Printf logs a message at info level.
func Printf(format string, a ...any) {
	helper.Info(fmt.Sprintf(format, a...))
}

// Printw logs a message at info level.
func Printw(keyvals ...any) {
	helper.Info("", normalizeKeyValues(keyvals)...)
}

// Debug logs a message at debug level.
func Debug(a ...any) {
	helper.Debug(fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...any) {
	helper.Debug(fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func Debugw(keyvals ...any) {
	helper.Debug("", normalizeKeyValues(keyvals)...)
}

// Info logs a message at info level.
func Info(a ...any) {
	helper.Info(fmt.Sprint(a...))
}

// Infof logs a message at info level.
func Infof(format string, a ...any) {
	helper.Info(fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func Infow(keyvals ...any) {
	helper.Info("", normalizeKeyValues(keyvals)...)
}

// Warn logs a message at warn level.
func Warn(a ...any) {
	helper.Warn(fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...any) {
	helper.Warn(fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func Warnw(keyvals ...any) {
	helper.Warn("", normalizeKeyValues(keyvals)...)
}

// Error logs a message at error level.
func Error(a ...any) {
	helper.Error(fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func Errorf(format string, a ...any) {
	helper.Error(fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func Errorw(keyvals ...any) {
	helper.Error("", normalizeKeyValues(keyvals)...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...any) {
	helper.Log(context.Background(), log.LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...any) {
	helper.Log(context.Background(), log.LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func Fatalw(keyvals ...any) {
	helper.Log(context.Background(), log.LevelFatal, "", normalizeKeyValues(keyvals)...)
	os.Exit(1)
}

func normalizeKeyValues(keyvals []any) []any {
	if len(keyvals)%2 != 0 {
		keyvals = append(append([]any(nil), keyvals...), "KEYVALS UNPAIRED")
	}
	return keyvals
}
