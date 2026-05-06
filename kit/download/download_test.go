package downloadpkg

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestdataPath = "./_output_testdata"
)

// go test -v -count 1 ./download -run TestStreamDownload
func TestStreamDownload(t *testing.T) {
	// 创建本地测试 HTTP 服务器
	testContent := []byte("hello world test content for download")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(testContent)
	}))
	defer server.Close()

	outputPath := filepath.Join(TestdataPath, "test_download.bin")
	// 清理
	defer func() { _ = os.RemoveAll(TestdataPath) }()

	t.Run("正常下载", func(t *testing.T) {
		got, err := StreamDownload(context.Background(), &DownloadParam{
			URL:        server.URL + "/test.bin",
			OutputPath: outputPath,
			HTTPClient: server.Client(),
			BufferSize: 4,
		})
		require.Nil(t, err)
		assert.Equal(t, outputPath, got.FilePath)

		// 验证文件内容
		data, err := os.ReadFile(outputPath)
		require.Nil(t, err)
		assert.Equal(t, testContent, data)
	})

	t.Run("空输出路径", func(t *testing.T) {
		_, err := StreamDownload(context.Background(), &DownloadParam{
			URL:        server.URL + "/test.bin",
			OutputPath: "",
		})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "output path is empty")
	})

	t.Run("服务器返回404", func(t *testing.T) {
		server404 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server404.Close()

		target := filepath.Join(TestdataPath, "notfound.bin")
		_, err := StreamDownload(context.Background(), &DownloadParam{
			URL:        server404.URL + "/notfound",
			OutputPath: target,
		})
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "status code: 404")
		_, statErr := os.Stat(target)
		assert.True(t, os.IsNotExist(statErr))
	})
}

// go test -v -count 1 ./download -run TestCheckOrCreateDir
func TestCheckOrCreateDir(t *testing.T) {
	testDir := filepath.Join(TestdataPath, "subdir", "nested")
	defer func() { _ = os.RemoveAll(TestdataPath) }()

	err := CheckOrCreateDir(filepath.Join(testDir, "file.txt"))
	require.Nil(t, err)

	info, err := os.Stat(testDir)
	require.Nil(t, err)
	assert.True(t, info.IsDir())
}
