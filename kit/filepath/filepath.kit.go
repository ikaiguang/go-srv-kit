package filepathpkg

import (
	"io/fs"
	"os"
	"path/filepath"

	filepkg "github.com/ikaiguang/go-srv-kit/kit/file"
)

// WalkDir 遍历所有的目录与文件（使用 filepath.WalkDir + fs.DirEntry）
func WalkDir(rootPath string) (filePaths []string, dirEntries []fs.DirEntry, err error) {
	dirFn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		filePaths = append(filePaths, path)
		dirEntries = append(dirEntries, d)
		return nil
	}
	err = filepath.WalkDir(rootPath, dirFn)
	return filePaths, dirEntries, err
}

// Deprecated: 拼写错误，使用 WalkDir 替代。
// 注意：返回类型从 []os.FileInfo 变为 []fs.DirEntry，此兼容函数保留原签名。
func WaldDir(rootPath string) (filePaths []string, fileInfos []os.FileInfo, err error) {
	dirFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		filePaths = append(filePaths, path)
		fileInfos = append(fileInfos, info)
		return nil
	}
	err = filepath.Walk(rootPath, dirFn)
	return filePaths, fileInfos, err
}

// ReadDirEntries 读取目录内容（使用 os.ReadDir）
func ReadDirEntries(rootPath string) ([]fs.DirEntry, error) {
	return os.ReadDir(rootPath)
}

// Deprecated: 使用 ReadDirEntries 替代
func ReadDir(rootPath string) (fileInfos []os.FileInfo, err error) {
	d, err := os.Open(rootPath)
	if err != nil {
		return fileInfos, err
	}
	defer func() { _ = d.Close() }()

	fileInfos, err = d.Readdir(-1)
	if err != nil {
		return fileInfos, err
	}
	return fileInfos, err
}

// CreateDir 创建目录
func CreateDir(destDir string) (err error) {
	// 创建目录
	err = os.MkdirAll(destDir, filepkg.DefaultFileMode)
	if err != nil {
		return err
	}
	return err
}

// RenewDir 重建目录
func RenewDir(destDir string) (err error) {
	_, err = os.Stat(destDir)
	if err == nil {
		// 删除已存在
		err = os.RemoveAll(destDir)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		err = nil // 文件不存在
	} else {
		return err
	}

	// 创建目录
	err = os.MkdirAll(destDir, filepkg.DefaultFileMode)
	if err != nil {
		return err
	}
	return err
}
