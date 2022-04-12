package gormutil

import (
	"context"
	"regexp"

	pageutil "github.com/ikaiguang/go-srv-kit/kit/page"
	"gorm.io/gorm"
)

var (
	// regColumn 正则表达式:列
	regColumn = regexp.MustCompile("^[A-Za-z-_]+$")
)

// Page 分页
func Page(db *gorm.DB, pageOption *pageutil.PageOption) *gorm.DB {
	// limit offset
	db = db.Limit(pageOption.Limit).Offset(pageOption.Offset)

	// where
	for i := range pageOption.Where {
		// Where("columnName = ?", columnData)
		//whereStr := pageOption.Where[i].Column + " " + pageOption.Where[i].Operator + " " + pageOption.Where[i].Placeholder
		db = db.Where(
			"? ? ?",
			pageOption.Where[i].Column,
			pageOption.Where[i].Operator,
			pageOption.Where[i].Data,
		)
	}

	// order
	for i := range pageOption.Order {
		column := pageOption.Order[i].Column
		if !regColumn.MatchString(pageOption.Order[i].Column) {
			//column = pageutil.DefaultOrderColumn
			column = "bad_order_with_invalid_column"
			if db.Logger != nil {
				db.Logger.Error(context.Background(), "invalid column(", pageOption.Order[i].Column, ")")
			}
		}
		//orderStr := pageOption.Order[i].Column + " " + pageOption.Order[i].Direction,
		db = db.Order(column + " " + pageutil.ParseOrderDirection(pageOption.Order[i].Direction))
	}
	return db
}
