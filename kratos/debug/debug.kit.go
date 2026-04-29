package debugpkg

// Print .
func Print(a ...any) {
	handler.Debug(a...)
}

// Println .
func Println(a ...any) {
	handler.Debug(a...)
}

// Printf .
func Printf(format string, a ...any) {
	handler.Debugf(format, a...)
}

// Printw .
func Printw(keyvals ...any) {
	handler.Debugw(keyvals...)
}

// Debug .
func Debug(a ...any) {
	handler.Debug(a...)
}

// Debugf .
func Debugf(format string, a ...any) {
	handler.Debugf(format, a...)
}

// Debugw .
func Debugw(keyvals ...any) {
	handler.Debugw(keyvals...)
}

// Info .
func Info(a ...any) {
	handler.Info(a...)
}

// Infof .
func Infof(format string, a ...any) {
	handler.Infof(format, a...)
}

// Infow .
func Infow(keyvals ...any) {
	handler.Infow(keyvals...)
}

// Warn .
func Warn(a ...any) {
	handler.Warn(a...)
}

// Warnf .
func Warnf(format string, a ...any) {
	handler.Warnf(format, a...)
}

// Warnw .
func Warnw(keyvals ...any) {
	handler.Warnw(keyvals...)
}

// Error .
func Error(a ...any) {
	handler.Error(a...)
}

// Errorf .
func Errorf(format string, a ...any) {
	handler.Errorf(format, a...)
}

// Errorw .
func Errorw(keyvals ...any) {
	handler.Errorw(keyvals...)
}

// Fatal .
func Fatal(a ...any) {
	handler.Fatal(a...)
}

// Fatalf .
func Fatalf(format string, a ...any) {
	handler.Fatalf(format, a...)
}

// Fatalw .
func Fatalw(keyvals ...any) {
	handler.Fatalw(keyvals...)
}
