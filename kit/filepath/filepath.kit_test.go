package filepathpkg

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count 1 ./filepath -run TestWalkDir
func TestWalkDir(t *testing.T) {
	rootPath := "./../"
	fp, de, err := WalkDir(rootPath)
	require.Nil(t, err)
	require.NotEmpty(t, fp)
	require.Equal(t, len(fp), len(de))
	for i := range de {
		relPath, err := filepath.Rel(rootPath, fp[i])
		require.Nil(t, err)
		t.Logf("name=%s rel=%s path=%s\n", de[i].Name(), relPath, fp[i])
	}
}

// go test -v -count 1 ./filepath -run TestWalkDir_NonExistentPath
func TestWalkDir_NonExistentPath(t *testing.T) {
	rootPath := "./non_existent_path"
	_, _, err := WalkDir(rootPath)
	require.NotNil(t, err)
}

// go test -v -count 1 ./filepath -run TestWaldDir
func TestWaldDir(t *testing.T) {
	rootPath := "./../"
	fp, fi, err := WaldDir(rootPath)
	require.Nil(t, err)
	require.NotEmpty(t, fp)
	require.Equal(t, len(fp), len(fi))
	for i := range fi {
		relPath, err := filepath.Rel(rootPath, fp[i])
		require.Nil(t, err)
		t.Logf("name=%s rel=%s path=%s\n", fi[i].Name(), relPath, fp[i])
	}
}

// go test -v -count 1 ./filepath -run TestReadDirEntries
func TestReadDirEntries(t *testing.T) {
	rootPath := "./../"
	de, err := ReadDirEntries(rootPath)
	require.Nil(t, err)
	require.NotEmpty(t, de)
	for i := range de {
		t.Logf("name=%s isDir=%v", de[i].Name(), de[i].IsDir())
	}
}

// go test -v -count 1 ./filepath -run TestReadDirEntries_NonExistentPath
func TestReadDirEntries_NonExistentPath(t *testing.T) {
	rootPath := "./non_existent_path"
	_, err := ReadDirEntries(rootPath)
	require.NotNil(t, err)
}

// go test -v -count 1 ./filepath -run TestReadDir
func TestReadDir(t *testing.T) {
	rootPath := "./../"
	fi, err := ReadDir(rootPath)
	require.Nil(t, err)
	require.NotEmpty(t, fi)
	for i := range fi {
		t.Logf("name=%s", fi[i].Name())
	}
}
