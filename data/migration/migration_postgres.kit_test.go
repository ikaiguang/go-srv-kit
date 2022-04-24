package migrationuitl

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go test -v -count=1 ./data/migration -test.run=TestMigrateALL_Postgres
func TestMigrateALL_Postgres(t *testing.T) {
	dbConn, err := newPostgresDB()
	require.Nil(t, err)

	var (
		normalMg = NewCreateTable(dbConn.Migrator(), &Migration{})
		testMg   = NewCreateTable(dbConn.Migrator(), &TestMigration{})
	)
	var args = []struct {
		name string
		repo MigrationRepo
	}{
		{
			name: "#创建表：" + normalMg.Identifier(),
			repo: normalMg,
		},
		{
			name: "#创建表：" + testMg.Identifier(),
			repo: testMg,
		},
	}

	for i := range args {
		MustRegisterMigrate(args[i].repo)
	}

	// 迁移
	err = MigrateALL()
	require.Nil(t, err)
	err = Migrate(testMg.Identifier(), normalMg.Identifier())
	require.Nil(t, err)
	err = MigrateRepos(testMg, normalMg)
	require.Nil(t, err)
	// 回滚
	err = RollbackALL()
	require.Nil(t, err)
	err = Rollback(testMg.Identifier(), normalMg.Identifier())
	require.Nil(t, err)
	err = RollbackRepos(testMg, normalMg)
	require.Nil(t, err)
}

// newPostgresDB ...
func newPostgresDB() (*gorm.DB, error) {
	var (
		err error
		opt = &gorm.Config{
			PrepareStmt:                              true,
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger: logger.New(log.New(os.Stderr, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			}),
		}
	)
	dsn := "host=127.0.0.1 user=postgres password=Postgres.123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, err := gorm.Open(postgres.Open(dsn), opt)
	if err != nil {
		err = fmt.Errorf("请先配置数据库，错误信息：%s", err.Error())
		return dbConn, err
	}
	return dbConn, err
}
