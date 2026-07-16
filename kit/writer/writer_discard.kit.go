package writerpkg

import (
	"io"
)

type discardWriter struct{}

// NewDiscardWriter returns a writer that successfully discards all data.
func NewDiscardWriter() (io.Writer, error) {
	return &discardWriter{}, nil
}

// Deprecated: use NewDiscardWriter instead.
func NewDummyWriter() (io.Writer, error) {
	return NewDiscardWriter()
}

func (d *discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (d *discardWriter) Close() error {
	return nil
}
