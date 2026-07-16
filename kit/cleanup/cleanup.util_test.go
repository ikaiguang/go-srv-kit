package cleanuputil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendOrCleanupAndLegacyWrapper(t *testing.T) {
	called := false
	manager, err := AppendOrCleanup(nil, func() { called = true }, nil)
	require.NoError(t, err)
	require.NotNil(t, manager)
	manager.Cleanup()
	require.True(t, called)

	wantErr := errors.New("setup failed")
	manager, err = Merge(nil, nil, wantErr)
	require.ErrorIs(t, err, wantErr)
	require.Nil(t, manager)
}
