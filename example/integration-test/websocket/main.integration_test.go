package testwebsocket

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/ikaiguang/go-srv-kit/example/pkg/setup"
)

func TestMain(m *testing.M) {
	var err error

	configPath := "../../configs"
	engineHandler, err := setuppkg.New(setuppkg.WithConfigPath(configPath))
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = engineHandler.Close() }()

	code := m.Run()

	// 初始化必要逻辑
	os.Exit(code)
}
