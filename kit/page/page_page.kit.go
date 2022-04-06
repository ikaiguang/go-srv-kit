package pageutil

import (
	"google.golang.org/protobuf/proto"

	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

var (
	pageHandler = &page{}
)

// page 分页模式
type page struct{}

// ParsePageRequest 解析分页请求
func (s *page) ParsePageRequest(pageRequest *pagev1.PageRequest) *pagev1.PageRequest {
	if pageRequest == nil {
		return DefaultPageRequest()
	}
	return s.parsePageRequest(pageRequest)
}

// ConvertToPageOption 转换为分页选项
func (s *page) ConvertToPageOption(pageRequest *pagev1.PageRequest) *Options {
	opts := &Options{
		Where:  []*Where{},
		Order:  []*pagev1.PagingOrder{},
		Limit:  pageRequest.PageSize,
		Offset: pageRequest.PageSize * (pageRequest.Page - 1),
	}
	return opts
}

// parsePageRequest 解析分页请求
func (s *page) parsePageRequest(pageRequest *pagev1.PageRequest) *pagev1.PageRequest {
	pageRequest = proto.Clone(pageRequest).(*pagev1.PageRequest)

	pageRequest.Page = ParsePage(pageRequest.Page)
	pageRequest.PageSize = ParsePageSize(pageRequest.PageSize)

	return pageRequest
}
