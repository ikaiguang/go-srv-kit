package migrationuitl

import (
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbConn *gorm.DB
)

// TestMigration test
// 文档地址：https://gorm.io/docs/models.html
type TestMigration struct {
	Id                uint64     `gorm:"COLUMN:id;primaryKey;type:bigint unsigned auto_increment;comment:ID"`
	ColumnUniqueIndex string     `gorm:"COLUMN:column_unique_index;uniqueIndex;type:string;size:255;not null;default:'';comment:唯一索引"`
	ColumnIndex       string     `gorm:"COLUMN:column_index;index;type:string;size:255;not null;default:0;comment:普通索引"`
	ColumnBool        bool       `gorm:"COLUMN:column_bool;type:bool;size:1;not null;default:0;comment:布尔值"`
	ColumnInt         int32      `gorm:"COLUMN:column_int;type:int;not null;default:0;comment:整型"`
	ColumnUint        uint32     `gorm:"COLUMN:column_uint;type:int unsigned;not null;default:0;comment:整型：无符号"`
	ColumnInt64       int64      `gorm:"COLUMN:column_int64;type:bigint;not null;default:0;comment:整型：64位"`
	ColumnUint64      uint64     `gorm:"COLUMN:column_uint64;type:bigint unsigned;not null;default:0;comment:整型：无符号64位"`
	ColumnFloat64     float64    `gorm:"COLUMN:column_float64;type:decimal(30,10) unsigned;not null;default:0;comment:浮点型：64位"`
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
func (s *TestMigration) TableName() string {
	return "srv_test_migration"
}

func TestMain(m *testing.M) {
	var (
		err error
		dsn = "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=30s&parseTime=True"
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
	dbConn, err = gorm.Open(mysql.Open(dsn), opt)
	if err != nil {
		log.Fatalf("==> 请先配置数据库，错误信息：%v\n", err)
	}

	//dbConn = dbConn.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	dbConn = gormutil.SetOption(dbConn, gormutil.OptionKeyTableOptions, gormutil.OptionValueEngineInnoDB)

	// 运行 & 退出
	os.Exit(m.Run())
}
