package main

import (
	"flag"
	"fmt"
	consulpkg "github.com/ikaiguang/go-srv-kit/data/consul"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	configutil "github.com/ikaiguang/go-srv-kit/service/config"
	consulutil "github.com/ikaiguang/go-srv-kit/service/consul"
	storeutil "github.com/ikaiguang/go-srv-kit/service/store"
	"path/filepath"
	"runtime"
)

const (
	serverNameSuffix = "-service"
)

var (
	configPath string
	sourceDir  string
	storeDir   string
)

func init() {
	flag.StringVar(&configPath, "consul_config", "", "consul config path, eg: -consul_config ./configs")
	flag.StringVar(&sourceDir, "source_dir", "", "store source path, eg: -source_dir path/to/source_dir")
	flag.StringVar(&storeDir, "store_dir", "", "custom store path, eg: -store_dir project_name/service_name/store_dir")
}

func currentPath() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	}()

	// source
	sourcePath := sourceDir
	if sourceDir == "" {
		err = errorpkg.ErrorBadRequest("请配置资源目录：source_dir")
		panic(err)
	}
	if !filepath.IsAbs(sourceDir) {
		sourcePath = filepath.Join(currentPath(), sourceDir)
	}

	// 配置
	confPath := configPath
	if confPath == "" {
		err = errorpkg.ErrorBadRequest("请配置consul config")
		panic(err)
	}
	if !filepath.IsAbs(confPath) {
		confPath = filepath.Join(currentPath(), confPath)
	}
	bootConfig, err := configutil.LoadingFile(confPath)
	if err != nil {
		panic(err)
	}
	if bootConfig.GetConsul() == nil {
		e := errorpkg.ErrorBadRequest("请先配置Consul配置再试")
		panic(e)
	}

	// store dir
	if storeDir == "" {
		bs, err := configutil.LoadingFile(sourcePath)
		if err != nil {
			e := errorpkg.ErrorBadRequest(err.Error())
			panic(e)
		}
		storeDir = apputil.Path(apputil.ToAppConfig(bs.GetApp()))
	}
	if storeDir == "" {
		e := errorpkg.ErrorBadRequest("请配置存储路径：store_dir")
		panic(e)
	}

	// consul
	cc, err := consulpkg.NewClient(consulutil.ToConsulConfig(bootConfig.GetConsul()))
	if err != nil {
		e := errorpkg.ErrorInternalServer(err.Error())
		panic(e)
	}
	err = storeutil.StoreInConsul(cc, sourcePath, storeDir)
	if err != nil {
		panic(err)
	}
}
