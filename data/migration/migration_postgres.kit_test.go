package migrationutil

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
		normalMg = NewCreateTable(dbConn.Migrator(), Version, &Migration{})
		testMg   = NewCreateTable(dbConn.Migrator(), Version, &TestMigration{})
	)
	var args = []struct {
		name string
		repo MigrationInterface
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

	for _, data := range args {
		t.Run(data.name, func(t *testing.T) {
			err := data.repo.Up()
			require.Nil(t, err)
			err = data.repo.Down()
			require.Nil(t, err)
		})
	}
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
