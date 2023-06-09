package pagepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./kit/page -test.run=TestPaginate_ConvertToPageOption
func TestPaginate_ConvertToPageOption(t *testing.T) {
	defaultRequest := DefaultPageRequest()
	pageRequestFor10 := &PageRequest{
		Page:     10,
		PageSize: 10,
	}

	tests := []struct {
		name        string
		pageRequest *PageRequest
		want        *PageOption
	}{
		{
			name:        "#准备分页选项#defaut",
			pageRequest: defaultRequest,
			want: &PageOption{
				Limit:  int(defaultRequest.PageSize),
				Offset: int(defaultRequest.PageSize * (defaultRequest.Page - 1)),
			},
		},
		{
			name:        "#准备分页选项#每页10条之第10页",
			pageRequest: pageRequestFor10,
			want: &PageOption{
				Limit:  int(pageRequestFor10.PageSize),
				Offset: int(pageRequestFor10.PageSize * (pageRequestFor10.Page - 1)),
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
		given *PageRequest
		want  *PageRequest
	}{
		{
			name:  "#解析分页请求#nil",
			given: nil,
			want: &PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#default",
			given: &PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
			want: &PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#zero",
			given: &PageRequest{
				Page:     0,
				PageSize: 0,
			},
			want: &PageRequest{
				Page:     DefaultPageNumber,
				PageSize: DefaultPageSize,
			},
		},
		{
			name: "#解析分页请求#custom",
			given: &PageRequest{
				Page:     2,
				PageSize: 300,
			},
			want: &PageRequest{
				Page:     2,
				PageSize: 300,
			},
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got, _ := ParsePageRequest(param.given)
			assert.Equal(t, param.want.Page, got.Page, "Page")
			assert.Equal(t, param.want.PageSize, got.PageSize, "PageSize")
		})
	}
}

// go test -v -count=1 ./kit/page -test.run=TestPaginate_DefaultPageRequest
func TestPaginate_DefaultPageRequest(t *testing.T) {
	tests := []struct {
		name string
		want *PageRequest
	}{
		{
			name: "#默认的分页请求",
			want: &PageRequest{
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
		})
	}
}
