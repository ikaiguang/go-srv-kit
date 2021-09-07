package logutil

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -test.run=TestNewFileLogger
func TestNewFileLogger(t *testing.T) {
	cfg := &ConfigFile{
		Enable:     true,
		Level:      log.LevelDebug,
		CallerSkip: 0,

		Dir:      "./../bin/",
		Filename: "rotation",

		//RotateTime: time.Second * 1,
		RotateSize: 50 << 20, // 50M : 50 << 20

		StorageCounter: 2,
		//StorageAge: time.Hour,
	}
	logImpl, err := NewFileLogger(cfg)
	require.Nil(t, err)

	logHandler := log.NewHelper(logImpl)

	logHandler.Error("log level error")
	logHandler.Debug("log level debug")
	logHandler.Info("log level info")
	logHandler.Error("log level error")
	logHandler.Info("a", "b")
	logHandler.Info("a", "b", "c")
	logHandler.Infof("%s%s", "a", "b")
	//Infof("%s%s", "a", "b", "c") // [build failed]
	logHandler.Infow("key", "value")
	logHandler.Infow("key", "value", "remain")
}
