package setup

import (
	"sync"
)

var (
	// 初始化
	initMutex sync.Once
	packages  Packages
)

// GetPackages 获取初始化后的Packages
func GetPackages() (Packages, error) {
	if err := Init(); err != nil {
		return nil, err
	}
	return packages, nil
}

// Init 启动与配置与设置存储Packages
func Init() (err error) {
	initMutex.Do(func() {
		packages, err = Setup()
		if err != nil {
			initMutex = sync.Once{}
		}
	})
	if err != nil {
		return err
	}
	if packages != nil {
		return err
	}

	packages, err = Setup()
	if err != nil {
		return err
	}
	return err
}

// Close .
func Close() error {
	if packages == nil {
		return nil
	}
	return packages.Close()
}
