package gormutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./data/gorm -test.run=TestPaginate_ParseDirection
func TestPaginate_ParseDirection(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "#解析分页排序方向#unknown",
			given: "unknown",
			want:  DefaultOrderDesc,
		},
		{
			name:  "#解析分页排序方向#desc",
			given: "desc",
			want:  DefaultOrderDesc,
		},
		{
			name:  "#解析分页排序方向#DESC",
			given: "DESC",
			want:  DefaultOrderDesc,
		},
		{
			name:  "#解析分页排序方向#DeSC",
			given: "DeSC",
			want:  DefaultOrderDesc,
		},
		{
			name:  "#解析分页排序方向#asc",
			given: "asc",
			want:  DefaultOrderAsc,
		},
		{
			name:  "#解析分页排序方向#ASC",
			given: "ASC",
			want:  DefaultOrderAsc,
		},
		{
			name:  "#解析分页排序方向#AsC",
			given: "AsC",
			want:  DefaultOrderAsc,
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := ParseOrderDirection(param.given)
			require.Equal(t, param.want, got, "Direction")
		})
	}
}
