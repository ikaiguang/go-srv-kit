package writerpkg

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestNewRotateFile(t *testing.T) {
	conf := &ConfigRotate{
		Dir:            t.TempDir(),
		Filename:       "test",
		RotateSize:     1,
		StorageCounter: 10,
		Compress:       true,
	}
	writer, err := NewRotateFile(conf)
	require.NoError(t, err)

	logger, ok := writer.(*lumberjack.Logger)
	require.True(t, ok)
	assert.Equal(t, filepath.Join(conf.Dir, "test.log"), logger.Filename)
	assert.Equal(t, 1, logger.MaxSize)
	assert.Equal(t, 10, logger.MaxBackups)
	assert.Equal(t, 0, logger.MaxAge)
	assert.True(t, logger.LocalTime)
	assert.True(t, logger.Compress)
}

func TestNewRotateFileWithFilenameSuffix(t *testing.T) {
	conf := &ConfigRotate{
		Dir:        t.TempDir(),
		Filename:   "test",
		RotateSize: DefaultRotationSize,
	}
	writer, err := NewRotateFile(conf, WithFilenameSuffix(".json.log"))
	require.NoError(t, err)

	logger, ok := writer.(*lumberjack.Logger)
	require.True(t, ok)
	assert.Equal(t, filepath.Join(conf.Dir, "test.json.log"), logger.Filename)
}

func TestNewRotateFileRejectsOnlyRotateTime(t *testing.T) {
	writer, err := NewRotateFile(&ConfigRotate{
		Dir:        t.TempDir(),
		Filename:   "test",
		RotateTime: time.Hour,
	})
	require.Error(t, err)
	require.Nil(t, writer)
	assert.Contains(t, err.Error(), "rotate time is not supported")
}

func TestNewRotateFileStorageAge(t *testing.T) {
	writer, err := NewRotateFile(&ConfigRotate{
		Dir:        t.TempDir(),
		Filename:   "test",
		RotateSize: DefaultRotationSize,
		StorageAge: 25 * time.Hour,
	})
	require.NoError(t, err)

	logger, ok := writer.(*lumberjack.Logger)
	require.True(t, ok)
	assert.Equal(t, 2, logger.MaxAge)
	assert.Equal(t, 0, logger.MaxBackups)
}

func TestNewRotateFileDefaultConfig(t *testing.T) {
	writer, err := NewRotateFile(&ConfigRotate{
		Dir:      t.TempDir(),
		Filename: "test",
	})
	require.NoError(t, err)

	logger, ok := writer.(*lumberjack.Logger)
	require.True(t, ok)
	assert.Equal(t, 100, logger.MaxSize)
	assert.Equal(t, 30, logger.MaxAge)
	assert.Equal(t, 0, logger.MaxBackups)
}

func TestNewRotateFileRejectsNilConfig(t *testing.T) {
	writer, err := NewRotateFile(nil)
	require.Error(t, err)
	require.Nil(t, writer)
}

func TestNewRotateFileRotatesBySizeWithTimestampBackup(t *testing.T) {
	dir := t.TempDir()
	writer, err := NewRotateFile(&ConfigRotate{
		Dir:        dir,
		Filename:   "test",
		RotateSize: 1,
	})
	require.NoError(t, err)

	payload := strings.Repeat("a", 256*1024)
	for i := 0; i < 5; i++ {
		_, err = io.WriteString(writer, payload)
		require.NoError(t, err)
	}

	logger, ok := writer.(*lumberjack.Logger)
	require.True(t, ok)
	require.NoError(t, logger.Close())

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)

	var hasActiveFile bool
	var hasTimestampBackup bool
	for _, entry := range entries {
		name := entry.Name()
		switch {
		case name == "test.log":
			hasActiveFile = true
		case strings.HasPrefix(name, "test-") && strings.HasSuffix(name, ".log"):
			hasTimestampBackup = true
		}
	}
	assert.True(t, hasActiveFile)
	assert.True(t, hasTimestampBackup)
}

func TestBytesToMegabytes(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want int
	}{
		{name: "zero", size: 0, want: 0},
		{name: "one byte", size: 1, want: 1},
		{name: "one megabyte", size: bytesPerMegabyte, want: 1},
		{name: "one megabyte plus one", size: bytesPerMegabyte + 1, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bytesToMegabytes(tt.size))
		})
	}
}

func TestDurationToDays(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		want     int
	}{
		{name: "zero", duration: 0, want: 0},
		{name: "one hour", duration: time.Hour, want: 1},
		{name: "one day", duration: 24 * time.Hour, want: 1},
		{name: "one day plus one hour", duration: 25 * time.Hour, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, durationToDays(tt.duration))
		})
	}
}
