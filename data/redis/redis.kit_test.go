package redisutil

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

var (
	redisConfig = &confv1.Data_Redis{
		Addr:            "127.0.0.1:6379",
		Username:        "",
		Password:        "",
		Db:              0,
		DialTimeout:     durationpb.New(time.Second * 3),
		ReadTimeout:     durationpb.New(time.Second * 3),
		WriteTimeout:    durationpb.New(time.Second * 3),
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(time.Minute * 30),
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
	}
)

// go test -v ./data/redis/ -count=1 -test.run=TestNewDB_Xxx
func TestNewDB_Xxx(t *testing.T) {
	db, err := NewDB(redisConfig)
	require.Nil(t, err)

	ctx := context.Background()

	tests := []struct {
		name  string
		key   string
		value string
		want  string
	}{
		{
			name:  "#set-foo1",
			key:   "foo1",
			value: "bar1",
			want:  "bar1",
		},
		{
			name:  "#set-foo1",
			key:   "foo2",
			value: "bar2",
			want:  "bar2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setRes, setErr := db.Set(ctx, tt.key, tt.value, 0).Result()
			require.Nil(t, setErr)
			t.Log(setRes)
			gotCmd := db.Get(ctx, tt.key)
			require.Nil(t, gotCmd.Err())
			require.Equal(t, tt.want, gotCmd.Val())
		})
	}
}
