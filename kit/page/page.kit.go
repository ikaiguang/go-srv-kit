package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

const (
	DefaultPageNumber = 1  // goto page number : which page (default : 1)
	DefaultPageSize   = 20 // show records number (default : 20)
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
	Limit  int
	Offset int
}

// ConvertToPageOption 转换为分页选项
func ConvertToPageOption(pageRequest *pagev1.PageRequest) *PageOption {
	opts := &PageOption{
		Limit:  int(pageRequest.PageSize),
		Offset: int(pageRequest.PageSize * (pageRequest.Page - 1)),
	}
	return opts
}
