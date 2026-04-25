package sortpkg

import (
	"cmp"
	"slices"
)

// Sort 对有序类型切片排序
func Sort[T cmp.Ordered](s []T) {
	slices.Sort(s)
}

// SortFunc 使用自定义比较函数排序
func SortFunc[T any](s []T, less func(a, b T) int) {
	slices.SortFunc(s, less)
}
