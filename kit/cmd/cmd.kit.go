package cmdpkg

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	bufferpkg "github.com/ikaiguang/go-kit/buffer"
)

// RunCommandContext 运行命令（支持 Context）
func RunCommandContext(ctx context.Context, command string, args []string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, command, args...)

	slog.DebugContext(ctx, "cmd", slog.String("command", command), slog.Any("args", args))

	return run(cmd)
}

// RunCommandWithWorkDirContext 运行命令（支持 Context 和工作目录）
func RunCommandWithWorkDirContext(ctx context.Context, workDir, command string, args []string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = workDir

	slog.DebugContext(ctx, "workdir", slog.String("dir", workDir), slog.String("command", command), slog.Any("args", args))

	return run(cmd)
}

// Deprecated: 使用 RunCommandContext 替代
func RunCommand(command string, args []string) (output []byte, err error) {
	return RunCommandContext(context.Background(), command, args)
}

// Deprecated: 使用 RunCommandWithWorkDirContext 替代
func RunCommandWithWorkDir(workDir, command string, args []string) (output []byte, err error) {
	return RunCommandWithWorkDirContext(context.Background(), workDir, command, args)
}

// run 运行命令
func run(cmdHandler *exec.Cmd) (output []byte, err error) {
	var (
		stdout = bufferpkg.GetBuffer()
		stderr = bufferpkg.GetBuffer()
	)
	defer bufferpkg.PutBuffer(stdout)
	defer bufferpkg.PutBuffer(stderr)

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
