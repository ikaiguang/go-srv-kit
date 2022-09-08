package setup

import (
	"sync"

	setuppkg "github.com/ikaiguang/go-srv-kit/example/pkg/setup"
)

// Engine ...
type Engine interface {
	setuppkg.Engine
}

// engines ...
type engines struct {
	setuppkg.Engine
}

var (
	// initEngineMutex 初始化
	initEngineMutex sync.Once
	engineInstance  *engines
)

// Init 启动与配置与设置存储Packages
func Init(opts ...setuppkg.Option) (err error) {
	initEngineMutex.Do(func() {
		var e setuppkg.Engine
		e, err = setuppkg.New(opts...)
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
