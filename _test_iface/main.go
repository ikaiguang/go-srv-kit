package main

type Level1 int8

type LoggerAny interface {
	Log(level any, keyvals ...any) error
}

type myLogger struct{}

func (m *myLogger) Log(level Level1, keyvals ...any) error { return nil }

func main() {
	var l LoggerAny = &myLogger{}
	_ = l
}
