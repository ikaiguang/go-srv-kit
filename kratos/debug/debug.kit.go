package debugpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v3/log"
)

// Print .
func Print(a ...any) {
	handler.Debug(fmt.Sprint(a...))
}

// Println .
func Println(a ...any) {
	handler.Debug(fmt.Sprintln(a...))
}

// Printf .
func Printf(format string, a ...any) {
	handler.Debug(fmt.Sprintf(format, a...))
}

// Printw .
func Printw(keyvals ...any) {
	handler.Debug("", normalizeKeyValues(keyvals)...)
}

// Debug .
func Debug(a ...any) {
	handler.Debug(fmt.Sprint(a...))
}

// Debugf .
func Debugf(format string, a ...any) {
	handler.Debug(fmt.Sprintf(format, a...))
}

// Debugw .
func Debugw(keyvals ...any) {
	handler.Debug("", normalizeKeyValues(keyvals)...)
}

// Info .
func Info(a ...any) {
	handler.Info(fmt.Sprint(a...))
}

// Infof .
func Infof(format string, a ...any) {
	handler.Info(fmt.Sprintf(format, a...))
}

// Infow .
func Infow(keyvals ...any) {
	handler.Info("", normalizeKeyValues(keyvals)...)
}

// Warn .
func Warn(a ...any) {
	handler.Warn(fmt.Sprint(a...))
}

// Warnf .
func Warnf(format string, a ...any) {
	handler.Warn(fmt.Sprintf(format, a...))
}

// Warnw .
func Warnw(keyvals ...any) {
	handler.Warn("", normalizeKeyValues(keyvals)...)
}

// Error .
func Error(a ...any) {
	handler.Error(fmt.Sprint(a...))
}

// Errorf .
func Errorf(format string, a ...any) {
	handler.Error(fmt.Sprintf(format, a...))
}

// Errorw .
func Errorw(keyvals ...any) {
	handler.Error("", normalizeKeyValues(keyvals)...)
}

// Fatal .
func Fatal(a ...any) {
	handler.Log(context.Background(), log.LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf .
func Fatalf(format string, a ...any) {
	handler.Log(context.Background(), log.LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw .
func Fatalw(keyvals ...any) {
	handler.Log(context.Background(), log.LevelFatal, "", normalizeKeyValues(keyvals)...)
	os.Exit(1)
}

func normalizeKeyValues(keyvals []any) []any {
	if len(keyvals)%2 != 0 {
		keyvals = append(append([]any(nil), keyvals...), "KEYVALS UNPAIRED")
	}
	return keyvals
}
