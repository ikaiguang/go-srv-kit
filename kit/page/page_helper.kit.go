package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

// DefaultPageRequest 默认分页请求
func DefaultPageRequest() *pagev1.PageRequest {
	return &pagev1.PageRequest{
		Page:     DefaultPageNumber,
		PageSize: DefaultPageSize,
	}
}

// CalcShowFrom 计算：分页显示开始位置
func CalcShowFrom(pageNumber, pageSize uint32) uint32 {
	return (pageNumber-1)*pageSize + 1
}

// CalcShowTo 计算：分页显示结束位置 长度
func CalcShowTo(showNumer, resultLength uint32) uint32 {
	if resultLength <= 1 {
		return showNumer
	}
	return showNumer + resultLength - 1
}

// HasNextPage 是否有下一页
func HasNextPage(totalPage, pageNumber uint32) bool {
	return totalPage > pageNumber
}
