package logutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"

	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"
	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
)

// go test -v ./log/ -count=1 -test.run=TestNewFileLogger_Xxx
func TestNewFileLogger_Xxx(t *testing.T) {
	cfg := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip,

		Dir:      "./../bin/",
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
		WithTimeFormat(timeutil.YmdHmsMillisecond),
	)
	require.Nil(t, err)
	defer func() { _ = logImpl.Sync() }()

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

// go test -v ./log/ -count=1 -test.run=TestNewFileLogger_WithWriter
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
	writerConfig := &writerutil.ConfigRotate{
		Dir:            cfg.Dir,
		Filename:       cfg.Filename,
		RotateTime:     time.Second,
		StorageCounter: 3,
	}
	writer, err := writerutil.NewRotateFile(writerConfig)
	require.Nil(t, err)

	logImpl, err := NewFileLogger(
		cfg,
		WithWriter(writer),
		//WithFilenameSuffix("_xxx"),
		//WithFilenameSuffix("_xxx.%Y%m%d%H%M%S.log"),
		WithLoggerKey(map[LoggerKey]string{LoggerKeyTime: "date"}),
		WithTimeFormat(timeutil.YmdHmsMillisecond),
	)
	require.Nil(t, err)
	defer func() { _ = logImpl.Sync() }()
	logHandler := log.NewHelper(logImpl)

	total := int(writerConfig.StorageCounter + 1)
	for i := 0; i < total; i++ {
		str := fmt.Sprintf("第 %d 行", i+1)
		logHandler.Info(str)

		time.Sleep(time.Second)
	}
}
