package configutil

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	stdlog "log"
)

func LoadingFile(filePath string, loadingOpts ...Option) (*configpb.Bootstrap, error) {
	stdlog.Println("|==================== LOADING FILE CONFIGURATION : START ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== LOADING FILE CONFIGURATION : END ====================|")
	loadOpts := &options{}
	for i := range loadingOpts {
		loadingOpts[i](loadOpts)
	}

	p, err := apputil.RuntimePath()
	if err != nil {
		return nil, err
	}
	stdlog.Println("|*** INFO: program running path: ", p)

	var configOpts []config.Option
	stdlog.Println("|*** LOADING: file configuration path: ", filePath)
	configOpts = append(configOpts, config.WithSource(file.NewSource(filePath)))

	handler := config.New(configOpts...)
	defer func() {
		stdlog.Println("|*** LOADING: COMPLETE : file configuration path: ", filePath)
		_ = handler.Close()
	}()

	// 加载配置
	if err = handler.Load(); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}

	// 读取配置文件
	conf := &configpb.Bootstrap{}
	if err = handler.Scan(conf); err != nil {
		err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
		return nil, err
	}
	for i := range loadOpts.configs {
		if loadOpts.configs[i] == nil {
			continue
		}
		if err = handler.Scan(loadOpts.configs[i]); err != nil {
			err = errorpkg.WithStack(errorpkg.ErrorInternalError(err.Error()))
			return nil, err
		}
	}
	return conf, nil
}
