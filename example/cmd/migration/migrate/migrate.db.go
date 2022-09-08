package dbmigrate

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	dbv1_0_0 "github.com/ikaiguang/go-srv-kit/example/cmd/migration/v1.0.0"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"

	stdlog "log"

	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// Run 运行迁移
func Run(opts ...Option) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}

	// 初始化
	if err := setup.Init(); err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// 模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	if opt.closeEngine {
		defer func() { _ = setup.Close() }()
	}

	// 数据库链接
	dbConn, err := engineHandler.GetMySQLGormDB()
	//dbConn, err := engineHandler.GetPostgresGormDB()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	migrateRepo := migrationuitl.NewMigrateRepo(dbConn)

	// 初始化迁移记录
	if err = migrateRepo.InitializeMigrationSchema(); err != nil {
		logutil.Fatalw("migrateHandler.InitializeMigrationSchema failed", err)
	}

	// v1.0.0
	if err = dbv1_0_0.Upgrade(dbConn, migrateRepo); err != nil {
		logutil.Fatalw("dbv1_0_0.Upgrade failed", err)
	}
}
