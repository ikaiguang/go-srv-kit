package sortutil

import (
	"testing"
)

func TestInt32s(t *testing.T) {
	// int32
	var i32Slice = []int32{1, 3, 2}
	Int32s(i32Slice)
	t.Log("int32 sort : ", i32Slice)

	// int64
	var i64Slice = []int64{1, 3, 2}
	Int64s(i64Slice)
	t.Log("int64 sort : ", i64Slice)
}
