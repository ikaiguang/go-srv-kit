# config - 配置管理

`config/` 负责加载和解析服务配置，支持从本地文件和 Consul 加载。

## 包名

```go
import configutil "github.com/ikaiguang/go-srv-kit/service/config"
```

## 配置加载

```go
// 从文件加载配置
conf, err := configutil.Loading(configFilePath)

// 从文件加载，支持 Consul 远程配置
conf, err := configutil.Loading(configFilePath,
    configutil.WithConsulConfigLoader(configutil.NewConsulConfigLoader()),
)
```

## 配置访问

提供便捷的配置项访问函数：

```go
appConfig := configutil.AppConfig(bootstrap)
mysqlConfig := configutil.MysqlConfig(bootstrap)
redisConfig := configutil.RedisConfig(bootstrap)
logConfig := configutil.LogConfig(bootstrap)
// ...
```

## 配置文件格式

配置文件使用 YAML 格式，结构由 `api/config/config.proto` 定义。

参考示例：`testdata/ping-service/configs/config_all.yaml`

## 文件说明

| 文件 | 说明 |
|------|------|
| `config.util.go` | 配置项便捷访问函数 |
| `config_loading.util.go` | 配置加载核心逻辑 |
| `config_from_file.util.go` | 从本地文件加载 |
| `config_from_consul.util.go` | 从 Consul 加载 |
| `config_watch.util.go` | 配置热更新监听 |
| `config_option.util.go` | 加载选项 |
| `config_example.yaml` | 配置示例 |
