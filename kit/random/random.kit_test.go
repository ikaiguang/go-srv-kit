package randompkg

import "testing"

func TestStrings(t *testing.T) {
	t.Log(Strings(3))
	t.Log(Strings(3))
}

func TestLetter(t *testing.T) {
	t.Log(Letter(3))
	t.Log(Letter(3))
}

func TestNumeric(t *testing.T) {
	t.Log(Numeric(3))
	t.Log(Numeric(3))
}

func TestBetweenNumeric(t *testing.T) {
	t.Log(NumericBetween(1, 5))
	t.Log(NumericBetween(5, 1))
}
