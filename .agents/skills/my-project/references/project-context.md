# 项目上下文

## 仓库标识

- Go：`1.25.9`

## 分层模型

基于该工具包实现业务服务时，遵循：

1. `Service` 层负责 HTTP/gRPC handler 和 DTO 转换
2. `Biz` 层负责业务逻辑和仓储接口
3. `Data` 层负责仓储实现和外部访问

常见数据流：

`Proto -> DTO -> BO -> PO`

不要让 `Service` 层直接调用 `Data` 层。

## 示例入口

需要具体例子时，优先读这些文件：

- `testdata/ping-service/cmd/ping-service/main.go`
- `testdata/ping-service/cmd/ping-service/export/wire.go`
- `testdata/ping-service/internal/service/service/ping.service.go`
- `testdata/ping-service/internal/biz/biz/ping.biz.go`
- `testdata/ping-service/internal/data/data/ping.data.go`

## 工具包关键文件

- `server_all_in_one.util.go`：服务端构建入口

## 启动形态

典型启动流程：

1. 导出服务专属配置和认证白名单
2. 通过 Wire 装配业务服务
3. 使用工具包基础设施创建 all-in-one server
