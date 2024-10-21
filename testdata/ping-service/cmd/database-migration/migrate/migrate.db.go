package dbmigrate

import (
	"context"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	logpkg "github.com/ikaiguang/go-srv-kit/kratos/log"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	stdlog "log"
)

// Run 运行迁移
func Run(launcherManager setuputil.LauncherManager, opts ...Option) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}

	// 关闭
	if opt.closeEngine {
		defer func() { _ = launcherManager.Close() }()
	}

	// 数据库链接
	//dbConn, err := setuputil.GetMysqlDBConn(launcherManager)
	dbConn, err := setuputil.GetPostgresDBConn(launcherManager)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	var (
		ctx         = context.Background()
		migrateRepo = migrationpkg.NewMigrateRepo(dbConn)
	)

	// 初始化迁移记录
	if err = migrateRepo.InitializeSchema(ctx); err != nil {
		logpkg.Fatalf("migrateHandler.InitializeSchema failed: %+v", err)
	}

	// v1.0.0
	//if err = dbv1_0_0.Upgrade(ctx, dbConn, migrateRepo); err != nil {
	//	logpkg.Fatalf("dbv1_0_0.Upgrade failed: %+v", err)
	//}
}
