package debugpkg

import "context"

// PrintWithContext .
func PrintWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Debug(a...)
}

// PrintlnWithContext .
func PrintlnWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Debug(a...)
}

// PrintfWithContext .
func PrintfWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Debugf(format, a...)
}

// PrintwWithContext .
func PrintwWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Debugw(keyvals...)
}

// DebugWithContext .
func DebugWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Debug(a...)
}

// DebugfWithContext .
func DebugfWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Debugf(format, a...)
}

// DebugwWithContext .
func DebugwWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Debugw(keyvals...)
}

// InfoWithContext .
func InfoWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Info(a...)
}

// InfofWithContext .
func InfofWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Infof(format, a...)
}

// InfowWithContext .
func InfowWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Infow(keyvals...)
}

// WarnWithContext .
func WarnWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Warn(a...)
}

// WarnfWithContext .
func WarnfWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Warnf(format, a...)
}

// WarnwWithContext .
func WarnwWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Warnw(keyvals...)
}

// ErrorWithContext .
func ErrorWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Error(a...)
}

// ErrorfWithContext .
func ErrorfWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Errorf(format, a...)
}

// ErrorwWithContext .
func ErrorwWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Errorw(keyvals...)
}

// FatalWithContext .
func FatalWithContext(ctx context.Context, a ...interface{}) {
	handler.WithContext(ctx).Fatal(a...)
}

// FatalfWithContext .
func FatalfWithContext(ctx context.Context, format string, a ...interface{}) {
	handler.WithContext(ctx).Fatalf(format, a...)
}

// FatalwWithContext .
func FatalwWithContext(ctx context.Context, keyvals ...interface{}) {
	handler.WithContext(ctx).Fatalw(keyvals...)
}
