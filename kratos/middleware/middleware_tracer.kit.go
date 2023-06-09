package middlewarepkg

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// SetTracerProvider set trace provider
// serviceNameKey == apputil.ID(appConfig)
func SetTracerProvider(serviceNameKey string, exporter *jaeger.Exporter) error {
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceNameKey),
			//attribute.String("env", appConfig.RuntimeEnv),
			//attribute.String("version", appConfig.Version),
			//attribute.String("branch", appConfig.EnvBranch),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
