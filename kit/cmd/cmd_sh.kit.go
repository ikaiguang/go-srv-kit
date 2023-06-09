package cmdpkg

import (
	"runtime"
	"strings"
)

const (
	// LinuxShellBin 执行脚本
	LinuxShellBin   string = "/bin/sh -c" // mac & linux
	WindowsShellBin string = "cmd.exe /C" // windows
)

// ExecShell 执行二进制
func ExecShell() []string {
	shellBin := LinuxShellBin
	if runtime.GOOS == "windows" {
		shellBin = WindowsShellBin
	}
	return strings.Split(strings.TrimSpace(shellBin), " ")
}
