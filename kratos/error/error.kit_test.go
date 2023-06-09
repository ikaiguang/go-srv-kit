package errorpkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./kratos/error -test.run=TestNew
func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		code       int
		reason     string
		wantCode   int
		wantReason string
	}{
		{
			name:       "code:1_reason:reason",
			code:       99,
			reason:     "reason",
			wantCode:   99,
			wantReason: "reason",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := genTestError(tt.code, tt.reason, tt.reason)
			t.Logf("%+v", err)
			cause := FromError(err)
			require.NotNil(t, cause)
			// code
			require.Equal(t, tt.wantCode, int(cause.Code))
			require.Equal(t, tt.wantCode, Code(err))
			// reason
			require.Equal(t, tt.wantReason, cause.Reason)
			require.Equal(t, tt.wantReason, Reason(err))
		})
	}
}

func genTestError(code int, reason, message string) error {
	return New(code, reason, message)
}
