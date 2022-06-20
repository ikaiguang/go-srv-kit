package gormutil

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbConn   *gorm.DB
	psqlConn *gorm.DB
)

// User 用户
type User struct {
	Id   uint64 `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID"`
	Name string `gorm:"column:name;uniqueIndex;type:string;size:255;not null;default:'';comment:name+唯一索引"`
	Age  int    `gorm:"column:age;type:int;not null;default:0;comment:年龄"`
}

// TableName 表名
func (s *User) TableName() string {
	return "users"
}

func TestMain(m *testing.M) {
	// initMySQLDB 初始化数据库 MySQL
	initMySQLDB()

	// initPostgresDB 初始化数据库 Postgres
	initPostgresDB()

	// 运行 & 退出
	os.Exit(m.Run())
}

// ========== 批量插入 ==========

var _ BatchInsertRepo = new(UserSlice)

// UserSlice 用户切片
type UserSlice []*User

// TableName 表名
func (s *UserSlice) TableName() string {
	if len(*s) > 0 {
		return (*s)[0].TableName()
	}
	return (&User{}).TableName()
}

// Len 长度
func (s *UserSlice) Len() int {
	return len(*s)
}

// InsertColumns 插入的列
func (s *UserSlice) InsertColumns() (columnList []string, placeholder string) {
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
func (s *UserSlice) InsertValues(args *BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].Name)
		prepareData = append(prepareData, dataModels[index].Age)
	}
	return prepareData, placeholderSlice
}

// ConflictActionForMySQL 更新冲突时的操作 (MySQL)
func (s *UserSlice) ConflictActionForMySQL() *BatchInsertConflictActionReq {
	updateColumns := []string{
		"name=" + DefaultBatchInsertConflictAlias + ".name",
		"age=" + DefaultBatchInsertConflictAlias + ".age+1",
	}

	req := DefaultBatchInsertConflictActionForMySQL
	req.OnConflictValueAlias = "AS " + DefaultBatchInsertConflictAlias
	req.OnConflictTarget = "ON DUPLICATE KEY"
	req.OnConflictAction = "UPDATE " + strings.Join(updateColumns, ", ")

	return &req
}

// ConflictActionForPostgres 更新冲突时的操作 (Postgres)
func (s *UserSlice) ConflictActionForPostgres() *BatchInsertConflictActionReq {
	updateColumns := []string{
		"name=" + DefaultBatchInsertConflictAlias + ".name",
		"age=" + DefaultBatchInsertConflictAlias + ".age+?",
	}

	req := DefaultBatchInsertConflictActionPostgres
	req.OnConflictValueAlias = ""
	req.OnConflictTarget = "ON CONFLICT(name)"
	req.OnConflictAction = "DO UPDATE SET " + strings.Join(updateColumns, ", ")
	req.OnConflictPrepareData = []interface{}{2}

	return &req
}

// initMySQLDB ...
func initMySQLDB() {
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
	db, err := gorm.Open(mysql.Open(dsn), opt)
	if err != nil {
		log.Fatalf("==> mysql 请先配置数据库，错误信息：%v\n", err)
	}

	// migration
	userModel := &User{}
	if !db.Migrator().HasTable(userModel) {
		err = db.Migrator().AutoMigrate(userModel)
		if err != nil {
			panic(err)
		}
	}
	dbConn = db
}

// initPostgresDB ...
func initPostgresDB() {
	var (
		err error
		dsn = "host=127.0.0.1 user=postgres password=Postgres.123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
	db, err := gorm.Open(postgres.Open(dsn), opt)
	if err != nil {
		log.Fatalf("==> postgres 请先配置数据库，错误信息：%v\n", err)
	}

	// migration
	userModel := &User{}
	if !db.Migrator().HasTable(userModel) {
		err = db.Migrator().AutoMigrate(userModel)
		if err != nil {
			panic(err)
		}
	}
	psqlConn = db
}
