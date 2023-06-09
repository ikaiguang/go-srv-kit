package logpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -test.run=TestNewMultiLogger
func TestNewMultiLogger(t *testing.T) {
	// std logger
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLogger.Close() }()

	// file logger
	fileLoggerConfig := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,

		Dir:      "./../bin/",
		Filename: "rotation",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		StorageCounter: 2,
		//StorageAge: time.Hour,
	}
	fileLogger, err := NewFileLogger(fileLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = fileLogger.Close() }()

	logger := NewMultiLogger(stdLogger, fileLogger)

	logHandler := log.NewHelper(logger)
	logHandler.Error("log level error")
	logHandler.Debug("log level debug")
	logHandler.Info("log level info")
	logHandler.Error("log level error")
	logHandler.Info("a", "b")
	logHandler.Info("a", "b", "c")
	logHandler.Infof("%s%s", "a", "b")
	// [build failed] Infof call needs 2 args but has 3 args
	//Infof("%s%s", "a", "b", "c")
	logHandler.Infow("key", "value")
	logHandler.Infow("key", "value", "remain")
}

// go test -v ./log/ -count=1 -test.run=TestParseLevel
func TestParseLevel(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  log.Level
	}{
		{
			name:  "#DEBUG",
			level: "DEBUG",
			want:  log.LevelDebug,
		},
		{
			name:  "#INFO",
			level: "INFO",
			want:  log.LevelInfo,
		},
		{
			name:  "#WARN",
			level: "WARN",
			want:  log.LevelWarn,
		},
		{
			name:  "#ERROR",
			level: "ERROR",
			want:  log.LevelError,
		},
		{
			name:  "#FATAL",
			level: "FATAL",
			want:  log.LevelFatal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lv := ParseLevel(tt.level)
			require.Equal(t, tt.want, lv)
		})
	}
}
