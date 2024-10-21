package loggerutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"google.golang.org/protobuf/types/known/durationpb"
	"os"
	"testing"
	"time"
)

var (
	appConfig = &configpb.App{
		ProjectName:   "go-micro-saas",
		ServerName:    "test-service",
		ServerEnv:     "develop",
		ServerVersion: "v1.0.0",
		Metadata:      nil,
	}
	logConfig = &configpb.Log{
		Console: &configpb.Log_Console{
			Enable: true,
			Level:  "DEBUG",
		},
		File: &configpb.Log_File{
			Enable:         true,
			Level:          "DEBUG",
			Dir:            "./runtime/logs",
			Filename:       "test",
			RotateTime:     durationpb.New(time.Hour * 24),
			RotateSize:     52428800,
			StorageAge:     durationpb.New(time.Hour * 24 * 30),
			StorageCounter: 10086,
		},
	}
	handler LoggerManager
)

func TestMain(m *testing.M) {
	var err error
	handler, err = NewLoggerManager(logConfig, appConfig)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
