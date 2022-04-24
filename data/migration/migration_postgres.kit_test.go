package migrationuitl

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

// go test -v -count=1 ./data/migration -test.run=TestMigrateALL_Postgres
func TestMigrateALL_Postgres(t *testing.T) {
	dbConn, err := newPostgresDB()
	require.Nil(t, err)

	var (
		normalMg = NewCreateTable(dbConn.Migrator(), &MigrationForPostgres{})
		testMg   = NewCreateTable(dbConn.Migrator(), &TestMigrationForPostgres{})
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

// MigrationForPostgres 数据库迁移
// 文档地址：https://gorm.io/docs/models.html
// MySQL 支持 auto_increment
// Postgres : serial or bigserial
type MigrationForPostgres struct {
	Id                 uint64    `gorm:"COLUMN:id;primaryKey;type:bigserial;comment:ID"`
	MigrationKey       string    `gorm:"COLUMN:migration_key;uniqueIndex;type:string;size:255;not null;default:'';comment:迁移key：唯一"`
	MigrationBatch     uint      `gorm:"COLUMN:migration_batch;type:int;not null;default:0;comment:迁移批次"`
	MigrationDesc      string    `gorm:"COLUMN:migration_desc;type:text;not null;comment:迁移描述"`
	MigrationExtraInfo string    `gorm:"COLUMN:migration_extra_info;type:json;not null;comment:迁移：额外信息"`
	CreatedTime        time.Time `gorm:"COLUMN:created_time;type:time;not null;autoCreateTime:milli;comment:创建时间"`
	UpdatedTime        time.Time `gorm:"COLUMN:updated_time;type:time;not null;autoUpdateTime:milli;comment:更新时间"`
}

// TableName 表名
func (s *MigrationForPostgres) TableName() string {
	return DefaultMigrationTableName
}

// TestMigrationForPostgres test
// 文档地址：https://gorm.io/docs/models.html
// MySQL 支持 unsigned
// Postgres 不支持 unsigned
// MySQL 支持 auto_increment
// Postgres : serial or bigserial
type TestMigrationForPostgres struct {
	Id                uint64     `gorm:"COLUMN:id;primaryKey;type:bigserial;comment:ID"`
	ColumnUniqueIndex string     `gorm:"COLUMN:column_unique_index;uniqueIndex;type:string;size:255;not null;default:'';comment:唯一索引"`
	ColumnIndex       string     `gorm:"COLUMN:column_index;index;type:string;size:255;not null;default:0;comment:普通索引"`
	ColumnBool        bool       `gorm:"COLUMN:column_bool;type:bool;size:1;not null;default:0;comment:布尔值"`
	ColumnInt         int32      `gorm:"COLUMN:column_int;type:int;not null;default:0;comment:整型"`
	ColumnUint        uint32     `gorm:"COLUMN:column_uint;type:int;not null;default:0;comment:整型：无符号"`
	ColumnInt64       int64      `gorm:"COLUMN:column_int64;type:bigint;not null;default:0;comment:整型：64位"`
	ColumnUint64      uint64     `gorm:"COLUMN:column_uint64;type:bigint;not null;default:0;comment:整型：无符号64位"`
	ColumnFloat64     float64    `gorm:"COLUMN:column_float64;type:decimal(30,10);not null;default:0;comment:浮点型：64位"`
	ColumnString      string     `gorm:"COLUMN:column_string;type:string;size:255;not null;default:'';comment:字符串"`
	ColumnText        string     `gorm:"COLUMN:column_text;type:text;not null;comment:文本"`
	ColumnJSON        string     `gorm:"COLUMN:column_json;type:json;not null;comment:JSON"`
	ColumnBytes       string     `gorm:"COLUMN:column_bytes;type:bytes;not null;comment:字节"`
	CreatedTime       time.Time  `gorm:"COLUMN:column_created_time;type:time;not null;autoCreateTime:milli;comment:创建时间"`
	ColumnUpdatedTime time.Time  `gorm:"COLUMN:column_updated_time;type:time;not null;autoUpdateTime:milli;comment:更新时间"`
	ColumnDeletedTime *time.Time `gorm:"COLUMN:column_deleted_time;type:time;default:null;comment:删除时间"`

	// IgnoreReadWrite ignore this field when write and read with struct
	IgnoreReadWrite string `gorm:"-"`
	// IgnoreAll ignore this field when write/read and migrate with struct
	IgnoreAll string `gorm:"-:all"`
	// IgnoreMigration ignore this field when migrate with struct
	IgnoreMigration string `gorm:"-:migration"`
}

// TableName 表名
func (s *TestMigrationForPostgres) TableName() string {
	return "srv_test_migration"
}

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
