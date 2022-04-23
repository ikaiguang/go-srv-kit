package migrationuitl

import (
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
	Id                uint64     `gorm:"COLUMN:id;primaryKey;type:uint;autoIncrement;size:20;comment:'ID'"`
	ColumnUniqueIndex string     `gorm:"COLUMN:column_unique_index;uniqueIndex;type:string;size:255;not null;default:'';comment:'唯一索引'"`
	ColumnIndex       string     `gorm:"COLUMN:column_index;index;type:string;not null;default:0;comment:'普通索引'"`
	ColumnBool        bool       `gorm:"COLUMN:column_bool;type:bool;size:1;not null;default:'0';comment:'布尔值'"`
	ColumnInt         int32      `gorm:"COLUMN:column_int;type:int;size:11;not null;default:'0';comment:'整型'"`
	ColumnUint        uint32     `gorm:"COLUMN:column_uint;type:uint;size:11;not null;default:'0';comment:'整型：无符号'"`
	ColumnInt64       int64      `gorm:"COLUMN:column_int64;type:int;size:20;not null;default:'0';comment:'整型：64位'"`
	ColumnUint64      uint64     `gorm:"COLUMN:column_uint64;type:uint;size:20;not null;default:'0';comment:'整型：无符号64位'"`
	ColumnFloat64     float64    `gorm:"COLUMN:column_float64;type:float;size:26;precision:6;not null;default:'0';comment:'浮点型：64位'"`
	ColumnString      string     `gorm:"COLUMN:column_string;type:string;size:255;not null;default:'';comment:'字符串'"`
	ColumnString2     string     `gorm:"COLUMN:column_string;type:string;size:65535;not null;default:'';comment:'字符串'"`
	ColumnTime        time.Time  `gorm:"COLUMN:column_time;type:time;not null;autoCreateTime:milli;comment:'时间'"`
	ColumnText        string     `gorm:"COLUMN:column_text;type:text;default:'';comment:'文本'"`
	CreatedTime       time.Time  `gorm:"COLUMN:column_created_time;type:time;not null;autoCreateTime:milli"`
	ColumnUpdatedTime time.Time  `gorm:"COLUMN:column_updated_time;type:time;not null;autoUpdateTime:milli"`
	ColumnDeletedTime *time.Time `gorm:"COLUMN:column_deleted_time;type:time;default:null"`

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
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
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

	// 运行 & 退出
	os.Exit(m.Run())
}
