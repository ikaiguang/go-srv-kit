package filepathpkg

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./pkg/filepath -run TestWaldDir
func TestWaldDir(t *testing.T) {
	rootPath := "./../../pkg"
	fp, fi, err := WaldDir(rootPath)
	require.Nil(t, err)
	for i := range fi {
		relPath, err := filepath.Rel(rootPath, fp[i])
		require.Nil(t, err)
		t.Logf("name=%s rel=%s path=%s\n", fi[i].Name(), relPath, fp[i])
	}
}

// go test -v -count=1 ./kit/filepath -run TestReadDir
func TestReadDir(t *testing.T) {
	rootPath := "./../"
	fi, err := ReadDir(rootPath)
	require.Nil(t, err)
	for i := range fi {
		t.Logf("name=%s", fi[i].Name())
	}
}
