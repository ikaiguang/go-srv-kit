package configutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	"testing"
)

// go test -v -count 1 ./config/ -run TestLoadingConfigFromConsul
func TestLoadingConfigFromConsul(t *testing.T) {
	type args struct {
		appConfig *configpb.App
		opts      []Option
	}

	otherConfig := &configpb.TestingConfig{}
	appConfig := &configpb.App{
		ConfigMethod:         CONFIG_METHOD_CONSUL,
		ConfigPathForGeneral: "go-micro-saas/general-config",
		ConfigPathForServer:  "go-micro-saas/ping-service/production/v1.0.0",
	}
	tests := []struct {
		name            string
		args            args
		want            *configpb.Bootstrap
		wantErr         bool
		withOtherConfig bool
	}{
		{
			name: "#loadingForConsul", args: args{
				appConfig: appConfig,
				opts:      []Option{WithOtherConfig(otherConfig)},
			},
			want:            nil,
			wantErr:         false,
			withOtherConfig: true,
		},
	}

	cfg := &configpb.Consul{
		Enable:             true,
		Address:            "127.0.0.1:8500",
		InsecureSkipVerify: true,
	}
	cc, err := newConsulClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadingConfigFromConsul(cc, tt.args.appConfig, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadingConfigFromConsul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("LoadingConfigFromConsul() got = %v, want %v", got, tt.want)
			//}
			if got.GetApp().GetServerName() == "" {
				t.Fatal("==> got.GetApp().GetServerName() is empty")
			}
			t.Log("==> got.GetApp().GetServerName(): ", got.GetApp().GetServerName())
			if tt.withOtherConfig {
				t.Logf("OtherConfig: %#v\n", otherConfig.GetTestdata())
			}
		})
	}
}
