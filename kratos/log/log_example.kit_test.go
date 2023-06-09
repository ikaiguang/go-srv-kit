package logpkg

import (
	"github.com/go-kratos/kratos/v2/log"
)

func ExampleNewMultiLogger() {
	// 查看 TestNewMultiLogger(nil)
	// std logger
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = stdLogger.Close() }()

	// file logger
	fileLoggerConfig := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	fileLogger, err := NewFileLogger(fileLoggerConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = fileLogger.Close() }()

	logger := NewMultiLogger(stdLogger, fileLogger)
	logHandler := log.NewHelper(logger)
	logHandler.Error("error")
}

func ExampleNewFileLogger() {
	// 查看 TestNewFileLogger(nil)
	// file logger
	fileLoggerConfig := &ConfigFile{
		Level: log.LevelDebug,
		// CallerSkip = DefaultCallerSkip
		CallerSkip: DefaultCallerSkip,

		Dir:      "./../bin/",
		Filename: "rotation",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		StorageCounter: 2,
		//StorageAge: time.Hour,
	}
	fileLogger, err := NewFileLogger(fileLoggerConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = fileLogger.Close() }()

	logHandler := log.NewHelper(fileLogger)
	logHandler.Error("error")
}

func ExampleNewStdLogger() {
	// 查看 TestNewStdLogger(nil)
	// std logger
	stdLoggerConfig := &ConfigStd{
		Level: log.LevelDebug,
		// CallerSkip = DefaultCallerSkip
		CallerSkip: DefaultCallerSkip,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	if err != nil {
		panic(err)
	}
	defer func() { _ = stdLogger.Close() }()

	logHandler := log.NewHelper(stdLogger)
	logHandler.Error("error")
}
