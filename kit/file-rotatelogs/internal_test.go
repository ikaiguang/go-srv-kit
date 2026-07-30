package rotatelogs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

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

func TestWriteContinuesAfterPreviousFileCloseError(t *testing.T) {
	rl, err := New(
		filepath.Join(t.TempDir(), "app.log"),
		WithRotationSize(1),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer rl.Close()

	if _, err := rl.Write([]byte("first")); err != nil {
		t.Fatal(err)
	}
	if err := rl.outFh.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := rl.Write([]byte("second")); err != nil {
		t.Fatalf("write should use the newly rotated file: %v", err)
	}

	content, err := os.ReadFile(rl.CurrentFileName())
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "second" {
		t.Fatalf("rotated file content = %q, want %q", content, "second")
	}
}

func TestNewValidatesRequiredValues(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		options []Option
	}{
		{name: "empty pattern"},
		{name: "nil clock", pattern: "app.log", options: []Option{WithClock(nil)}},
		{name: "nil location", pattern: "app.log", options: []Option{WithLocation(nil)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.pattern, tt.options...); err == nil {
				t.Fatal("New should reject invalid input")
			}
		})
	}
}

func TestNewIgnoresNilOption(t *testing.T) {
	rl, err := New(filepath.Join(t.TempDir(), "app.log"), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer rl.Close()

	if _, err := rl.Write([]byte("test")); err != nil {
		t.Fatal(err)
	}
}

func TestRotationCountPurgesGenerationalFiles(t *testing.T) {
	base := filepath.Join(t.TempDir(), "app.log")
	rl, err := New(base, WithRotationCount(2))
	if err != nil {
		t.Fatal(err)
	}
	defer rl.Close()

	if _, err := rl.Write([]byte("initial")); err != nil {
		t.Fatal(err)
	}
	for range 4 {
		if err := rl.Rotate(); err != nil {
			t.Fatal(err)
		}
		if _, err := rl.Write([]byte("next")); err != nil {
			t.Fatal(err)
		}
	}

	matches, err := filepath.Glob(base + "*")
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("retained files = %v, want 2 files", matches)
	}
}

func TestRotationCountKeepsNewestByModificationTime(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "app.log")
	older := base + ".10"
	newer := base + ".2"
	if err := os.WriteFile(older, []byte("older"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(newer, []byte("newer"), 0644); err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	if err := os.Chtimes(older, now.Add(-2*time.Hour), now.Add(-2*time.Hour)); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(newer, now.Add(-time.Hour), now.Add(-time.Hour)); err != nil {
		t.Fatal(err)
	}

	rl, err := New(base, WithRotationCount(2))
	if err != nil {
		t.Fatal(err)
	}
	defer rl.Close()
	if _, err := rl.Write([]byte("current")); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(older); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("oldest file should be removed, stat error = %v", err)
	}
	if _, err := os.Stat(newer); err != nil {
		t.Fatalf("newer generation should be retained: %v", err)
	}
	if _, err := os.Stat(base); err != nil {
		t.Fatalf("current file should be retained: %v", err)
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
			WithClock(clockFn(func() time.Time { return xt })),
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
