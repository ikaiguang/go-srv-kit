package debugutil

// Print .
func Print(msg interface{}) {
	handler.Debug(msg)
}

// Println .
func Println(msg interface{}) {
	handler.Debug(msg)
}

// Printf .
func Printf(format string, a ...interface{}) {
	handler.Debugf(format, a...)
}

// Printw .
func Printw(keyvals ...interface{}) {
	handler.Fatalw(keyvals...)
}

// Debug .
func Debug(msg interface{}) {
	handler.Debug(msg)
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
func Info(msg interface{}) {
	handler.Info(msg)
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
func Warn(msg interface{}) {
	handler.Warn(msg)
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
func Error(msg interface{}) {
	handler.Error(msg)
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
func Fatal(msg interface{}) {
	handler.Fatal(msg)
}

// Fatalf .
func Fatalf(format string, a ...interface{}) {
	handler.Fatalf(format, a...)
}

// Fatalw .
func Fatalw(keyvals ...interface{}) {
	handler.Fatalw(keyvals...)
}
