# store - 配置存储

`store/` 提供将本地配置文件批量存储到外部存储（如 Consul KV）的工具。

## 包名

```go
import storeutil "github.com/ikaiguang/go-srv-kit/service/store"
```

## 使用

### 存储到 Consul

```go
err := storeutil.StoreInConsul(consulClient, sourceDir, storeDir)
```

- `sourceDir`：本地配置文件目录
- `storeDir`：Consul KV 存储路径前缀

### 命令行工具

参考 `testdata/ping-service/cmd/store-configuration/` 和 `testdata/configuration/`。
