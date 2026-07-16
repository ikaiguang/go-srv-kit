package zippkg

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdataDir = "./_testdata_temp"

func setup(t *testing.T) {
	t.Helper()
	require.Nil(t, os.MkdirAll(testdataDir, 0755))
}

func teardown() {
	_ = os.RemoveAll(testdataDir)
}

// go test -v -count 1 ./kit/zip -run TestZipFile
func TestZipFile(t *testing.T) {
	setup(t)
	defer teardown()

	// 创建源文件
	srcFile := filepath.Join(testdataDir, "test.txt")
	require.Nil(t, os.WriteFile(srcFile, []byte("hello zip"), 0644))

	zipPath := filepath.Join(testdataDir, "test.txt.zip")
	err := ZipFile(srcFile, zipPath)
	require.Nil(t, err)

	// 验证 zip 文件存在
	info, err := os.Stat(zipPath)
	require.Nil(t, err)
	assert.Greater(t, info.Size(), int64(0))
}

func TestZipFile_CreateDestDir(t *testing.T) {
	setup(t)
	defer teardown()

	srcFile := filepath.Join(testdataDir, "test.txt")
	require.Nil(t, os.WriteFile(srcFile, []byte("hello zip"), 0644))

	zipPath := filepath.Join(testdataDir, "nested", "test.txt.zip")
	err := ZipFile(srcFile, zipPath)
	require.Nil(t, err)

	info, err := os.Stat(zipPath)
	require.Nil(t, err)
	assert.Greater(t, info.Size(), int64(0))
}

// go test -v -count 1 ./kit/zip -run TestZip_Directory
func TestZip_Directory(t *testing.T) {
	setup(t)
	defer teardown()

	// 创建目录结构
	subDir := filepath.Join(testdataDir, "src", "sub")
	require.Nil(t, os.MkdirAll(subDir, 0755))
	require.Nil(t, os.WriteFile(filepath.Join(testdataDir, "src", "a.txt"), []byte("file a"), 0644))
	require.Nil(t, os.WriteFile(filepath.Join(subDir, "b.txt"), []byte("file b"), 0644))

	zipPath := filepath.Join(testdataDir, "src.zip")
	err := Zip(filepath.Join(testdataDir, "src"), zipPath)
	require.Nil(t, err)

	info, err := os.Stat(zipPath)
	require.Nil(t, err)
	assert.Greater(t, info.Size(), int64(0))
}

func TestZipExcludesDestinationInsideSource(t *testing.T) {
	root := t.TempDir()
	sourceDir := filepath.Join(root, "source")
	require.NoError(t, os.MkdirAll(sourceDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(sourceDir, "input.txt"), []byte("input"), 0o644))

	zipPath := filepath.Join(sourceDir, "archive.zip")
	require.NoError(t, Zip(sourceDir, zipPath))

	reader, err := zip.OpenReader(zipPath)
	require.NoError(t, err)
	defer func() { _ = reader.Close() }()

	var names []string
	for _, entry := range reader.File {
		names = append(names, entry.Name)
	}
	assert.Contains(t, names, "input.txt")
	assert.NotContains(t, names, "archive.zip")
}

func TestZipFileRejectsSourceAsDestination(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "input.txt")
	require.NoError(t, os.WriteFile(filePath, []byte("input"), 0o644))

	err := ZipFile(filePath, filePath)
	require.Error(t, err)
	content, readErr := os.ReadFile(filePath)
	require.NoError(t, readErr)
	assert.Equal(t, []byte("input"), content)
}

// go test -v -count 1 ./kit/zip -run TestZip_SingleFile
func TestZip_SingleFile(t *testing.T) {
	setup(t)
	defer teardown()

	srcFile := filepath.Join(testdataDir, "single.txt")
	require.Nil(t, os.WriteFile(srcFile, []byte("single file"), 0644))

	zipPath := filepath.Join(testdataDir, "single.txt.zip")
	err := Zip(srcFile, zipPath)
	require.Nil(t, err)

	info, err := os.Stat(zipPath)
	require.Nil(t, err)
	assert.Greater(t, info.Size(), int64(0))
}

// go test -v -count 1 ./kit/zip -run TestUnzip
func TestUnzip(t *testing.T) {
	setup(t)
	defer teardown()

	// 先创建一个 zip 文件
	srcFile := filepath.Join(testdataDir, "unzip_test.txt")
	content := []byte("content for unzip test")
	require.Nil(t, os.WriteFile(srcFile, content, 0644))

	zipPath := filepath.Join(testdataDir, "unzip_test.zip")
	require.Nil(t, ZipFile(srcFile, zipPath))

	// 解压
	unzipDir := filepath.Join(testdataDir, "unzipped")
	require.Nil(t, os.MkdirAll(unzipDir, 0755))
	err := Unzip(zipPath, unzipDir)
	require.Nil(t, err)

	// 验证解压后的文件
	data, err := os.ReadFile(filepath.Join(unzipDir, "unzip_test.txt"))
	require.Nil(t, err)
	assert.Equal(t, content, data)
}

// go test -v -count 1 ./kit/zip -run TestZipFile_SrcNotExist
func TestZipFile_SrcNotExist(t *testing.T) {
	setup(t)
	defer teardown()

	zipPath := filepath.Join(testdataDir, "notexist.zip")
	err := ZipFile("/nonexistent/file.txt", zipPath)
	assert.NotNil(t, err)
}

// go test -v -count 1 ./kit/zip -run TestUnzip_ZipNotExist
func TestUnzip_ZipNotExist(t *testing.T) {
	err := Unzip("/nonexistent/file.zip", testdataDir)
	assert.NotNil(t, err)
}

func TestUnzip_RejectsZipSlip(t *testing.T) {
	setup(t)
	defer teardown()

	zipPath := filepath.Join(testdataDir, "evil.zip")
	file, err := os.Create(zipPath)
	require.Nil(t, err)
	zipWriter := zip.NewWriter(file)
	w, err := zipWriter.Create("../evil.txt")
	require.Nil(t, err)
	_, err = w.Write([]byte("evil"))
	require.Nil(t, err)
	require.Nil(t, zipWriter.Close())
	require.Nil(t, file.Close())

	unzipDir := filepath.Join(testdataDir, "unzipped")
	err = Unzip(zipPath, unzipDir)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "illegal file path")

	_, err = os.Stat(filepath.Join(testdataDir, "evil.txt"))
	assert.True(t, os.IsNotExist(err))
}

func TestUnzipCreatesImplicitParentDirectories(t *testing.T) {
	root := t.TempDir()
	zipPath := filepath.Join(root, "nested.zip")
	file, err := os.Create(zipPath)
	require.NoError(t, err)
	zipWriter := zip.NewWriter(file)
	entry, err := zipWriter.Create("nested/path/file.txt")
	require.NoError(t, err)
	_, err = entry.Write([]byte("nested"))
	require.NoError(t, err)
	require.NoError(t, zipWriter.Close())
	require.NoError(t, file.Close())

	dest := filepath.Join(root, "output")
	require.NoError(t, Unzip(zipPath, dest))
	content, err := os.ReadFile(filepath.Join(dest, "nested", "path", "file.txt"))
	require.NoError(t, err)
	assert.Equal(t, []byte("nested"), content)

	require.Error(t, ExtractZipEntry(nil, dest))
}

func TestUnzipRejectsDestinationSymlinkEscape(t *testing.T) {
	root := t.TempDir()
	dest := filepath.Join(root, "output")
	outside := filepath.Join(root, "outside")
	require.NoError(t, os.MkdirAll(dest, 0o755))
	require.NoError(t, os.MkdirAll(outside, 0o755))
	if err := os.Symlink(outside, filepath.Join(dest, "nested")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}

	zipPath := filepath.Join(root, "symlink-escape.zip")
	file, err := os.Create(zipPath)
	require.NoError(t, err)
	zipWriter := zip.NewWriter(file)
	entry, err := zipWriter.Create("nested/escape.txt")
	require.NoError(t, err)
	_, err = entry.Write([]byte("escape"))
	require.NoError(t, err)
	require.NoError(t, zipWriter.Close())
	require.NoError(t, file.Close())

	err = Unzip(zipPath, dest)
	require.Error(t, err)
	_, statErr := os.Stat(filepath.Join(outside, "escape.txt"))
	assert.True(t, os.IsNotExist(statErr))
}
