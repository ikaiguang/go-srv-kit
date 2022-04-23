package fileutil

import (
	"io"
	"os"
	"path/filepath"
)

const (
	// DefaultFileMode 文件权限
	DefaultFileMode os.FileMode = 0744

	// DefaultRuntimeDir 临时目录
	DefaultRuntimeDir = "runtime"
)

// CopyFile 复制文件
func CopyFile(from, to string) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()

	dest, err := os.Create(to)
	if err != nil {
		return err
	}
	defer func() { _ = dest.Close() }()

	if _, err := io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}

// MoveFileToDir 移动文件到目录
func MoveFileToDir(filePath, fileDir string) (targetPath string, err error) {
	targetPath = filepath.Join(fileDir, filepath.Base(filePath))
	if err = os.Rename(filePath, targetPath); err != nil {
		return targetPath, err
	}
	return targetPath, err
}
