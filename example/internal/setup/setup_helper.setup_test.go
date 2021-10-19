package setup

import (
	"testing"

	"github.com/stretchr/testify/require"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
)

// go test -v ./example/internal/setup/ -count=1 -test.run=TestGetPackages -conf=./../../configs
func TestGetPackages(t *testing.T) {
	packages, err := GetPackages()
	require.Nil(t, err)
	require.NotNil(t, packages)
	packages, err = GetPackages()
	require.Nil(t, err)
	require.NotNil(t, packages)

	// env
	loghelper.Infof("env = %v", packages.Env())

	// debug
	debugutil.Println("*** | ==> debug util print")

	// 日志
	loghelper.Info("*** | ==> log helper info")
}
