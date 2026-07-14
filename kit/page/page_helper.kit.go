package pagepkg

// HasNextPage 是否有下一页
func HasNextPage(pageResponse *PageResponse) bool {
	if pageResponse == nil {
		return false
	}
	return pageResponse.TotalPage > pageResponse.Page
}

// CalcPageResponse 计算分页响应
func CalcPageResponse(pageRequest *PageRequest, totalNumber uint32) *PageResponse {
	if pageRequest == nil {
		pageRequest = DefaultPageRequest()
	}
	pageResponse := &PageResponse{
		TotalNumber: totalNumber,
		Page:        pageRequest.Page,
		PageSize:    pageRequest.PageSize,
	}
	if pageResponse.TotalNumber <= 0 || pageResponse.PageSize <= 0 {
		return pageResponse
	}
	// TotalPage
	pageResponse.TotalPage = pageResponse.TotalNumber / pageResponse.PageSize
	if pageResponse.TotalNumber%pageResponse.PageSize != 0 {
		pageResponse.TotalPage += 1
	}
	return pageResponse
}

// CalcShowFrom 计算：分页显示开始位置
func CalcShowFrom(pageNumber, pageSize uint32) uint32 {
	pageNumber = ParsePage(pageNumber)
	pageSize = ParsePageSize(pageSize)
	return (pageNumber-1)*pageSize + 1
}

// CalcShowTo 计算：分页显示结束位置 长度
func CalcShowTo(showFromNumber, resultLength uint32) uint32 {
	if resultLength <= 1 {
		return showFromNumber
	}
	return showFromNumber + resultLength - 1
}
