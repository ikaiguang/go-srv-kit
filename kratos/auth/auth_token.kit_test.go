package authpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

// go test -v -count 1 ./kratos/auth -run TestNewAuthRepo
func TestNewAuthRepo(t *testing.T) {
	type args struct {
		logger log.Logger
		config Config
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
				logger: log.DefaultLogger,
				config: Config{
					SignCrypto:    NewSignEncryptor("SignKey"),
					RefreshCrypto: NewCBCCipher("RefreshKey"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthRepo(tt.args.config, tt.args.logger, nil)
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
