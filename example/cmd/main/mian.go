package main

import (
	stdlog "log"
	"os"

	setup "github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			stdlog.Printf("%+v\n", err)
		}
	}()

	// 初始化
	err = setup.Setup()
	if err != nil {
		stdlog.Println()
		os.Exit(1)
	}
}
