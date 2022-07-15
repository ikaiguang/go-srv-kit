package sliceutil

import (
	"reflect"
)

// Reverse 反转slice顺序
// panic if s is not a slice
func Reverse(slice interface{}) {
	reflectValue := reflect.ValueOf(slice)

	swap := reflect.Swapper(slice)

	for i, j := 0, reflectValue.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// InStringSlice 在数组内
func InStringSlice(s []string, dest string) bool {
	for i := range s {
		if s[i] == dest {
			return true
		}
	}
	return false
}
