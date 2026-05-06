package ospkg

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./os -run TestIsWindows
func TestIsWindows(t *testing.T) {
	expected := runtime.GOOS == "windows"
	assert.Equal(t, expected, IsWindows())
}
