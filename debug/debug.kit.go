package debuguitl

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
