package downloadpkg

import (
	"context"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	TestdataPath = "./_output_testdata"
)

// go test -v -count 1 ./kit/download -run TestStreamDownload
func TestStreamDownload(t *testing.T) {
	type args struct {
		ctx   context.Context
		param *DownloadParam
	}
	tests := []struct {
		name    string
		args    args
		want    *DownloadReply
		wantErr bool
	}{
		{
			name: "test stream download",
			args: args{
				ctx: context.Background(),
				param: &DownloadParam{
					URL:        "http://a.com/test.mp4",
					OutputPath: filepath.Join(TestdataPath, "test.mp4"),
				},
			},
			want: &DownloadReply{
				FilePath: filepath.Join(TestdataPath, "test.mp4"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StreamDownload(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("StreamDownload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamDownload() got = %v, want %v", got, tt.want)
			}
		})
	}
}
