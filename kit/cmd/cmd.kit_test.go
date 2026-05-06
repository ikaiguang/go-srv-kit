package cmdpkg

import (
	"context"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecShell(t *testing.T) {
	got := ExecShell()
	require.NotEmpty(t, got)
	if runtime.GOOS == "windows" {
		assert.Equal(t, []string{"cmd.exe", "/C"}, got)
		return
	}
	assert.Equal(t, []string{"/bin/sh", "-c"}, got)
}

func TestRunCommandContext(t *testing.T) {
	ctx := context.Background()
	var (
		command string
		args    []string
	)
	if runtime.GOOS == "windows" {
		command = "cmd.exe"
		args = []string{"/C", "echo hello"}
	} else {
		command = "sh"
		args = []string{"-c", "printf hello"}
	}

	got, err := RunCommandContext(ctx, command, args)
	require.NoError(t, err)
	assert.Equal(t, "hello", strings.TrimSpace(string(got)))
}

func TestRunCommandWithWorkDirContext(t *testing.T) {
	ctx := context.Background()
	workDir := t.TempDir()
	var (
		command string
		args    []string
	)
	if runtime.GOOS == "windows" {
		command = "cmd.exe"
		args = []string{"/C", "cd"}
	} else {
		command = "pwd"
	}

	got, err := RunCommandWithWorkDirContext(ctx, workDir, command, args)
	require.NoError(t, err)
	assert.Equal(t, workDir, strings.TrimSpace(string(got)))
}

func TestRunCommandContextErrorKeepsUnderlyingError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	defer cancel()

	_, err := RunCommandContext(ctx, "command-not-exist-go-kit-test", nil)
	require.Error(t, err)
	assert.NotEmpty(t, err.Error())
}
