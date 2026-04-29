package gormpkg

import (
	"testing"

	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
	"github.com/stretchr/testify/require"
)

// go test -v ./data/gorm/ -count=1 -run TestPaginatorArgs
func TestPaginatorArgs(t *testing.T) {
	tests := []struct {
		name     string
		page     uint32
		pageSize uint32
	}{
		{
			name:     "#初始化分页参数#第1页",
			page:     1,
			pageSize: 10,
		},
		{
			name:     "#初始化分页参数#第2页",
			page:     2,
			pageSize: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := InitPaginatorArgs(tt.page, tt.pageSize)
			require.NotNil(t, args)
			require.NotNil(t, args.PageRequest)
			require.NotNil(t, args.PageOption)
		})
	}
}

// go test -v ./data/gorm/ -count=1 -run TestCalcPageResponse
func TestCalcPageResponse(t *testing.T) {
	tests := []struct {
		name      string
		page      uint32
		pageSize  uint32
		total     uint32
		wantPage  uint32
		wantTotal uint32
	}{
		{
			name:      "#计算分页响应#总数为0",
			page:      1,
			pageSize:  10,
			total:     0,
			wantPage:  1,
			wantTotal: 0,
		},
		{
			name:      "#计算分页响应#总数小于页大小",
			page:      1,
			pageSize:  10,
			total:     5,
			wantPage:  1,
			wantTotal: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageReq := &pagepkg.PageRequest{
				Page:     tt.page,
				PageSize: tt.pageSize,
			}
			req, _ := pagepkg.ParsePageRequest(pageReq)
			resp := pagepkg.CalcPageResponse(req, tt.total)
			require.NotNil(t, resp)
			require.Equal(t, tt.wantTotal, resp.TotalNumber)
		})
	}
}
