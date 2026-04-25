package slicepkg

import "slices"

// ReverseSlice 原地反转切片
func ReverseSlice[T any](s []T) {
	slices.Reverse(s)
}

// Contains 检查元素是否存在于切片中
func Contains[T comparable](s []T, v T) bool {
	return slices.Contains(s, v)
}
