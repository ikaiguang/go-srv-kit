//go:build manual

package writerpkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// go test -v -count 1 ./writer -run TestManualWriteRotateFile
func TestManualWriteRotateFile(t *testing.T) {
	logDir := filepath.Join(".", "manual-logs")
	if err := os.RemoveAll(logDir); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(logDir, 0755); err != nil {
		t.Fatal(err)
	}

	writeSmallLog(t, logDir)
	writeRotateLog(t, logDir)
}

func writeSmallLog(t *testing.T, logDir string) {
	t.Helper()

	writer, err := NewRotateFile(&ConfigRotate{
		Dir:        logDir,
		Filename:   "app-3k",
		RotateSize: 1 << 20,
	})
	if err != nil {
		t.Fatal(err)
	}

	line := strings.Repeat("x", 1024)
	for i := 0; i < 3; i++ {
		if _, err = fmt.Fprintf(writer, "line=%d payload=%s\n", i+1, line); err != nil {
			t.Fatal(err)
		}
	}
}

func writeRotateLog(t *testing.T, logDir string) {
	t.Helper()

	writer, err := NewRotateFile(&ConfigRotate{
		Dir:            logDir,
		Filename:       "app-rotate",
		RotateSize:     1 << 20,
		StorageCounter: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	payload := strings.Repeat("r", 256*1024)
	for i := 0; i < 5; i++ {
		if _, err = io.WriteString(writer, payload+"\n"); err != nil {
			t.Fatal(err)
		}
	}
}
