package migrationpkg

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go test -v -count=1 ./data/migration -test.run=TestMigrateALL_MySQL
func TestMigrateALL_MySQL(t *testing.T) {
	dbConn, err := newMysqlDB()
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
	dbConn = gormpkg.SetOption(dbConn, gormpkg.OptionKeyTableOptions, gormpkg.OptionValueEngineInnoDB)

	return dbConn, err
}
