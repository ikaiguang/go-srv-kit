package gormutil

import (
	"gorm.io/gorm"

	pageutil "github.com/ikaiguang/go-srv-kit/kit/page"
)

// PaginatorArgs 列表参数
type PaginatorArgs struct {
	// PageOption 分页
	PageOption *pageutil.PageOption
	// PageOrders 排序
	PageOrders []*Order
	// PageWheres 条件
	PageWheres []*Where
}

// Paginator 分页
func Paginator(db *gorm.DB, pageOption *pageutil.PageOption) *gorm.DB {
	// limit offset
	return db.Limit(pageOption.Limit).Offset(pageOption.Offset)
}
