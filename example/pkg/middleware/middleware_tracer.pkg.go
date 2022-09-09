package middlewarepkg

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	setuppkg "github.com/ikaiguang/go-srv-kit/example/pkg/setup"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
)

// SetTracerProvider set trace provider
func SetTracerProvider(engineHandler setuppkg.Engine) error {
	appConfig := engineHandler.AppConfig()

	// Create the Jaeger exporter
	exp, err := engineHandler.GetJaegerTraceExporter()
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(apputil.ID(appConfig)),
			//attribute.String("env", appConfig.Env),
			//attribute.String("version", appConfig.Version),
			//attribute.String("branch", appConfig.EnvBranch),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
