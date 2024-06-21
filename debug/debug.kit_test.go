package debugpkg

import (
	"os"
	"testing"

	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
)

func TestMain(m *testing.M) {
	closer, err := Setup()
	if err != nil {
		panic(err)
	}
	defer func() { _ = closer.Close() }()

	code := m.Run()

	os.Exit(code)
}

// go test -v ./debug/ -count=1 -test.run=TestDebug
func TestDebug(t *testing.T) {
	var msg = struct {
		Name string
		Age  int
	}{
		Name: "zhang",
		Age:  30,
	}

	Debug(msg)
}

// go test -v ./debug/ -count=1 -test.run=TestDebugf
func TestDebugf(t *testing.T) {
	Debugf("%+v", errorpkg.WithStack(errorpkg.ErrorBadRequest("error 1")))
	Debugf("%+v", errorpkg.WithStack(errorpkg.ErrorBadRequest("error 2")))
}

// go test -v ./debug/ -count=1 -test.run=TestFatal
func TestFatal(t *testing.T) {
	//Fatal("==> fatal")
}
