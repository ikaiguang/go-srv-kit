package gormpkg

import (
	"gorm.io/gorm"

	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

// PaginatorArgs 列表参数
type PaginatorArgs struct {
	// PageRequest 分页
	PageRequest *pagepkg.PageRequest
	// PageOption 分页
	PageOption *pagepkg.PageOption
	// PageOrders 排序
	PageOrders []*Order
	// PageWheres 条件
	PageWheres []*Where
}

func InitPaginatorArgs(page, pageSize uint32) *PaginatorArgs {
	res := &PaginatorArgs{}
	res.PageRequest, res.PageOption = pagepkg.ParsePageRequestArgs(page, pageSize)
	return res
}

func InitPaginatorArgsByPageRequest(pageRequest *pagepkg.PageRequest) *PaginatorArgs {
	res := &PaginatorArgs{}
	res.PageRequest, res.PageOption = pagepkg.ParsePageRequest(pageRequest)
	return res
}

// Paginator 分页
func Paginator(db *gorm.DB, pageOption *pagepkg.PageOption) *gorm.DB {
	// limit offset
	return db.Limit(pageOption.Limit).Offset(pageOption.Offset)
}
