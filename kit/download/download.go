package downloadpkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	defaultDownloadBufferSize = 32 << 10
	maxDownloadBufferSize     = 16 << 20
)

var replaceFileMu sync.Mutex

type DownloadParam struct {
	URL             string
	OutputPath      string
	FileSizeChannel chan<- int64
	HTTPClient      *http.Client
	BufferSize      int
}

type DownloadReply struct {
	FilePath string
}

// EnsureOutputDirectory creates the parent directory for outputPath.
func EnsureOutputDirectory(outputPath string) error {
	dir := filepath.Dir(outputPath)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// Deprecated: use EnsureOutputDirectory instead.
func CheckOrCreateDir(outputPath string) error { return EnsureOutputDirectory(outputPath) }

func StreamDownload(ctx context.Context, param *DownloadParam) (*DownloadReply, error) {
	if param == nil {
		return nil, errors.New("download param is nil")
	}
	if param.OutputPath == "" {
		return nil, errors.New("output path is empty")
	}
	if param.BufferSize > maxDownloadBufferSize {
		return nil, fmt.Errorf("buffer size %d exceeds maximum %d", param.BufferSize, maxDownloadBufferSize)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	// 确保保存目录存在
	if err := EnsureOutputDirectory(param.OutputPath); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, param.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	httpClient := param.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status code: %d", resp.StatusCode)
	}

	// Use a unique temporary file in the destination directory. Keeping it on
	// the same filesystem preserves atomic rename semantics and avoids clashes
	// between concurrent downloads of the same target.
	tmpPattern := "." + filepath.Base(param.OutputPath) + ".*.tmp"
	outFile, err := os.CreateTemp(filepath.Dir(param.OutputPath), tmpPattern)
	if err != nil {
		return nil, fmt.Errorf("create file failed: %w", err)
	}
	tmpPath := outFile.Name()
	if err := outFile.Chmod(0o644); err != nil {
		_ = outFile.Close()
		_ = os.Remove(tmpPath)
		return nil, fmt.Errorf("set temporary file mode failed: %w", err)
	}
	defer func() {
		_ = outFile.Close()
		_ = os.Remove(tmpPath)
	}()

	// 流式下载并写入文件
	bufferSize := param.BufferSize
	if bufferSize <= 0 {
		bufferSize = defaultDownloadBufferSize
	}
	buffer := make([]byte, bufferSize)
	var totalBytes int64
	for {
		// 从响应体读取数据
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// 写入本地文件
			written, writeErr := outFile.Write(buffer[:n])
			if writeErr != nil {
				return nil, fmt.Errorf("write file failed: %w", writeErr)
			}
			if written != n {
				return nil, fmt.Errorf("write file failed: %w", io.ErrShortWrite)
			}
			totalBytes += int64(n)

			// 打印下载进度（每下载1MB打印一次）
			//if totalBytes%(1024*1024) < int64(n) {
			//	fmt.Printf("已下载: %.2f MB\n", float64(totalBytes)/1024/1024)
			//}
			// 发送文件大小到通道
			if param.FileSizeChannel != nil {
				select {
				case param.FileSizeChannel <- totalBytes:
				default:
				}
			}
		}

		// 检查是否读取完毕
		if err != nil {
			if err == io.EOF {
				break // 正常结束
			}
			return nil, fmt.Errorf("download failed: %w", err)
		}
	}

	if err := outFile.Close(); err != nil {
		return nil, fmt.Errorf("close file failed: %w", err)
	}
	if err := replaceFile(tmpPath, param.OutputPath); err != nil {
		return nil, fmt.Errorf("rename file failed: %w", err)
	}

	//fmt.Printf("下载完成! 文件保存至: %s, 大小: %.2f MB\n", param.OutputPath, float64(totalBytes)/1024/1024)
	return &DownloadReply{
		FilePath: param.OutputPath,
	}, nil
}

func replaceFile(sourcePath, targetPath string) error {
	replaceFileMu.Lock()
	defer replaceFileMu.Unlock()

	if runtime.GOOS == "windows" {
		if err := os.Remove(targetPath); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return os.Rename(sourcePath, targetPath)
}
