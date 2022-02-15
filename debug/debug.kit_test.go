package debugutil

import (
	"os"
	"testing"

	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	syncFn, err := Setup()
	if err != nil {
		panic(err)
	}
	defer func() { _ = syncFn() }()

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
	Debugf("%+v", errors.New("error 1"))
	Debugf("%+v", errors.New("error 2"))
}

// go test -v ./debug/ -count=1 -test.run=TestFatal
func TestFatal(t *testing.T) {
	//Fatal("==> fatal")
}
