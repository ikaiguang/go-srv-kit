package jaegerpkg

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"google.golang.org/protobuf/types/known/durationpb"
	"testing"
	"time"
)

// go test -v ./data/jaeger/ -count=1 -run TestNewJaegerExporter_Xxx
func TestNewJaegerExporter_Xxx(t *testing.T) {
	type args struct {
		conf *Config
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *otlptrace.Exporter
		wantErr bool
	}{
		{
			name: "#NewJaegerExporter_GRPC",
			args: args{
				conf: &Config{
					Kind:              KingGRPC,
					Addr:              "my-jaeger-hostname:4317",
					IsInsecure:        true,
					Timeout:           durationpb.New(time.Second * 30),
					WithHttpBasicAuth: false,
					Username:          "",
					Password:          "",
				},
				opts: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "#NewJaegerExporter_HTTP",
			args: args{
				conf: &Config{
					Kind:              KingGRPC,
					Addr:              "my-jaeger-hostname:4318",
					IsInsecure:        true,
					Timeout:           durationpb.New(time.Second * 30),
					WithHttpBasicAuth: false,
					Username:          "",
					Password:          "",
				},
				opts: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJaegerExporter(tt.args.conf, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJaegerExporter() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewJaegerExporter() got = %v, want %v", got, tt.want)
			//}
			t.Logf("==> got: %#v\n", got)
		})
	}
}
