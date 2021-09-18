package logutil

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -test.run=TestSetup_Xxx
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
		oneLoggerFn = func(addSkip int) (log.Logger, error) {
			stdLoggerConfig := &ConfigStd{
				Level:      log.LevelDebug,
				CallerSkip: DefaultCallerSkip + addSkip,
			}
			return NewStdLogger(stdLoggerConfig)
		}
		multiLoggerFn = func(addSkip int) (log.Logger, error) {
			logger1, err := oneLoggerFn(addSkip)
			if err != nil {
				return nil, err
			}
			logger2, err := oneLoggerFn(addSkip)
			if err != nil {
				return nil, err
			}
			return log.MultiLogger(logger1, logger2), nil
		}
	)
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				logger log.Logger
				err    error
			)
			if tt.isMulti {
				logger, err = multiLoggerFn(tt.addSkip)
			} else {
				logger, err = oneLoggerFn(tt.addSkip)
			}
			require.Nil(t, err)
			if tt.hasWith {
				logger = log.With(logger, "caller", log.Caller(DefaultCallerValuer+tt.callerSkip))
			}

			Setup(logger)

			Infof("第 %d 个", i+1)
		})
	}
}

// go test -v ./log/ -count=1 -test.run=TestSetup_OneLogger_Xxx
func TestSetup_OneLogger_Xxx(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)

	// CallerSkip: DefaultCallerSkip + 2,
	//stdLogger = log.With(stdLogger, "caller", log.Caller(DefaultCallerValuer+2))

	Setup(stdLogger)

	Debug("TestSetup_OneLogger Then Debug")
}

// go test -v ./log/ -count=1 -test.run=TestSetup_OneLogger_With
func TestSetup_OneLogger_With(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)

	// CallerSkip: DefaultCallerSkip + 2,
	stdLogger = log.With(stdLogger, "caller", log.Caller(DefaultCallerValuer+2))

	Setup(stdLogger)

	Debug("TestSetup_OneLogger_With Then Debug")
}

// go test -v ./log/ -count=1 -test.run=TestSetup_MultiLogger
func TestSetup_MultiLogger(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2, // +2
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)

	// two
	stdLoggerConfig2 := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2, // +2
	}
	stdLogger2, err := NewStdLogger(stdLoggerConfig2)
	require.Nil(t, err)

	multiLogger := log.MultiLogger(stdLogger, stdLogger2)
	multiLogger = log.With(multiLogger, "caller", log.Caller(DefaultCallerValuer+2))

	Setup(multiLogger)

	Debug("TestSetup_MultiLogger Then Debug")
}
