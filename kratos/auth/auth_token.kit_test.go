package authpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

// go test -v -count 1 ./kratos/auth -run TestNewAuthRepo
func TestNewAuthRepo(t *testing.T) {
	type args struct {
		redisCC redis.UniversalClient
		logger  log.Logger
		config  Config
	}
	tests := []struct {
		name    string
		args    args
		want    AuthRepo
		wantErr bool
	}{
		{
			name: "#test",
			args: args{
				redisCC: &redis.Client{},
				logger:  log.DefaultLogger,
				config: Config{
					SignCrypto:    NewSignEncryptor("SignKey"),
					RefreshCrypto: NewCBCCipher("RefreshKey"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenM := NewTokenManager(tt.args.logger, tt.args.redisCC, nil, &testLocker{})
			got, err := NewAuthRepo(tt.args.config, tt.args.logger, tokenM)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("NewAuthRepo() 返回 nil，期望非 nil")
			}
		})
	}
}
