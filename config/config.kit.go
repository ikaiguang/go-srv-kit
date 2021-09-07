package configutil

import "flag"

const (
	_defaultConfigFilepath = "./configs"
)

var (
	// 配置文件 所在的目录
	configFilepath string
)

func init() {
	flag.StringVar(&configFilepath, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// Setup .
func Setup() {
	if configFilepath == "" {
		configFilepath = _defaultConfigFilepath
	}
}
