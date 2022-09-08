package setup

import (
	"sync"

	setuputil "github.com/ikaiguang/go-srv-kit/setup"
)

// Engine ...
type Engine interface {
	setuputil.Engine
}

// engines ...
type engines struct {
	setuputil.Engine
}

var (
	// initEngineMutex 初始化
	initEngineMutex sync.Once
	engineInstance  *engines
)

// Init 启动与配置与设置存储Packages
func Init(opts ...setuputil.Option) (err error) {
	initEngineMutex.Do(func() {
		var e setuputil.Engine
		e, err = setuputil.New(opts...)
		engineInstance = &engines{
			Engine: e,
		}
	})
	if err != nil {
		engineInstance = nil
		initEngineMutex = sync.Once{}
		return err
	}
	return err
}

// GetEngine 获取初始化后的引擎模块
func GetEngine() (Engine, error) {
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
	if err := engineInstance.Engine.Close(); err != nil {
		return err
	}
	return nil
}
