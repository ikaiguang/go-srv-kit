package writerpkg

import (
	"io"
)

// dummy .
type dummy struct{}

// NewDummyWriter 假的
func NewDummyWriter() (io.Writer, error) {
	return &dummy{}, nil
}

// Write .
func (d *dummy) Write(p []byte) (int, error) {
	return 0, nil
}

func (d *dummy) Close() error {
	return nil
}
