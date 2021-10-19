package mysqlutil

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	writerutil "github.com/ikaiguang/go-srv-kit/writer"
)

var (
	mysqlConfig = &confv1.Data_MySQL{
		Dsn:             "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=30s&parseTime=True",
		SlowThreshold:   durationpb.New(time.Millisecond * 100),
		LoggerEnable:    true,
		LoggerLevel:     "INFO",
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(time.Minute * 30),
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
	}
)

// go test -v ./mysql/ -count=1 -test.run=TestNewDB_Xxx
func TestNewDB_Xxx(t *testing.T) {
	db, err := NewMysqlDB(mysqlConfig)
	require.Nil(t, err)

	// res
	type Model struct {
		ID   uint64 `gorm:"column:id;primary_key;" json:"id"`
		UUID string `gorm:"column:uuid" json:"uuid"`
	}
	res := &Model{}

	// select
	err = db.Table("ikg_test").Where("id = ?", 1).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log("select result ", err)
		}
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v\n", res)
}

// go test -v ./mysql/ -count=1 -test.run=TestNewDB_WithWriters
func TestNewDB_WithWriters(t *testing.T) {
	writerConfig := &writerutil.ConfigRotate{
		Dir:      ".",
		Filename: "test",

		RotateTime:     time.Hour,
		StorageCounter: 10,
	}
	fileWriter, err := writerutil.NewRotateFile(writerConfig)
	require.Nil(t, err)

	//opt := WithWriters(NewStdWriter(), NewWriter(fileWriter), NewJSONWriter(fileWriter))
	opt := WithWriters(NewStdWriter(), NewJSONWriter(fileWriter))
	db, err := NewMysqlDB(mysqlConfig, opt)
	require.Nil(t, err)

	// res
	type Model struct {
		ID   uint64 `gorm:"column:id;primary_key;" json:"id"`
		UUID string `gorm:"column:uuid" json:"uuid"`
	}
	res := &Model{}

	// select
	err = db.Table("ikg_test").Where("id = ?", 1).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log("select result ", err)
		}
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v\n", res)
}

// go test -v ./mysql/ -count=1 -test.run=TestDefaultGorm_Xxx
func TestDefaultGorm_Xxx(t *testing.T) {
	// 拨号
	dialect := mysql.Open(mysqlConfig.Dsn)

	loggerHandler := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel:                  logger.Info,
		SlowThreshold:             time.Millisecond * 100,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
	})

	opt := &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   loggerHandler,
	}

	db, err := gorm.Open(dialect, opt)
	require.Nil(t, err)

	// res
	type Model struct {
		ID   uint64 `gorm:"column:id;primary_key;" json:"id"`
		UUID string `gorm:"column:uuid" json:"uuid"`
	}
	res := &Model{}

	// select
	err = db.Table("ikg_test").Where("id = ?", 1).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Log("select result ", err)
		}
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v\n", res)
}
