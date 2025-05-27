package pagetestdata

import (
	"testing"

	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
)

// go test -v -count=1 ./kit/page/testdata -run TestPage_Page
func TestPaging_Page(t *testing.T) {
	data := []struct {
		name        string
		pageRequest *pagepkg.PageRequest
	}{
		{
			name: "#页码分页：页码为：1",
			pageRequest: &pagepkg.PageRequest{
				Page:     1,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：2",
			pageRequest: &pagepkg.PageRequest{
				Page:     2,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：5",
			pageRequest: &pagepkg.PageRequest{
				Page:     5,
				PageSize: 2,
			},
		},
		{
			name: "#页码分页：页码为：6",
			pageRequest: &pagepkg.PageRequest{
				Page:     6,
				PageSize: 2,
			},
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			pageRequest, pageOptions := pagepkg.ParsePageRequest(tt.pageRequest)
			_ = pageRequest
			_ = pageOptions
		})
	}
}

// go test -v -count=1 ./kit/page/testdata -run TestPaging_Cursor
func TestPaging_Cursor(t *testing.T) {

}
