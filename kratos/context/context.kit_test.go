package contextpkg

import (
	"context"
	stdhttp "net/http"
	"sync"
	"testing"

	"github.com/go-kratos/kratos/v3/transport"
	grpcTransport "github.com/go-kratos/kratos/v3/transport/grpc"
	httpTransport "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/stretchr/testify/require"
)

type testHeader map[string]string

func (h testHeader) Get(key string) string      { return h[key] }
func (h testHeader) Set(key, value string)      { h[key] = value }
func (h testHeader) Add(key, value string)      { h[key] = value }
func (h testHeader) Keys() []string             { return nil }
func (h testHeader) Values(key string) []string { return []string{h[key]} }

type testHTTPTransport struct {
	request *stdhttp.Request
}

func TestTrustedPlatformConcurrentAccess(_ *testing.T) {
	defer SetTrustedPlatform("x-kit-remote-addr")
	var wg sync.WaitGroup
	for range 16 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			SetTrustedPlatform("x-forwarded-for")
		}()
		go func() {
			defer wg.Done()
			_ = trustedPlatform()
		}()
	}
	wg.Wait()
}

func TestClientIPFromHTTPIgnoresInvalidTrustedHeader(t *testing.T) {
	defer SetTrustedPlatform("x-kit-remote-addr")
	SetTrustedPlatform("x-trusted-client-ip")
	request, err := stdhttp.NewRequest(stdhttp.MethodGet, "/", nil)
	require.NoError(t, err)
	request.RemoteAddr = "192.0.2.10:1234"
	request.Header.Set("x-trusted-client-ip", "not-an-ip")

	require.Equal(t, "192.0.2.10", ClientIPFromHTTP(context.Background(), request))
}

func TestClientIPFromHTTPAcceptsNilRequest(t *testing.T) {
	require.Empty(t, ClientIPFromHTTP(context.Background(), nil))
}

func (testHTTPTransport) Kind() transport.Kind            { return transport.KindHTTP }
func (testHTTPTransport) Endpoint() string                { return "http://example.test" }
func (testHTTPTransport) Operation() string               { return "/test" }
func (testHTTPTransport) RequestHeader() transport.Header { return testHeader{} }
func (testHTTPTransport) ReplyHeader() transport.Header   { return testHeader{} }
func (tr testHTTPTransport) Request() *stdhttp.Request    { return tr.request }
func (testHTTPTransport) PathTemplate() string            { return "/test" }

func TestMatchClientContexts(t *testing.T) {
	httpTr := testHTTPTransport{request: &stdhttp.Request{}}
	httpCtx := transport.NewClientContext(context.Background(), httpTr)
	gotHTTP, ok := MatchHTTPClientContext(httpCtx)
	require.True(t, ok)
	require.Equal(t, httpTr.request, gotHTTP.Request())

	grpcTr := &grpcTransport.Transport{}
	grpcCtx := transport.NewClientContext(context.Background(), grpcTr)
	gotGRPC, ok := MatchGRPCClientContext(grpcCtx)
	require.True(t, ok)
	require.Same(t, grpcTr, gotGRPC)

	_, ok = MatchHTTPClientContext(transport.NewServerContext(context.Background(), httpTr))
	require.False(t, ok)
	var _ httpTransport.Transporter = httpTr
}
