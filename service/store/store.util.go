package storeutil

import (
	stdlog "log"
	"os"
	"path/filepath"

	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
	ospkg "github.com/ikaiguang/go-srv-kit/kit/os"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

type StoreManager interface {
	Save(sourceDir, storeDir string) error
}

// ReadStoreFiles 读取文件
func ReadStoreFiles(sourceDir, storeDir string) (map[string][]byte, error) {
	fs, err := filepathpkg.ReadDir(sourceDir)
	if err != nil {
		e := errorpkg.ErrorInternalError("failed to read source directory")
		return nil, errorpkg.Wrap(e, err)
	}
	if len(fs) == 0 {
		e := errorpkg.ErrorBadRequest("no config files found in source directory: sourceDir=%s", sourceDir)
		return nil, errorpkg.WithStack(e)
	}
	if storeDir == "" {
		e := errorpkg.ErrorBadRequest("store directory path is required: storeDir")
		return nil, errorpkg.WithStack(e)
	}
	configDataM := make(map[string][]byte)
	for i := range fs {
		if fs[i].IsDir() {
			continue
		}
		destPath := filepath.Join(storeDir, fs[i].Name())
		filePath := filepath.Join(sourceDir, fs[i].Name())
		if ospkg.IsWindows() {
			destPath = filepath.ToSlash(destPath)
		}
		stdlog.Println("|*** 读取文件：", filePath)
		content, err := os.ReadFile(filePath)
		if err != nil {
			e := errorpkg.ErrorInternalError("failed to read file")
			return nil, errorpkg.Wrap(e, err)
		}
		configDataM[destPath] = content
	}
	return configDataM, nil
}
