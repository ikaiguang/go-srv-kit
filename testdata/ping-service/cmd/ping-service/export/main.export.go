package serviceexporter

import (
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	setupv2 "github.com/ikaiguang/go-srv-kit/service/setup_v2"
	pingapi "github.com/ikaiguang/go-srv-kit/testdata/ping-service/api/ping-service"
	dbmigrate "github.com/ikaiguang/go-srv-kit/testdata/ping-service/cmd/database-migration/migrate"
	"github.com/ikaiguang/go-srv-kit/testdata/ping-service/internal/conf"
)

func ExportServiceConfig() []configutil.Option {
	return conf.LoadServiceConfig()
}

func ExportAuthWhitelist() []map[string]middlewareutil.TransportServiceKind {
	return []map[string]middlewareutil.TransportServiceKind{
		pingapi.GetAuthWhiteList(),
	}
}

func ExportServices(launcherManager setupv2.LauncherManager, serverManager serverutil.ServerManager) (cleanuputil.CleanupManager, error) {
	hs, err := serverManager.GetHTTPServer()
	if err != nil {
		return nil, err
	}
	gs, err := serverManager.GetGRPCServer()
	if err != nil {
		return nil, err
	}
	return exportServices(launcherManager, hs, gs)
	//return serverutil.MergeCleanup(exportServices(launcherManager, hs, gs))
}

func ExportDatabaseMigration() []dbutil.MigrationFunc {
	return []dbutil.MigrationFunc{
		dbmigrate.Run,
	}
}
