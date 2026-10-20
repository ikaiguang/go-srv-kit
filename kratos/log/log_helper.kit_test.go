package logpkg

import (
	"io"
	"log/slog"
	"sync"
	"testing"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/stretchr/testify/require"
)

func TestSetupConcurrentLogging(_ *testing.T) {
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range 100 {
				Setup(NewNopLogger())
			}
		}()
		go func() {
			defer wg.Done()
			for range 100 {
				Info("concurrent")
			}
		}()
	}
	wg.Wait()
}

// go test -v -count 1 ./log/helper/ -run TestSetup_Xxx
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
			addSkip:    0,
			callerSkip: 0,
			hasWith:    false,
			isMulti:    false,
		},
		{
			name:       "#Setup_OneLogger_With",
			addSkip:    1,
			callerSkip: 1,
			hasWith:    true,
			isMulti:    false,
		},
		{
			name:       "#Setup_MultiLogger",
			addSkip:    1,
			callerSkip: 1,
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
		oneLoggerFn = func(addSkip int) (Logger, io.Closer, error) {
			stdLoggerConfig := &ConfigStd{
				Level:      log.LevelDebug,
				CallerSkip: DefaultCallerSkip + addSkip,
			}
			// 在for中Sync
			logger, err := NewStdLogger(stdLoggerConfig)
			if err != nil {
				return logger, nil, err
			}
			return logger, logger, err
		}
		multiLoggerFn = func(addSkip int) (Logger, []io.Closer, error) {
			var closeFnSlice []io.Closer
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

			return NewMultiLogger(logger1, logger2), closeFnSlice, nil
		}
	)
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				logger      Logger
				syncFn      io.Closer
				syncFnSlice []io.Closer
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
				logger = slog.New(logger.Handler()).With("caller_skip", DefaultCallerValuer+tt.callerSkip)
			}

			Setup(logger)

			Infof("第 %d 个", i+1)

			for fnIndex := range syncFnSlice {
				_ = syncFnSlice[fnIndex].Close()
			}
		})
	}
}

// go test -v ./log/helper/ -count=1 -run TestSetup_OneLogger_Xxx
func TestSetup_OneLogger_Xxx(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 1,
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLogger.Close() }()

	// CallerSkip: DefaultCallerSkip + 2,
	//stdLogger = log.With(stdLogger, "caller", log.Caller(DefaultCallerValuer+2))

	Setup(stdLogger)

	Debug("TestSetup_OneLogger Then Debug")
}

// go test -v ./log/helper/ -count=1 -run TestSetup_OneLogger_With
func TestSetup_OneLogger_With(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2,
	}
	stdLoggerHandler, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLoggerHandler.Close() }()

	// CallerSkip: DefaultCallerSkip + 2,
	var stdLogger Logger = stdLoggerHandler
	stdLogger = slog.New(stdLogger.Handler()).With("caller_skip", DefaultCallerValuer+2)

	Setup(stdLogger)

	Debug("TestSetup_OneLogger_With Then Debug")
}

// go test -v ./log/helper/ -count=1 -run TestSetup_MultiLogger
func TestSetup_MultiLogger(t *testing.T) {
	stdLoggerConfig := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2, // +2
	}
	stdLogger, err := NewStdLogger(stdLoggerConfig)
	require.Nil(t, err)
	defer func() { _ = stdLogger.Close() }()

	// two
	stdLoggerConfig2 := &ConfigStd{
		Level:      log.LevelDebug,
		CallerSkip: DefaultCallerSkip + 2, // +2
	}
	stdLogger2, err := NewStdLogger(stdLoggerConfig2)
	require.Nil(t, err)
	defer func() { _ = stdLogger2.Close() }()

	multiLogger := NewMultiLogger(stdLogger, stdLogger2)
	multiLogger = multiLogger.With("caller_skip", DefaultCallerValuer+2)

	Setup(multiLogger)

	Debug("TestSetup_MultiLogger Then Debug")
}
