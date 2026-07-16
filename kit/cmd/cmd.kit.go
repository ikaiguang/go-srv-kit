package cmdpkg

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/valyala/bytebufferpool"
)

// RunCommandContext 运行命令（支持 Context）
func RunCommandContext(ctx context.Context, command string, args []string) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	cmd := exec.CommandContext(ctx, command, args...)

	slog.DebugContext(ctx, "cmd", slog.String("command", command), slog.Any("args", args))

	return run(cmd)
}

// RunCommandInDirContext runs a command in workDir with context cancellation.
func RunCommandInDirContext(ctx context.Context, workDir, command string, args []string) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = workDir

	slog.DebugContext(ctx, "workdir", slog.String("dir", workDir), slog.String("command", command), slog.Any("args", args))

	return run(cmd)
}

// Deprecated: use RunCommandInDirContext instead.
func RunCommandWithWorkDirContext(ctx context.Context, workDir, command string, args []string) ([]byte, error) {
	return RunCommandInDirContext(ctx, workDir, command, args)
}

// Deprecated: 使用 RunCommandContext 替代
func RunCommand(command string, args []string) (output []byte, err error) {
	return RunCommandContext(context.Background(), command, args)
}

// RunCommandInDir runs a command in workDir.
func RunCommandInDir(workDir, command string, args []string) ([]byte, error) {
	return RunCommandInDirContext(context.Background(), workDir, command, args)
}

// Deprecated: use RunCommandInDir instead.
func RunCommandWithWorkDir(workDir, command string, args []string) (output []byte, err error) {
	return RunCommandInDir(workDir, command, args)
}

// run 运行命令
func run(cmdHandler *exec.Cmd) (output []byte, err error) {
	var (
		stdout = bytebufferpool.Get()
		stderr = bytebufferpool.Get()
	)
	defer bytebufferpool.Put(stdout)
	defer bytebufferpool.Put(stderr)

	cmdHandler.Stdout = stdout
	cmdHandler.Stderr = stderr

	// run
	if err = cmdHandler.Run(); err != nil {
		errText := strings.TrimSpace(stderr.String())
		if errText == "" {
			return output, err
		}
		return output, fmt.Errorf("%w: %s", err, errText)
	}

	// 在归还 buffer 前复制数据，避免数据竞争
	result := make([]byte, stdout.Len())
	copy(result, stdout.Bytes())
	return result, err
}
