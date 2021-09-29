package mysqlutil

import (
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/gorm"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	writerutil "github.com/ikaiguang/go-srv-kit/writer"
)

var (
	conf = &confv1.Data_MySQL{
		Dsn:             "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&timeout=30s&parseTime=True",
		LoggerEnable:    true,
		LoggerLevel:     "INFO",
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(time.Minute * 30),
	}
)

// go test -v ./mysql/ -count=1 -test.run=TestNewDB_Xxx
func TestNewDB_Xxx(t *testing.T) {
	db, err := NewDB(conf)
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

	opt := WithWriters(NewStdWriter(), NewJSONWriter(fileWriter))
	db, err := NewDB(conf, opt)
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
