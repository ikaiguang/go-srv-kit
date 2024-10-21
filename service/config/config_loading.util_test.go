package configutil

import (
	configpb "github.com/go-micro-saas/service-kit/api/config"
	"os"
	"testing"
)

// go test -v -count=1 ./config/ -test.run=TestLoading_Config
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

// go test -v -count=1 ./config/ -test.run=TestCurrentPath
func TestCurrentPath(t *testing.T) {
	// get $GOPATH
	gopath := os.Getenv("GOPATH")
	// get $GOPATH/src/github.com/go-micro-saas/service-kit/config
	tests := []struct {
		name string
		want string
	}{
		{
			name: "#TestCurrentPath",
			want: gopath + "/src/github.com/go-micro-saas/service-kit/config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentPath(); got != tt.want {
				t.Errorf("CurrentPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
