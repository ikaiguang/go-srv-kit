package timeutil

import (
	"testing"
)

func TestFormat(t *testing.T) {
	tNow := now()

	t.Log(tNow.Year())
	t.Log(FormatRFC3339(tNow))
}
