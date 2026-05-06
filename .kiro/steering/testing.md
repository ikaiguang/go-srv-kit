---
inclusion: fileMatch
fileMatchPattern: "**/*_test.go"
---

# 测试规范

## 文件组织

测试文件与源文件同目录：`{name}_test.go`

## 测试命名

| 类型 | 格式 | 示例 |
|------|------|------|
| 测试函数 | `Test{函数名}` | `TestGetPingMessage` |
| 子测试 | `t.Run("场景", ...)` | `TestGetPingMessage/Success` |

## 分层测试

Service 层：mock Business 层接口（gomock）
Business 层：mock Repository 接口
Data 层：testcontainers 或 SQLite 内存数据库

```go
// Service 层测试示例
ctrl := gomock.NewController(t)
defer ctrl.Finish()
mockBiz := mock.NewMockPingBizRepo(ctrl)
mockBiz.EXPECT().GetPingMessage(gomock.Any(), gomock.Any()).Return(&bo.Result{}, nil)
service := NewPingService(logger, mockBiz)
resp, err := service.GetPingMessage(context.Background(), req)
assert.Nil(t, err)
```

## 运行测试

```bash
go test ./...                          # 全部测试
go test -run TestGetPingMessage ./...  # 指定测试
go test -coverprofile=coverage.out ./... # 覆盖率
go test -bench=. -benchmem            # 基准测试
```

## Mock 生成

```bash
mockgen -source=internal/biz/repo/ping.repo.go -destination=internal/biz/repo/mock/ping.repo.go
```

## 覆盖率目标

- 核心业务逻辑：≥ 80%
- 工具函数：≥ 90%
- 整体：≥ 70%
