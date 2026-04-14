package main

import (
	"flag"
	stdlog "log"

	dbutil "github.com/ikaiguang/go-srv-kit/service/database"
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

	launcher, err := setuputil.NewLauncherManager(flagconf)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	db, err := setuputil.GetRecommendDBConn(launcher)
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	dbmigrate.Run(db, dbutil.WithClose())
}
