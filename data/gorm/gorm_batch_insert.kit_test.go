package gormpkg

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// go test -v ./data/gorm/ -count=1 -run TestBatchInsert_ForMySQL
func TestBatchInsert_ForMySQL(t *testing.T) {
	dataModels := getDataModels()

	var args = []struct {
		name string
		opts []BatchInsertOption
	}{
		//{
		//	name: "#batch_insert_for_mysql",
		//	opts: nil,
		//},
		{
			name: "#batch_insert_for_mysql_#_with_ignore",
			opts: []BatchInsertOption{
				WithBatchInsertIgnore(),
			},
		},
		{
			name: "#batch_insert_for_mysql_#_with_conflict_action",
			opts: []BatchInsertOption{
				WithBatchInsertConflictAction(dataModels.ConflictActionForMySQL()),
			},
		},
	}

	for _, data := range args {
		t.Run(data.name, func(t *testing.T) {
			err := BatchInsert(dbConn, &dataModels, data.opts...)
			require.Nil(t, err)
		})
	}
}

// go test -v ./data/gorm/ -count=1 -run TestBatchInsert_ForPostgres
func TestBatchInsert_ForPostgres(t *testing.T) {
	dataModels := getDataModels()

	var args = []struct {
		name string
		opts []BatchInsertOption
	}{
		//{
		//	name: "#batch_insert_for_mysql",
		//	opts: nil,
		//},
		//{
		//	// 无IGNORE语法
		//	name: "#batch_insert_for_mysql_#_with_ignore",
		//	//opts: []BatchInsertOption{
		//	//	WithBatchInsertIgnore(),
		//	//},
		//},
		{
			name: "#batch_insert_for_mysql_#_with_conflict_action",
			opts: []BatchInsertOption{
				WithBatchInsertConflictAction(dataModels.ConflictActionForPostgres()),
			},
		},
	}

	for _, data := range args {
		t.Run(data.name, func(t *testing.T) {
			err := BatchInsert(psqlConn, &dataModels, data.opts...)
			require.Nil(t, err)
		})
	}
}

// getDataModels 获取数据模型
func getDataModels() UserSlice {
	var (
		now        = time.Now().Format(time.RFC3339)
		userTotal  = 10
		userModels = make([]*User, userTotal)
	)
	_ = now
	for i := 0; i < userTotal; i++ {
		userModels[i] = &User{
			//Name: "user_" + now + "_" + strconv.Itoa(i),
			Name: "user_" + strconv.Itoa(i),
			Age:  i + 1,
		}
	}
	return userModels
}
