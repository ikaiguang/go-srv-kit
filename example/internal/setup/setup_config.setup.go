package setuphandler

import (
	"flag"

	pkgerrors "github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/config"
)

const (
	_defaultConfigFilepath = "./configs"
)

var (
	// 配置文件 所在的目录
	_configFilepath string
)

func init() {
	flag.StringVar(&_configFilepath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// getConfigHandler 获取配置手柄
func (s *setup) getConfigHandler() (handler config.Config, err error) {
	handler = config.New()

	if err = handler.Load(); err != nil {
		err = pkgerrors.WithStack(err)
		return
	}
	return
}
