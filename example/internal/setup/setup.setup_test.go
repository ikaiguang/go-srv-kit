package setuphandler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./example/internal/setup/ -count=1 -test.run=TestSetup
func TestSetup(t *testing.T) {
	err := Setup()
	require.Nil(t, err)
}
