package sliceutil

import "testing"

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
