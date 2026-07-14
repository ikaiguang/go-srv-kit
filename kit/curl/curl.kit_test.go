package curlpkg

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	req, err := NewRequest(http.MethodPut, "http://example.com", bytes.NewBufferString("body"))
	require.NoError(t, err)

	assert.Equal(t, http.MethodPut, req.Method)
	assert.Equal(t, "http://example.com", req.URL.String())

	postReq, err := NewPostRequest("http://example.com/post", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodPost, postReq.Method)

	getReq, err := NewGetRequest("http://example.com/get", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, getReq.Method)

	putReq, err := NewPutRequest("http://example.com/put", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, putReq.Method)

	patchReq, err := NewPatchRequest("http://example.com/patch", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodPatch, patchReq.Method)

	deleteReq, err := NewDeleteRequest("http://example.com/delete", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodDelete, deleteReq.Method)
}

func TestNewRequestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req, err := NewRequestContext(ctx, http.MethodDelete, "http://example.com/delete", nil)
	require.NoError(t, err)
	assert.Equal(t, http.MethodDelete, req.Method)
	assert.Equal(t, context.Canceled, req.Context().Err())

	req, err = NewPutRequestContext(nil, "http://example.com/put", bytes.NewBufferString("body"))
	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, req.Method)
	assert.NotNil(t, req.Context())
}

func TestNewHTTPClient(t *testing.T) {
	client := NewHTTPClient(WithTimeout(2 * time.Second))
	assert.Equal(t, 2*time.Second, client.Timeout)
	assert.Nil(t, client.Transport)

	insecureClient := NewHTTPClient(WithInsecureSkipVerify())
	tr, ok := insecureClient.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, tr.TLSClientConfig)
	assert.True(t, tr.TLSClientConfig.InsecureSkipVerify)
}

func TestDoAndDoWithClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	req, err := NewPostRequest(server.URL, bytes.NewBufferString("body"))
	require.NoError(t, err)

	code, body, err := Do(req, WithTimeout(time.Second))
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, []byte("ok"), body)

	req, err = NewPostRequest(server.URL, bytes.NewBufferString("body"))
	require.NoError(t, err)
	code, body, err = DoWithClient(server.Client(), req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, []byte("ok"), body)
}

func TestDefault(t *testing.T) {
	originalTransport := http.DefaultTransport
	defer func() { http.DefaultTransport = originalTransport }()

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("default"))
	}))
	defer server.Close()

	http.DefaultTransport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	req, err := NewGetRequest(server.URL, nil)
	require.NoError(t, err)

	code, body, err := Default(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, code)
	assert.Equal(t, []byte("default"), body)
}

func TestResponse(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusNoContent,
		Body:       io.NopCloser(errReader{}),
	}

	code, body, err := response(resp)
	require.Error(t, err)
	assert.Equal(t, http.StatusNoContent, code)
	assert.Empty(t, body)
}

func TestIsSuccessCodeAndErrRequestFailure(t *testing.T) {
	assert.True(t, IsSuccessCode(http.StatusOK))
	assert.True(t, IsSuccessCode(http.StatusNoContent))
	assert.False(t, IsSuccessCode(http.StatusMultipleChoices))
	assert.EqualError(t, ErrRequestFailure(http.StatusBadGateway), "request failure; code=502")
}

type errReader struct{}

func (errReader) Read(_ []byte) (int, error) {
	return 0, errors.New("read failed")
}
