package gormutil

import (
	"bytes"
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
	Age  int    `gorm:"COLUMN:age"`            // age
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
		log.Fatalf("==> 请先配置数据库，错误信息：%v\n", err)
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

var _ BatchInsertRepo = new(UserSlice)

// UserSlice 用户切片
type UserSlice []*User

// TableName 表名
func (s UserSlice) TableName() string {
	if len(s) > 0 {
		return s[0].TableName()
	}
	return (&User{}).TableName()
}

// Len 长度
func (s UserSlice) Len() int {
	return len(s)
}

// InsertColumns 插入的列
func (s UserSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	insertColumn := []string{
		"name", "age",
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(insertColumn))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return insertColumn, string(insertPlaceholderBytes)
}

// InsertValues 插入的值
func (s UserSlice) InsertValues(args *BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := ([]*User(s))[args.StepStart:args.StepEnd]
	for _, dataModel := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModel.Name)
		prepareData = append(prepareData, dataModel.Age)
	}
	return prepareData, placeholderSlice
}
