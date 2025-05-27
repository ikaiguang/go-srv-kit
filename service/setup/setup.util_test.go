package setuputil

import (
	configtestdata "github.com/ikaiguang/go-srv-kit/testdata/configs"
	"testing"
)

// go test -v -count=1 ./setup/ -run TestNewLauncherManager
func TestNewLauncherManager(t *testing.T) {
	confPath := configtestdata.ConfigPath()
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "#testingSetup",
			args:    args{configFilePath: confPath},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLauncherManager(tt.args.configFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %+v, wantErr %v", err, tt.wantErr)
				t.FailNow()
			}
			err = Close(got)
			if err != nil {
				t.Errorf("Close() error = %+v", err)
				t.FailNow()
			}
		})
	}
}
