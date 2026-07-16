package rotatelogs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
)

func TestRotateLogsWriteLifecycle(t *testing.T) {
	rl, err := New(filepath.Join(t.TempDir(), "app-%Y%m%d.log"))
	if err != nil {
		t.Fatal(err)
	}

	n, err := rl.Write([]byte("first"))
	if err != nil {
		t.Fatalf("first write failed: %v", err)
	}
	if n != len("first") {
		t.Fatalf("first write length = %d, want %d", n, len("first"))
	}
	if err := rl.Close(); err != nil {
		t.Fatalf("close failed: %v", err)
	}
	if _, err := rl.Write([]byte("after-close")); !errors.Is(err, os.ErrClosed) {
		t.Fatalf("write after close error = %v, want os.ErrClosed", err)
	}
	if err := rl.Rotate(); !errors.Is(err, os.ErrClosed) {
		t.Fatalf("rotate after close error = %v, want os.ErrClosed", err)
	}
	if err := rl.Close(); err != nil {
		t.Fatalf("second close should be idempotent: %v", err)
	}
}

func TestRotateLogsClosePropagatesFileError(t *testing.T) {
	rl, err := New(filepath.Join(t.TempDir(), "app-%Y%m%d.log"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := rl.Write([]byte("first")); err != nil {
		t.Fatal(err)
	}
	if err := rl.outFh.Close(); err != nil {
		t.Fatal(err)
	}
	if err := rl.Close(); err == nil {
		t.Fatal("expected Close to propagate the underlying file error")
	}
}

func TestGenFilename(t *testing.T) {
	// Mock time
	ts := []time.Time{
		time.Time{},
		(time.Time{}).Add(24 * time.Hour),
	}

	for _, xt := range ts {
		rl, err := New(
			"/path/to/%Y/%m/%d",
			WithClock(clockwork.NewFakeClockAt(xt)),
		)
		if !assert.NoError(t, err, "New should succeed") {
			return
		}

		defer rl.Close()

		fn := rl.genFilename()
		expected := fmt.Sprintf("/path/to/%04d/%02d/%02d",
			xt.Year(),
			xt.Month(),
			xt.Day(),
		)

		if !assert.Equal(t, expected, fn) {
			return
		}
	}
}
