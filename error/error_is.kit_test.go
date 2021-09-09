package errorutil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./error/ -count=1 -test.run=TestIs_Xxx
func TestIs_Xxx(t *testing.T) {
	tests := []struct {
		name       string
		error      error
		wantCode   int
		wantReason string
		want       bool
	}{
		{
			name:       "#IsCode",
			error:      New(http.StatusBadRequest, "StatusBadRequest", "StatusBadRequest"),
			wantCode:   http.StatusBadRequest,
			wantReason: "StatusBadRequest",
			want:       true,
		},
		{
			name:       "#IsReason",
			error:      New(http.StatusBadRequest, "StatusBadRequest", "StatusBadRequest"),
			wantCode:   http.StatusBadRequest,
			wantReason: "StatusBadRequest",
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
