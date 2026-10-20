package errorpkg

import (
	stderrors "errors"
	"testing"

	"github.com/go-kratos/kratos/v3/errors"
	"github.com/stretchr/testify/require"
)

func TestWrapKvsPreservesMetadata(t *testing.T) {
	base := errors.New(500, "TEST", "test")
	err := WrapKvs(base, "request_id", "request-1")

	require.Equal(t, "request-1", err.GetMetadata()["request_id"])
}

func TestErrorSupportsStandardUnwrap(t *testing.T) {
	base := errors.New(500, "TEST", "test")
	wrapped := WithStack(base)
	var target *errors.Error

	require.True(t, stderrors.As(wrapped, &target))
	require.Same(t, base, target)
}

func TestWrapKeepsCauseOutOfPublicMetadata(t *testing.T) {
	base := errors.New(500, "TEST", "test")
	cause := stderrors.New("database address")
	wrapped := Wrap(base, cause)

	require.True(t, stderrors.Is(wrapped, cause))
	require.NotContains(t, wrapped.GetMetadata(), EnumMDErrorKey+".0")
}

func TestCallStackAppliesSkipBeforeDepth(t *testing.T) {
	err := newNestedError()
	full := CallStackWithSkipAndDepth(err, 0, 4)
	skipped := CallStackWithSkipAndDepth(err, 1, 3)

	require.Len(t, full, 4)
	require.Len(t, skipped, 3)
	require.Equal(t, full[1:], skipped)
	require.Empty(t, CallStackWithSkipAndDepth(err, len(Stack(err)), 1))
}

func newNestedError() error {
	return New(500, "TEST", "test")
}

type cyclicError struct{}

func (e *cyclicError) Error() string { return "cycle" }
func (e *cyclicError) Cause() error  { return e }

func TestCauseStopsOnCycles(t *testing.T) {
	err := &cyclicError{}
	require.Same(t, err, Cause(err))
}
