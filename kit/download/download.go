package downloadpkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

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

func CheckOrCreateDir(outputPath string) error {
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

func StreamDownload(ctx context.Context, param *DownloadParam) (*DownloadReply, error) {
	if param.OutputPath == "" {
		return nil, errors.New("output path is empty")
	}
	// 确保保存目录存在
	if err := CheckOrCreateDir(param.OutputPath); err != nil {
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

	tmpPath := param.OutputPath + ".tmp"
	_ = os.Remove(tmpPath)
	// 创建本地临时文件，下载完成后再替换目标文件
	outFile, err := os.Create(tmpPath)
	if err != nil {
		return nil, fmt.Errorf("create file failed: %w", err)
	}
	defer func() {
		_ = outFile.Close()
		_ = os.Remove(tmpPath)
	}()

	// 流式下载并写入文件
	bufferSize := param.BufferSize
	if bufferSize <= 0 {
		bufferSize = 32 * 1024
	}
	buffer := make([]byte, bufferSize)
	var totalBytes int64
	for {
		// 从响应体读取数据
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			// 写入本地文件
			if _, writeErr := outFile.Write(buffer[:n]); writeErr != nil {
				return nil, fmt.Errorf("write file failed: %w", writeErr)
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
				err = nil
				break // 正常结束
			}
			return nil, fmt.Errorf("download failed: %w", err)
		}
	}

	if err := outFile.Close(); err != nil {
		return nil, fmt.Errorf("close file failed: %w", err)
	}
	if err := os.Rename(tmpPath, param.OutputPath); err != nil {
		return nil, fmt.Errorf("rename file failed: %w", err)
	}

	//fmt.Printf("下载完成! 文件保存至: %s, 大小: %.2f MB\n", param.OutputPath, float64(totalBytes)/1024/1024)
	return &DownloadReply{
		FilePath: param.OutputPath,
	}, nil
}
