package jaegerpkg

import (
	"context"
	"net"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func TestNewExporterRejectsInvalidConfig(t *testing.T) {
	if _, err := NewExporter(nil); err == nil {
		t.Fatal("NewExporter(nil) error = nil")
	}
	if _, err := NewExporter(&Config{}); err == nil {
		t.Fatal("NewExporter() invalid address error = nil")
	}
}

func TestNewExporter(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen() error = %v", err)
	}
	defer listener.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			conn, acceptErr := listener.Accept()
			if acceptErr != nil {
				return
			}
			_ = conn.Close()
		}
	}()

	for _, kind := range []Kind{KindHTTP, KindGRPC} {
		t.Run(string(kind), func(t *testing.T) {
			exporter, exportErr := NewExporter(&Config{
				Kind:       string(kind),
				Addr:       listener.Addr().String(),
				IsInsecure: true,
				Timeout:    durationpb.New(time.Second),
			})
			if exportErr != nil {
				t.Fatalf("NewExporter() error = %v", exportErr)
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			if shutdownErr := exporter.Shutdown(ctx); shutdownErr != nil {
				t.Fatalf("Exporter.Shutdown() error = %v", shutdownErr)
			}
		})
	}

	_ = listener.Close()
	<-done
}

func TestDirectExporterConstructorsRejectNilConfig(t *testing.T) {
	if _, err := NewHTTPExporter(nil); err == nil {
		t.Fatal("NewHTTPExporter(nil) error = nil")
	}
	if _, err := NewGRPCExporter(nil); err == nil {
		t.Fatal("NewGRPCExporter(nil) error = nil")
	}
}
