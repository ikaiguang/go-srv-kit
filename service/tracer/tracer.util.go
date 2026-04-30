package tracerutil

import (
	stdlog "log"

	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
)

func InitTracerWithOptions(appConfig *configpb.App, opts ...middlewarepkg.TracerOption) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)), opts...)
}

func InitTracer(appConfig *configpb.App) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)))
}
