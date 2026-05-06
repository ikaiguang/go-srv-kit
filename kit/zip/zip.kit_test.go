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

// go test -v -count 1 ./zip -run TestZipFile
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

// go test -v -count 1 ./zip -run TestZipFile_CreateDestDir
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

// go test -v -count 1 ./zip -run TestZip_Directory
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

// go test -v -count 1 ./zip -run TestZip_SingleFile
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

// go test -v -count 1 ./zip -run TestUnzip
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

// go test -v -count 1 ./zip -run TestZipFile_SrcNotExist
func TestZipFile_SrcNotExist(t *testing.T) {
	setup(t)
	defer teardown()

	zipPath := filepath.Join(testdataDir, "notexist.zip")
	err := ZipFile("/nonexistent/file.txt", zipPath)
	assert.NotNil(t, err)
}

// go test -v -count 1 ./zip -run TestUnzip_ZipNotExist
func TestUnzip_ZipNotExist(t *testing.T) {
	err := Unzip("/nonexistent/file.zip", testdataDir)
	assert.NotNil(t, err)
}

// go test -v -count 1 ./zip -run TestUnzip_RejectsZipSlip
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
