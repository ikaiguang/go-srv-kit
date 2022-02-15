package loghelper

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"

	"github.com/ikaiguang/go-srv-kit/log"
)

// go test -v ./log/helper/ -count=1 -test.run=TestSetup_Xxx
func TestSetup_Xxx(t *testing.T) {
	tests := []struct {
		name       string
		addSkip    int
		callerSkip int
		hasWith    bool
		isMulti    bool
	}{
		{
			name:       "#Setup_OneLogger",
			addSkip:    1,
			callerSkip: 2,
			hasWith:    false,
			isMulti:    false,
		},
		{
			name:       "#Setup_OneLogger_With",
			addSkip:    2,
			callerSkip: 2,
			hasWith:    true,
			isMulti:    false,
		},
		{
			name:       "#Setup_MultiLogger",
			addSkip:    2,
			callerSkip: 2,
			hasWith:    false,
			isMulti:    true,
		},
		{
			name:       "#Setup_MultiLogger_With",
			addSkip:    2,
			callerSkip: 2,
			hasWith:    true,
			isMulti:    true,
		},
	}

	var (
		oneLoggerFn = func(addSkip int) (log.Logger, func() error, error) {
			stdLoggerConfig := &logutil.ConfigStd{
				Level:      log.LevelDebug,
				CallerSkip: logutil.DefaultCallerSkip + addSkip,
			}
			// 在for中Sync
			logger, err := logutil.NewStdLogger(stdLoggerConfig)
			if err != nil {
				return logger, nil, err
			}
			return logger, logger.Sync, err
		}
		multiLoggerFn = func(addSkip int) (log.Logger, []func() error, error) {
			var closeFnSlice []func() error
			logger1, syncFn1, err := oneLoggerFn(addSkip)
			if err != nil {
				return logger1, closeFnSlice, err
			}
			closeFnSlice = append(closeFnSlice, syncFn1)

			logger2, syncFn2, err := oneLoggerFn(addSkip)
			if err != nil {
				return logger2, closeFnSlice, err
			}
			closeFnSlice = append(closeFnSlice, syncFn2)

			return log.MultiLogger(logger1, logger2), closeFnSlice, nil
		}
	)
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				logger      log.Logger
				syncFn      func() error
				syncFnSlice []func() error
				err         error
			)
			if tt.isMulti {
				logger, syncFnSlice, err = multiLoggerFn(tt.addSkip)
			} else {
				logger, syncFn, err = oneLoggerFn(tt.addSkip)
				syncFnSlice = append(syncFnSlice, syncFn)
			}
			require.Nil(t, err)

			if tt.hasWith {
				logger = log.With(logger, "caller", log.Caller(logutil.DefaultCallerValuer+tt.callerSkip))
			}

			Setup(logger)

			Infof("第 %d 个", i+1)

			for fnIndex := range syncFnSlice {
				_ = syncFnSlice[fnIndex]()
			}
		})
	}
}

// go test -v ./log/helper/ -count=1 -test.run=TestSetup_OneLogger_Xxx
func TestSetup_OneLogger_Xxx(t *testing.T) {
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip + 1,
	}
	stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLogger.Sync() }()

	// CallerSkip: DefaultCallerSkip + 2,
	//stdLogger = log.With(stdLogger, "caller", log.Caller(DefaultCallerValuer+2))

	Setup(stdLogger)

	Debug("TestSetup_OneLogger Then Debug")
}

// go test -v ./log/helper/ -count=1 -test.run=TestSetup_OneLogger_With
func TestSetup_OneLogger_With(t *testing.T) {
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip + 2,
	}
	stdLoggerHandler, err := logutil.NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLoggerHandler.Sync() }()

	// CallerSkip: DefaultCallerSkip + 2,
	var stdLogger log.Logger = stdLoggerHandler
	stdLogger = log.With(stdLogger, "caller", log.Caller(logutil.DefaultCallerValuer+2))

	Setup(stdLogger)

	Debug("TestSetup_OneLogger_With Then Debug")
}

// go test -v ./log/helper/ -count=1 -test.run=TestSetup_MultiLogger
func TestSetup_MultiLogger(t *testing.T) {
	stdLoggerConfig := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip + 2, // +2
	}
	stdLogger, err := logutil.NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLogger.Sync() }()

	// two
	stdLoggerConfig2 := &logutil.ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: logutil.DefaultCallerSkip + 2, // +2
	}
	stdLogger2, err := logutil.NewStdLogger(stdLoggerConfig2)
	require.Nil(t, err)
	defer func() { _ = stdLogger2.Sync() }()

	multiLogger := log.MultiLogger(stdLogger, stdLogger2)
	multiLogger = log.With(multiLogger, "caller", log.Caller(logutil.DefaultCallerValuer+2))

	Setup(multiLogger)

	Debug("TestSetup_MultiLogger Then Debug")
}
