# 测试规范

## 测试文件组织

```
internal/service/service/
├── ping.service.go
└── ping.service_test.go

internal/biz/biz/
├── ping.biz.go
└── ping.biz_test.go

internal/data/data/
├── ping.data.go
└── ping.data_test.go
```

## 测试命名约定

| 类型 | 命名格式 | 示例 |
|------|----------|------|
| 测试文件 | `{源文件名}_test.go` | `ping.service_test.go` |
| 测试函数 | `Test{函数名}` | `TestGetPingMessage` |
| 子测试 | `func(t *testing.T)` | `TestGetPingMessage/Success` |

## 单元测试规范

### Service Layer 测试

```go
func TestPingService_GetPingMessage(t *testing.T) {
    // 1. 准备 mock
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockBiz := mock.NewMockPingBizRepo(ctrl)

    // 2. 准备测试数据
    req := &pb.GetPingMessageReq{Message: "test"}
    mockBiz.EXPECT().GetPingMessage(gomock.Any(), gomock.Any()).Return(&bo.GetPingMessageResult{
        Message: "pong",
    }, nil)

    // 3. 执行测试
    service := NewPingService(logger, mockBiz)
    resp, err := service.GetPingMessage(context.Background(), req)

    // 4. 断言
    assert.Nil(t, err)
    assert.Equal(t, "pong", resp.Message)
}
```

### Business Layer 测试

```go
func TestPingBiz_GetPingMessage(t *testing.T) {
    // 使用 mock repository
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mock.NewMockPingBizRepo(ctrl)

    biz := NewPingBiz(logger, mockRepo)

    result, err := biz.GetPingMessage(context.Background(), &bo.GetPingMessageParam{
        Message: "test",
    })

    assert.Nil(t, err)
    assert.NotNil(t, result)
}
```

### Data Layer 测试

```go
func TestPingData_GetPingMessage(t *testing.T) {
    // 使用 testcontainers 或 SQLite 内存数据库
    db, err := testutil.SetupTestDB(t)
    assert.Nil(t, err)

    data := NewPingData(logger, db)

    result, err := data.GetPingMessage(context.Background(), &bo.GetPingMessageParam{})

    assert.Nil(t, err)
    assert.NotNil(t, result)
}
```

## 集成测试

### HTTP 测试

```go
func TestPingService_HTTP(t *testing.T) {
    // 1. 启动测试服务器
    hs := testutil.StartTestHTTPServer(t, service)

    // 2. 发送请求
    resp, err := http.Get(fmt.Sprintf("http://%s/api/v1/ping/say_hello", hs.Addr))

    // 3. 断言
    assert.Nil(t, err)
    assert.Equal(t, 200, resp.StatusCode)
}
```

### gRPC 测试

```go
func TestPingService_gRPC(t *testing.T) {
    // 1. 启动测试服务器
    gs := testutil.StartTestGRPCServer(t, service)

    // 2. 创建客户端
    conn, err := grpc.Dial(gs.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    client := v1.NewSrvPingClient(conn)

    // 3. 调用
    resp, err := client.GetPingMessage(context.Background(), &v1.GetPingMessageReq{})

    // 4. 断言
    assert.Nil(t, err)
    assert.NotNil(t, resp)
}
```

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./testdata/ping-service/internal/service/...

# 运行特定测试函数
go test -run TestGetPingMessage ./testdata/ping-service/internal/service/...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# 运行基准测试
go test -bench=. -benchmem
```

## Mock 使用

使用 `gomock` 生成 mock：

```bash
# 安装 mockgen
go install github.com/golang/mock/mockgen@latest

# 生成 mock
mockgen -source=internal/biz/repo/ping.repo.go -destination=internal/biz/repo/mock/ping.repo.go
```

## 测试数据管理

- 使用 `testdata/` 目录存放测试数据文件
- 测试配置使用独立的 `configs/test.yaml`
- 数据库使用 testcontainers 或内存数据库

## 测试覆盖率目标

- 核心业务逻辑：≥ 80%
- 工具函数：≥ 90%
- 整体目标：≥ 70%
