package gormutil

import (
	"context"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

// BatchInsertValueArgs 批量插入值的参数
type BatchInsertValueArgs struct {
	// StepStart 开始步长：索引
	StepStart int
	// StepEnd 结束步长：索引
	StepEnd int
	// InsertColumns 插入的列；例："id"，"name"，"age"
	InsertColumnStr string
	// InsertPlaceholder 列的占位符；例："?, ?, ?"
	InsertPlaceholder string
}

// BatchInsertRepo 批量插入
type BatchInsertRepo interface {
	// TableName 表名
	TableName() string
	// Len 数据长度
	// 例子：length := len([]*User{})
	Len() int
	// InsertColumns 插入的列
	// @param columnList 插入的列名数组；例：[]string{"id"，"name"，"age"}
	// @param placeholder 列的占位符；例："?, ?, ?"
	// 在实现此方法时：需要自行拼接占位符；
	InsertColumns() (columnList []string, placeholder string)
	// InsertValues 插入的值
	// @result prepareData 插入的值；例：[]interface{}{1, "张三", 18, 2, "李四", 20, 3, "小明", 30}
	// @result prepareDataLen 插入的占位符；例：[]string{"(?, ?, ?)", "(?, ?, ?)", "(?, ?, ?)"}
	InsertValues(args *BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string)
}

// BatchInsert 批量插入
func BatchInsert(db *gorm.DB, repo BatchInsertRepo) error {
	if repo.Len() == 0 {
		if db.Logger != nil {
			db.Logger.Info(context.Background(), "insert data is empty")
		}
		return nil
	}

	// insert columns
	insertColumnList, insertPlaceholder := repo.InsertColumns()
	insertColumnStr := strings.Join(insertColumnList, ", ")
	columnLen := len(insertColumnList)

	// sql : 1390 Prepared statement contains too many placeholders[65535(2^16-1)]
	// insert channelLen records at a time
	channelLen := 65535 / columnLen
	channelCount := int(math.Ceil(float64(repo.Len()) / float64(channelLen)))

	// insert
	for i := 1; i <= channelCount; i++ {
		start := (i - 1) * channelLen
		end := i * channelLen
		if end > repo.Len() {
			end = repo.Len()
		}

		// insert
		args := &BatchInsertValueArgs{
			StepStart:         start,
			StepEnd:           end,
			InsertColumnStr:   insertColumnStr,
			InsertPlaceholder: insertPlaceholder,
		}
		if err := insertIntoTable(db, repo, args); err != nil {
			return err
		}
	}
	return nil
}

// insertIntoTable into table
func insertIntoTable(dbConn *gorm.DB, repo BatchInsertRepo, args *BatchInsertValueArgs) (err error) {
	// SQL
	insertSQL := fmt.Sprintf("INSERT INTO %s(%s) VALUES ", repo.TableName(), args.InsertColumnStr)

	// prepare data
	//var (
	//	prepareData []interface{}
	//	placeholderSlice []string
	//)
	//for _, dataModel := range dataModels {
	//	// placeholder
	//	placeholderSlice = append(placeholderSlice, "("+placeholder+")")
	//
	//	// prepare data
	//	prepareData = append(prepareData, m.insertValues(dataModel)...)
	//}

	// prepare data
	prepareData, placeholderSlice := repo.InsertValues(args)

	// SQL
	insertSQL += strings.Join(placeholderSlice, ", ")

	// insert
	if err = dbConn.Exec(insertSQL, prepareData...).Error; err != nil {
		return err
	}
	return
}
