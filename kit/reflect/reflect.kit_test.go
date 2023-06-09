package reflectpkg

import "testing"

func TestIsDefaultValue(t *testing.T) {
	var i complex64

	if !IsDefaultValue(i) {
		t.Error("error func")
	}
}
