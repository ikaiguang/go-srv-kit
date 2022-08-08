package main

import (
	stdlog "log"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	routes "github.com/ikaiguang/go-srv-kit/example/internal/route"
	servers "github.com/ikaiguang/go-srv-kit/example/internal/server"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
)

// go run ./example/cmd/main/... -conf=./example/configs
func main() {
	var err error

	// 初始化
	err = setup.Init()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 引擎模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = setup.Close() }()

	// loadingAppSettingEnv 加载程序配置
	if err = loadingAppSettingEnv(engineHandler); err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// 启动程序
	stdlog.Println()
	stdlog.Println("|==================== 加载程序 开始 ====================|")

	// 启动
	app, err := newApp(engineHandler)
	stdlog.Println("|==================== 加载程序 结束 ====================|")
	if err := app.Run(); err != nil {
		debugutil.Fatalf("%+v\n", err)
		return
	}
}

// loadingAppSettingEnv 加载计划任务
func loadingAppSettingEnv(engineHandler setup.Engine) error {
	// 计划任务
	settingConfig := engineHandler.AppSettingConfig()
	stdlog.Println()
	stdlog.Println("|==================== 加载计划任务 开始 ====================|")

	// 数据库迁移
	if settingConfig != nil && settingConfig.EnableMigrateDb {
		//stdlog.Printf("|*** 加载程序计划任务：数据库迁移")
		//dbmigrate.Run()
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
	stdlog.Println("|==================== 加载计划任务 结束 ====================|")
	return nil
}

// newApp .
func newApp(engineHandler setup.Engine) (app *kratos.App, err error) {
	// 主机
	hostname, _ := os.Hostname()

	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		return app, err
	}
	log.SetLogger(logger)

	// 加载程序运行环境
	err = loadingAppDependentEnv(engineHandler)
	if err != nil {
		return app, err
	}

	// 服务
	hs, err := servers.NewHTTPServer(engineHandler)
	if err != nil {
		return app, err
	}
	gs, err := servers.NewGRPCServer(engineHandler)
	if err != nil {
		return app, err
	}

	// 路由
	err = routes.RegisterRoutes(engineHandler, hs, gs)
	if err != nil {
		return app, err
	}

	// app
	appConfig := engineHandler.AppConfig()
	app = kratos.New(
		kratos.ID(hostname),
		kratos.Name(appConfig.Name),
		kratos.Version(appConfig.Version),
		kratos.Metadata(appConfig.Metadata),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
	)
	return app, err
}

// loadingAppDependentEnv 加载程序运行环境
func loadingAppDependentEnv(engineHandler setup.Engine) error {
	//stdlog.Printf("|*** 加载程序运行环境：XXX")
	return nil
}
