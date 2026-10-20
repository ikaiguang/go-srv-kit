package middlewarepkg

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCORSDoesNotCombineWildcardAndCredentials(t *testing.T) {
	handler := NewCORS()(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, _ *stdhttp.Request) {
		w.WriteHeader(stdhttp.StatusNoContent)
	}))
	req := httptest.NewRequest(stdhttp.MethodGet, "/", nil)
	req.Header.Set("Origin", "https://example.test")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	require.Equal(t, "*", response.Header().Get("Access-Control-Allow-Origin"))
	require.Empty(t, response.Header().Get("Access-Control-Allow-Credentials"))
}
