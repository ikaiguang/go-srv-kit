package logutil

import (
	"github.com/go-kratos/kratos/v2/log"
)

func ExampleNewMultiLogger() {
	// std logger
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	if err != nil {
		panic(err)
	}

	// file logger
	fileLoggerConfig := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	fileLogger, err := NewFileLogger(fileLoggerConfig)
	if err != nil {
		panic(err)
	}

	logger := NewMultiLogger(stdLogger, fileLogger)
	logHandler := log.NewHelper(logger)
	logHandler.Error("error")
}

func ExampleNewFileLogger() {
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

	logHandler := log.NewHelper(fileLogger)
	logHandler.Error("error")
}

func ExampleNewStdLogger() {
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

	logHandler := log.NewHelper(stdLogger)
	logHandler.Error("error")
}
