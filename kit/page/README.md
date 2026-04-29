# page

`kit/page/` 提供分页相关的 Proto 定义和解析辅助。

## 包含内容

- `page.kit.proto`：分页请求/响应相关 Proto 定义
- `page.kit.pb.go`：生成后的 Go 结构
- `page_parser.kit.go`：分页参数解析
- `page_helper.kit.go`：分页辅助函数

## 生成

如需重新生成分页 Proto：

```bash
kratos proto client kit/page/page.kit.proto
```
