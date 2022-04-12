package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

const (
	DefaultPageNumber    = 1      // goto page number : which page (default : 1)
	DefaultPageSize      = 20     // show records number (default : 20)
	DefaultPlaceholder   = "?"    // param placeholder
	DefaultOrderColumn   = "id"   // default order column
	DefaultDirectionAsc  = "asc"  // order direction : asc
	DefaultDirectionDesc = "desc" // order direction : desc
)

// DefaultPageRequest 默认分页请求
func DefaultPageRequest() *pagev1.PageRequest {
	return &pagev1.PageRequest{
		Page:     DefaultPageNumber,
		PageSize: DefaultPageSize,
	}
}

// PageOption .
type PageOption struct {
	Where  []*PageWhere
	Order  []*pagev1.PageOrder
	Limit  int
	Offset int
}

// ConvertToPageOption 转换为分页选项
func ConvertToPageOption(pageRequest *pagev1.PageRequest) *PageOption {
	opts := &PageOption{
		Where:  []*PageWhere{},
		Order:  []*pagev1.PageOrder{},
		Limit:  int(pageRequest.PageSize),
		Offset: int(pageRequest.PageSize * (pageRequest.Page - 1)),
	}
	return opts
}

// PageWhere 分页条件；例：where id = ?(where id = 1)
type PageWhere struct {
	// Column 字段
	Column string
	// Operator 运算符
	Operator string
	// Placeholder 占位符
	Placeholder string
	// Data 数据
	Data interface{}
}
