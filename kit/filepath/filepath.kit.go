package filepathpkg

import (
	"os"
	"path/filepath"

	filepkg "github.com/ikaiguang/go-srv-kit/kit/file"
)

// WaldDir 遍历所有的目录与文件
func WaldDir(rootPath string) (filePaths []string, fileInfos []os.FileInfo, err error) {
	//filepath.WalkDir()
	//dirFn := func(filePath string, d fs.DirEntry, err error) (fnErr error) {
	//}
	dirFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		filePaths = append(filePaths, path)
		fileInfos = append(fileInfos, info)
		return err
	}

	err = filepath.Walk(rootPath, dirFn)
	if err != nil {
		return filePaths, fileInfos, err
	}
	return filePaths, fileInfos, err
}

// ReadDir 当前目录与文件
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
