package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

const (
	DefaultCurrentPageNumber = 0      // current page number : which page (default : 1)
	DefaultPageNumber        = 1      // goto page number : which page (default : 1)
	DefaultPageSize          = 20     // show records number (default : 15)
	DefaultCursorValue       = 0      // cursor value (default : 0)
	DefaultWherePlaceholder  = "?"    // where param placeholder
	DefaultOrderColumn       = "id"   // default order column
	DefaultDirectionAsc      = "asc"  // order direction : asc
	DefaultDirectionDesc     = "desc" // order direction : desc
)

// Options .
type Options struct {
	Where  []*Where
	Order  []*pagev1.PagingOrder
	Limit  uint32
	Offset uint32
}

// Where 分页条件；例：where id = ?(where id = 1)
type Where struct {
	// Column 字段
	Column string
	// Symbol 条件
	Symbol string
	// Placeholder 占位符
	Placeholder string
	// Data 数据
	Data interface{}
}
