package setup

import (
	"sync"

	setuppkg "github.com/ikaiguang/go-srv-kit/example/pkg/setup"
)

var (
	// initEngineMutex 初始化
	initEngineMutex sync.Once
	engineInstance  setuppkg.Engine
)

// Init 启动与配置与设置存储Packages
func Init(opts ...setuppkg.Option) (err error) {
	initEngineMutex.Do(func() {
		engineInstance, err = setuppkg.New(opts...)
	})
	if err != nil {
		initEngineMutex = sync.Once{}
		return err
	}
	return err
}

// GetEngine 获取初始化后的引擎模块
func GetEngine() (setuppkg.Engine, error) {
	if err := Init(); err != nil {
		return nil, err
	}
	return engineInstance, nil
}

// Close .
func Close() error {
	if engineInstance == nil {
		return nil
	}
	return engineInstance.Close()
}
