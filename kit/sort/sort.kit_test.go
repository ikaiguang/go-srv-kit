package sortpkg

import (
	"cmp"
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

func TestSort(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		s := []int32{3, 1, 2}
		Sort(s)
		if s[0] != 1 || s[1] != 2 || s[2] != 3 {
			t.Errorf("Sort[int32] 结果错误: %v", s)
		}
	})

	t.Run("int64", func(t *testing.T) {
		s := []int64{30, 10, 20}
		Sort(s)
		if s[0] != 10 || s[1] != 20 || s[2] != 30 {
			t.Errorf("Sort[int64] 结果错误: %v", s)
		}
	})

	t.Run("float64", func(t *testing.T) {
		s := []float64{3.3, 1.1, 2.2}
		Sort(s)
		if s[0] != 1.1 || s[1] != 2.2 || s[2] != 3.3 {
			t.Errorf("Sort[float64] 结果错误: %v", s)
		}
	})

	t.Run("string", func(t *testing.T) {
		s := []string{"c", "a", "b"}
		Sort(s)
		if s[0] != "a" || s[1] != "b" || s[2] != "c" {
			t.Errorf("Sort[string] 结果错误: %v", s)
		}
	})

	t.Run("空切片", func(t *testing.T) {
		var s []int
		Sort(s)
		if len(s) != 0 {
			t.Errorf("空切片排序后长度应为 0, 实际: %d", len(s))
		}
	})

	t.Run("nil切片", func(t *testing.T) {
		var s []int32
		Sort(s) // nil 切片排序不应 panic
	})
}

func TestSortFunc(t *testing.T) {
	t.Run("自定义比较-降序", func(t *testing.T) {
		s := []int{1, 3, 2}
		SortFunc(s, func(a, b int) int {
			return cmp.Compare(b, a) // 降序
		})
		if s[0] != 3 || s[1] != 2 || s[2] != 1 {
			t.Errorf("SortFunc 降序结果错误: %v", s)
		}
	})

	t.Run("空切片", func(t *testing.T) {
		var s []int
		SortFunc(s, func(a, b int) int {
			return cmp.Compare(a, b)
		})
		if len(s) != 0 {
			t.Errorf("空切片排序后长度应为 0, 实际: %d", len(s))
		}
	})
}

func TestDeprecatedConsistency(t *testing.T) {
	t.Run("Int32s与Sort[int32]结果一致", func(t *testing.T) {
		s1 := []int32{5, 3, 1, 4, 2}
		s2 := []int32{5, 3, 1, 4, 2}
		Int32s(s1)
		Sort(s2)
		for i := range s1 {
			if s1[i] != s2[i] {
				t.Errorf("Int32s 与 Sort[int32] 结果不一致: %v vs %v", s1, s2)
				break
			}
		}
	})

	t.Run("Int64s与Sort[int64]结果一致", func(t *testing.T) {
		s1 := []int64{5, 3, 1, 4, 2}
		s2 := []int64{5, 3, 1, 4, 2}
		Int64s(s1)
		Sort(s2)
		for i := range s1 {
			if s1[i] != s2[i] {
				t.Errorf("Int64s 与 Sort[int64] 结果不一致: %v vs %v", s1, s2)
				break
			}
		}
	})
}
