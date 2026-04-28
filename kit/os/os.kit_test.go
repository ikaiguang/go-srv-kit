package ospkg

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWindows(t *testing.T) {
	expected := runtime.GOOS == "windows"
	assert.Equal(t, expected, IsWindows())
}
