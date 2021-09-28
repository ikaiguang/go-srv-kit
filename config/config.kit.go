package configutil

import (
	"flag"

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

// Setup .
func Setup(opts ...config.Option) {
	if _configFilepath == "" {
		_configFilepath = _defaultConfigFilepath
	}

	//handler := config.New(opts...)
}
