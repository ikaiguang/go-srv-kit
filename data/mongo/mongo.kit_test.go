//go:build ignore

package mongopkg

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/protobuf/types/known/durationpb"
)

func getTestMongoURI() string {
	if uri := os.Getenv("DB_MONGO_URI"); uri != "" {
		return uri
	}
	return "mongodb://mongo:Mongo.123456@my-mongo:27017/admin"
}

var (
	dbConfig = &Config{
		Debug:             true,
		AppName:           "mongo:test",
		Hosts:             nil,
		Addr:              getTestMongoURI(),
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

// kratosLoggerAdapter 将 kratos log.Logger 适配为本包的 Logger 接口
func kratosLoggerAdapter(l log.Logger) Logger {
	return LogAdapter(func(level Level, keyvals ...any) error {
		return l.Log(log.Level(level), keyvals...)
	})
}

// go test -v ./data/mongo/ -count=1 -run TestNewMongoClient
func TestNewMongoClient(t *testing.T) {
	type args struct {
		config *Config
		logger Logger
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
				logger: kratosLoggerAdapter(log.DefaultLogger),
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
