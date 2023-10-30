package middlewarepkg

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TracerExporterType string

const (
	TracerExporterTypeUnknown TracerExporterType = "UNKNOWN"
	TracerExporterTypeJaeger  TracerExporterType = "JAEGER"
)

// tracerOptions ...
type tracerOptions struct {
	jaegerExporterType     TracerExporterType
	jaegerExporterInstance *jaeger.Exporter
}

// TracerOption is config option.
type TracerOption func(*tracerOptions)

// WithTracerJaegerExporter with config writer.
func WithTracerJaegerExporter(exporter *jaeger.Exporter) TracerOption {
	return func(o *tracerOptions) {
		o.jaegerExporterType = TracerExporterTypeJaeger
		o.jaegerExporterInstance = exporter
	}
}

// SetTracer set trace provider
// serviceNameKey == apputil.ID(appConfig)
func SetTracer(serviceNameKey string, opts ...TracerOption) error {
	opt := &tracerOptions{}
	for i := range opts {
		opts[i](opt)
	}

	providerOptions := []tracesdk.TracerProviderOption{
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceNameKey),
			//attribute.String("env", appConfig.RuntimeEnv),
			//attribute.String("version", appConfig.Version),
			//attribute.String("branch", appConfig.EnvBranch),
		)),
	}
	switch opt.jaegerExporterType {
	case TracerExporterTypeJaeger:
		if opt.jaegerExporterInstance != nil {
			// Always be sure to batch in production.
			providerOptions = append(providerOptions, tracesdk.WithBatcher(opt.jaegerExporterInstance))
		}
	}
	tp := tracesdk.NewTracerProvider(providerOptions...)
	otel.SetTracerProvider(tp)
	return nil
}

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
