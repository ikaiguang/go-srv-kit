package pagetestdata

import (
	"testing"

	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
	pageutil "github.com/ikaiguang/go-srv-kit/kit/page"
)

// go test -v -count=1 ./kit/page/testdata -test.run=TestPage_Page
func TestPaging_Page(t *testing.T) {
	data := []struct {
		name        string
		pageRequest *pagev1.PageRequest
	}{
		{
			name: "#页码分页：页码为：1",
			pageRequest: &pagev1.PageRequest{
				Page:     1,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：2",
			pageRequest: &pagev1.PageRequest{
				Page:     2,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：5",
			pageRequest: &pagev1.PageRequest{
				Page:     5,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：6",
			pageRequest: &pagev1.PageRequest{
				Page:     6,
				PageSize: 2,
			},
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			pageRequest, pageOptions := pageutil.ParsePageRequest(tt.pageRequest)
			_ = pageRequest
			_ = pageOptions
		})
	}
}

// go test -v -count=1 ./kit/page/testdata -test.run=TestPaging_Cursor
func TestPaging_Cursor(t *testing.T) {

}
