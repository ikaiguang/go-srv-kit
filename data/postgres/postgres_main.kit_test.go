package psqlpkg

import (
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用户（测试用）
type User struct {
	Id   uint64 `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID"`
	Name string `gorm:"column:name;uniqueIndex;type:string;size:255;not null;default:'';comment:name+唯一索引"`
	Age  int    `gorm:"column:age;type:int;not null;default:0;comment:年龄"`
}

// TableName 表名
func (s *User) TableName() string {
	return "users"
}

var dbConn *gorm.DB

func TestMain(m *testing.M) {
	initPostgresDB()
	os.Exit(m.Run())
}

// initPostgresDB 初始化 Postgres 测试数据库
func initPostgresDB() {
	var (
		err error
		dsn = getTestPostgresDSN()
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
	dbConn = db
}
