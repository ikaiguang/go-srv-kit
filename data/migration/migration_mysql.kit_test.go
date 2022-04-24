package migrationuitl

import (
	"fmt"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

// go test -v -count=1 ./data/migration -test.run=TestMigrateALL_MySQL
func TestMigrateALL_MySQL(t *testing.T) {
	dbConn, err := newMysqlDB()
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

// newMysqlDB ...
func newMysqlDB() (*gorm.DB, error) {
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
	dsn := "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=30s&parseTime=True"
	dbConn, err := gorm.Open(mysql.Open(dsn), opt)
	if err != nil {
		err = fmt.Errorf("请先配置数据库，错误信息：%s", err.Error())
		return dbConn, err
	}
	dbConn = dbConn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	dbConn = gormutil.SetOption(dbConn, gormutil.OptionKeyTableOptions, gormutil.OptionValueEngineInnoDB)

	return dbConn, err
}
