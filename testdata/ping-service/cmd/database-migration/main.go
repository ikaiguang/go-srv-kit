package main

import (
	"flag"
	stdlog "log"

	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
	postgresutil "github.com/ikaiguang/go-srv-kit/service/postgres"
	setuputil "github.com/ikaiguang/go-srv-kit/service/setup"
	dbmigrate "github.com/ikaiguang/go-srv-kit/testdata/ping-service/cmd/database-migration/migrate"
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

// go run ./cmd/store-configuration/... -conf=./configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	launcher, cleanup, err := setuputil.NewWithCleanupOptions(flagconf, nil, postgresutil.WithSetup())
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	defer cleanup()

	db, err := postgresutil.GetDB(launcher)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	dbmigrate.Run(db, dbutil.WithClose())
}
