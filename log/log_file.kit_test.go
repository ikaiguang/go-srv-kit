package logutil

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -test.run=TestNewFileLogger
func TestNewFileLogger(t *testing.T) {
	cfg := &ConfigFile{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip,

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
	// [build failed] Infof call needs 2 args but has 3 args
	//Infof("%s%s", "a", "b", "c")
	logHandler.Infow("key", "value")
	logHandler.Infow("key", "value", "remain")
}
