package setup

import (
	"testing"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/stretchr/testify/require"
)

// go test -v ./example/internal/setup/ -count=1 -test.run=TestNewUpPackages
func TestNewUpPackages(t *testing.T) {
	// config
	confPath := "./../../configs"
	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	configHandler, err := NewConfiguration(opts...)
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}
	t.Log("*** | env：", configHandler.Env())

	// up
	upHandler := NewModules(configHandler)

	// db
	db, err := upHandler.MysqlGormDB()
	require.Nil(t, err)
	require.NotNil(t, db)

	// redis
	redisCC, err := upHandler.RedisClient()
	require.Nil(t, err)
	require.NotNil(t, redisCC)
}
