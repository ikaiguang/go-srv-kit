package stringpkg

import "testing"

func TestToSnakeString(t *testing.T) {
	camelString := "XxYy"
	snakeString := "xx_yy"

	got := ToSnake(camelString)
	// check value
	if got != snakeString {
		t.Errorf("testing : ToSnake error : ToSnake(%s) != %s", camelString, snakeString)
		return
	}
	t.Log(got)
	t.Log(ToSnake("X_Y_Z"))
}

func TestToCamelString(t *testing.T) {
	camelString := "XxYy"
	snakeString := "xx_yy"

	got := ToCamel(snakeString)
	// check value
	if got != camelString {
		t.Errorf("testing : ToCamel error : ToCamel(%s) != %s", snakeString, camelString)
		return
	}
	t.Log(got)
	t.Log(ToCamel("a__b__c"))
}
