package tracerutil

import (
	stdlog "log"

	middlewarepkg "github.com/ikaiguang/go-kratos-kit/middleware"
	configpb "github.com/ikaiguang/go-service-kit/api/config"
	apputil "github.com/ikaiguang/go-service-kit/app"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
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
