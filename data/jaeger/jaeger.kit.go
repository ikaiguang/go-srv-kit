package jaegerpkg

import (
	"context"
	"fmt"
	connectionpkg "github.com/ikaiguang/go-srv-kit/kit/connection"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Kind string

const (
	KingHTTP Kind = "http"
	KingGRPC Kind = "grpc"
)

// Config jaeger config
type Config struct {
	Kind              Kind
	Addr              string
	IsInsecure        bool
	WithHttpBasicAuth bool
	Username          string
	Password          string
	Timeout           *durationpb.Duration
}

// NewJaegerExporter ...
func NewJaegerExporter(conf *Config, opts ...Option) (*otlptrace.Exporter, error) {
	return NewExporter(conf, opts...)
}

// NewExporter jaeger.Exporter
func NewExporter(conf *Config, opts ...Option) (*otlptrace.Exporter, error) {
	isValidConnection, err := connectionpkg.CheckEndpointValidity(conf.Addr)
	if err != nil {
		err = fmt.Errorf("address error : %w", err)
		return nil, err
	}
	if !isValidConnection {
		err = fmt.Errorf("address error : invalid connection")
		return nil, err
	}
	if conf.Kind == KingHTTP {
		return NewHTTPExporter(conf)
	}
	return NewGRPCExporter(conf)
}

func NewHTTPExporter(conf *Config) (*otlptrace.Exporter, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(conf.Addr),
	}
	if conf.IsInsecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	if conf.Timeout.AsDuration() > 0 {
		opts = append(opts, otlptracehttp.WithTimeout(conf.Timeout.AsDuration()))
	}
	exp, err := otlptracehttp.New(context.Background(), opts...)
	if err != nil {
		err = fmt.Errorf("new http exporter error : %w", err)
		return nil, err
	}
	return exp, nil
}

func NewGRPCExporter(conf *Config) (*otlptrace.Exporter, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(conf.Addr),
	}
	if conf.IsInsecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	if conf.Timeout.AsDuration() > 0 {
		opts = append(opts, otlptracegrpc.WithTimeout(conf.Timeout.AsDuration()))
	}
	exp, err := otlptracegrpc.New(context.Background(), opts...)
	if err != nil {
		err = fmt.Errorf("new grpc exporter error : %w", err)
		return nil, err
	}
	return exp, nil
}
