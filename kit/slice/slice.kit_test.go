package slicepkg

import "testing"

// go test -v -count 1 ./slice -run TestReverse
func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	last := s[len(s)-1]

	Reverse(s)

	if s[0] != last {
		t.Error("reverse func incorrect")
		return
	}
	t.Log(s)
}

// go test -v -count 1 ./slice -run TestReverseSlice
func TestReverseSlice(t *testing.T) {
	t.Run("int切片反转", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		ReverseSlice(s)
		expected := []int{5, 4, 3, 2, 1}
		for i, v := range s {
			if v != expected[i] {
				t.Errorf("ReverseSlice: got %v, want %v", s, expected)
				return
			}
		}
	})

	t.Run("string切片反转", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		ReverseSlice(s)
		expected := []string{"c", "b", "a"}
		for i, v := range s {
			if v != expected[i] {
				t.Errorf("ReverseSlice: got %v, want %v", s, expected)
				return
			}
		}
	})

	t.Run("空切片不panic", func(t *testing.T) {
		s := []int{}
		ReverseSlice(s)
		if len(s) != 0 {
			t.Error("ReverseSlice: empty slice should remain empty")
		}
	})

	t.Run("单元素切片", func(t *testing.T) {
		s := []int{42}
		ReverseSlice(s)
		if s[0] != 42 {
			t.Errorf("ReverseSlice: got %d, want 42", s[0])
		}
	})
}

// go test -v -count 1 ./slice -run TestContains
func TestContains(t *testing.T) {
	t.Run("存在的元素", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		if !Contains(s, "b") {
			t.Error("Contains: should find 'b'")
		}
	})

	t.Run("不存在的元素", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		if Contains(s, "d") {
			t.Error("Contains: should not find 'd'")
		}
	})

	t.Run("空切片返回false", func(t *testing.T) {
		s := []int{}
		if Contains(s, 1) {
			t.Error("Contains: empty slice should return false")
		}
	})

	t.Run("nil切片返回false", func(t *testing.T) {
		var s []int
		if Contains(s, 1) {
			t.Error("Contains: nil slice should return false")
		}
	})

	t.Run("int类型查找", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		if !Contains(s, 3) {
			t.Error("Contains: should find 3")
		}
		if Contains(s, 6) {
			t.Error("Contains: should not find 6")
		}
	})
}

// go test -v -count 1 ./slice -run TestInStringSlice_Deprecated
func TestInStringSlice_Deprecated(t *testing.T) {
	s := []string{"hello", "world"}
	if !InStringSlice(s, "hello") {
		t.Error("InStringSlice: should find 'hello'")
	}
	if InStringSlice(s, "foo") {
		t.Error("InStringSlice: should not find 'foo'")
	}
}
