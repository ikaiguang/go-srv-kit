package loggerutil

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

// go test -v -count=1 ./setup/logger -test.run=Test_loggerManager_GetLoggers
func Test_loggerManager_GetLoggers(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "#info",
			msg:     "info",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := handler.GetLogger()
			if err != nil {
				t.Errorf("GetLogger() err = %v\n", err)
				t.FailNow()
			}
			logHandler := log.NewHelper(logger)
			logHandler.Info(tt.msg)
		})
	}
}
