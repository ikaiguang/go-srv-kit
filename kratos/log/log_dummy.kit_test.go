package logpkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
)

// go test -v ./log/ -count=1 -test.run=TestNewDummyLogger
func TestNewDummyLogger(t *testing.T) {
	logImpl, err := NewDummyLogger()
	require.Nil(t, err)

	logHandler := log.NewHelper(logImpl)
	logHandler.Error("err")
}
