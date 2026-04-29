package jaegerpkg

import (
	"os"
	"testing"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"google.golang.org/protobuf/types/known/durationpb"
)

func getTestJaegerGRPCAddr() string {
	if addr := os.Getenv("JAEGER_GRPC_ADDR"); addr != "" {
		return addr
	}
	return "my-jaeger:4317"
}

func getTestJaegerHTTPAddr() string {
	if addr := os.Getenv("JAEGER_HTTP_ADDR"); addr != "" {
		return addr
	}
	return "my-jaeger:4318"
}

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
					Kind:              KindGRPC,
					Addr:              getTestJaegerGRPCAddr(),
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
					Kind:              KindGRPC,
					Addr:              getTestJaegerHTTPAddr(),
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
