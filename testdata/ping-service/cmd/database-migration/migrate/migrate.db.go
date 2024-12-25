package dbmigrate

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	"gorm.io/gorm"
)

// Run 运行迁移
func Run(dbConn *gorm.DB, opts ...dbutil.MigrationOption) {
	opt := dbutil.DefaultMigrationOptions()
	for i := range opts {
		opts[i](opt)
	}

	// 关闭
	if opt.Close {
		defer func() {
			db, err := dbConn.DB()
			if err != nil {
				return
			}
			_ = db.Close()
		}()
	}

	// migrateHandler 迁移手柄
	var (
		ctx         = context.Background()
		migrateRepo = migrationpkg.NewMigrateRepo(dbConn)
		logHandler  = log.NewHelper(log.With(opt.Logger, "module", "database/migration"))
	)

	// 初始化迁移记录
	if err := migrateRepo.InitializeSchema(ctx); err != nil {
		logHandler.Fatalf("migrateHandler.InitializeSchema failed: %+v", err)
	}

	// v1.0.0
	//if err := dbv1_0_0.Upgrade(ctx, dbConn, migrateRepo); err != nil {
	//	logHandler.Fatalf("dbv1_0_0.Upgrade failed: %+v", err)
	//}
}
