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

	// InsertPlaceholder 列的占位符；例："?, ?, ?"
	InsertPlaceholder string
	// InsertSQL 入库的SQL；INSERT INTO ... VALUES (...) AS alias
	InsertSQL string
	// ConflictActionSQL 存在冲突，执行冲突动作
	//insertSQL += "ON DUPLICATE KEY UPDATE " + strings.Join(s.UpdateColumnFromMtdReportData(), ",")
	//insertSQL += "ON CONFLICT (id) DO UPDATE SET column_2= CONCAT(test_table.column_2, excluded.column_2)"
	ConflictActionSQL   string
	ConflictPrepareData []interface{}
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
func BatchInsert(db *gorm.DB, repo BatchInsertRepo, opts ...BatchInsertOption) error {
	return BatchInsertWithContext(context.Background(), db, repo, opts...)
}

// BatchInsertWithContext 批量插入
func BatchInsertWithContext(ctx context.Context, db *gorm.DB, repo BatchInsertRepo, opts ...BatchInsertOption) error {
	if repo.Len() == 0 {
		if db.Logger != nil {
			db.Logger.Info(ctx, "insert data is empty")
		}
		return nil
	}

	// 选项
	opt := &batchInsertOptions{}
	for i := range opts {
		opts[i](opt)
	}

	// insert columns
	insertColumnList, insertPlaceholder := repo.InsertColumns()
	insertColumnStr := strings.Join(insertColumnList, ", ")

	// SQL
	insertSQL := "INSERT"
	if opt.isInsertIgnore {
		insertSQL += " IGNORE"
	}
	insertSQL = fmt.Sprintf(insertSQL+" INTO %s(%s) VALUES ", repo.TableName(), insertColumnStr)

	// ON CONFLICT；
	// MySQL : ON DUPLICATE KEY UPDATE；
	//insertSQL += "ON DUPLICATE KEY UPDATE " + strings.Join(s.UpdateColumnFromMtdReportData(), ",")
	// Postgres : ON CONFLICT : PostgresSQL V9.5 以上可用。
	//insertSQL += "ON CONFLICT (id) DO UPDATE SET column_2= CONCAT(test_table.column_2, excluded.column_2)"
	conflictActionSQL := ""
	if opt.withConflictAction {
		conflictActionSQL += " " + opt.onConflictValueAlias + " "
		conflictActionSQL += opt.onConflictTarget + " "
		conflictActionSQL += opt.onConflictAction + " "
	}

	// insert args
	args := &BatchInsertValueArgs{
		StepStart:           0,
		StepEnd:             0,
		InsertPlaceholder:   insertPlaceholder,
		InsertSQL:           insertSQL,
		ConflictActionSQL:   conflictActionSQL,
		ConflictPrepareData: opt.onConflictPrepareData,
	}

	// sql : 1390 Prepared statement contains too many placeholders[65535(2^16-1)]
	// insert channelLen records at a time
	columnLen := len(insertColumnList)
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
		args.StepStart = start
		args.StepEnd = end
		if err := insertIntoTable(ctx, db, repo, args); err != nil {
			return err
		}
	}
	return nil
}

// insertIntoTable into table
func insertIntoTable(ctx context.Context, dbConn *gorm.DB, repo BatchInsertRepo, args *BatchInsertValueArgs) (err error) {
	// SQL
	insertSQL := args.InsertSQL

	// prepare data
	prepareData, placeholderSlice := repo.InsertValues(args)
	if len(args.ConflictPrepareData) > 0 {
		prepareData = append(prepareData, args.ConflictPrepareData...)
	}

	// SQL
	insertSQL += strings.Join(placeholderSlice, ", ")
	insertSQL += args.ConflictActionSQL

	// insert
	if err = dbConn.WithContext(ctx).Exec(insertSQL, prepareData...).Error; err != nil {
		return err
	}
	return
}
