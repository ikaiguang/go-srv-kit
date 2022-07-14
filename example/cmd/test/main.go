package main

import (
	"flag"
	stdlog "log"

	websockettest "github.com/ikaiguang/go-srv-kit/example/cmd/test/websocket"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

/*
# 运行测试
go run example/cmd/test/main.go -conf=./example/configs -run-test=websocket

*/
const (
	RunFlagUsage = `设置运行测试指令;
等于 空值 不运行；
等于 websocket 测试websocket;
`
	RunFlagWebsocket = "websocket"
)

var (
	runTestFlag = flag.String("run-test", "", RunFlagUsage)
)

// runTest 运行测试
func runTest(engineHandler setup.Engine) {
	if !flag.Parsed() {
		flag.Parse()
	}

	switch *runTestFlag {
	case RunFlagWebsocket:
		// 测试ws
		websockettest.RunTestWebsocket()
	default:
		stdlog.Printf("| *** 未匹配到运行程序：%s\n", *runTestFlag)
	}
	stdlog.Println("| *** Done!")
}

func main() {
	var err error

	// 初始化
	err = setup.Init()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 包
	packages, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = setup.Close() }()

	stdlog.Println("| *** IsDebugMode = ", packages.IsDebugMode())

	// 启动程序
	stdlog.Println()
	stdlog.Println("|==================== 启动程序 开始 ====================|")
	defer stdlog.Println("|==================== 启动程序 结束 ====================|")

	// 运行测试
	runTest(packages)
}
