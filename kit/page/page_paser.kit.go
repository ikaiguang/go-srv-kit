package pagepkg

import (
	"google.golang.org/protobuf/proto"
)

// ParsePageRequest 解析页码分页请求
func ParsePageRequest(pageRequest *PageRequest) (*PageRequest, *PageOption) {
	if pageRequest == nil {
		pageRequest = DefaultPageRequest()
	} else {
		pageRequest = proto.Clone(pageRequest).(*PageRequest)
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
