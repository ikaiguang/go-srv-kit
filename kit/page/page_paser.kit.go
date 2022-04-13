package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
	"google.golang.org/protobuf/proto"
)

// ParsePageRequest 解析页码分页请求
func ParsePageRequest(pageRequest *pagev1.PageRequest) (*pagev1.PageRequest, *PageOption) {
	if pageRequest == nil {
		pageRequest = DefaultPageRequest()
	} else {
		pageRequest = proto.Clone(pageRequest).(*pagev1.PageRequest)
		pageRequest.Page = ParsePage(pageRequest.Page)
		pageRequest.PageSize = ParsePageSize(pageRequest.PageSize)
	}

	opt := ConvertToPageOption(pageRequest)
	return pageRequest, opt
}

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
