package pathpkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./path/ -count=1 -run TestPath
func TestPath(t *testing.T) {
	currentPath := filepath.Join(
		os.Getenv("GOPATH"),
		"src",
		"github.com/ikaiguang/go-srv-kit/path",
	)
	path := Path()
	t.Log("==> path :", path)
	require.Equal(t, path, currentPath)
}
