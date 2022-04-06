package pageutil

import (
	"strings"

	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

// ParsePageRequest 解析页码分页请求
func ParsePageRequest(pageRequest *pagev1.PageRequest) (*pagev1.PageRequest, *Options) {
	pageRequest = pageHandler.ParsePageRequest(pageRequest)
	opt := pageHandler.ConvertToPageOption(pageRequest)
	return pageRequest, opt
}

// ParseCursorRequest 解析游标分页请求
//func ParseCursorRequest(cursorRequest *pagev1.CursorRequest) (*pagev1.CursorRequest, *Options) {
//	cursorRequest = cursorHandler.ParsePageRequest(cursorRequest)
//	opt := pageHandler.ConvertToPageOption(cursorRequest)
//	return cursorRequest, opt
//}

// ParsePage 页码
func ParsePage(pageNumber uint32) uint32 {
	if pageNumber < 1 {
		return DefaultPageNumber
	}
	return pageNumber
}

// ParsePageSize 每页显示多少条
func ParsePageSize(pageSize uint32) uint32 {
	if pageSize < 1 {
		return DefaultPageSize
	}
	return pageSize
}

// ParseOrderDirection 排序方向
func ParseOrderDirection(orderDirection string) string {
	if orderDirection = strings.ToLower(orderDirection); orderDirection == DefaultDirectionAsc {
		return DefaultDirectionAsc
	}
	return DefaultDirectionDesc
}
