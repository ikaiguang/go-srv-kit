package logpkg

import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsyncWriterCloseFlushesAndRejectsWrites(t *testing.T) {
	var output bytes.Buffer
	writer := NewAsyncWriter(&output, 128)
	for range 32 {
		_, err := writer.Write([]byte("line\n"))
		require.NoError(t, err)
	}

	writer.Close()
	require.Equal(t, 32, bytes.Count(output.Bytes(), []byte("line\n")))
	_, err := writer.Write([]byte("late"))
	require.True(t, errors.Is(err, os.ErrClosed))
}

func TestAsyncWriterConcurrentWriteAndClose(_ *testing.T) {
	writer := NewAsyncWriter(io.Discard, 128)
	var writers sync.WaitGroup
	for range 16 {
		writers.Add(1)
		go func() {
			defer writers.Done()
			for range 100 {
				_, _ = writer.Write([]byte("line\n"))
			}
		}()
	}

	writer.Close()
	writers.Wait()
}
