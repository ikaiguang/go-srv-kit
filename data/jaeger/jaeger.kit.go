package jaegerutil

import (
	"go.opentelemetry.io/otel/exporters/jaeger"

	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
)

// NewJaegerExporter ...
func NewJaegerExporter(conf *confv1.Data_JaegerTrace, opts ...Option) (*jaeger.Exporter, error) {
	return NewExporter(conf, opts...)
}

// NewExporter jaeger.Exporter
func NewExporter(conf *confv1.Data_JaegerTrace, opts ...Option) (*jaeger.Exporter, error) {
	var jaegerOptions []jaeger.CollectorEndpointOption
	if conf.Endpoint != "" {
		jaegerOptions = append(jaegerOptions, jaeger.WithEndpoint(conf.Endpoint))
	}
	if conf.Username != "" {
		jaegerOptions = append(jaegerOptions, jaeger.WithUsername(conf.Username))
	}
	if conf.Password != "" {
		jaegerOptions = append(jaegerOptions, jaeger.WithPassword(conf.Password))
	}

	return jaeger.New(jaeger.WithCollectorEndpoint(jaegerOptions...))
}
