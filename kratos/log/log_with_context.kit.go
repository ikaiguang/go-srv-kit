package logpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v3/log"
)

// LogWithContext Print log by level and keyvals.
func LogWithContext(ctx context.Context, level log.Level, keyvals ...any) {
	helper.Log(ctx, level, "", normalizeKeyValues(keyvals)...)
}

// DebugWithContext logs a message at debug level.
func DebugWithContext(ctx context.Context, a ...any) {
	helper.DebugContext(ctx, fmt.Sprint(a...))
}

// DebugfWithContext logs a message at debug level.
func DebugfWithContext(ctx context.Context, format string, a ...any) {
	helper.DebugContext(ctx, fmt.Sprintf(format, a...))
}

// DebugwWithContext logs a message at debug level.
func DebugwWithContext(ctx context.Context, keyvals ...any) {
	helper.DebugContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// InfoWithContext logs a message at info level.
func InfoWithContext(ctx context.Context, a ...any) {
	helper.InfoContext(ctx, fmt.Sprint(a...))
}

// InfofWithContext logs a message at info level.
func InfofWithContext(ctx context.Context, format string, a ...any) {
	helper.InfoContext(ctx, fmt.Sprintf(format, a...))
}

// InfowWithContext logs a message at info level.
func InfowWithContext(ctx context.Context, keyvals ...any) {
	helper.InfoContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// WarnWithContext logs a message at warn level.
func WarnWithContext(ctx context.Context, a ...any) {
	helper.WarnContext(ctx, fmt.Sprint(a...))
}

// WarnfWithContext logs a message at warnf level.
func WarnfWithContext(ctx context.Context, format string, a ...any) {
	helper.WarnContext(ctx, fmt.Sprintf(format, a...))
}

// WarnwWithContext logs a message at warnf level.
func WarnwWithContext(ctx context.Context, keyvals ...any) {
	helper.WarnContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// ErrorWithContext logs a message at error level.
func ErrorWithContext(ctx context.Context, a ...any) {
	helper.ErrorContext(ctx, fmt.Sprint(a...))
}

// ErrorfWithContext logs a message at error level.
func ErrorfWithContext(ctx context.Context, format string, a ...any) {
	helper.ErrorContext(ctx, fmt.Sprintf(format, a...))
}

// ErrorwWithContext logs a message at error level.
func ErrorwWithContext(ctx context.Context, keyvals ...any) {
	helper.ErrorContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// FatalWithContext logs a message at fatal level.
func FatalWithContext(ctx context.Context, a ...any) {
	helper.Log(ctx, log.LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// FatalfWithContext logs a message at fatal level.
func FatalfWithContext(ctx context.Context, format string, a ...any) {
	helper.Log(ctx, log.LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// FatalwWithContext logs a message at fatal level.
func FatalwWithContext(ctx context.Context, keyvals ...any) {
	helper.Log(ctx, log.LevelFatal, "", normalizeKeyValues(keyvals)...)
	os.Exit(1)
}
