package psqlpkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/gorm"
)

var (
	dbConfig = &Config{
		Dsn:             "host=127.0.0.1 user=postgres password=Postgres.123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		SlowThreshold:   durationpb.New(time.Millisecond * 100),
		LoggerEnable:    true,
		LoggerLevel:     "INFO",
		ConnMaxActive:   100,
		ConnMaxLifetime: durationpb.New(time.Minute * 30),
		ConnMaxIdle:     10,
		ConnMaxIdleTime: durationpb.New(time.Hour),
	}
)

// go test -v ./data/postgres/ -count=1 -run TestNewDB_Xxx
//
// ===== 创建用户：修改密码：分配权限 =====
// CREATE DATABASE test ENCODING = 'utf8';
// CREATE USER postgres WITH PASSWORD 'Postgres.123456';
// ALTER USER postgres WITH PASSWORD 'Postgres.123456';
// GRANT ALL PRIVILEGES ON DATABASE test TO postgres;
func TestNewDB_Xxx(t *testing.T) {
	db, err := NewPostgresDB(dbConfig)
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
