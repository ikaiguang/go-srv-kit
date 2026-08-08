package mongopkg

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	dbConfig = &Config{
		Debug:             true,
		AppName:           "mongo:test",
		Hosts:             nil,
		Addr:              os.Getenv("MONGO_TEST_URI"),
		MaxPoolSize:       100,
		MinPoolSize:       2,
		MaxConnecting:     10,
		ConnectTimeout:    durationpb.New(time.Second * 3),
		Timeout:           durationpb.New(time.Second * 3),
		HeartbeatInterval: durationpb.New(time.Second * 3),
		MaxConnIdleTime:   durationpb.New(time.Second * 60),
		SlowThreshold:     durationpb.New(time.Millisecond * 100),
	}
)

// go test -v ./data/mongo/ -count=1 -run TestNewMongoClient
func TestNewMongoClient(t *testing.T) {
	if dbConfig.Addr == "" {
		t.Skip("set MONGO_TEST_URI to run the MongoDB integration test")
	}

	type args struct {
		config *Config
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    *mongo.Client
		wantErr bool
	}{
		{
			name: "#testNewClient",
			args: args{
				config: dbConfig,
				logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoClient(tt.args.config, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewMongoClient() got = %v, want %v", got, tt.want)
			//}
			defer func() { _ = got.Disconnect(context.Background()) }()
		})
	}
}

func TestNewMongoClientRejectsNilConfig(t *testing.T) {
	client, err := NewMongoClient(nil, nil)
	if err == nil {
		t.Fatal("NewMongoClient(nil, nil) error = nil, want error")
	}
	if client != nil {
		t.Fatalf("NewMongoClient(nil, nil) client = %v, want nil", client)
	}
}
