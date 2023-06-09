package cmdpkg

import (
	"fmt"
	"os/exec"

	debugpkg "github.com/ikaiguang/go-srv-kit/debug"
	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
)

// RunCommand 运行命令
func RunCommand(command string, args []string) (output []byte, err error) {
	cmd := exec.Command(command, args...)

	debugpkg.Debugw("cmd", command, "args", args)

	return run(cmd)
}

// RunCommandWithWorkDir 运行命令
// @param workDir specifies the working directory of the command.
func RunCommandWithWorkDir(workDir, command string, args []string) (output []byte, err error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workDir

	debugpkg.Debugw("workdir", workDir, "cmd", command, "args", args)

	return run(cmd)
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
		err = fmt.Errorf("%s", stderr.Bytes())
		return output, err
	}
	return stdout.Bytes(), err
}
