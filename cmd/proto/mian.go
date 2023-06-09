package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	debugpkg "github.com/ikaiguang/go-srv-kit/debug"
	bufferpkg "github.com/ikaiguang/go-srv-kit/kit/buffer"
	cmdpkg "github.com/ikaiguang/go-srv-kit/kit/cmd"
	filepkg "github.com/ikaiguang/go-srv-kit/kit/file"
	filepathpkg "github.com/ikaiguang/go-srv-kit/kit/filepath"
	pkgerrors "github.com/pkg/errors"
)

const (
	// LinuxShellBin 执行脚本
	LinuxShellBin   string = "/bin/sh -c" // mac & linux
	WindowsShellBin string = "cmd.exe /C" // windows

	ExecScriptFilename = "proto_script.sh"

	_uninitialized = "uninitialized"
)

var protoPath = flag.String("path", _uninitialized, "proto 路径")

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *protoPath == _uninitialized {
		fmt.Println("==> 请添加启动参数：-path")
		return
	}
	_, _ = debugpkg.Setup()

	var dirs = []string{
		*protoPath,
	}
	err := genManyProtoScriptFileAndExecScript(dirs)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
}

var (
	_defaultHandler = &Proto{}
)

// genManyProtoScriptFileAndExecScript 生成协议执行脚本文件 与 执行脚本
func genManyProtoScriptFileAndExecScript(dirs []string) (err error) {
	for i := range dirs {
		if _, err = genProtoScriptFileAndExecScript(dirs[i]); err != nil {
			return err
		}
	}
	return err
}

// genProtoScriptFileAndExecScript 生成协议执行脚本文件 与 执行脚本
func genProtoScriptFileAndExecScript(dir string) (scriptFilePath string, err error) {
	// 协议文件
	protoFiles, err := _defaultHandler.FindProtoFiles(dir)
	if err != nil {
		return scriptFilePath, err
	}
	if len(protoFiles) == 0 {
		err = pkgerrors.Errorf("==> 未找到proto文件；Path=%s", dir)
		return scriptFilePath, err
	}

	// 执行脚本
	scripts := _defaultHandler.GenProtoExecScripts(protoFiles)

	// 生成协议执行脚本文件路径 & 生成协议执行脚本文件
	scriptFilePath = _defaultHandler.GenProtoScriptFilePath(dir, ExecScriptFilename)
	err = _defaultHandler.GenProtoScriptFile(scriptFilePath, scripts)
	if err != nil {
		return scriptFilePath, err
	}

	// 执行二进制
	execBinSlice := _defaultHandler.CmdExecBin()
	command := execBinSlice[0]
	args := execBinSlice[1:]
	for i := range scripts {
		newArgs := append(args, scripts[i])
		out, err := cmdpkg.RunCommand(command, newArgs)
		if err != nil {
			err = pkgerrors.WithStack(err)
			return scriptFilePath, err
		}
		if strings.Contains(string(out), "exit status 1") {
			err = fmt.Errorf("\n\tscript : %s \n\terror : %s", scripts[i], out)
			err = pkgerrors.WithStack(err)
			return scriptFilePath, err
		}
		fmt.Println("==> Exec : ", scripts[i])
		fmt.Println("==> Output : ", string(out))
	}
	return scriptFilePath, err
}

// Proto 协议
type Proto struct{}

// CmdExecBin 执行二进制
func (s *Proto) CmdExecBin() []string {
	shellBin := LinuxShellBin
	if runtime.GOOS == "windows" {
		shellBin = WindowsShellBin
	}
	return strings.Split(strings.TrimSpace(shellBin), " ")
}

// CmdKratosClient kratos proto client xxx.proto
func (s *Proto) CmdKratosClient() []string {
	return []string{
		"kratos", "proto", "client",
		"--proto_path=.", "--proto_path=$GOPATH/src",
	}
}

// GenProtoScriptFile 生成协议执行脚本文件
func (s *Proto) GenProtoScriptFile(filename string, scripts []string) (err error) {
	buf := bufferpkg.GetBuffer()
	defer bufferpkg.PutBuffer(buf)

	buf.WriteString("#!/bin/bash\n\n")

	for i := range scripts {
		buf.WriteString(scripts[i])
		buf.WriteString("\n")
	}

	err = os.WriteFile(filename, buf.Bytes(), filepkg.DefaultFileMode)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return err
	}
	return err
}

// GenProtoScriptFilePath 生成协议执行脚本文件路径
func (s *Proto) GenProtoScriptFilePath(dir string, filename string) string {
	return filepath.Join(dir, filename)
}

// GenProtoExecScripts 生成协议执行脚本
func (s *Proto) GenProtoExecScripts(protoFiles []string) (scripts []string) {
	scripts = make([]string, len(protoFiles))
	for i := range protoFiles {
		scripts[i] = strings.Join(s.CmdKratosClient(), " ") + " " + protoFiles[i]
	}
	return
}

// FindProtoFiles 查找协议文件
func (s *Proto) FindProtoFiles(dir string) (protoFiles []string, err error) {
	filePaths, _, err := filepathpkg.WaldDir(dir)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return protoFiles, err
	}

	for i := range filePaths {
		if !strings.HasSuffix(filePaths[i], ".proto") {
			continue
		}
		protoFiles = append(protoFiles, filePaths[i])
	}
	sort.Strings(filePaths)
	return protoFiles, err
}
