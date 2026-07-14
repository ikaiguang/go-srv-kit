# cmd

执行外部命令，支持 `context.Context` 和工作目录。

## 基础用法

```go
out, err := cmdpkg.RunCommandContext(ctx, "go", []string{"version"})
out, err = cmdpkg.RunCommandWithWorkDirContext(ctx, ".", "go", []string{"test", "./cmd"})
```

## 注意事项

- 命令和参数要分开传递，不要拼接未校验的用户输入。
- 长耗时命令必须使用带超时的 context。

## 验证

```bash
go test ./cmd
```
