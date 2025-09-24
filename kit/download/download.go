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

	resp, err := http.Get(param.URL)
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

	// 创建本地文件
	outFile, err := os.Create(param.OutputPath)
	if err != nil {
		return nil, fmt.Errorf("create file failed: %w", err)
	}
	defer func() {
		_ = outFile.Close()
	}()

	// 流式下载并写入文件
	buffer := make([]byte, 8*1024*1024) // 缓冲区(M)
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

	//fmt.Printf("下载完成! 文件保存至: %s, 大小: %.2f MB\n", param.OutputPath, float64(totalBytes)/1024/1024)
	return &DownloadReply{
		FilePath: param.OutputPath,
	}, nil
}
