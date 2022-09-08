package setuppkg

import (
	pkgerrors "github.com/pkg/errors"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdlog "log"
	"sync"

	jaegerutil "github.com/ikaiguang/go-srv-kit/data/jaeger"
)

// GetJaegerTraceExporter jaegerTrace
func (s *engines) GetJaegerTraceExporter() (*jaeger.Exporter, error) {
	var err error
	s.jaegerTraceExporterMutex.Do(func() {
		s.jaegerTraceExporter, err = s.loadingJaegerTraceExporter()
	})
	if err != nil {
		s.jaegerTraceExporterMutex = sync.Once{}
	}
	return s.jaegerTraceExporter, err
}

// loadingJaegerTraceExporter jaegerTrace
func (s *engines) loadingJaegerTraceExporter() (*jaeger.Exporter, error) {
	if s.Config.JaegerTracerConfig() == nil {
		stdlog.Println("|*** 加载：JaegerTrace：未初始化")
		return nil, pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Println("|*** 加载：JaegerTrace：...")

	return jaegerutil.NewJaegerExporter(s.Config.JaegerTracerConfig())
}
