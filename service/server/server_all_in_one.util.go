package serverutil

import (
	"github.com/go-kratos/kratos/v2"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	cleanuputil "github.com/ikaiguang/go-srv-kit/service/cleanup"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	middlewareutil "github.com/ikaiguang/go-srv-kit/service/middleware"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	stdlog "log"
)

func RunServer(app *kratos.App, cleanup func()) {
	defer func() {
		if cleanup != nil {
			cleanup()
		}
	}()
	// start
	if err := app.Run(); err != nil {
		stdlog.Fatalf("==> app.Run failed: %+v\n", err)
	}
}

type ServiceExporter func(launcherManager setuputil.LauncherManager, serverManager ServerManager) (cleanuputil.CleanupManager, error)

func AllInOneServer(
	configFilePath string,
	configOpts []configutil.Option,
	services []ServiceExporter,
	authWhitelist []map[string]middlewareutil.TransportServiceKind,
) (*kratos.App, func(), error) {
	if len(services) == 0 {
		e := errorpkg.ErrorBadRequest("services cannot be empty")
		return nil, nil, errorpkg.WithStack(e)
	}

	var (
		err            error
		cleanupManager = cleanuputil.NewCleanupManager()
	)
	defer func() {
		if err != nil {
			cleanupManager.Cleanup()
		}
	}()

	// launcher
	launcherManager, cleanup, err := setuputil.NewLauncherManagerWithCleanup(configFilePath, configOpts...)
	if err != nil {
		return nil, nil, err
	}
	cleanupManager.Append(cleanup)

	// whitelist
	whitelist := middlewareutil.MergeWhitelist(authWhitelist...)
	srvManager, err := NewServerManager(launcherManager, whitelist)
	if err != nil {
		return nil, nil, err
	}

	// services
	for i := range services {
		srvCleanup, serviceErr := services[i](launcherManager, srvManager)
		if serviceErr != nil {
			err = serviceErr
			return nil, nil, err
		}
		cleanupManager.Append(srvCleanup.Cleanup)
	}

	app, err := srvManager.GetApp()
	if err != nil {
		return nil, nil, err
	}
	stopApp := func() {
		closeErr := app.Stop()
		if closeErr != nil {
			stdlog.Printf("==> app.Stop failed: %+v\n", closeErr)
		}
	}
	cleanupManager.Append(stopApp)
	return app, cleanupManager.Cleanup, nil
}
