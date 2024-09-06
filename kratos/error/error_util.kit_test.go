package errorpkg

import (
	"github.com/go-kratos/kratos/v2/errors"
	"testing"
)

// go test -v -count=1 ./kratos/error -test.run=TestCause
func TestCause(t *testing.T) {
	e := ErrorRecordNotFound("testdata")
	err := WithStack(e)
	t.Logf("err: %+v\n", err)
	cErr := Cause(err)
	var e2 *errors.Error
	ok2 := errors.As(cErr, &e2)
	if !ok2 {
		t.Errorf("not errors.Error")
		t.FailNow()
	}
	err2 := FormatError(err)
	t.Logf("err2: %+v\n", err2)
	ok3 := IsCustomError(err, ERROR_RECORD_NOT_FOUND)
	if !ok3 {
		t.Errorf("not ERROR_RECORD_NOT_FOUND")
		t.FailNow()
	}
}
