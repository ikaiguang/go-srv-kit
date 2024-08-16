package mongopkg

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/durationpb"
	"testing"
	"time"
)

var (
	dbConfig = &Config{
		Debug:             true,
		AppName:           "mongo:test",
		Hosts:             nil,
		Addr:              "mongodb://mongo:Mongo.123456@my-mongo-hostname:27017/admin",
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

// go test -v ./data/mongo/ -count=1 -test.run=TestNewMongoClient
func TestNewMongoClient(t *testing.T) {
	type args struct {
		config *Config
		logger log.Logger
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
				logger: log.DefaultLogger,
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
