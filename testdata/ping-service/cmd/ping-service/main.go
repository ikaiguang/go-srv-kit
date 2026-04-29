package main

import (
	"flag"
	clientutil "github.com/ikaiguang/go-srv-kit/service/cluster_service_api"
	serverutil "github.com/ikaiguang/go-srv-kit/service/server"
	serviceexporter "github.com/ikaiguang/go-srv-kit/testdata/ping-service/cmd/ping-service/export"
	stdlog "log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	configOpts := serviceexporter.ExportServiceConfig()
	whitelist := serviceexporter.ExportAuthWhitelist()
	services := []serverutil.ServiceExporter{serviceexporter.ExportServices}
	setupOpts := []serverutil.Option{
		serverutil.WithSetupOptions(
			clientutil.WithSetup(),
		),
	}

	app, cleanup, err := serverutil.AllInOneServer(flagconf, configOpts, services, whitelist, setupOpts...)
	if err != nil {
		stdlog.Fatalf("==> runservices.GetServerApp failed: %+v\n", err)
	}
	serverutil.RunServer(app, cleanup)
}
