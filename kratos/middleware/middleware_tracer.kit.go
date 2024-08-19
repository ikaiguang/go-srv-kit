package middlewarepkg

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TracerExporterType string

const (
	TracerExporterTypeUnknown TracerExporterType = "unknown"
	TracerExporterTypeJaeger  TracerExporterType = "jaeger"
)

// tracerOptions ...
type tracerOptions struct {
	exporterType   TracerExporterType
	jaegerExporter *otlptrace.Exporter
}

// TracerOption is config option.
type TracerOption func(*tracerOptions)

// WithTracerJaegerExporter with config writer.
func WithTracerJaegerExporter(exporter *otlptrace.Exporter) TracerOption {
	return func(o *tracerOptions) {
		o.exporterType = TracerExporterTypeJaeger
		o.jaegerExporter = exporter
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
	switch opt.exporterType {
	case TracerExporterTypeJaeger:
		if opt.jaegerExporter != nil {
			// Always be sure to batch in production.
			providerOptions = append(providerOptions, tracesdk.WithBatcher(opt.jaegerExporter))
		}
	}
	tp := tracesdk.NewTracerProvider(providerOptions...)
	otel.SetTracerProvider(tp)
	return nil
}

// SetTracerProvider set trace provider
// serviceNameKey == apputil.ID(appConfig)
func SetTracerProvider(serviceNameKey string, exporter *otlptrace.Exporter) error {
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
