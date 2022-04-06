package pageutil

import (
	"testing"

	pagev1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./kit/page -test.run=TestPaginate_MakePageOptions
func TestPaginate_MakePageOptions(t *testing.T) {
	defaultRequest := DefaultPageRequest()
	pageRequestFor10 := &pagev1.PageRequest{
		Page:     10,
		PageSize: 10,
	}

	tests := []struct {
		name        string
		pageRequest *pagev1.PageRequest
		want        *Options
	}{
		{
			name:        "#准备分页选项#defaut",
			pageRequest: defaultRequest,
			want: &Options{
				Limit:  defaultRequest.PageSize,
				Offset: defaultRequest.PageSize * (defaultRequest.Page - 1),
			},
		},
		{
			name:        "#准备分页选项#每页10条之第10页",
			pageRequest: pageRequestFor10,
			want: &Options{
				Limit:  pageRequestFor10.PageSize,
				Offset: pageRequestFor10.PageSize * (pageRequestFor10.Page - 1),
			},
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			_, got := ParsePageRequest(param.pageRequest)
			assert.Equal(t, param.want.Limit, got.Limit, "Limit")
			assert.Equal(t, param.want.Offset, got.Offset, "Offset")
		})
	}
}

// go test -v -count=1 ./kit/page -test.run=TestPaginate_ParsePageRequest
func TestPaginate_ParsePageRequest(t *testing.T) {
	tests := []struct {
		name  string
		given *pagev1.PageRequest
		want  *pagev1.PageRequest
	}{
		{
			name:  "#解析分页请求#nil",
			given: nil,
			want: &pagev1.PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#default",
			given: &pagev1.PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
			want: &pagev1.PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#zero",
			given: &pagev1.PageRequest{
				Page:     0,
				PageSize: 0,
			},
			want: &pagev1.PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#custom",
			given: &pagev1.PageRequest{
				Page:     2,
				PageSize: 300,
			},
			want: &pagev1.PageRequest{
				Page:     2,
				PageSize: 300,
			},
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := pageHandler.ParsePageRequest(param.given)
			assert.Equal(t, param.want.Page, got.Page, "Page")
			assert.Equal(t, param.want.PageSize, got.PageSize, "PageSize")
			assert.Equal(t, len(param.want.OrderByArray), len(got.OrderByArray), "OrderByArray")
		})
	}
}

// go test -v -count=1 ./kit/page -test.run=TestPaginate_ParseDirection
func TestPaginate_ParseDirection(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "#解析分页排序方向#unknown",
			given: "unknown",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#desc",
			given: "desc",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#DESC",
			given: "DESC",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#DeSC",
			given: "DeSC",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#asc",
			given: "asc",
			want:  DefaultDirectionAsc,
		},
		{
			name:  "#解析分页排序方向#ASC",
			given: "ASC",
			want:  DefaultDirectionAsc,
		},
		{
			name:  "#解析分页排序方向#AsC",
			given: "AsC",
			want:  DefaultDirectionAsc,
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := ParseOrderDirection(param.given)
			assert.Equal(t, param.want, got, "Direction")
		})
	}
}

// go test -v -count=1 ./kit/page -test.run=TestPaginate_DefaultPageRequest
func TestPaginate_DefaultPageRequest(t *testing.T) {
	tests := []struct {
		name string
		want *pagev1.PageRequest
	}{
		{
			name: "#默认的分页请求",
			want: &pagev1.PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
	}

	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := DefaultPageRequest()
			assert.Equal(t, param.want.Page, got.Page, "Page")
			assert.Equal(t, param.want.PageSize, got.PageSize, "PageSize")
			assert.Equal(t, len(param.want.OrderByArray), len(got.OrderByArray), "OrderByArray")
		})
	}
}
