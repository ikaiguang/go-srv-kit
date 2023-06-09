package logpkg

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// LogWithContext Print log by level and keyvals.
func LogWithContext(ctx context.Context, level log.Level, keyvals ...interface{}) {
	helper.WithContext(ctx).Log(level, keyvals...)
}

// DebugWithContext logs a message at debug level.
func DebugWithContext(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Debug(a...)
}

// DebugfWithContext logs a message at debug level.
func DebugfWithContext(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Debugf(format, a...)
}

// DebugwWithContext logs a message at debug level.
func DebugwWithContext(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Debugw(keyvals...)
}

// InfoWithContext logs a message at info level.
func InfoWithContext(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Info(a...)
}

// InfofWithContext logs a message at info level.
func InfofWithContext(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Infof(format, a...)
}

// InfowWithContext logs a message at info level.
func InfowWithContext(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Infow(keyvals...)
}

// WarnWithContext logs a message at warn level.
func WarnWithContext(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Warn(a...)
}

// WarnfWithContext logs a message at warnf level.
func WarnfWithContext(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Warnf(format, a...)
}

// WarnwWithContext logs a message at warnf level.
func WarnwWithContext(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Warnw(keyvals...)
}

// ErrorWithContext logs a message at error level.
func ErrorWithContext(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Error(a...)
}

// ErrorfWithContext logs a message at error level.
func ErrorfWithContext(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Errorf(format, a...)
}

// ErrorwWithContext logs a message at error level.
func ErrorwWithContext(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Errorw(keyvals...)
}

// FatalWithContext logs a message at fatal level.
func FatalWithContext(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Fatal(a...)
}

// FatalfWithContext logs a message at fatal level.
func FatalfWithContext(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Fatalf(format, a...)
}

// FatalwWithContext logs a message at fatal level.
func FatalwWithContext(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Fatalw(keyvals...)
}
