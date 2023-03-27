package main

import (
	stdlog "log"

	dbmigrate "github.com/ikaiguang/go-srv-kit/example/cmd/migration/migrate"
	servers "github.com/ikaiguang/go-srv-kit/example/internal/server"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

// go run ./example/cmd/main/... -conf=./example/configs
// go run ./cmd/main/... -conf=./configs
func main() {
	var err error

	// 初始化
	err = setup.Init()
	if err != nil {
		stdlog.Fatalf("setup.Init : %+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 引擎模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("setup.GetEngine : %+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = setup.Close() }()

	// loadingBaseSettingEnv 加载程序配置
	if err = loadingBaseSettingEnv(engineHandler); err != nil {
		stdlog.Fatalf("loadingBaseSettingEnv : %+v\n", err)
		return
	}

	// 加载程序运行环境
	err = loadingAppDependentEnv(engineHandler)
	if err != nil {
		stdlog.Fatalf("loadingAppDependentEnv : %+v\n", err)
		return
	}

	// 启动程序
	stdlog.Println()
	stdlog.Println("|==================== 加载程序 开始 ====================|")
	app, err := servers.NewApp(engineHandler)
	if err != nil {
		stdlog.Fatalf("loadingAppDependentEnv : %+v\n", err)
		return
	}
	stdlog.Println("|==================== 加载程序 结束 ====================|")
	stdlog.Println("|==================== 程序运行中... ====================|")
	if err = app.Run(); err != nil {
		stdlog.Fatalf("app.Run %+v\n", err)
		return
	}
}

// loadingBaseSettingEnv 加载计划任务
func loadingBaseSettingEnv(engineHandler setup.Engine) error {
	// 计划任务
	settingConfig := engineHandler.BaseSettingConfig()
	stdlog.Println()
	stdlog.Println("|==================== 加载计划任务 开始 ====================|")
	defer stdlog.Println("|==================== 加载计划任务 结束 ====================|")

	// 数据库迁移
	if settingConfig != nil && settingConfig.EnableMigrateDb {
		stdlog.Printf("|*** 加载程序计划任务：数据库迁移")
		dbmigrate.Run()
	}
	// 定时任务
	//if scheduleConfig != nil && scheduleConfig.EnableScheduleCron {
	//	stdlog.Printf("|*** 加载程序计划任务：定时任务")
	//	cronHandler, err := crons.InitCron()
	//	if err != nil {
	//		debugutil.Fatalf("%+v\n", err)
	//		return
	//	}
	//	cronHandler.Start()
	//	defer cronHandler.Stop()
	//}
	return nil
}

// loadingAppDependentEnv 加载程序运行环境
func loadingAppDependentEnv(engineHandler setup.Engine) error {
	stdlog.Println()
	stdlog.Println("|==================== 加载程序运行环境 开始 ====================|")
	defer stdlog.Println("|==================== 加载程序运行环境 结束 ====================|")
	return nil
}
