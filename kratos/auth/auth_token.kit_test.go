package authpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

// go test -v -count=1 ./auth -test.run=TestNewAuthRepo
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
					SignKey: "abc",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthRepo(tt.args.redisCC, tt.args.logger, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAuthRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("got: ", got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewAuthRepo() got = %v, want %v", got, tt.want)
			//}
		})
	}
	t.Log("testdata: ", ERROR_UNAUTHORIZED.String())
	t.Log("testdata: ", ErrBlacklist())
}
