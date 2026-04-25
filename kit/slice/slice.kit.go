package slicepkg

import (
	"reflect"
)

// Deprecated: 使用 ReverseSlice 替代。
// 注意：Reverse 使用 reflect，性能较差且缺乏编译期类型检查。
func Reverse(slice interface{}) {
	reflectValue := reflect.ValueOf(slice)

	swap := reflect.Swapper(slice)

	for i, j := 0, reflectValue.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// Deprecated: 使用 Contains 替代
func InStringSlice(s []string, dest string) bool {
	return Contains(s, dest)
}
