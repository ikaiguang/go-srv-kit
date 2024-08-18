package jaegerpkg

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func setTracerProvider(ctx context.Context) error {
	// Create the Jaeger exporter
	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("jaeger:4317"), otlptracegrpc.WithInsecure())
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		//tracesdk.WithResource(resource.NewSchemaless(
		//	semconv.ServiceNameKey.String(Name),
		//	attribute.String("env", "dev"),
		//)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
