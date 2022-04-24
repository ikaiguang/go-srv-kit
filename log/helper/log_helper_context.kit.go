package loghelper

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// ContextLog Print log by level and keyvals.
func ContextLog(ctx context.Context, level log.Level, keyvals ...interface{}) {
	helper.WithContext(ctx).Log(level, keyvals...)
}

// ContextDebug logs a message at debug level.
func ContextDebug(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Debug(a...)
}

// ContextDebugf logs a message at debug level.
func ContextDebugf(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Debugf(format, a...)
}

// ContextDebugw logs a message at debug level.
func ContextDebugw(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Debugw(keyvals...)
}

// ContextInfo logs a message at info level.
func ContextInfo(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Info(a...)
}

// ContextInfof logs a message at info level.
func ContextInfof(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Infof(format, a...)
}

// ContextInfow logs a message at info level.
func ContextInfow(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Infow(keyvals...)
}

// ContextWarn logs a message at warn level.
func ContextWarn(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Warn(a...)
}

// ContextWarnf logs a message at warnf level.
func ContextWarnf(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Warnf(format, a...)
}

// ContextWarnw logs a message at warnf level.
func ContextWarnw(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Warnw(keyvals...)
}

// ContextError logs a message at error level.
func ContextError(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Error(a...)
}

// ContextErrorf logs a message at error level.
func ContextErrorf(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Errorf(format, a...)
}

// ContextErrorw logs a message at error level.
func ContextErrorw(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Errorw(keyvals...)
}

// ContextFatal logs a message at fatal level.
func ContextFatal(ctx context.Context, a ...interface{}) {
	helper.WithContext(ctx).Fatal(a...)
}

// ContextFatalf logs a message at fatal level.
func ContextFatalf(ctx context.Context, format string, a ...interface{}) {
	helper.WithContext(ctx).Fatalf(format, a...)
}

// ContextFatalw logs a message at fatal level.
func ContextFatalw(ctx context.Context, keyvals ...interface{}) {
	helper.WithContext(ctx).Fatalw(keyvals...)
}
