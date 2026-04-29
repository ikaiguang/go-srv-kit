package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"strings"
	"testing"
)

// go test -v -count 1 ./config/ -run TestLoading_Config
func TestLoading_Config(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *configpb.Bootstrap
		wantErr bool
	}{
		{
			name:    "#testLoadingConfig",
			args:    args{filePath: "config_example.yaml"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Loading(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Loading() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Loading() got = %v, want %v", got, tt.want)
			//}
			if got.GetApp().GetServerName() == "" {
				t.Fatal("==> got.GetApp().GetServerName() is empty")
			}
			t.Log("==> got.GetApp().GetServerName(): ", got.GetApp().GetServerName())
		})
	}
}

// go test -v -count 1 ./config/ -run TestCurrentPath
func TestCurrentPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "#TestCurrentPath",
			want: "/service/config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.ReplaceAll(CurrentPath(), "\\", "/")
			if !strings.HasSuffix(got, tt.want) {
				t.Errorf("CurrentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
