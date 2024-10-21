package main

import (
	"flag"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	consulutil "github.com/ikaiguang/go-srv-kit/service/consul"
	storeutil "github.com/ikaiguang/go-srv-kit/service/store"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf  string
	sourceDir string
	storeDir  string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&sourceDir, "source_dir", "", "store source path, eg: -source_dir path/to/source_dir")
	flag.StringVar(&storeDir, "store_dir", "", "custom store path, eg: -store_dir project_name/service_name/store_dir")
}

// go run ./cmd/store-configuration/... -conf=./configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	bootstrap, err := configutil.LoadingFile(flagconf)
	if err != nil {
		panic(err)
	}

	consulManager, err := consulutil.NewConsulManager(bootstrap.GetConsul())
	if err != nil {
		panic(err)
	}
	defer func() { _ = consulManager.Close() }()
	consulClient, err := consulManager.GetClient()
	if err != nil {
		panic(err)
	}

	if sourceDir == "" {
		sourceDir = flagconf
	}
	if storeDir == "" {
		storeDir = apputil.Path(apputil.ToAppConfig(bootstrap.GetApp()))
	}
	err = storeutil.StoreInConsul(consulClient, sourceDir, storeDir)
	if err != nil {
		panic(err)
	}
}
