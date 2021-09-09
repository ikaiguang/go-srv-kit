package errorutil

import (
	stderrors "errors"
	"net/http"
	"testing"
)

// go test -v ./error/ -count=1 -test.run=TestCaller
func TestCaller(t *testing.T) {
	tests := []struct {
		name        string
		error       error
		callerDepth int
	}{
		{
			name:  "#errorutil",
			error: wrap(http.StatusNotFound, "reason", "message"),
		},
		{
			name:  "#strerrors",
			error: stderrors.New("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callers := Caller(tt.error)
			for i := range callers {
				t.Log(callers[i])
				//t.Logf("%q", callers[i])
			}
		})
	}
}
