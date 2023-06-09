package errorpkg

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./error/ -count=1 -test.run=TestHTTP_Error
func TestHTTP_Error(t *testing.T) {
	tests := []struct {
		name       string
		error      error
		wantCode   int
		wantReason string
		want       bool
	}{
		{
			name:       "#StatusNotFound",
			error:      NotFound("NotFound", "NotFound"),
			wantCode:   http.StatusNotFound,
			wantReason: "NotFound",
			want:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, IsCode(tt.error, tt.wantCode))
			require.Equal(t, tt.want, IsReason(tt.error, tt.wantReason))
		})
	}
}
