package logpkg

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestNewMultiLoggerWritesToEveryHandler(t *testing.T) {
	var first bytes.Buffer
	var second bytes.Buffer
	logger := NewMultiLogger(
		slog.New(slog.NewJSONHandler(&first, nil)),
		slog.New(slog.NewJSONHandler(&second, nil)),
	)

	logger.Info("ready", "module", "kratos")

	require.Contains(t, first.String(), `"msg":"ready"`)
	require.Contains(t, second.String(), `"module":"kratos"`)
}

func TestZapHandlerPreservesSlogGroups(t *testing.T) {
	core, observed := observer.New(zap.DebugLevel)
	logger := slog.New(&zapHandler{logger: zap.New(core)})

	logger.WithGroup("request").Info("received", "id", 7)

	require.Equal(t, int64(7), observed.All()[0].ContextMap()["request.id"])
}

func TestLoggerConstructorsRejectNilConfig(t *testing.T) {
	stdLogger, err := NewStdLogger(nil)
	require.Nil(t, stdLogger)
	require.ErrorIs(t, err, errNilConfig)

	fileLogger, err := NewFileLogger(nil)
	require.Nil(t, fileLogger)
	require.ErrorIs(t, err, errNilConfig)
}

func TestToSlogAcceptsTypedNilLogger(t *testing.T) {
	var logger *slog.Logger
	require.NotNil(t, ToSlog(logger))
	var stdLogger *Std
	require.NotNil(t, ToSlog(stdLogger))
}

// go test -v ./log/ -count=1 -run TestNewMultiLogger
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

	Setup(logger)
	Error("log level error")
	Debug("log level debug")
	Info("log level info")
	Error("log level error")
	Info("a", "b")
	Info("a", "b", "c")
	Infof("%s%s", "a", "b")
	// [build failed] Infof call needs 2 args but has 3 args
	//Infof("%s%s", "a", "b", "c")
	Infow("key", "value")
	Infow("key", "value", "remain")
}

// go test -v ./log/ -count=1 -run TestParseLevel
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
