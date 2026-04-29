# conf - 服务私有配置

存放服务私有的 Proto 配置定义。

`config.conf.proto` 定义了服务特有的配置结构（区别于 `api/config/` 中的全局配置）。

## 代码生成

```bash
make protoc-config-protobuf
```

> 不要手动修改 `*.pb.go`、`*.pb.validate.go` 等生成文件。
