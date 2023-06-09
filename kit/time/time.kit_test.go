package timepkg

import (
	"testing"
	"time"
)

// go test -v -count=1 ./kit/time -run=TestFormat
func TestFormat(t *testing.T) {
	tNow := now()

	t.Log("tNow.Year() :", tNow.Year())
	t.Log(FormatRFC3339(tNow))
	t.Logf("==> time.RFC3339(%s) = %s\n", time.RFC3339, tNow.Format(time.RFC3339))
	t.Logf("==> YmdHmsTZ(%s) = %s\n", YmdHmsTZ, tNow.Format(YmdHmsTZ))
}
