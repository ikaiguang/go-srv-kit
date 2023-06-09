package debugpkg

// Print .
func Print(a ...interface{}) {
	handler.Debug(a...)
}

// Println .
func Println(a ...interface{}) {
	handler.Debug(a...)
}

// Printf .
func Printf(format string, a ...interface{}) {
	handler.Debugf(format, a...)
}

// Printw .
func Printw(keyvals ...interface{}) {
	handler.Debugw(keyvals...)
}

// Debug .
func Debug(a ...interface{}) {
	handler.Debug(a...)
}

// Debugf .
func Debugf(format string, a ...interface{}) {
	handler.Debugf(format, a...)
}

// Debugw .
func Debugw(keyvals ...interface{}) {
	handler.Debugw(keyvals...)
}

// Info .
func Info(a ...interface{}) {
	handler.Info(a...)
}

// Infof .
func Infof(format string, a ...interface{}) {
	handler.Infof(format, a...)
}

// Infow .
func Infow(keyvals ...interface{}) {
	handler.Infow(keyvals...)
}

// Warn .
func Warn(a ...interface{}) {
	handler.Warn(a...)
}

// Warnf .
func Warnf(format string, a ...interface{}) {
	handler.Warnf(format, a...)
}

// Warnw .
func Warnw(keyvals ...interface{}) {
	handler.Warnw(keyvals...)
}

// Error .
func Error(a ...interface{}) {
	handler.Error(a...)
}

// Errorf .
func Errorf(format string, a ...interface{}) {
	handler.Errorf(format, a...)
}

// Errorw .
func Errorw(keyvals ...interface{}) {
	handler.Errorw(keyvals...)
}

// Fatal .
func Fatal(a ...interface{}) {
	handler.Fatal(a...)
}

// Fatalf .
func Fatalf(format string, a ...interface{}) {
	handler.Fatalf(format, a...)
}

// Fatalw .
func Fatalw(keyvals ...interface{}) {
	handler.Fatalw(keyvals...)
}
