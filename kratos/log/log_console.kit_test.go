package logpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -run TestNewStdLogger
func TestNewStdLogger(t *testing.T) {
	cfg := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip,
	}
	logImpl, err := NewStdLogger(cfg)
	require.Nil(t, err)
	defer func() { _ = logImpl.Close() }()

	Setup(logImpl)
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
	/*
		2021-07-28T10:01:48.915 ERROR   log/log.toolkit_test.go:17      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "log level error"}
		2021-07-28T10:01:48.915 DEBUG   log/log.toolkit_test.go:18      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "log level debug"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:19      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "log level info"}
		2021-07-28T10:01:48.915 ERROR   log/log.toolkit_test.go:20      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "log level error"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:21      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "ab"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:22      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "abc"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:23      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"msg": "ab"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:25      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"key": "value"}
		2021-07-28T10:01:48.915 INFO    log/log.toolkit_test.go:26      github.com/ikaiguang/go-srv-toolkit/log.TestNewStdLogger
		        {"key": "value", "remain": "KEYVALS UNPAIRED"}
	*/
}

// go test -v ./log/ -count=1 -run TestKratos_NewStdLogger
//func TestKratos_NewStdLogger(t *testing.T) {
//	logImpl := log.NewStdLogger(os.Stderr)
//
//	Setup(logImpl)
//
//	Error("log level error")
//	Debug("log level debug")
//	Info("log level info")
//	Error("log level error")
//	Info("a", "b")
//	Info("a", "b", "c")
//	Infof("%s%s", "a", "b")
//	// [build failed] Infof call needs 2 args but has 3 args
//	//Infof("%s%s", "a", "b", "c")
//	Infow("key", "value")
//	Infow("key", "value", "remain")
//	/*
//		ERROR msg=log level error
//		DEBUG msg=log level debug
//		INFO msg=log level info
//		ERROR msg=log level error
//		INFO msg=ab
//		INFO msg=abc
//		INFO msg=ab
//		INFO key=value
//		INFO key=value remain=KEYVALS UNPAIRED
//	*/
//}
