package debugpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v3/log"
)

// Print .
func Print(a ...any) {
	currentHandler().Debug(fmt.Sprint(a...))
}

// Println .
func Println(a ...any) {
	currentHandler().Debug(fmt.Sprintln(a...))
}

// Printf .
func Printf(format string, a ...any) {
	currentHandler().Debug(fmt.Sprintf(format, a...))
}

// Printw .
func Printw(keyvals ...any) {
	currentHandler().Debug("", normalizeKeyValues(keyvals)...)
}

// Debug .
func Debug(a ...any) {
	currentHandler().Debug(fmt.Sprint(a...))
}

// Debugf .
func Debugf(format string, a ...any) {
	currentHandler().Debug(fmt.Sprintf(format, a...))
}

// Debugw .
func Debugw(keyvals ...any) {
	currentHandler().Debug("", normalizeKeyValues(keyvals)...)
}

// Info .
func Info(a ...any) {
	currentHandler().Info(fmt.Sprint(a...))
}

// Infof .
func Infof(format string, a ...any) {
	currentHandler().Info(fmt.Sprintf(format, a...))
}

// Infow .
func Infow(keyvals ...any) {
	currentHandler().Info("", normalizeKeyValues(keyvals)...)
}

// Warn .
func Warn(a ...any) {
	currentHandler().Warn(fmt.Sprint(a...))
}

// Warnf .
func Warnf(format string, a ...any) {
	currentHandler().Warn(fmt.Sprintf(format, a...))
}

// Warnw .
func Warnw(keyvals ...any) {
	currentHandler().Warn("", normalizeKeyValues(keyvals)...)
}

// Error .
func Error(a ...any) {
	currentHandler().Error(fmt.Sprint(a...))
}

// Errorf .
func Errorf(format string, a ...any) {
	currentHandler().Error(fmt.Sprintf(format, a...))
}

// Errorw .
func Errorw(keyvals ...any) {
	currentHandler().Error("", normalizeKeyValues(keyvals)...)
}

// Fatal .
func Fatal(a ...any) {
	currentHandler().Log(context.Background(), log.LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf .
func Fatalf(format string, a ...any) {
	currentHandler().Log(context.Background(), log.LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw .
func Fatalw(keyvals ...any) {
	currentHandler().Log(context.Background(), log.LevelFatal, "", normalizeKeyValues(keyvals)...)
	os.Exit(1)
}

func normalizeKeyValues(keyvals []any) []any {
	if len(keyvals)%2 != 0 {
		keyvals = append(append([]any(nil), keyvals...), "KEYVALS UNPAIRED")
	}
	return keyvals
}
