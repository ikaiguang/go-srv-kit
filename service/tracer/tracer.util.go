package tracerutil

import (
	configpb "github.com/ikaiguang/go-srv-kit/api/config"
	middlewarepkg "github.com/ikaiguang/go-srv-kit/kratos/middleware"
	apputil "github.com/ikaiguang/go-srv-kit/service/app"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	stdlog "log"
)

func InitTracerWithJaegerExporter(appConfig *configpb.App, exp *otlptrace.Exporter) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	// Create the Jaeger exporter
	var opts = []middlewarepkg.TracerOption{
		middlewarepkg.WithTracerJaegerExporter(exp),
	}
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)), opts...)
}

func InitTracer(appConfig *configpb.App) error {
	stdlog.Println("|*** LOADING: Tracer: ...")
	return middlewarepkg.SetTracer(apputil.ID(apputil.ToAppConfig(appConfig)))
}
