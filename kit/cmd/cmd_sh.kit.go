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

// ShellCommandArgs returns the platform shell executable and command flag.
func ShellCommandArgs() []string {
	shellBin := LinuxShellBin
	if runtime.GOOS == "windows" {
		shellBin = WindowsShellBin
	}
	return strings.Split(strings.TrimSpace(shellBin), " ")
}

// Deprecated: use ShellCommandArgs instead.
func ExecShell() []string { return ShellCommandArgs() }
