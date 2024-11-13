package loggerutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
)

var (
	_defaultAppConfig = configpb.App{
		ProjectName:   "go-srv-kit",
		ServerName:    "test-service",
		ServerEnv:     "develop",
		ServerVersion: "v1.0.0",
		Metadata:      nil,
	}
	_defaultLogConfig = &configpb.Log{
		Console: &configpb.Log_Console{
			Enable: true,
			Level:  "DEBUG",
		},
		File: &configpb.Log_File{
			Enable:         false,
			Level:          "DEBUG",
			Dir:            "./runtime/logs",
			Filename:       "test",
			RotateTime:     durationpb.New(time.Hour * 24),
			RotateSize:     52428800,
			StorageAge:     durationpb.New(time.Hour * 24 * 30),
			StorageCounter: 10086,
		},
	}
)
