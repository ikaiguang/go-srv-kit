package setup

import (
	"sync"
)

var (
	// initModulesMutex 初始化
	initModulesMutex sync.Once
	modulesInstance  Modules
)

// Init 启动与配置与设置存储Packages
func Init() (err error) {
	initModulesMutex.Do(func() {
		modulesInstance, err = Setup()
	})
	if err != nil {
		initModulesMutex = sync.Once{}
		return err
	}
	if modulesInstance != nil {
		return err
	}

	modulesInstance, err = Setup()
	if err != nil {
		return err
	}
	return err
}

// GetModules 获取初始化后的模块
func GetModules() (Modules, error) {
	if err := Init(); err != nil {
		return nil, err
	}
	return modulesInstance, nil
}

// Close .
func Close() error {
	if modulesInstance == nil {
		return nil
	}
	return modulesInstance.Close()
}
