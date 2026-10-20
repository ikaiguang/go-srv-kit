package apppkg

import (
	"errors"
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	kratoserrors "github.com/go-kratos/kratos/v3/errors"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/v3/error"
	"github.com/stretchr/testify/require"
)

type trackingResponseWriter struct {
	header      stdhttp.Header
	wroteHeader bool
}

func TestErrorResponseEncoderRemovesInternalErrorMetadata(t *testing.T) {
	RegisterCodec()
	request := httptest.NewRequest(stdhttp.MethodGet, "/", nil)
	request.Header.Set("Accept", "application/json")
	response := httptest.NewRecorder()
	base := kratoserrors.New(stdhttp.StatusInternalServerError, "INTERNAL", "request failed")
	err := errorpkg.WrapKvs(
		base,
		"request_id", "request-1",
		errorpkg.EnumMDErrorKey+".0", "secret database address",
	)

	ErrorResponseEncoder(response, request, err)

	body := response.Body.String()
	require.NotContains(t, body, "secret database address")
	require.NotContains(t, body, `"error.0"`)
	require.Contains(t, body, "request_id")
}

func TestErrorResponseEncoderNormalizesInvalidStatus(t *testing.T) {
	RegisterCodec()
	request := httptest.NewRequest(stdhttp.MethodGet, "/", nil)
	request.Header.Set("Accept", "application/json")
	response := httptest.NewRecorder()

	require.NotPanics(t, func() {
		ErrorResponseEncoder(response, request, kratoserrors.New(0, "INVALID", "invalid"))
	})
	require.Equal(t, stdhttp.StatusInternalServerError, response.Code)
}

func TestErrorResponseEncoderRedactsUnknownInternalMessage(t *testing.T) {
	RegisterCodec()
	request := httptest.NewRequest(stdhttp.MethodGet, "/", nil)
	request.Header.Set("Accept", "application/json")
	response := httptest.NewRecorder()

	ErrorResponseEncoder(response, request, errors.New("secret database address"))

	require.NotContains(t, response.Body.String(), "secret database address")
	require.Contains(t, response.Body.String(), "internal server error")
}

func (w *trackingResponseWriter) Header() stdhttp.Header { return w.header }
func (w *trackingResponseWriter) Write(data []byte) (int, error) {
	w.wroteHeader = true
	return len(data), nil
}
func (w *trackingResponseWriter) WriteHeader(int) { w.wroteHeader = true }

type failingJSONMarshaler struct{}

func (failingJSONMarshaler) MarshalJSON() ([]byte, error) {
	return nil, errors.New("marshal failed")
}

func TestSuccessResponseEncoderMarshalsBeforeWritingHeader(t *testing.T) {
	RegisterCodec()
	request, err := stdhttp.NewRequest(stdhttp.MethodGet, "/", nil)
	require.NoError(t, err)
	request.Header.Set("Accept", "application/json")
	response := &trackingResponseWriter{header: make(stdhttp.Header)}

	err = SuccessResponseEncoder(response, request, failingJSONMarshaler{})
	require.Error(t, err)
	require.False(t, response.wroteHeader)
}
