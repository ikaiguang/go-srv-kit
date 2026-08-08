package gormpkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./data/gorm -run TestPaginate_ParseDirection
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

func TestIsValidColumnName(t *testing.T) {
	for _, field := range []string{"id", "users.id", "schema.users.id"} {
		require.True(t, IsValidColumnName(field), "field = %q", field)
	}
	for _, field := range []string{"", ".id", "users.", "users..id", "id desc", "id;drop"} {
		require.False(t, IsValidColumnName(field), "field = %q", field)
	}
}
