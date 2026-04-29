package mysqlpkg

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
	"time"

	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	"github.com/stretchr/testify/require"
)

var _ gormpkg.BatchInsertRepo = new(UserSlice)

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
	insertColumn := []string{
		"name", "age",
	}
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(insertColumn))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))
	return insertColumn, string(insertPlaceholderBytes)
}

// InsertValues 插入的值
func (s *UserSlice) InsertValues(args *gormpkg.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")
		prepareData = append(prepareData, dataModels[index].Name)
		prepareData = append(prepareData, dataModels[index].Age)
	}
	return prepareData, placeholderSlice
}

// ConflictActionForMySQL 更新冲突时的操作 (MySQL)
func (s *UserSlice) ConflictActionForMySQL() *gormpkg.BatchInsertConflictActionReq {
	updateColumns := []string{
		"name=" + gormpkg.DefaultBatchInsertConflictAlias + ".name",
		"age=" + gormpkg.DefaultBatchInsertConflictAlias + ".age+1",
	}

	req := gormpkg.DefaultBatchInsertConflictActionForMySQL
	req.OnConflictValueAlias = "AS " + gormpkg.DefaultBatchInsertConflictAlias
	req.OnConflictTarget = "ON DUPLICATE KEY"
	req.OnConflictAction = "UPDATE " + strings.Join(updateColumns, ", ")

	return &req
}

// getDataModels 获取测试数据
func getDataModels() UserSlice {
	var (
		userTotal  = 10
		userModels = make([]*User, userTotal)
	)
	_ = time.Now()
	for i := 0; i < userTotal; i++ {
		userModels[i] = &User{
			Name: "user_" + strconv.Itoa(i),
			Age:  i + 1,
		}
	}
	return userModels
}

// go test -v ./data/mysql/ -count=1 -run TestBatchInsert_ForMySQL
func TestBatchInsert_ForMySQL(t *testing.T) {
	dataModels := getDataModels()

	var args = []struct {
		name string
		opts []gormpkg.BatchInsertOption
	}{
		{
			name: "#batch_insert_for_mysql_#_with_ignore",
			opts: []gormpkg.BatchInsertOption{
				gormpkg.WithBatchInsertIgnore(),
			},
		},
		{
			name: "#batch_insert_for_mysql_#_with_conflict_action",
			opts: []gormpkg.BatchInsertOption{
				gormpkg.WithBatchInsertConflictAction(dataModels.ConflictActionForMySQL()),
			},
		},
	}

	for _, data := range args {
		t.Run(data.name, func(t *testing.T) {
			err := gormpkg.BatchInsert(dbConn, &dataModels, data.opts...)
			require.Nil(t, err)
		})
	}
}
