# download

通过 HTTP 流式下载文件，并自动检查或创建目标目录。

## 基础用法

```go
res, err := downloadpkg.StreamDownload(ctx, &downloadpkg.DownloadParam{
	URL:        "https://example.com/file.zip",
	OutputPath: "./runtime/file.zip",
	BufferSize: 32 * 1024,
})
```

## 注意事项

- 下载外部资源时使用带超时的 context。
- 输出路径由调用方控制，避免覆盖重要文件。
- 下载会先写入 `.tmp` 临时文件，完成后再替换目标文件。
- 测试或定制传输时可通过 `HTTPClient` 注入客户端。

## 验证

```bash
go test ./download
```
