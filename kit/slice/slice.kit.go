package sliceutil

import (
	"reflect"
)

// Reverse reverse slice
// panic if s is not a slice
func Reverse(slice interface{}) {
	reflectValue := reflect.ValueOf(slice)

	swap := reflect.Swapper(slice)

	for i, j := 0, reflectValue.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
