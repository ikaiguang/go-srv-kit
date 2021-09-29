package setuphandler

import (
	"testing"

	loghelper "github.com/ikaiguang/go-srv-kit/log/helper"
)

// go test -v ./example/internal/setup/ -count=1 -test.run=TestSetup -conf=./../../configs
func TestSetup(t *testing.T) {
	err := Setup()
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}

	loghelper.Info("*** | ==> Info")
}
