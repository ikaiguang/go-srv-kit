package setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
)

// go test -v ./example/internal/setup/ -count=1 -test.run=TestSetup -conf=./../../configs
func TestSetup(t *testing.T) {
	modulesHandler, err := Setup()
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}
	defer func() { _ = modulesHandler.Close() }()

	ctx := context.Background()

	// env
	loghelper.Infof("env = %v", modulesHandler.Env())

	// debug
	debugutil.Println("*** | ==> debug util print")

	// 日志
	loghelper.Info("*** | ==> log helper info")

	// db
	db, err := modulesHandler.MysqlGormDB()
	require.Nil(t, err)
	require.NotNil(t, db)
	type DBRes struct {
		DBName string `gorm:"column:db_name"`
	}
	var dbRes DBRes
	err = db.WithContext(ctx).Raw("SELECT DATABASE() AS db_name").Scan(&dbRes).Error
	require.Nil(t, err)
	t.Logf("db res : %+v\n", dbRes)

	// redis
	redisCC, err := modulesHandler.RedisClient()
	require.Nil(t, err)
	require.NotNil(t, redisCC)
	redisKey := "test-foo"
	redisValue := "test-bar"
	err = redisCC.Set(ctx, redisKey, redisValue, 0).Err()
	require.Nil(t, err)
	redisGotValue, err := redisCC.Get(ctx, redisKey).Result()
	require.Nil(t, err)
	require.Equal(t, redisValue, redisGotValue)
	t.Logf("redis res : %+v\n", redisGotValue)
}

// go test -v ./example/internal/setup/ -count=1 -test.run=TestGetModules -conf=./../../configs
func TestGetModules(t *testing.T) {
	modulesHandler, err := GetModules()
	require.Nil(t, err)
	require.NotNil(t, modulesHandler)
	modulesHandler, err = GetModules()
	require.Nil(t, err)
	require.NotNil(t, modulesHandler)

	// env
	loghelper.Infof("env = %v", modulesHandler.Env())

	// debug
	debugutil.Println("*** | ==> debug util print")

	// 日志
	loghelper.Info("*** | ==> log helper info")
}
