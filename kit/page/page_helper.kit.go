package pageutil

import (
	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
)

// HasNextPage 是否有下一页
func HasNextPage(pageResponse *pagev1.PageResponse) bool {
	return pageResponse.TotalPage > int64(pageResponse.Page)
}

// CalcPageResponse 计算分页响应
func CalcPageResponse(pageRequest *pagev1.PageRequest, totalNumber int64) *pagev1.PageResponse {
	pageResponse := &pagev1.PageResponse{
		TotalNumber: totalNumber,
		Page:        pageRequest.Page,
		PageSize:    pageRequest.PageSize,
	}
	if pageResponse.TotalNumber <= 0 || pageResponse.PageSize <= 0 {
		return pageResponse
	}
	// TotalPage
	pageResponse.TotalPage = pageResponse.TotalNumber / int64(pageResponse.PageSize)
	if pageResponse.TotalNumber%int64(pageResponse.PageSize) != 0 {
		pageResponse.TotalPage += 1
	}
	return pageResponse
}

// CalcShowFrom 计算：分页显示开始位置
func CalcShowFrom(pageNumber, pageSize uint32) uint32 {
	return (pageNumber-1)*pageSize + 1
}

// CalcShowTo 计算：分页显示结束位置 长度
func CalcShowTo(showFromNumer, resultLength uint32) uint32 {
	if resultLength <= 1 {
		return showFromNumer
	}
	return showFromNumer + resultLength - 1
}
