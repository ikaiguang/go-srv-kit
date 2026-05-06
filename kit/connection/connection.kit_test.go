package connectionpkg

import (
	"errors"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsWebSocketConn(t *testing.T) {
	t.Run("nil request", func(t *testing.T) {
		assert.False(t, IsWebSocketConn(nil))
	})

	t.Run("websocket headers", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "http://example.com/ws", nil)
		require.NoError(t, err)
		req.Header.Set("Connection", "keep-alive, Upgrade")
		req.Header.Set("Upgrade", "websocket")

		assert.True(t, IsWebSocketConn(req))
	})

	t.Run("missing upgrade", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "http://example.com/ws", nil)
		require.NoError(t, err)

		assert.False(t, IsWebSocketConn(req))
	})
}

func TestIsValidConnection(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer func() { _ = listener.Close() }()

	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			_ = conn.Close()
		}
	}()

	got, err := IsValidConnection(listener.Addr().String())
	require.NoError(t, err)
	assert.True(t, got)
	<-done
}

func TestCheckEndpointValidity(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer func() { _ = listener.Close() }()

	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, acceptErr := listener.Accept()
		if acceptErr == nil {
			_ = conn.Close()
		}
	}()

	got, err := CheckEndpointValidity("http://" + listener.Addr().String())
	require.NoError(t, err)
	assert.True(t, got)
	<-done
}

func TestIsConnCloseErr(t *testing.T) {
	assert.False(t, IsConnCloseErr(errors.New("use of closed network connection")))

	err := &net.OpError{Err: errors.New("use of closed network connection")}
	assert.True(t, IsConnCloseErr(err))
}
