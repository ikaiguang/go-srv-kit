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

// PageOptions .
type PageOptions struct {
	Where  []*PageWhere
	Order  []*pagev1.PageOrder
	Limit  uint32
	Offset uint32
}

// ConvertToPageOption 转换为分页选项
func ConvertToPageOption(pageRequest *pagev1.PageRequest) *PageOptions {
	opts := &PageOptions{
		Where:  []*PageWhere{},
		Order:  []*pagev1.PageOrder{},
		Limit:  pageRequest.PageSize,
		Offset: pageRequest.PageSize * (pageRequest.Page - 1),
	}
	return opts
}

// PageWhere 分页条件；例：where id = ?(where id = 1)
type PageWhere struct {
	// Column 字段
	Column string
	// Condition 条件
	Condition string
	// Placeholder 占位符
	Placeholder string
	// Data 数据
	Data interface{}
}
