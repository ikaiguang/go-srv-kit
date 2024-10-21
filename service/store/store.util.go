package storeutil

import (
	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"os"
	"path/filepath"
)

type StoreManager interface {
	Save(sourceDir, storeDir string) error
}

// ReadStoreFiles 读取文件
func ReadStoreFiles(sourceDir, storeDir string) (map[string][]byte, error) {
	fs, err := filepathpkg.ReadDir(sourceDir)
	if err != nil {
		e := errorpkg.ErrorInternalError(err.Error())
		return nil, errorpkg.WithStack(e)
	}
	if len(fs) == 0 {
		e := errorpkg.ErrorBadRequest("资源目录查无配置文件：sourceDir=%s", sourceDir)
		return nil, errorpkg.WithStack(e)
	}
	if storeDir == "" {
		e := errorpkg.ErrorBadRequest("请配置存储路径：storeDir")
		return nil, errorpkg.WithStack(e)
	}
	configDataM := make(map[string][]byte)
	for i := range fs {
		if fs[i].IsDir() {
			continue
		}
		destPath := filepath.Join(storeDir, fs[i].Name())
		filePath := filepath.Join(sourceDir, fs[i].Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			e := errorpkg.ErrorInternalError(err.Error())
			return nil, errorpkg.WithStack(e)
		}
		configDataM[destPath] = content
	}
	return configDataM, nil
}
