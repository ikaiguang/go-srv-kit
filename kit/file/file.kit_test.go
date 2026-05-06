package filepkg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdataDir = "./_testdata_temp"

// go test -v -count 1 ./kit/file -run TestMoveFileToDir
func TestMoveFileToDir(t *testing.T) {
	// 准备临时目录和文件
	srcDir := filepath.Join(testdataDir, "src")
	dstDir := filepath.Join(testdataDir, "dst")
	require.Nil(t, os.MkdirAll(srcDir, 0755))
	require.Nil(t, os.MkdirAll(dstDir, 0755))
	defer func() { _ = os.RemoveAll(testdataDir) }()

	// 创建源文件
	srcFile := filepath.Join(srcDir, "test.txt")
	require.Nil(t, os.WriteFile(srcFile, []byte("hello"), 0644))

	// 移动文件
	targetPath, err := MoveFileToDir(srcFile, dstDir)
	require.Nil(t, err)
	assert.Equal(t, filepath.Join(dstDir, "test.txt"), targetPath)

	// 验证目标文件存在
	data, err := os.ReadFile(targetPath)
	require.Nil(t, err)
	assert.Equal(t, []byte("hello"), data)

	// 验证源文件已不存在
	_, err = os.Stat(srcFile)
	assert.True(t, os.IsNotExist(err))
}

// go test -v -count 1 ./kit/file -run TestCopyFile
func TestCopyFile(t *testing.T) {
	require.Nil(t, os.MkdirAll(testdataDir, 0755))
	defer func() { _ = os.RemoveAll(testdataDir) }()

	srcFile := filepath.Join(testdataDir, "src.txt")
	dstFile := filepath.Join(testdataDir, "dst.txt")
	content := []byte("copy test content")

	require.Nil(t, os.WriteFile(srcFile, content, 0644))

	err := CopyFile(srcFile, dstFile)
	require.Nil(t, err)

	data, err := os.ReadFile(dstFile)
	require.Nil(t, err)
	assert.Equal(t, content, data)
}

func TestCopyFile_CreateDestDir(t *testing.T) {
	require.Nil(t, os.MkdirAll(testdataDir, 0755))
	defer func() { _ = os.RemoveAll(testdataDir) }()

	srcFile := filepath.Join(testdataDir, "src.txt")
	dstFile := filepath.Join(testdataDir, "nested", "dst.txt")
	content := []byte("copy test content")

	require.Nil(t, os.WriteFile(srcFile, content, 0644))
	require.Nil(t, CopyFile(srcFile, dstFile))

	data, err := os.ReadFile(dstFile)
	require.Nil(t, err)
	assert.Equal(t, content, data)
}

func TestMoveFileToDir_CreateDestDir(t *testing.T) {
	require.Nil(t, os.MkdirAll(testdataDir, 0755))
	defer func() { _ = os.RemoveAll(testdataDir) }()

	srcFile := filepath.Join(testdataDir, "src.txt")
	dstDir := filepath.Join(testdataDir, "nested", "dst")
	require.Nil(t, os.WriteFile(srcFile, []byte("move"), 0644))

	targetPath, err := MoveFileToDir(srcFile, dstDir)
	require.Nil(t, err)
	assert.Equal(t, filepath.Join(dstDir, "src.txt"), targetPath)
}

// go test -v -count 1 ./kit/file -run TestCopyFile_SrcNotExist
func TestCopyFile_SrcNotExist(t *testing.T) {
	err := CopyFile("/nonexistent/file.txt", "/tmp/dst.txt")
	assert.NotNil(t, err)
}

// go test -v -count 1 ./kit/file -run TestMoveFileToDir_SrcNotExist
func TestMoveFileToDir_SrcNotExist(t *testing.T) {
	_, err := MoveFileToDir("/nonexistent/file.txt", "/tmp")
	assert.NotNil(t, err)
}
