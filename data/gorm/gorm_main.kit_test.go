package gormutil

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

// User 用户
type User struct {
	Id   int64  `gorm:"PRIMARY_KEY;COLUMN:id"` // id
	Name string `gorm:"COLUMN:name"`           // name
	Age  int64  `gorm:"COLUMN:age"`            // age
}

// TableName 表名
func (s *User) TableName() string {
	return "users"
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
		panic(err)
	}

	// select
	userModel := &User{}
	err = dbConn.Table(userModel.TableName()).
		Where("id > ?", 0).
		First(userModel).Error
	if err != nil {
		panic(err)
	}
	if userModel.Id == 0 {
		panic("请先配置：测试数据库")
	}

	// 运行 & 退出
	os.Exit(m.Run())
}

// ========== 批量插入 ==========

// UserSlice 用户切片
type UserSlice []*User
