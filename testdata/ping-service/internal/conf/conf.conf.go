package conf

import configutil "github.com/ikaiguang/go-srv-kit/service/config"

var (
	serviceConfig = &ServiceConfig{}
)

// LoadServiceConfig 加载服务配置
// 由 setuputil.NewLauncherManager 进行加载赋值
func LoadServiceConfig() []configutil.Option {
	return []configutil.Option{
		configutil.WithOtherConfig(serviceConfig),
	}
}
