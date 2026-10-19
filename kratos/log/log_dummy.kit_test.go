package logpkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -run TestNewDummyLogger
func TestNewDummyLogger(t *testing.T) {
	logImpl, err := NewDummyLogger()
	require.Nil(t, err)

	Setup(logImpl)
	Error("err")
}
