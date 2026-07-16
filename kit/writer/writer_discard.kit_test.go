package writerpkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscardWriterReportsConsumedBytes(t *testing.T) {
	w, err := NewDiscardWriter()
	require.NoError(t, err)

	p := []byte("discard me")
	n, err := w.Write(p)
	require.NoError(t, err)
	require.Equal(t, len(p), n)

	legacy, err := NewDummyWriter()
	require.NoError(t, err)
	n, err = legacy.Write(p)
	require.NoError(t, err)
	require.Equal(t, len(p), n)
}
