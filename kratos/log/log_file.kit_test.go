package logpkg

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
	"github.com/stretchr/testify/require"

	writerpkg "github.com/ikaiguang/go-srv-kit/kit/writer"
)

// go test -v ./log/ -count=1 -run TestNewFileLogger_Xxx
func TestNewFileLogger_Xxx(t *testing.T) {
	cfg := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip,

		Dir:      "./runtime/logs",
		Filename: "rotation",

		RotateTime: time.Second * 1,
		//RotateSize: 50 << 20, // 50M : 50 << 20

		StorageCounter: 2,
		//StorageAge: time.Hour,
	}
	logImpl, err := NewFileLogger(
		cfg,
		//WithFilenameSuffix("_xxx"),
		WithFilenameSuffix("_xxx.%Y%m%d%H%M%S.log"),
		WithLoggerKey(map[LoggerKey]string{LoggerKeyTime: "date"}),
		WithTimeFormat(timepkg.YmdHmsMillisecond),
	)
	require.Nil(t, err)
	defer func() { _ = logImpl.Close() }()

	logHandler := log.NewHelper(logImpl)
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

// go test -v ./log/ -count=1 -run TestNewFileLogger_WithWriter
func TestNewFileLogger_WithWriter(t *testing.T) {
	cfg := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip,

		Dir:      "./../bin/",
		Filename: "writer",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		StorageCounter: 2,
		//StorageAge: time.Hour,
	}
	writerConfig := &writerpkg.ConfigRotate{
		Dir:            cfg.Dir,
		Filename:       cfg.Filename,
		RotateTime:     time.Second,
		StorageCounter: 3,
	}
	writer, err := writerpkg.NewRotateFile(writerConfig)
	require.Nil(t, err)

	logImpl, err := NewFileLogger(
		cfg,
		WithWriter(writer),
		//WithFilenameSuffix("_xxx"),
		//WithFilenameSuffix("_xxx.%Y%m%d%H%M%S.log"),
		WithLoggerKey(map[LoggerKey]string{LoggerKeyTime: "date"}),
		WithTimeFormat(timepkg.YmdHmsMillisecond),
	)
	require.Nil(t, err)
	defer func() { _ = logImpl.Close() }()
	logHandler := log.NewHelper(logImpl)

	total := int(writerConfig.StorageCounter + 1)
	for i := 0; i < total; i++ {
		str := fmt.Sprintf("第 %d 行", i+1)
		logHandler.Info(str)

		time.Sleep(time.Second)
	}
}
