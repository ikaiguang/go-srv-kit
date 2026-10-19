package debugpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v3/log"
)

// PrintWithContext .
func PrintWithContext(ctx context.Context, a ...any) {
	handler.DebugContext(ctx, fmt.Sprint(a...))
}

// PrintlnWithContext .
func PrintlnWithContext(ctx context.Context, a ...any) {
	handler.DebugContext(ctx, fmt.Sprintln(a...))
}

// PrintfWithContext .
func PrintfWithContext(ctx context.Context, format string, a ...any) {
	handler.DebugContext(ctx, fmt.Sprintf(format, a...))
}

// PrintwWithContext .
func PrintwWithContext(ctx context.Context, keyvals ...any) {
	handler.DebugContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// DebugWithContext .
func DebugWithContext(ctx context.Context, a ...any) {
	handler.DebugContext(ctx, fmt.Sprint(a...))
}

// DebugfWithContext .
func DebugfWithContext(ctx context.Context, format string, a ...any) {
	handler.DebugContext(ctx, fmt.Sprintf(format, a...))
}

// DebugwWithContext .
func DebugwWithContext(ctx context.Context, keyvals ...any) {
	handler.DebugContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// InfoWithContext .
func InfoWithContext(ctx context.Context, a ...any) {
	handler.InfoContext(ctx, fmt.Sprint(a...))
}

// InfofWithContext .
func InfofWithContext(ctx context.Context, format string, a ...any) {
	handler.InfoContext(ctx, fmt.Sprintf(format, a...))
}

// InfowWithContext .
func InfowWithContext(ctx context.Context, keyvals ...any) {
	handler.InfoContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// WarnWithContext .
func WarnWithContext(ctx context.Context, a ...any) {
	handler.WarnContext(ctx, fmt.Sprint(a...))
}

// WarnfWithContext .
func WarnfWithContext(ctx context.Context, format string, a ...any) {
	handler.WarnContext(ctx, fmt.Sprintf(format, a...))
}

// WarnwWithContext .
func WarnwWithContext(ctx context.Context, keyvals ...any) {
	handler.WarnContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// ErrorWithContext .
func ErrorWithContext(ctx context.Context, a ...any) {
	handler.ErrorContext(ctx, fmt.Sprint(a...))
}

// ErrorfWithContext .
func ErrorfWithContext(ctx context.Context, format string, a ...any) {
	handler.ErrorContext(ctx, fmt.Sprintf(format, a...))
}

// ErrorwWithContext .
func ErrorwWithContext(ctx context.Context, keyvals ...any) {
	handler.ErrorContext(ctx, "", normalizeKeyValues(keyvals)...)
}

// FatalWithContext .
func FatalWithContext(ctx context.Context, a ...any) {
	handler.Log(ctx, log.LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// FatalfWithContext .
func FatalfWithContext(ctx context.Context, format string, a ...any) {
	handler.Log(ctx, log.LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// FatalwWithContext .
func FatalwWithContext(ctx context.Context, keyvals ...any) {
	handler.Log(ctx, log.LevelFatal, "", normalizeKeyValues(keyvals)...)
	os.Exit(1)
}
