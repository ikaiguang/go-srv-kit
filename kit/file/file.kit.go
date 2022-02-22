package fileutil

import (
	"os"
	"path/filepath"
)

const (
	// DefaultFileMode 文件权限
	DefaultFileMode os.FileMode = 0744

	// RuntimeDir 临时目录
	DefaultRuntimeDir = "runtime"
)

// MoveFileToDir 移动文件到目录
func MoveFileToDir(filePath, fileDir string) (targetPath string, err error) {
	targetPath = filepath.Join(fileDir, filepath.Base(filePath))
	if err = os.Rename(filePath, targetPath); err != nil {
		return targetPath, err
	}
	return targetPath, err
}
