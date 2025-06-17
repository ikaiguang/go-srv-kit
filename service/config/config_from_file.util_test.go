package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"testing"
)

// go test -v -count 1 ./config/ -run TestLoadingFile
func TestLoadingFile(t *testing.T) {
	type args struct {
		filePath string
		opts     []Option
	}

	otherConfig := &configpb.TestingConfig{}
	tests := []struct {
		name            string
		args            args
		want            *configpb.Bootstrap
		wantErr         bool
		withOtherConfig bool
	}{
		{
			name: "#TestLoadingFile",
			args: args{
				filePath: "config_example.yaml",
				opts:     []Option{WithOtherConfig(otherConfig)},
			},
			want:            nil,
			wantErr:         false,
			withOtherConfig: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadingFile(tt.args.filePath, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadingFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("LoadingFile() got = %v, want %v", got, tt.want)
			//}
			t.Logf("Boostrap.App: %#v\n", got.App)
			if tt.withOtherConfig {
				t.Logf("OtherConfig: %#v\n", otherConfig.GetTestdata())
			}
		})
	}
}
