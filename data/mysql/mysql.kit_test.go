package mysqlutil

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	writerutil "github.com/ikaiguang/go-srv-kit/kit/writer"
)

var (
	dbConfig = &confv1.Data_MySQL{
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

// go test -v ./data/mysql/ -count=1 -test.run=TestNewDB_Xxx
func TestNewDB_Xxx(t *testing.T) {
	db, err := NewMysqlDB(dbConfig)
	require.Nil(t, err)

	// testDBConn
	testDBConn(t, db)
}

// go test -v ./data/mysql/ -count=1 -test.run=TestNewDB_WithWriters
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
	opt := gormutil.WithWriters(gormutil.NewStdWriter(), gormutil.NewJSONWriter(fileWriter))
	db, err := NewMysqlDB(dbConfig, opt)
	require.Nil(t, err)

	// testDBConn
	testDBConn(t, db)
}

// go test -v ./data/mysql/ -count=1 -test.run=TestDefaultGorm_Xxx
func TestDefaultGorm_Xxx(t *testing.T) {
	// 拨号
	dialect := mysql.Open(dbConfig.Dsn)

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

	// testDBConn
	testDBConn(t, db)
}

// testDBConn .
func testDBConn(t *testing.T, db *gorm.DB) {
	stdDB, err := db.DB()
	if err != nil {
		t.Error(err)
		return
	}

	err = stdDB.Ping()
	if err != nil {
		t.Error(err)
		return
	}
}
